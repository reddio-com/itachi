package app

import (
	"context"
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

func StartUpChain(poaCfg *poa.PoaConfig, crCfg *config.Config) (*kernel.Kernel, error) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	chain, err := InitItachi(poaCfg, crCfg)
	if err != nil {
		return nil, err
	}
	rpcSrv := starknetrpc.StartUpStarknetRPC(chain, crCfg)

	// Subscribe to L1
	l1Client, err := l1.NewL1(chain, crCfg)
	if err != nil {
		panic(err)
	}
	l1Client.Run(context.Background(), rpcSrv)

	utils.StartUpPprof(crCfg)
	chain.Startup()

	return chain, nil
}

func InitItachi(poaCfg *poa.PoaConfig, crCfg *config.Config) (*kernel.Kernel, error) {
	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain := startup.InitDefaultKernel(
		poaTri, cairoTri,
	)
	return chain, nil
}
