package main

import (
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
	"itachi/cmd/node/app"
)

func main() {
	startup.InitDefaultKernelConfig()
	poaCfg := poa.DefaultCfg(0)
	cairoCfg := cairo.LoadCairoCfg("./conf/cairo_cfg.toml")

	app.StartUpChain(poaCfg, cairoCfg)
}
