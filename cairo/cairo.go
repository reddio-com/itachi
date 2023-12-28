package cairo

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/db/pebble"
	"github.com/NethermindEth/juno/node"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/tripod"
)

type Cairo struct {
	*tripod.Tripod
	cairoVM vm.VM
	state   core.StateReader
	cfg     *Config
}

func NewCairo(cfg *Config) *Cairo {
	state, err := newState(cfg)
	if err != nil {
		logrus.Fatal("init state for Cairo failed: ", err)
	}
	cairoVM, err := newVM(cfg)
	if err != nil {
		logrus.Fatal("init cairoVM failed: ", err)
	}

	cairo := &Cairo{
		Tripod:  tripod.NewTripod(),
		cairoVM: cairoVM,
		state:   state,
		cfg:     cfg,
	}

	cairo.SetWritings(cairo.AddTxn)
	cairo.SetReadings(cairo.Call)

	return cairo
}

func newVM(cfg *Config) (vm.VM, error) {
	log, err := utils.NewZapLogger(utils.LogLevel(cfg.LogLevel), cfg.Colour)
	if err != nil {
		return nil, err
	}
	return node.NewThrottledVM(vm.New(log), cfg.MaxVMs, cfg.MaxVMQueue), nil
}

func newState(cfg *Config) (core.StateReader, error) {
	dbLog, err := utils.NewZapLogger(utils.ERROR, cfg.Colour)
	if err != nil {
		return nil, err
	}
	db, err := pebble.New(cfg.Path, cfg.Cache, cfg.MaxOpenFiles, dbLog)
	if err != nil {
		return nil, err
	}
	txn, err := db.NewTransaction(true)
	if err != nil {
		return nil, err
	}
	return core.NewState(txn), nil
}

func (c *Cairo) AddTxn(ctx *context.WriteContext) error {

}

func (c *Cairo) Call(ctx *context.ReadContext) {

}

func (c *Cairo) call(
	contractAddr, classHash, selector *felt.Felt,
	calldata []felt.Felt,
	blockNumber, blockTimestamp uint64,
) ([]*felt.Felt, error) {
	return c.cairoVM.Call(contractAddr, classHash, selector, calldata, blockNumber, blockTimestamp, c.state, utils.Network(c.cfg.Network))
}

// FIXME: should implement startup.TxnExecute
func (c *Cairo) Execute(txns []core.Transaction, declaredClasses []core.Class, blockNumber, blockTimestamp uint64,
	sequencerAddress *felt.Felt, paidFeesOnL1 []*felt.Felt,
	skipChargeFee, skipValidate, errOnRevert bool, gasPriceWEI *felt.Felt, gasPriceSTRK *felt.Felt, legacyTraceJSON bool,
) ([]*felt.Felt, []vm.TransactionTrace, error) {
	return c.cairoVM.Execute(txns, declaredClasses, blockNumber, blockTimestamp, sequencerAddress,
		c.state, utils.Network(c.cfg.Network), paidFeesOnL1, skipChargeFee, skipValidate, errOnRevert, gasPriceWEI, gasPriceSTRK, legacyTraceJSON)
}
