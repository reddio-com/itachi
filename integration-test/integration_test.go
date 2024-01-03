package integration_test

import (
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
	"testing"
	"time"
)

var chain *kernel.Kernel

func TestIntegration(t *testing.T) {
	go func() {
		runItachiMockVM()
		time.AfterFunc(30*time.Second, chain.Stop)
	}()

}

func runItachiMockVM() {
	startup.InitDefaultConfig()
	poaCfg := poa.DefaultCfg(0)
	crCfg := cairo.DefaultCfg()

	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain = startup.InitDefaultKernel(
		poaTri, cairoTri,
	)
	chain.WithExecuteFn(cairoTri.Execute)
	chain.Startup()
}
