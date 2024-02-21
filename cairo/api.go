package cairo

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/yu-org/yu/core/context"
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
		ctx.JsonOk(NonceResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}
	nonce, err := c.cairoState.ContractNonce(&nq.Addr)
	if err != nil {
		ctx.JsonOk(NonceResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
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
