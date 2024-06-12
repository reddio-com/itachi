package evm

import (
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type CallRequest struct {
	Input    []byte         `json:"input"`
	Address  common.Address `json:"address"`
	Origin   common.Address `json:"origin"`
	GasLimit uint64         `json:"gasLimit"`
	Value    *big.Int       `json:"value"`
}

type CallResponse struct {
	Ret         []byte         `json:"ret"`
	LeftOverGas uint64         `json:"leftOverGas"`
	Err         *jsonrpc.Error `json:"err"`
}

type TxRequest struct {
	Input    []byte         `json:"input"`
	Code     []byte         `json:"code"`
	Origin   common.Address `json:"origin"`
	GasLimit uint64         `json:"gasLimit"`
	Value    *big.Int       `json:"value"`
}

type CreateRequest struct {
	Input    []byte         `json:"input"`
	Origin   common.Address `json:"origin"`
	GasLimit uint64         `json:"gasLimit"`
	Value    *big.Int       `json:"value"`
}
