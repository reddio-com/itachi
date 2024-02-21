package cairo

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
)

type CallRequest struct {
	ContractAddr *felt.Felt  `json:"contract_addr"`
	Selector     *felt.Felt  `json:"selector"`
	Calldata     []felt.Felt `json:"calldata"`
}

type CallResponse struct {
	ReturnData []*felt.Felt `json:"return_data"`
}

type TxRequest struct {
	Tx              *rpc.BroadcastedTransaction `json:"tx"`
	GasPriceWEI     *felt.Felt                  `json:"gas_price_wei"`
	GasPriceSTRK    *felt.Felt                  `json:"gas_price_strk"`
	LegacyTraceJson bool                        `json:"legacy_trace_json"`
}
