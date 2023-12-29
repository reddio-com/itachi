package cairo

type Config struct {
	Network int `toml:"network"`

	// pebble db
	Path         string `toml:"path"`
	Cache        uint   `toml:"cache"`
	MaxOpenFiles int    `toml:"max_open_files"`

	// cairoVM
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
