package cairo

import (
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
)

// TODO: will use in cairo tripod

type CairoState struct {
	PendingState *junostate.PendingStateWriter
	state        *core.State
}

func NewCairoState(cfg *Config) (*CairoState, error) {
	state, err := newState(cfg)
	if err != nil {
		return nil, err
	}
	pendingState := junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), state)
	return &CairoState{
		PendingState: pendingState,
		state:        state,
	}, nil
}

func (cs *CairoState) Commit(blockNum uint64) error {
	stateDiff, newClasses := cs.PendingState.StateDiffAndClasses()
	err := cs.state.Update(blockNum, stateDiff, newClasses)
	if err != nil {
		return err
	}
	cs.PendingState = junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), cs.state)
	return nil
}
