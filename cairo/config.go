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
}
