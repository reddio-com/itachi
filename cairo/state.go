package cairo

import (
	"fmt"
	"itachi/cairo/config"

	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	junopebble "github.com/NethermindEth/juno/db/pebble"
	"github.com/NethermindEth/juno/encoder"
	"github.com/NethermindEth/juno/utils"
	"github.com/cockroachdb/pebble"
	"github.com/yu-org/yu/core/types"
)

type CairoState struct {
	*junostate.PendingStateWriter
	*core.State
	db *pebble.DB
}

func NewCairoState(cfg *config.Config) (*CairoState, error) {
	state, err := newState(cfg)
	if err != nil {
		return nil, err
	}
	db, err := newCairoStateDB(cfg)
	if err != nil {
		return nil, err
	}
	pendingState := junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), state)
	return &CairoState{
		PendingStateWriter: pendingState,
		State:              state,
		db:                 db,
	}, nil
}

// Create a StarkNet pebble database to store extra data
func newCairoStateDB(cfg *config.Config) (*pebble.DB, error) {
	db, err := pebble.Open(cfg.StateDiffDbPath, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (cs *CairoState) Commit(block *types.Block) (*felt.Felt, error) {
	blockNum := uint64(block.Height)

	oldRoot, err := cs.State.Root()
	if err != nil {
		return nil, err
	}
	stateDiff, newClasses := cs.StateDiffAndClasses()
	err = cs.State.Update(blockNum, stateDiff, newClasses)
	if err != nil {
		return nil, err
	}
	cs.PendingStateWriter = junostate.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), cs.State)

	newRoot, err := cs.State.Root()
	if err != nil {
		return nil, err
	}
	if oldRoot.Cmp(newRoot) == 0 {
		return newRoot, nil
	}

	stateUpdate := &core.StateUpdate{
		BlockHash: new(felt.Felt).SetBytes(block.Hash.Bytes()),
		NewRoot:   newRoot,
		OldRoot:   oldRoot,
		StateDiff: stateDiff,
	}
	storeStateUpdate(cs.db, blockNum, stateUpdate)

	// if err := cs.db.Close(); err != nil {
	// 	return nil, err
	// }
	return newRoot, nil
}

func newState(cfg *config.Config) (*core.State, error) {
	dbLog, err := utils.NewZapLogger(utils.ERROR, cfg.Colour)
	if err != nil {
		return nil, err
	}
	db, err := junopebble.New(cfg.DbPath, cfg.DbCache, cfg.DbMaxOpenFiles, dbLog)
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

func storeStateUpdate(db *pebble.DB, blockNumber uint64, update *core.StateUpdate) error {
	stateUpdateBytes, err := encoder.Marshal(update)
	if err != nil {
		return nil
	}
	numBytes := core.MarshalBlockNumber(blockNumber)
	key := []byte(fmt.Sprintf("StateUpdate:%d", numBytes))

	return db.Set(key, stateUpdateBytes, pebble.Sync)
}

func (cs *CairoState) GetStateUpdateByNumber(blockNumber uint64) (*core.StateUpdate, error) {
	numBytes := core.MarshalBlockNumber(blockNumber)
	var update *core.StateUpdate
	key := []byte(fmt.Sprintf("StateUpdate:%d", numBytes))
	val, closer, err := cs.db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	if err := encoder.Unmarshal(val, &update); err != nil {
		return nil, err
	}
	return update, nil
}

// func (cs *CairoState) GetStateUpdateByHash(hash *felt.Felt) (*core.StateUpdate, error) {
// 	var update *core.StateUpdate
// 	key := []byte(fmt.Sprintf("StateUpdate:%s", hash.String()))

// }
