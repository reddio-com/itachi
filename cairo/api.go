package cairo

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/yu-org/yu/core/context"
	"net/http"
)

type NonceRequest struct {
	BlockID rpc.BlockID `json:"block_id"`
	Addr    felt.Felt   `json:"addr"`
}

type NonceResponse struct {
	Nonce *felt.Felt     `json:"nonce"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetNonce(ctx *context.ReadContext) {
	var nq NonceRequest
	err := ctx.BindJson(&nq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, NonceResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}

	var nonce *felt.Felt
	switch {
	case nq.BlockID.Latest:
		nonce, err = c.cairoState.ContractNonce(&nq.Addr)
	default:
		nonce, err = c.cairoState.ContractNonceAt(&nq.Addr, nq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, NonceResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	ctx.JsonOk(NonceResponse{Nonce: nonce})
}

type StorageRequest struct {
	BlockID rpc.BlockID `json:"block_id"`
	Addr    felt.Felt   `json:"addr"`
	Key     felt.Felt   `json:"key"`
}

type StorageResponse struct {
	Value *felt.Felt     `json:"value"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetStorage(ctx *context.ReadContext) {

}

type ClassRequest struct {
	BlockID   rpc.BlockID `json:"block_id"`
	ClassHash felt.Felt   `json:"class_hash"`
}

type ClassResponse struct {
	Class *rpc.Class     `json:"class"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetClass(ctx *context.ReadContext) {

}

type ClassHashRequest struct {
	BlockID rpc.BlockID `json:"block_id"`
	Addr    felt.Felt   `json:"addr"`
}

type ClassHashResponse struct {
	ClassHash *felt.Felt     `json:"class_hash"`
	Err       *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetClassHash(ctx *context.ReadContext) {

}
