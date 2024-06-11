package evm

import (
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/ethereum/go-ethereum/common"
)

type CallRequest struct {
	Input   []byte         `json:"input"`
	Address common.Address `json:"address"`
	Origin  common.Address `json:"origin"`
}

type CallResponse struct {
	Ret         []byte         `json:"ret"`
	LeftOverGas uint64         `json:"leftOverGas"`
	Err         *jsonrpc.Error `json:"err"`
}

type TxRequest struct {
	Input  []byte         `json:"input"`
	Code   []byte         `json:"code"`
	Origin common.Address `json:"origin"`
}

type CreateRequest struct {
	Input  []byte         `json:"input"`
	Origin common.Address `json:"origin"`
}
