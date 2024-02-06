package integration_test

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
	"github.com/stretchr/testify/assert"
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

func init() {
	startup.InitDefaultKernelConfig()
	poaCfg := poa.DefaultCfg(0)
	crCfg := cairo.DefaultCfg()

	chain = app.InitItachi(poaCfg, crCfg)
}

func TestIntegration(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		time.AfterFunc(15*time.Second, chain.Stop)
		chain.Startup()
		wg.Done()
	}()

	time.Sleep(3 * time.Second) // wait for starting up

	//err := addTxToItachi("AddDeployAccountTxn", simulateBcTx())
	//assert.NoError(t, err)
	//err = addTxToItachi("AddDeclareTxn", simulateBcTx())
	//assert.NoError(t, err)
	//err = addTxToItachi("AddInvokeTxn", simulateBcTx())
	//assert.NoError(t, err)
	//err = addTxToItachi("AddL1HandleTxn", simulateBcTx())
	//assert.NoError(t, err)

	err := addTxToItachi("AddTxn", simulateBcTx())
	assert.NoError(t, err)

	time.Sleep(5 * time.Second)

	retData, err := callItachi(
		"Call",
		&cairo.CallRequest{
			ContractAddr: &felt.Zero,
			ClassHash:    &felt.Zero,
			Selector:     &felt.Zero,
			Calldata:     []felt.Felt{felt.Zero},
		})
	assert.NoError(t, err)
	t.Logf("the return data of Call is %v", retData)

	wg.Wait()

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

func simulateBcTx() *rpc.BroadcastedTransaction {
	return &rpc.BroadcastedTransaction{
		Transaction: rpc.Transaction{
			Hash:            new(felt.Felt).SetUint64(1),
			Type:            rpc.TxnInvoke,
			Version:         new(core.TransactionVersion).SetUint64(1).AsFelt(),
			Nonce:           new(felt.Felt).SetUint64(10),
			MaxFee:          new(felt.Felt).SetUint64(6),
			ContractAddress: new(felt.Felt).SetUint64(7),
			SenderAddress:   new(felt.Felt).SetUint64(888),
			Signature: &[]*felt.Felt{
				new(felt.Felt).SetUint64(4),
				new(felt.Felt).SetUint64(5),
			},
			CallData: &[]*felt.Felt{
				new(felt.Felt).SetUint64(2),
				new(felt.Felt).SetUint64(3),
			},
			EntryPointSelector: new(felt.Felt).SetUint64(9),
			Tip:                new(felt.Felt).SetUint64(10),
		},
	}
}
