package config

import (
	"github.com/BurntSushi/toml"
	"github.com/NethermindEth/juno/utils"
	"github.com/sirupsen/logrus"
)

type Config struct {
	// blockchain network
	// Mainnet Network =  1
	// Goerli = 2
	// Goerli2 = 3
	// Integration = 4
	// Sepolia = 5
	// SepoliaIntegration =6
	Network int `toml:"network"`

	// pebble db
	DbPath         string `toml:"db_path"`
	DbCache        uint   `toml:"db_cache"`
	DbMaxOpenFiles int    `toml:"db_max_open_files"`

	// cairoVM
	MockVM     bool  `toml:"mock_vm"`
	MaxVMs     uint  `toml:"max_vms"`
	MaxVMQueue int32 `toml:"max_vm_queue"`
	LogLevel   int   `toml:"log_level"`
	Colour     bool  `toml:"colour"`
	// cairo VM execute
	SequencerAddr string `toml:"sequencer_addr"`
	SkipChargeFee bool   `toml:"skip_charge_fee"`
	SkipValidate  bool   `toml:"skip_validate"`
	ErrOnRevert   bool   `toml:"err_on_revert"`

	// map[ClassHash]ClassFilePath
	GenesisClasses map[string]string `toml:"genesis_classes"`
	// map[ContractAddress]ClassHash
	GenesisContracts map[string]string `toml:"genesis_contracts"`
	GenesisStorages  []*GenesisStorage `toml:"genesis_storages"`

	// RPC configs
	EnableStarknetRPC bool   `toml:"enable_starknet_rpc"`
	StarknetHost      string `toml:"starknet_host"`
	StarknetPort      string `toml:"starknet_port"`

	// L1 configs
	EnableL1           bool   `toml:"enable_l1"`
	EthClientAddress   string `toml:"eth_client_address"`
	EthContractAddress string `toml:"eth_contract_address"`

	// L2 eth client configs
	EnableL2      bool   `toml:"enable_l2"`
	ChainID       int64  `toml:"chain_id"`
	GasLimit      uint64 `toml:"gas_limit"`
	EthPrivateKey string `toml:"eth_private_key"`

	EnablePprof bool   `toml:"enable_pprof"`
	PprofAddr   string `toml:"pprof_addr"`
}

type GenesisStorage struct {
	ContractAddress string `toml:"contract_address"`
	Key             string `toml:"key"`
	Value           string `toml:"value"`
}

func LoadCairoCfg(fpath string) *Config {
	cfg := new(Config)
	_, err := toml.DecodeFile(fpath, cfg)
	if err != nil {
		logrus.Fatalf("load config file failed: %v", err)
	}
	return cfg
}

func DefaultCfg() *Config {
	return &Config{
		MockVM:         true,
		Network:        int(utils.Integration),
		DbPath:         "cairo_db",
		DbCache:        1,
		DbMaxOpenFiles: 3,
		MaxVMs:         3,
		MaxVMQueue:     3,
		LogLevel:       1,
		Colour:         false,
		// test addr
		SequencerAddr:     "0x46a89ae102987331d369645031b49c27738ed096f2789c24449966da4c6de6b",
		SkipChargeFee:     true,
		SkipValidate:      true,
		ErrOnRevert:       true,
		EnableStarknetRPC: false,
	}
}
