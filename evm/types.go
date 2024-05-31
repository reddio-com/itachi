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
	Ret         []byte         `json:"ret"`
	LeftOverGas uint64         `json:"leftOverGas"`
	Err         *jsonrpc.Error `json:"err"`
}

type TxRequest struct {
	Input []byte `json:"input"`
	Code  []byte `json:"code"`
}

type CreateRequest struct {
	Input []byte `json:"input"`
}

// type CreateResponse struct {
// 	Ret          []byte         `json:"ret"`
// 	ContractAddr common.Address `json:"contractAddr"`
// 	LeftOverGas  uint64         `json:"leftOverGas"`
// 	Err          error          `json:"err"`
// }
