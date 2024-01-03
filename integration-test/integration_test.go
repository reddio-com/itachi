package integration_test

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"
	"github.com/yu-org/yu/example/client/callchain"
	"itachi/cairo"
	"itachi/cmd/node/app"
	"sync"
	"testing"
	"time"
)

var chain *kernel.Kernel

func TestIntegration(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		runItachiMockVM()
		time.AfterFunc(30*time.Second, chain.Stop)
		wg.Done()
	}()
	// TODO: test call/invoke chain
	wg.Wait()

}

func runItachiMockVM() {
	startup.InitDefaultConfig()
	poaCfg := poa.DefaultCfg(0)
	crCfg := cairo.DefaultCfg()

	chain = app.InitYu(poaCfg, crCfg)
	chain.Startup()
}

const CairoTripod = "cairo"

func addTxToItachi(funcName string, tx *rpc.BroadcastedTransaction) error {
	txReq := &cairo.TxRequest{
		Tx:              tx,
		GasPriceWEI:     nil,
		GasPriceSTRK:    nil,
		LegacyTraceJson: false,
	}
	byt, err := json.Marshal(txReq)
	if err != nil {
		return err
	}
	return callchain.CallChainByWriting(&common.WrCall{
		TripodName: CairoTripod,
		FuncName:   funcName,
		Params:     string(byt),
	})
}

func callItachi(funcName string, callReq *cairo.CallRequest) ([]*felt.Felt, error) {
	byt, err := json.Marshal(callReq)
	if err != nil {
		return nil, err
	}
	resp, err := callchain.CallChainByReading(&common.RdCall{
		TripodName: CairoTripod,
		FuncName:   funcName,
		Params:     string(byt),
	})
	if err != nil {
		return nil, err
	}
	callResp := new(cairo.CallResponse)
	err = json.Unmarshal(resp, callResp)
	if err != nil {
		return nil, err
	}
	return callResp.ReturnData, nil
}
