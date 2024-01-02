package integration_test

import (
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
	"itachi/cairo"
	"itachi/cmd/node/app"
	"testing"
)

func TestItachiMockVM(t *testing.T) {
	startup.InitDefaultConfig()
	poaCfg := poa.DefaultCfg(0)
	crCfg := cairo.DefaultCfg()
	app.StartUpChain(poaCfg, crCfg)
}
