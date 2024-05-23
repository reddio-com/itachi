package evm

import (
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/ethereum/go-ethereum/common"
)

type CallRequest struct {
	Input   []byte         `json:"input"`
	Address common.Address `json:"address"`
}

type CallResponse struct {
	ReturnData []byte         `json:"return_data"`
	Err        *jsonrpc.Error `json:"err"`
}

type TxRequest struct {
	Input []byte `json:"input"`
	Code  []byte `json:"code"`
}
