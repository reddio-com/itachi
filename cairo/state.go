package cairo

import (
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/db/pebble"
	"github.com/NethermindEth/juno/utils"
	"itachi/cairo/config"
)

type CairoState struct {
	*junostate.PendingStateWriter
	*core.State
}

func NewCairoState(cfg *config.Config) (*CairoState, error) {
	state, err := newState(cfg)
	if err != nil {
		return nil, err
	}
	pendingState := junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), state)
	return &CairoState{
		PendingStateWriter: pendingState,
		State:              state,
	}, nil
}

func (cs *CairoState) Commit(blockNum uint64) (*felt.Felt, error) {
	stateDiff, newClasses := cs.StateDiffAndClasses()
	err := cs.State.Update(blockNum, stateDiff, newClasses)
	if err != nil {
		return nil, err
	}
	cs.PendingStateWriter = junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), cs.State)
	return cs.State.Root()
}

func newState(cfg *config.Config) (*core.State, error) {
	dbLog, err := utils.NewZapLogger(utils.ERROR, cfg.Colour)
	if err != nil {
		return nil, err
	}
	db, err := pebble.New(cfg.DbPath, cfg.DbCache, cfg.DbMaxOpenFiles, dbLog)
	if err != nil {
		return nil, err
	}
	txn, err := db.NewTransaction(true)
	if err != nil {
		return nil, err
	}
	state := core.NewState(txn)
	return state, nil
}
