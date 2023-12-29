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
	"github.com/yu-org/yu/core/types"
)

type Cairo struct {
	*tripod.Tripod
	cairoVM       vm.VM
	state         core.StateReader
	cfg           *Config
	sequencerAddr *felt.Felt
	network       utils.Network
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
	sequencerAddr, err := new(felt.Felt).SetString(cfg.SequencerAddr)
	if err != nil {
		logrus.Fatal("load sequencer address failed: ", err)
	}

	cairo := &Cairo{
		Tripod:        tripod.NewTripod(),
		cairoVM:       cairoVM,
		state:         state,
		cfg:           cfg,
		sequencerAddr: sequencerAddr,
		network:       utils.Network(cfg.Network),
	}

	cairo.SetWritings(cairo.AddDeployAccountTxn, cairo.AddDeclareTxn, cairo.AddInvokeTxn, cairo.AddL1HandleTxn)
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

func (c *Cairo) InitChain() {
	// TODO: init the genesis block
}

func (c *Cairo) CheckTxn(txn *types.SignedTxn) error {
	// TODO: check tx, if illegal, will not insert to txpool.
}

func (c *Cairo) AddDeployAccountTxn(ctx *context.WriteContext) error {
	return nil
}

func (c *Cairo) AddDeclareTxn(ctx *context.WriteContext) error {
	// TODO: check if contract is already declared
	return nil
}

func (c *Cairo) AddInvokeTxn(ctx *context.WriteContext) error {
	// TODO: check if contract is already declared
	return nil
}

func (c *Cairo) AddL1HandleTxn(ctx *context.WriteContext) error {
	return nil
}

func (c *Cairo) Call(ctx *context.ReadContext) {
	block, err := c.Chain.GetEndBlock()
	if err != nil {

	}
	blockNumber := uint64(block.Height)
	blockTimestamp := block.Timestamp

}
