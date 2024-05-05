package evm

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/ethereum/go-ethereum/triedb/hashdb"
	"github.com/ethereum/go-ethereum/triedb/pathdb"
	"path/filepath"
)

type State struct {
	cfg     *Config
	stateDB *state.StateDB
	trieDB  *triedb.Database
	snaps   *snapshot.Tree
	logger  *tracing.Hooks
}

func NewState(cfg *Config) (*State, error) {
	vmConfig := vm.Config{
		EnablePreimageRecording: cfg.EnablePreimageRecording,
	}
	if cfg.VMTrace != "" {
		var traceConfig json.RawMessage
		if cfg.VMTraceConfig != "" {
			traceConfig = json.RawMessage(cfg.VMTraceConfig)
		}
		t, err := tracers.LiveDirectory.New(cfg.VMTrace, traceConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create tracer %s: %v", cfg.VMTrace, err)
		}
		vmConfig.Tracer = t
	}

	db, err := rawdb.Open(rawdb.OpenOptions{
		Type:              cfg.DbType,
		Directory:         cfg.DbPath,
		AncientsDirectory: filepath.Join(cfg.DbPath, cfg.Ancient),
		Namespace:         cfg.NameSpace,
		Cache:             cfg.Cache,
		Handles:           cfg.Handles,
		ReadOnly:          false,
	})
	if err != nil {
		return nil, err
	}
	scheme, err := rawdb.ParseStateScheme(cfg.StateScheme, db)
	if err != nil {
		return nil, err
	}
	cacheCfg := &core.CacheConfig{
		TrieCleanLimit:      cfg.TrieCleanCache,
		TrieCleanNoPrefetch: cfg.NoPrefetch,
		TrieDirtyLimit:      cfg.TrieDirtyCache,
		TrieDirtyDisabled:   cfg.NoPruning,
		TrieTimeLimit:       cfg.TrieTimeout,
		SnapshotLimit:       cfg.SnapshotCache,
		Preimages:           cfg.Preimages,
		StateHistory:        cfg.StateHistory,
		StateScheme:         scheme,
	}

	trieDB := triedb.NewDatabase(db, trieConfig(cacheCfg, false))

	stateCache := state.NewDatabaseWithNodeDB(db, trieDB)

	snapCfg := snapshot.Config{
		CacheSize:  cacheCfg.SnapshotLimit,
		Recovery:   cfg.Recovery,
		NoBuild:    cfg.NoBuild,
		AsyncBuild: !cfg.SnapshotWait,
	}

	snaps, err := snapshot.New(snapCfg, db, trieDB, types.EmptyRootHash /* current block hash */)
	if err != nil {
		return nil, err
	}

	statedb, err := state.New(types.EmptyRootHash, stateCache, snaps)
	if err != nil {
		return nil, err
	}

	return &State{
		cfg:     cfg,
		stateDB: statedb,
		trieDB:  trieDB,
		snaps:   snaps,
		logger:  vmConfig.Tracer,
	}, nil
}

func trieConfig(c *core.CacheConfig, isVerkle bool) *triedb.Config {
	config := &triedb.Config{
		Preimages: c.Preimages,
		IsVerkle:  isVerkle,
	}
	if c.StateScheme == rawdb.HashScheme {
		config.HashDB = &hashdb.Config{
			CleanCacheSize: c.TrieCleanLimit * 1024 * 1024,
		}
	}
	if c.StateScheme == rawdb.PathScheme {
		config.PathDB = &pathdb.Config{
			StateHistory:   c.StateHistory,
			CleanCacheSize: c.TrieCleanLimit * 1024 * 1024,
			DirtyCacheSize: c.TrieDirtyLimit * 1024 * 1024,
		}
	}
	return config
}

func (s *State) Commit(blockNum uint64) (common.Hash, error) {
	stateRoot, err := s.stateDB.Commit(blockNum, true)
	if err != nil {
		return common.Hash{}, err
	}
	err = s.trieDB.Commit(stateRoot, true)
	return stateRoot, err
}
