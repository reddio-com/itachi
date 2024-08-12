package app

import (
	"itachi/cairo"
	"itachi/cairo/config"
	"itachi/cairo/l1"
	"itachi/cairo/starknetrpc"
	"itachi/utils"

	yuconfig "github.com/yu-org/yu/config"

	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"
)

func StartUpChain(poaCfg *poa.PoaConfig, crCfg *config.Config, yuCfg *yuconfig.KernelConf) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	chain := InitItachi(poaCfg, crCfg, yuCfg)

	// Starknet RPC server
	rpcSrv := starknetrpc.StartUpStarknetRPC(chain, crCfg)

	// Subscribe to L1
	l1.StartupL1(chain, crCfg, rpcSrv)

	utils.StartUpPprof(crCfg)
	chain.Startup()

}

func InitItachi(poaCfg *poa.PoaConfig, crCfg *config.Config, yuCfg *yuconfig.KernelConf) *kernel.Kernel {
	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain := startup.InitDefaultKernel(
		yuCfg, poaTri, cairoTri,
	)
	return chain
}
