package app

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
)

func StartUpChain(poaCfg *poa.PoaConfig, cairoCfg *cairo.Config) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()
	startup.DefaultStartup(
		poa.NewPoa(poaCfg),
		cairo.NewCairo(cairoCfg),
	)
}
