package main

import (
	"itachi/cairo/config"
	"itachi/cmd/node/app"

	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
)

func main() {
	startup.InitDefaultKernelConfig()
	poaCfg := poa.DefaultCfg(0)
	cairoCfg := config.LoadCairoCfg("./conf/cairo_cfg.toml")

	app.StartUpChain(poaCfg, cairoCfg)
}
