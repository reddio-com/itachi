package main

import (
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo/config"
	"itachi/cmd/node/app"
	"itachi/evm"
)

func main() {
	startup.InitDefaultKernelConfig()
	poaCfg := poa.DefaultCfg(0)
	cairoCfg := config.LoadCairoCfg("./conf/cairo_cfg.toml")

	app.StartUpChain(poaCfg, cairoCfg, &evm.Config{})
}
