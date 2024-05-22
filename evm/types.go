package evm

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
)

type CallRequest struct {
	ContractAddr *felt.Felt  `json:"contract_addr"`
	Selector     *felt.Felt  `json:"selector"`
	Calldata     []felt.Felt `json:"calldata"`
	// BlockID      BlockID     `json:"block_id"`
}

type CallResponse struct {
	ReturnData []*felt.Felt   `json:"return_data"`
	Err        *jsonrpc.Error `json:"err"`
}

type TxRequest struct {
	Input           []byte                      `json:"input"`
	Code            []byte						`json:"code"`
}