package app

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
)

func StartUpChain(poaCfg *poa.PoaConfig, crCfg *cairo.Config) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain := startup.InitDefaultKernel(
		poaTri, cairoTri,
	)
	chain.WithExecuteFn(cairoTri.Execute)
	chain.Startup()
}
