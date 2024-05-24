package app

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"

	"itachi/evm"
)

func StartUpEvmChain(poaCfg *poa.PoaConfig, crCfg *evm.Config) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	chain := InitEth(poaCfg, crCfg)
	// starknetrpc.StartUpStarknetRPC(chain, crCfg)
	// utils.StartUpPprof(crCfg)
	chain.Startup()
}

func InitEth(poaCfg *poa.PoaConfig, crCfg *evm.Config) *kernel.Kernel {
	poaTri := poa.NewPoa(poaCfg)
	solidityTri := evm.NewSolidity(crCfg)
	chain := startup.InitDefaultKernel(
		poaTri, solidityTri,
	)
	return chain
}
