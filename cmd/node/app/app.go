package app

import (
	"context"
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
	"itachi/cairo/config"
	"itachi/starknetrpc"
)

func StartUpChain(poaCfg *poa.PoaConfig, crCfg *config.Config) {
	figure.NewColorFigure("Itachi", "big", "green", false).Print()

	chain := InitItachi(poaCfg, crCfg)
	if crCfg.EnableStarknetRPC {
		rpcSrv, err := starknetrpc.NewStarknetRPC(chain, crCfg)
		if err != nil {
			logrus.Fatalf("init starknetRPC server failed, %v", err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			defer cancel()
			err = rpcSrv.Serve(ctx)
			if err != nil {
				logrus.Errorf("starknetRPC serves failed, %v", err)
			}
		}()
	}
	chain.Startup()
}

func InitItachi(poaCfg *poa.PoaConfig, crCfg *config.Config) *kernel.Kernel {
	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain := startup.InitDefaultKernel(
		poaTri, cairoTri,
	)
	chain.WithExecuteFn(cairoTri.TxnExecute)
	return chain
}
