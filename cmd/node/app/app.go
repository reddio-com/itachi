package app

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
	"itachi/cairo/config"
	"itachi/starknetrpc"
	"itachi/utils"
)

func StartUpChain(poaCfg *poa.PoaConfig, crCfg *config.Config) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	chain := InitItachi(poaCfg, crCfg)
	starknetrpc.StartUpStarknetRPC(chain, crCfg)
	utils.StartUpPprof(crCfg)
	chain.Startup()
}

func InitItachi(poaCfg *poa.PoaConfig, crCfg *config.Config) *kernel.Kernel {
	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain := startup.InitDefaultKernel(
		poaTri, cairoTri,
	)
	return chain
}
