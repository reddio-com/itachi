package cairo

import (
	"encoding/hex"
	"fmt"
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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/tripod"
	"github.com/yu-org/yu/core/types"
	. "github.com/yu-org/yu/core/types"
	"github.com/yu-org/yu/utils/log"
	"itachi/cairo/adapters"
	"itachi/cairo/config"
	"itachi/cairo/l1"
	"itachi/cairo/l1/contract"
	snos_ouput "itachi/cairo/snos-ouput"
	"math/big"
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

func (c *Cairo) StartBlock(block *Block) {

	logrus.Infof("Cairo TriPod StartBlock for height (%d)! ", block.Height)

}

func (c *Cairo) EndBlock(block *Block) {

	logrus.Infof("Cairo TriPod EndBlock for height (%d)! ", block.Height)

}

func (c *Cairo) FinalizeBlock(block *Block) {
	logrus.Infof("Cairo TriPod finalize block for height (%d)! ", block.Height)

	// for PrevStateRoot, get last finalized block
	compactBlock, err := c.Chain.LastFinalized()
	if err != nil {
		logrus.Fatal("get compactBlock for finalize block failed: ", err)
	}

	var starkReceipt *rpc.TransactionReceipt
	txns := block.Txns.ToArray()
	messagesToL1 := make([]*rpc.MsgToL1, 0)
	for t := 0; t < len(txns); t++ {
		txn := txns[t]
		receipt, _ := c.TxDB.GetReceipt(txn.TxnHash)
		receiptExtraByt := receipt.Extra
		err := encoder.Unmarshal(receiptExtraByt, &starkReceipt)
		if err != nil {
			// handle error
			logrus.Fatal("unmarshal starkReceipt failed: ", err)
		} else {
			messagesToL1 = append(messagesToL1, starkReceipt.MessagesSent...)
		}
	}
	// Adapt
	messageL2ToL1 := make([]*adapters.MessageL2ToL1, len(messagesToL1))
	for idx, msg := range messagesToL1 {
		messageL2ToL1[idx] = &adapters.MessageL2ToL1{
			From:    msg.From,
			To:      msg.To,
			Payload: msg.Payload,
		}
	}

	// todo 	messagesToL2 := make([]*rpc.MsgFromL1, 0)
	messagesToL2 := make([]*adapters.MessageL1ToL2, 0)

	num := uint64(block.Height)
	// init StarknetOsOutput by block
	snOsOutput := &snos_ouput.StarknetOsOutput{
		PrevStateRoot: new(felt.Felt).SetBytes(compactBlock.StateRoot.Bytes()),
		NewStateRoot:  new(felt.Felt).SetBytes(block.StateRoot.Bytes()),
		BlockNumber:   new(felt.Felt).SetUint64(num),
		BlockHash:     new(felt.Felt).SetBytes(block.Hash.Bytes()),
		ConfigHash:    new(felt.Felt).SetUint64(0),
		KzgDA:         new(felt.Felt).SetUint64(0),
		MessagesToL1:  messageL2ToL1,
		MessagesToL2:  messagesToL2,
	}
	// cairoState.UpdateStarknetOsOutput(snOsOutput)
	logrus.Infof("snOsOutput: %+v", snOsOutput)

	// send snOsOutput to L1 chain
	c.ethCallUpdateState(c.cairoState, snOsOutput)

	log.DoubleLineConsole.Info(fmt.Sprintf("Cairo Tripod finalize block, height=%d, hash=%s", block.Height, block.Hash.String()))

}

func (c *Cairo) ethCallUpdateState(cairoState *CairoState, snOsOutput *snos_ouput.StarknetOsOutput) {

	ethClient, err := l1.NewEthClient(c.cfg.EthClientAddress)
	if err != nil {
		logrus.Errorf("init ethClient failed: %s", err)
	}

	starknetCore, err := contract.NewStarknetCore(common.HexToAddress(c.cfg.EthContractAddress), ethClient)
	if err != nil {
		return
	}

	// encode snOsOutput to []*big.Int
	programOutput, err := snOsOutput.EncodeTo()
	if err != nil {
		logrus.Errorf("encode snOsOutput failed: %s", err)
		return
	}

	// compute onchainDataHash and onchainDataSize
	onchainDataHash, onchainDataSize, err := calculateOnchainData(programOutput)
	if err != nil {
		logrus.Errorf("calculate onchain data failed: %s", err)
		return
	}

	chainID := big.NewInt(c.cfg.ChainID)
	privateKeyHex := c.cfg.EthPrivateKey
	address := c.cfg.EthClientAddress
	gasLimit := c.cfg.GasLimit
	auth, err := l1.CreateAuth(ethClient, privateKeyHex, address, gasLimit, chainID)
	if err != nil {
		logrus.Errorf("create auth failed: %s", err)
		return
	}

	// call updateState
	tx, err := starknetCore.UpdateState(auth, programOutput, onchainDataHash, onchainDataSize)
	if err != nil {
		logrus.Errorf("call updateState failed: %s", err)
		return
	}

	// retrieve transaction hash and print.
	logrus.Infof("update state, tx size: %d bytes", tx.Size())
	txHash := tx.Hash()
	logrus.Infof("update state, tx hash: %s", txHash.Hex())

}

func calculateOnchainData(programOutput []*big.Int) (*big.Int, *big.Int, error) {
	var data []byte
	for _, output := range programOutput {
		data = append(data, output.Bytes()...)
	}
	onchainDataHash := crypto.Keccak256Hash(data)
	onchainDataHashBig := new(big.Int).SetBytes(onchainDataHash.Bytes())

	// compute onchainDataSize
	onchainDataSize := new(big.Int).SetInt64(int64(len(programOutput)))

	return onchainDataHashBig, onchainDataSize, nil
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
	if err != nil {
		logrus.Fatal("open StarkNet DB failed: ", err)
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
		cairo.GetBlockHashAndNumber, cairo.GetBlockNumber,
		cairo.GetTransactionByBlockIDAndIndex, cairo.GetBlockTransactionCount,
		cairo.GetStateUpdate,
	)

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

func (c *Cairo) PreHandleTxn(txn *types.SignedTxn) error {
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
	// The processing of blockHash below is necessary to convert the Hash into a Felt-compatible uint256 type,
	// using the official Starknet conversion method, ensuring that Felt and uint256 can be converted without loss.
	// Without this processing, users would experience a loss of 4 bits of information when performing a uint252->256 bit conversion using starknet rpc.
	block.Hash = new(felt.Felt).SetBytes(block.Hash.Bytes()).Bytes()
	stateRoot, err := c.cairoState.Commit(block)
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
