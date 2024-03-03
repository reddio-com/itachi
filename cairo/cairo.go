package cairo

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/node"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/tripod"
	"github.com/yu-org/yu/core/types"
	"itachi/cairo/config"
	"net/http"
)

type Cairo struct {
	*tripod.Tripod
	cairoVM       vm.VM
	cairoState    *CairoState
	cfg           *config.Config
	sequencerAddr *felt.Felt
	network       utils.Network
}

func NewCairo(cfg *config.Config) *Cairo {
	state, err := NewCairoState(cfg)
	if err != nil {
		logrus.Fatal("init cairoState for Cairo failed: ", err)
	}
	cairoVM, err := newVM(cfg)
	if err != nil {
		logrus.Fatal("init cairoVM failed: ", err)
	}
	sequencerAddr, err := new(felt.Felt).SetString(cfg.SequencerAddr)
	if err != nil {
		logrus.Fatal("load sequencer address failed: ", err)
	}

	cairo := &Cairo{
		Tripod:        tripod.NewTripod(),
		cairoVM:       cairoVM,
		cairoState:    state,
		cfg:           cfg,
		sequencerAddr: sequencerAddr,
		network:       utils.Network(cfg.Network),
	}

	cairo.SetWritings(cairo.ExecuteTxn)
	cairo.SetReadings(
		cairo.Call, cairo.GetClass, cairo.GetClassAt,
		cairo.GetClassHash, cairo.GetNonce, cairo.GetStorage,
		cairo.GetTransaction, cairo.GetTransactionStatus, cairo.GetReceipt,
	)
	cairo.SetInit(cairo)
	cairo.SetTxnChecker(cairo)

	return cairo
}

func newVM(cfg *config.Config) (vm.VM, error) {
	log, err := utils.NewZapLogger(utils.LogLevel(cfg.LogLevel), cfg.Colour)
	if err != nil {
		return nil, err
	}
	if cfg.MockVM {
		logrus.Info("Mock CairoVM!")
		return new(MockCairoVM), nil
	}
	return node.NewThrottledVM(vm.New(log), cfg.MaxVMs, cfg.MaxVMQueue), nil
}

func (c *Cairo) InitChain() {
	// init codec for juno types
	junostate.RegisterCoreTypesToEncoder()

	err := c.buildGenesisClasses()
	if err != nil {
		logrus.Fatal("build genesis classes failed: ", err)
	}
}

func (c *Cairo) CheckTxn(txn *types.SignedTxn) error {
	txReq := new(TxRequest)
	err := txn.BindJson(txReq)
	if err != nil {
		return err
	}
	// Replace the txHash with the Hash of starknet Txn
	txn.TxnHash = txReq.Tx.Hash.Bytes()

	if txReq.Tx.Type == rpc.TxnDeclare && txReq.Tx.Version.Cmp(new(felt.Felt).SetUint64(2)) != -1 {
		contractClass := make(map[string]any)
		if err := json.Unmarshal(txReq.Tx.ContractClass, &contractClass); err != nil {
			return fmt.Errorf("unmarshal contract class: %v", err)
		}
		sierraProg, ok := contractClass["sierra_program"]
		if !ok {
			return fmt.Errorf("{'sierra_program': ['Missing data for required field.']}")
		}

		sierraProgBytes, errIn := json.Marshal(sierraProg)
		if errIn != nil {
			return errIn
		}

		gwSierraProg, errIn := utils.Gzip64Encode(sierraProgBytes)
		if errIn != nil {
			return errIn
		}

		contractClass["sierra_program"] = gwSierraProg
		newContractClass, err := json.Marshal(contractClass)
		if err != nil {
			return fmt.Errorf("marshal revised contract class: %v", err)
		}
		txReq.Tx.ContractClass = newContractClass
	}

	newTxReqByt, err := json.Marshal(txReq)
	if err != nil {
		return err
	}

	txn.SetParams(string(newTxReqByt))

	return nil
}

func (c *Cairo) ExecuteTxn(ctx *context.WriteContext) error {
	txReq := new(TxRequest)
	err := ctx.BindJson(txReq)
	if err != nil {
		return err
	}

	tx, class, paidFeeOnL1, err := c.adaptBroadcastedTransaction(txReq.Tx)
	if err != nil {
		return err
	}

	var (
		starknetTxns = []core.Transaction{tx}
		classes      = []core.Class{class}
		paidFeesOnL1 = []*felt.Felt{paidFeeOnL1}
	)

	blockNumber := uint64(ctx.Block.Height)
	blockTimestamp := ctx.Block.Timestamp

	// FIXME: GasPriceWEI, GasPriceSTRK and legacyTraceJSON should be filled.
	actualFees, traces, err := c.execute(
		starknetTxns, classes, blockNumber, blockTimestamp,
		paidFeesOnL1, &felt.Zero, &felt.Zero, false,
	)
	if err != nil {
		return err
	}

	var starkReceipt *rpc.TransactionReceipt
	if len(traces) > 0 && len(actualFees) > 0 {
		starkReceipt = makeStarkReceipt(traces[0], ctx.Block, tx, actualFees[0])
	}

	receiptByt, err := json.Marshal(starkReceipt)
	if err != nil {
		return err
	}
	ctx.EmitExtra(receiptByt)
	return nil
}

