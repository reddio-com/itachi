package cairo

import (
	"encoding/json"
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/node"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/result"
	"github.com/yu-org/yu/core/tripod"
	"github.com/yu-org/yu/core/types"
	"net/http"
)

type Cairo struct {
	*tripod.Tripod
	cairoVM       vm.VM
	cairoState    *CairoState
	cfg           *Config
	sequencerAddr *felt.Felt
	network       utils.Network
}

func NewCairo(cfg *Config) *Cairo {
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

	cairo.SetWritings(cairo.ExecuteTxn /*cairo.AddDeployAccountTxn, cairo.AddDeclareTxn, cairo.AddInvokeTxn, cairo.AddL1HandleTxn*/)
	cairo.SetReadings(cairo.Call)
	cairo.SetInit(cairo)
	cairo.SetTxnChecker(cairo)

	return cairo
}

func newVM(cfg *Config) (vm.VM, error) {
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
	var (
		starknetTxns = make([]core.Transaction, 0)
		classes      = make([]core.Class, 0)
		paidFeesOnL1 = make([]*felt.Felt, 0)
	)
	txReq := new(TxRequest)
	err := ctx.BindJson(txReq)
	if err != nil {
		return err
	}
	tx, class, paidFeeOnL1, err := c.adaptBroadcastedTransaction(txReq.Tx)
	if err != nil {
		return err
	}
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
	callRequest := new(CallRequest)
	err := ctx.BindJson(callRequest)
	if err != nil {
		ctx.Err(
			http.StatusBadRequest,
			errors.Errorf("Json decoded CallRequest failed: %v", err),
		)
		return
	}

	block, err := c.GetCurrentBlock()
	if err != nil {
		ctx.Err(
			http.StatusInternalServerError,
			errors.Errorf("Get current block failed: %v", err),
		)
		return
	}
	blockNumber := uint64(block.Height)
	blockTimestamp := block.Timestamp

	retData, err := c.cairoVM.Call(
		callRequest.ContractAddr,
		callRequest.ClassHash,
		callRequest.Selector,
		callRequest.Calldata,
		blockNumber, blockTimestamp,
		c.cairoState.state, c.network,
	)
	if err != nil {
		ctx.Err(
			http.StatusInternalServerError,
			errors.Errorf("CairoVM call failed: %v", err),
		)
		return
	}

	ctx.JsonOk(&CallResponse{ReturnData: retData})
}

//func (c *Cairo) newPendingStateWriter() *junostate.PendingStateWriter {
//	return junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), c.cairoState)
//}
