package cairo

import (
	"encoding/hex"
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/encoder"
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
		cairo.GetClassHashAt, cairo.GetNonce, cairo.GetStorage,
		cairo.GetTransaction, cairo.GetTransactionStatus, cairo.GetReceipt,
		cairo.SimulateTransactions,
		cairo.GetBlockWithTxs, cairo.GetBlockWithTxHashes,
	)
	cairo.SetInit(cairo)
	cairo.SetTxnChecker(cairo)
	cairo.SetCommitter(cairo)

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

func (c *Cairo) InitChain(genesisBlock *types.Block) {
	// init codec for juno types
	junostate.RegisterCoreTypesToEncoder()

	stateRoot, err := c.buildGenesis()
	if err != nil {
		logrus.Fatal("build genesis classes failed: ", err)
	}
	genesisBlock.StateRoot = stateRoot.Bytes()
}

func (c *Cairo) CheckTxn(txn *types.SignedTxn) error {
	txReq := new(TxRequest)
	err := txn.BindJson(txReq)
	if err != nil {
		return err
	}
	starkTx, _, _, err := c.adaptBroadcastedTransaction(txReq.Tx)
	if err != nil {
		return err
	}

	// Replace the txHash with the Hash of starknet Txn
	txn.TxnHash = starkTx.Hash().Bytes()

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
		starknetTxns = make([]core.Transaction, 0)
		classes      = make([]core.Class, 0)
		paidFeesOnL1 = make([]*felt.Felt, 0)
	)
	if tx != nil {
		starknetTxns = append(starknetTxns, tx)
	}
	if class != nil {
		classes = append(classes, class)
	}
	if paidFeeOnL1 != nil {
		paidFeesOnL1 = append(paidFeesOnL1, paidFeeOnL1)
	}

	blockNumber := uint64(ctx.Block.Height)
	blockTimestamp := ctx.Block.Timestamp

	// FIXME: GasPriceWEI, GasPriceSTRK should be filled.
	gasPrice := new(felt.Felt).SetUint64(1)
	actualFees, traces, err := c.execute(
		starknetTxns, classes, blockNumber, blockTimestamp,
		paidFeesOnL1, gasPrice, gasPrice, txReq.LegacyTraceJson,
	)

	var starkReceipt *rpc.TransactionReceipt
	if len(traces) > 0 && len(actualFees) > 0 {
		starkReceipt = makeStarkReceipt(traces[0], ctx.Block, tx, actualFees[0])
	}
	if err != nil {
		// fmt.Printf("execute txn(%s) error: %v \n", tx.Hash(), err)
		starkReceipt = makeErrStarkReceipt(ctx.Block, tx, err)
	}

	//spew.Dump(starkReceipt)

	receiptByt, err := encoder.Marshal(starkReceipt)
	if err != nil {
		return err
	}
	ctx.EmitExtra(receiptByt)
	return nil
}

func (c *Cairo) Commit(block *types.Block) {
	blockNumber := uint64(block.Height)
	stateRoot, err := c.cairoState.Commit(blockNumber)
	if err != nil {
		logrus.Errorf("cairo commit failed on Block(%d), error: %v", blockNumber, err)
	}
	block.StateRoot = stateRoot.Bytes()
}