func (c *Cairo) Call(ctx *context.ReadContext) {
	callReq := new(CallRequest)
	err := ctx.BindJson(callReq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, CallResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}

	block, err := c.GetCurrentBlock()
	if err != nil {
		ctx.JsonOk(CallResponse{Err: rpc.ErrBlockNotFound})
		return
	}
	blockNumber := uint64(block.Height)
	blockTimestamp := block.Timestamp

	var classHash *felt.Felt

	switch {
	case callReq.BlockID.Latest:
		classHash, err = c.cairoState.ContractClassHash(callReq.ContractAddr)
	default:
		classHash, err = c.cairoState.ContractClassHashAt(callReq.ContractAddr, callReq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusBadRequest, CallResponse{Err: rpc.ErrContractNotFound})
		return
	}

	retData, err := c.cairoVM.Call(
		callReq.ContractAddr,
		classHash,
		callReq.Selector,
		callReq.Calldata,
		blockNumber, blockTimestamp,
		c.cairoState.State, c.network,
	)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, CallResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}

	ctx.JsonOk(CallResponse{ReturnData: retData})
}

func makeStarkReceipt(trace vm.TransactionTrace, block *types.Block, tx core.Transaction, amount *felt.Felt) *rpc.TransactionReceipt {
	starkReceipt := new(rpc.TransactionReceipt)
	switch {
	case trace.ExecuteInvocation != nil:
		starkReceipt = makeStarkReceiptFromInvocation(trace.ExecuteInvocation.FunctionInvocation)
	case trace.ValidateInvocation != nil:
		starkReceipt = makeStarkReceiptFromInvocation(trace.ValidateInvocation)
	}
	starkReceipt.RevertReason = trace.RevertReason()
	starkReceipt.Type = rpc.TransactionType(trace.Type)
	if trace.IsReverted() {
		starkReceipt.ExecutionStatus = rpc.TxnFailure
	} else {
		starkReceipt.ExecutionStatus = rpc.TxnSuccess
	}
	blockNum := uint64(block.Height)
	starkReceipt.BlockNumber = &blockNum
	starkReceipt.BlockHash = new(felt.Felt).SetBytes(block.Hash.Bytes())
	starkReceipt.FinalityStatus = rpc.TxnAcceptedOnL2
	starkReceipt.Hash = tx.Hash()
	starkReceipt.ActualFee = &rpc.FeePayment{
		Amount: amount,
		Unit:   feeUnit(tx),
	}

	switch v := tx.(type) {
	case *core.DeployTransaction:
		starkReceipt.ContractAddress = v.ContractAddress
	case *core.DeployAccountTransaction:
		starkReceipt.ContractAddress = v.ContractAddress
	case *core.L1HandlerTransaction:
		starkReceipt.MessageHash = "0x" + hex.EncodeToString(v.MessageHash())
	}

	return starkReceipt
}

func makeStarkReceiptFromInvocation(invocation *vm.FunctionInvocation) *rpc.TransactionReceipt {
	var starkReceipt rpc.TransactionReceipt
	for _, event := range invocation.Events {
		starkReceipt.Events = append(starkReceipt.Events, &rpc.Event{
			From: event.From,
			Keys: event.Keys,
			Data: event.Data,
		})
	}
	resources := invocation.ExecutionResources
	starkReceipt.ExecutionResources = &rpc.ExecutionResources{
		Steps:        resources.Steps,
		MemoryHoles:  resources.MemoryHoles,
		Pedersen:     resources.Pedersen,
		RangeCheck:   resources.RangeCheck,
		Bitwise:      resources.Bitwise,
		Ecsda:        resources.Ecdsa,
		EcOp:         resources.EcOp,
		Keccak:       resources.Keccak,
		Poseidon:     resources.Poseidon,
		SegmentArena: resources.SegmentArena,
	}

	for _, message := range invocation.Messages {
		starkReceipt.MessagesSent = append(starkReceipt.MessagesSent, &rpc.MsgToL1{
			From:    message.From,
			To:      common.HexToAddress(message.To),
			Payload: message.Payload,
		})
	}

	return &starkReceipt
}

func feeUnit(txn core.Transaction) rpc.FeeUnit {
	feeUnit := rpc.WEI
	version := txn.TxVersion()
	if !version.Is(0) && !version.Is(1) && !version.Is(2) {
		feeUnit = rpc.FRI
	}

	return feeUnit
}

//func (c *Cairo) newPendingStateWriter() *junostate.PendingStateWriter {
//	return junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), c.cairoState)
//}
