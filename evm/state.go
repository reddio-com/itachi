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
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/ethereum/go-ethereum/triedb/hashdb"
	"github.com/ethereum/go-ethereum/triedb/pathdb"
	"path/filepath"
)

type EthState struct {
	cfg        *Config
	stateCache state.Database
	trieDB     *triedb.Database
	snaps      *snapshot.Tree
	logger     *tracing.Hooks
}

func NewEthState(cfg *Config, currentStateRoot common.Hash) (*EthState, error) {
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

	cacheCfg, err := cacheConfig(cfg, db)
	if err != nil {
		return nil, err
	}
	snapCfg := snapsConfig(cfg)

	trieDB := triedb.NewDatabase(db, trieConfig(cacheCfg, false))
	stateCache := state.NewDatabaseWithNodeDB(db, trieDB)

	snaps, err := snapshot.New(snapCfg, db, trieDB, currentStateRoot)
	if err != nil {
		return nil, err
	}

	return &EthState{
		cfg:        cfg,
		stateCache: stateCache,
		trieDB:     trieDB,
		snaps:      snaps,
		logger:     vmConfig.Tracer,
	}, nil
}

func (s *EthState) GenesisStateDB() (*state.StateDB, error) {
	return state.New(types.EmptyRootHash, s.stateCache, nil)
}

func (s *EthState) NewStateDB(parentStateRoot common.Hash) (statedb *state.StateDB, err error) {
	statedb, err = state.New(parentStateRoot, s.stateCache, s.snaps)
	if err != nil {
		return
	}
	statedb.SetLogger(s.logger)
	// Enable prefetching to pull in trie node paths while processing transactions
	statedb.StartPrefetcher("chain")
	return
}

func (s *EthState) Commit(blockNum uint64, stateDB *state.StateDB) (common.Hash, error) {
	stateDB.StopPrefetcher()
	stateRoot, err := stateDB.Commit(blockNum, true)
	if err != nil {
		return common.Hash{}, err
	}
	err = s.trieDB.Commit(stateRoot, true)
	return stateRoot, err
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

func cacheConfig(cfg *Config, db ethdb.Database) (*core.CacheConfig, error) {
	scheme, err := rawdb.ParseStateScheme(cfg.StateScheme, db)
	if err != nil {
		return nil, err
	}
	return &core.CacheConfig{
		TrieCleanLimit:      cfg.TrieCleanCache,
		TrieCleanNoPrefetch: cfg.NoPrefetch,
		TrieDirtyLimit:      cfg.TrieDirtyCache,
		TrieDirtyDisabled:   cfg.NoPruning,
		TrieTimeLimit:       cfg.TrieTimeout,
		SnapshotLimit:       cfg.SnapshotCache,
		Preimages:           cfg.Preimages,
		StateHistory:        cfg.StateHistory,
		StateScheme:         scheme,
	}, nil
}

func snapsConfig(cfg *Config) snapshot.Config {
	return snapshot.Config{
		CacheSize:  cfg.SnapshotCache,
		Recovery:   cfg.Recovery,
		NoBuild:    cfg.NoBuild,
		AsyncBuild: !cfg.SnapshotWait,
	}
}
