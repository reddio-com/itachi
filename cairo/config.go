package cairo

import (
	"github.com/BurntSushi/toml"
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
	Path         string `toml:"path"`
	Cache        uint   `toml:"cache"`
	MaxOpenFiles int    `toml:"max_open_files"`

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
}

func LoadCfg(fpath string) *Config {
	cfg := new(Config)
	_, err := toml.DecodeFile(fpath, cfg)
	if err != nil {
		logrus.Fatalf("load config file failed: %v", err)
	}
	return cfg
}

func DefaultCfg() *Config {
	return &Config{
		MockVM:       true,
		Network:      6,
		Path:         "cairo_db",
		Cache:        1,
		MaxOpenFiles: 3,
		MaxVMs:       3,
		MaxVMQueue:   3,
		LogLevel:     1,
		Colour:       false,
		// test addr
		SequencerAddr: "0x46a89ae102987331d369645031b49c27738ed096f2789c24449966da4c6de6b",
		SkipChargeFee: true,
		SkipValidate:  true,
		ErrOnRevert:   true,
	}
}
