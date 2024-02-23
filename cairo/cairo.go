package cairo

import (
	"encoding/json"
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/node"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/result"
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
	cairo.SetReadings(cairo.Call, cairo.GetClass, cairo.GetClassHash, cairo.GetNonce, cairo.GetStorage)
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
	// TODO: check tx, if illegal, will not insert to txpool.
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
	starknetTxns = append(starknetTxns, tx)
	classes = append(classes, class)
	paidFeesOnL1 = append(paidFeesOnL1, paidFeeOnL1)

	blockNumber := uint64(ctx.Block.Height)
	blockTimestamp := ctx.Block.Timestamp
	blockHash := ctx.Block.Hash

	// FIXME: GasPriceWEI, GasPriceSTRK and legacyTraceJSON should be filled.
	_, traces, err := c.execute(
		starknetTxns, classes, blockNumber, blockTimestamp,
		paidFeesOnL1, &felt.Zero, &felt.Zero, false,
	)
	if err != nil {
		return err
	}

	for _, trace := range traces {
		if trace.ExecuteInvocation != nil {
			for _, event := range trace.ExecuteInvocation.Events {
				eventByt, terr := json.Marshal(event)
				if terr != nil {
					return terr
				}
				callerByt := trace.ExecuteInvocation.CallerAddress.Bytes()
				caller := common.BytesToAddress(callerByt[:])
				yuEvent := &result.Event{
					Caller:    &caller,
					BlockHash: blockHash,
					Height:    ctx.Block.Height,
					Value:     eventByt,
				}
				ctx.EmitJsonEvent(result.NewEvent(yuEvent))
			}
		}

	}
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

//func (c *Cairo) newPendingStateWriter() *junostate.PendingStateWriter {
//	return junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), c.cairoState)
//}