func (c *Cairo) Call(ctx *context.ReadContext) {
	callReq := new(CallRequest)
	err := ctx.BindJson(callReq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &CallResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}

	block, err := c.GetCurrentBlock()
	if err != nil {
		ctx.JsonOk(&CallResponse{Err: rpc.ErrBlockNotFound})
		return
	}
	blockNumber := uint64(block.Height)
	blockTimestamp := block.Timestamp

	var classHash *felt.Felt

	switch {
	case callReq.BlockID.Latest || callReq.BlockID.Pending:
		classHash, err = c.cairoState.ContractClassHash(callReq.ContractAddr)
	default:
		classHash, err = c.cairoState.ContractClassHashAt(callReq.ContractAddr, callReq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusBadRequest, &CallResponse{Err: rpc.ErrContractNotFound})
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
		ctx.Json(http.StatusInternalServerError, &CallResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	ctx.JsonOk(&CallResponse{ReturnData: retData})
}

func makeErrStarkReceipt(block *types.Block, tx core.Transaction, err error) *rpc.TransactionReceipt {
	starkReceipt := new(rpc.TransactionReceipt)
	starkReceipt.RevertReason = err.Error()
	starkReceipt.ExecutionStatus = rpc.TxnFailure
	blockNum := uint64(block.Height)
	starkReceipt.BlockNumber = &blockNum
	starkReceipt.BlockHash = new(felt.Felt).SetBytes(block.Hash.Bytes())
	starkReceipt.FinalityStatus = rpc.TxnAcceptedOnL2
	starkReceipt.Hash = tx.Hash()
	starkReceipt.ActualFee = &rpc.FeePayment{
		Amount: &felt.Zero,
		Unit:   feeUnit(tx),
	}
	starkReceipt.ExecutionResources = &rpc.ExecutionResources{}
	starkReceipt.Events = make([]*rpc.Event, 0)
	starkReceipt.MessagesSent = make([]*rpc.MsgToL1, 0)

	switch v := tx.(type) {
	case *core.DeployTransaction:
		starkReceipt.ContractAddress = v.ContractAddress
		starkReceipt.Type = rpc.TxnDeploy
	case *core.DeployAccountTransaction:
		starkReceipt.ContractAddress = v.ContractAddress
		starkReceipt.Type = rpc.TxnDeployAccount
	case *core.L1HandlerTransaction:
		starkReceipt.MessageHash = "0x" + hex.EncodeToString(v.MessageHash())
		starkReceipt.Type = rpc.TxnL1Handler
	case *core.DeclareTransaction:
		starkReceipt.Type = rpc.TxnDeclare
	case *core.InvokeTransaction:
		starkReceipt.Type = rpc.TxnInvoke
	}

	return starkReceipt
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
	//spew.Dump(invocation)

	var starkReceipt rpc.TransactionReceipt

	allOrderedEvents := allEvents(*invocation)
	receiptEvents := make([]*rpc.Event, 0)
	for _, orderedEvent := range allOrderedEvents {
		receiptEvents = append(receiptEvents, &rpc.Event{
			From: orderedEvent.From,
			Keys: orderedEvent.Keys,
			Data: orderedEvent.Data,
		})
	}
	starkReceipt.Events = receiptEvents

	resources := invocation.ExecutionResources
	if resources == nil {
		resources = &vm.ExecutionResources{}
	}
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

	allOrderedMsgs := allMessages(*invocation)
	receiptMsgs := make([]*rpc.MsgToL1, 0)
	for _, msg := range allOrderedMsgs {
		receiptMsgs = append(receiptMsgs, &rpc.MsgToL1{
			From:    msg.From,
			To:      common.HexToAddress(msg.To),
			Payload: msg.Payload,
		})
	}
	starkReceipt.MessagesSent = receiptMsgs

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

func allEvents(invocation vm.FunctionInvocation) []vm.OrderedEvent {
	events := make([]vm.OrderedEvent, 0)
	for i := range invocation.Calls {
		events = append(events, allEvents(invocation.Calls[i])...)
	}
	return append(events, utils.Map(invocation.Events, func(e vm.OrderedEvent) vm.OrderedEvent {
		e.From = &invocation.ContractAddress
		return e
	})...)
}

func allMessages(invocation vm.FunctionInvocation) []vm.OrderedL2toL1Message {
	messages := make([]vm.OrderedL2toL1Message, 0)
	for i := range invocation.Calls {
		messages = append(messages, allMessages(invocation.Calls[i])...)
	}
	return append(messages, utils.Map(invocation.Messages, func(e vm.OrderedL2toL1Message) vm.OrderedL2toL1Message {
		e.From = &invocation.ContractAddress
		return e
	})...)
}

//func (c *Cairo) newPendingStateWriter() *junostate.PendingStateWriter {
//	return junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), c.cairoState)
//}
