package app

import (
	"itachi/cairo"
	"itachi/cairo/config"
	"itachi/cairo/l1"
	"itachi/cairo/starknetrpc"
	"itachi/utils"

	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"
)

func StartUpChain(poaCfg *poa.PoaConfig, crCfg *config.Config) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	chain := InitItachi(poaCfg, crCfg)

	// Starknet RPC server
	rpcSrv := starknetrpc.StartUpStarknetRPC(chain, crCfg)

	// Subscribe to L1
	l1.StartupL1(chain, crCfg, rpcSrv)

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
