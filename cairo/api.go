package cairo

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/yu-org/yu/core/context"
	"net/http"
)

type TransactionRequest struct {
	Hash felt.Felt `json:"hash"`
}

type TransactionResponse struct {
	Tx  *rpc.Transaction `json:"tx"`
	Err *jsonrpc.Error   `json:"err"`
}

func (c *Cairo) GetTransaction(ctx *context.ReadContext) {
	var tq TransactionRequest
	err := ctx.BindJson(&tq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, TransactionResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}
	signedTx, err := c.TxDB.GetTxn(tq.Hash.Bytes())
	if err != nil {
		ctx.Json(http.StatusInternalServerError, TransactionResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	txReq := new(TxRequest)
	err = signedTx.BindJson(txReq)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, TransactionResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	ctx.JsonOk(TransactionResponse{Tx: &txReq.Tx.Transaction})
}

type ReceiptRequest struct {
	Hash felt.Felt `json:"hash"`
}

type ReceiptResponse struct {
	Receipt *rpc.TransactionReceipt `json:"receipt"`
	Err     *jsonrpc.Error          `json:"err"`
}

func (c *Cairo) GetReceipt(ctx *context.ReadContext) {
	var rq ReceiptRequest
	err := ctx.BindJson(&rq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, ReceiptResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}

	starkReceipt, err := c.getReceipt(rq.Hash)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, ReceiptResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	ctx.JsonOk(ReceiptResponse{Receipt: starkReceipt})
}

func (c *Cairo) getReceipt(hash felt.Felt) (*rpc.TransactionReceipt, error) {
	receipt, err := c.TxDB.GetReceipt(hash.Bytes())
	if err != nil {
		return nil, err
	}
	starkReceipt := new(rpc.TransactionReceipt)
	err = json.Unmarshal(receipt.Extra, starkReceipt)
	return starkReceipt, err
}

type TransactionStatusRequest struct {
	Hash felt.Felt `json:"hash"`
}

type TransactionStatusResponse struct {
	Status *rpc.TransactionStatus `json:"status"`
	Err    *jsonrpc.Error         `json:"err"`
}

func (c *Cairo) GetTransactionStatus(ctx *context.ReadContext) {
	var tr TransactionStatusRequest
	err := ctx.BindJson(&tr)
	if err != nil {
		ctx.Json(http.StatusBadRequest, TransactionStatusResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}
	starkReceipt, err := c.getReceipt(tr.Hash)
	if err != nil {
		// TODO: when ErrTxnHashNotFound, should fetch from ETH L1
		ctx.Json(http.StatusInternalServerError, TransactionStatusResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	ctx.JsonOk(TransactionStatusResponse{Status: &rpc.TransactionStatus{
		Finality:  rpc.TxnStatus(starkReceipt.FinalityStatus),
		Execution: starkReceipt.ExecutionStatus,
	}})
}

type NonceRequest struct {
	BlockID rpc.BlockID `json:"block_id"`
	Addr    *felt.Felt  `json:"addr"`
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
		nonce, err = c.cairoState.ContractNonce(nq.Addr)
	default:
		nonce, err = c.cairoState.ContractNonceAt(nq.Addr, nq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, NonceResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	ctx.JsonOk(NonceResponse{Nonce: nonce})
}

type ClassRequest struct {
	BlockID   rpc.BlockID `json:"block_id"`
	ClassHash *felt.Felt  `json:"class_hash"`
}

type ClassAtRequest struct {
	BlockID rpc.BlockID `json:"block_id"`
	Addr    *felt.Felt  `json:"addr"`
}

type ClassResponse struct {
	Class *rpc.Class     `json:"class"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetClass(ctx *context.ReadContext) {
	var cq ClassRequest
	err := ctx.BindJson(&cq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, ClassResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}

	c.getClass(ctx, &cq.BlockID, cq.ClassHash)
}

func (c *Cairo) GetClassAt(ctx *context.ReadContext) {
	var cq ClassAtRequest
	err := ctx.BindJson(&cq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, ClassResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}
	var classHash *felt.Felt
	switch {
	case cq.BlockID.Latest:
		classHash, err = c.cairoState.ContractClassHash(cq.Addr)
	default:
		classHash, err = c.cairoState.ContractClassHashAt(cq.Addr, cq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, ClassResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}

	c.getClass(ctx, &cq.BlockID, classHash)
}

func (c *Cairo) getClass(ctx *context.ReadContext, blockID *rpc.BlockID, classHash *felt.Felt) {
	class, err := c.cairoState.Class(classHash)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, ClassResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	if !blockID.Latest {
		if blockID.Number < class.At {
			ctx.Json(http.StatusBadRequest, ClassResponse{Err: rpc.ErrClassHashNotFound})
			return
		}
	}
	rpcClass := declaredClassToClass(class)
	if rpcClass != nil {
		ctx.JsonOk(ClassResponse{Class: rpcClass})
	} else {
		ctx.Json(http.StatusBadRequest, ClassResponse{Err: rpc.ErrClassHashNotFound})
	}
}

type ClassHashRequest struct {
	BlockID rpc.BlockID `json:"block_id"`
	Addr    *felt.Felt  `json:"addr"`
}

type ClassHashResponse struct {
	ClassHash *felt.Felt     `json:"class_hash"`
	Err       *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetClassHash(ctx *context.ReadContext) {
	var cq ClassHashRequest
	err := ctx.BindJson(&cq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, ClassHashResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}

	var classHash *felt.Felt
	switch {
	case cq.BlockID.Latest:
		classHash, err = c.cairoState.ContractClassHash(cq.Addr)
	default:
		classHash, err = c.cairoState.ContractClassHashAt(cq.Addr, cq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, ClassHashResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err)})
		return
	}
	ctx.JsonOk(ClassHashResponse{ClassHash: classHash})
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

func declaredClassToClass(declared *core.DeclaredClass) (rpcClass *rpc.Class) {
	switch c := declared.Class.(type) {
	case *core.Cairo0Class:
		rpcClass = &rpc.Class{
			Abi:         c.Abi,
			Program:     c.Program,
			EntryPoints: rpc.EntryPoints{},
		}

		rpcClass.EntryPoints.Constructor = make([]rpc.EntryPoint, 0, len(c.Constructors))
		for _, entryPoint := range c.Constructors {
			rpcClass.EntryPoints.Constructor = append(rpcClass.EntryPoints.Constructor, rpc.EntryPoint{
				Offset:   entryPoint.Offset,
				Selector: entryPoint.Selector,
			})
		}

		rpcClass.EntryPoints.L1Handler = make([]rpc.EntryPoint, 0, len(c.L1Handlers))
		for _, entryPoint := range c.L1Handlers {
			rpcClass.EntryPoints.L1Handler = append(rpcClass.EntryPoints.L1Handler, rpc.EntryPoint{
				Offset:   entryPoint.Offset,
				Selector: entryPoint.Selector,
			})
		}

		rpcClass.EntryPoints.External = make([]rpc.EntryPoint, 0, len(c.Externals))
		for _, entryPoint := range c.Externals {
			rpcClass.EntryPoints.External = append(rpcClass.EntryPoints.External, rpc.EntryPoint{
				Offset:   entryPoint.Offset,
				Selector: entryPoint.Selector,
			})
		}

	case *core.Cairo1Class:
		rpcClass = &rpc.Class{
			Abi:                  c.Abi,
			SierraProgram:        c.Program,
			ContractClassVersion: c.SemanticVersion,
			EntryPoints:          rpc.EntryPoints{},
		}

		rpcClass.EntryPoints.Constructor = make([]rpc.EntryPoint, 0, len(c.EntryPoints.Constructor))
		for _, entryPoint := range c.EntryPoints.Constructor {
			index := entryPoint.Index
			rpcClass.EntryPoints.Constructor = append(rpcClass.EntryPoints.Constructor, rpc.EntryPoint{
				Index:    &index,
				Selector: entryPoint.Selector,
			})
		}

		rpcClass.EntryPoints.L1Handler = make([]rpc.EntryPoint, 0, len(c.EntryPoints.L1Handler))
		for _, entryPoint := range c.EntryPoints.L1Handler {
			index := entryPoint.Index
			rpcClass.EntryPoints.L1Handler = append(rpcClass.EntryPoints.L1Handler, rpc.EntryPoint{
				Index:    &index,
				Selector: entryPoint.Selector,
			})
		}

		rpcClass.EntryPoints.External = make([]rpc.EntryPoint, 0, len(c.EntryPoints.External))
		for _, entryPoint := range c.EntryPoints.External {
			index := entryPoint.Index
			rpcClass.EntryPoints.External = append(rpcClass.EntryPoints.External, rpc.EntryPoint{
				Index:    &index,
				Selector: entryPoint.Selector,
			})
		}
	}
	return
}
