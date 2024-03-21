package starknetrpc

import (
	"context"
	"encoding/json"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/starknet.go/hash"
	sdk "github.com/NethermindEth/starknet.go/rpc"
	"github.com/yu-org/yu/common"
	yucore "github.com/yu-org/yu/core"
	yucontext "github.com/yu-org/yu/core/context"
	"itachi/cairo"
)

func (s *StarknetRPC) GetChainID() (*felt.Felt, *jsonrpc.Error) {
	return s.network.ChainID(), nil
}

func (s *StarknetRPC) AddTransaction(tx rpc.BroadcastedTransaction) (*rpc.AddTxResponse, *jsonrpc.Error) {
	return s.addTransaction(tx, false)
}

func (s *StarknetRPC) LegacyAddTransaction(tx rpc.BroadcastedTransaction) (*rpc.AddTxResponse, *jsonrpc.Error) {
	return s.addTransaction(tx, true)
}

func (s *StarknetRPC) addTransaction(tx rpc.BroadcastedTransaction, legacyTraceJson bool) (*rpc.AddTxResponse, *jsonrpc.Error) {
	txn := &tx
	if txn.ContractAddress == nil && txn.Type == rpc.TxnDeployAccount {
		txn.ContractAddress = core.ContractAddress(&felt.Zero, tx.ClassHash, tx.ContractAddressSalt, *tx.ConstructorCallData)
	}
	if txn.ClassHash == nil && txn.Type == rpc.TxnDeclare {
		var class sdk.ContractClass
		err := json.Unmarshal(tx.ContractClass, &class)
		if err != nil {
			return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err.Error())
		}
		txn.ClassHash, err = hash.ClassHash(class)
		if err != nil {
			return nil, jsonrpc.Err(jsonrpc.InvalidParams, err.Error())
		}
	}

	txReq := cairo.TxRequest{Tx: txn, LegacyTraceJson: legacyTraceJson}
	byt, err := json.Marshal(txReq)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())
	}
	signedWrCall := &yucore.SignedWrCall{
		Call: &common.WrCall{
			TripodName: CairoTripod,
			FuncName:   "ExecuteTxn",
			Params:     string(byt),
		},
	}
	err = s.chain.HandleTxn(signedWrCall)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err.Error())
	}

	bcTx, _, _, err := cairo.AdaptBroadcastedTransaction(txn, s.network)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err.Error())
	}

	return &rpc.AddTxResponse{
		TransactionHash: bcTx.Hash(),
		ContractAddress: txn.ContractAddress,
		ClassHash:       txn.ClassHash,
	}, nil
}

func (s *StarknetRPC) Call(call rpc.FunctionCall, id rpc.BlockID) ([]*felt.Felt, *jsonrpc.Error) {
	callReq := &cairo.CallRequest{
		ContractAddr: &call.ContractAddress,
		Selector:     &call.EntryPointSelector,
		Calldata:     call.Calldata,
		BlockID:      cairo.NewFromJunoBlockID(id),
	}
	resp, jsonErr := s.adaptChainRead(callReq, "Call")
	if jsonErr != nil {
		return nil, jsonErr
	}
	cr := resp.DataInterface.(*cairo.CallResponse)
	return cr.ReturnData, cr.Err
}

func (s *StarknetRPC) GetTransactionByHash(hash felt.Felt) (*rpc.Transaction, *jsonrpc.Error) {
	txReq := &cairo.TransactionRequest{Hash: hash}
	resp, jsonErr := s.adaptChainRead(txReq, "GetTransaction")
	if jsonErr != nil {
		return nil, jsonErr
	}
	tr := resp.DataInterface.(*cairo.TransactionResponse)
	return tr.Tx, tr.Err
}

func (s *StarknetRPC) GetTransactionStatus(ctx context.Context, hash felt.Felt) (*rpc.TransactionStatus, *jsonrpc.Error) {
	tsReq := &cairo.TransactionStatusRequest{Hash: hash}
	resp, jsonErr := s.adaptChainRead(tsReq, "GetTransactionStatus")
	if jsonErr != nil {
		return nil, jsonErr
	}
	tsr := resp.DataInterface.(*cairo.TransactionStatusResponse)
	return tsr.Status, tsr.Err
}

func (s *StarknetRPC) GetReceiptByHash(hash felt.Felt) (*rpc.TransactionReceipt, *jsonrpc.Error) {
	rcptReq := &cairo.ReceiptRequest{Hash: hash}
	resp, jsonErr := s.adaptChainRead(rcptReq, "GetReceipt")
	if jsonErr != nil {
		return nil, jsonErr
	}
	rr := resp.DataInterface.(*cairo.ReceiptResponse)
	return rr.Receipt, rr.Err
}

func (s *StarknetRPC) LegacyGetReceiptByHash(hash felt.Felt) (*rpc.TransactionReceipt, *jsonrpc.Error) {
	receipt, rpcErr := s.GetReceiptByHash(hash)
	if rpcErr != nil {
		return nil, rpcErr
	}
	receipt.ActualFee.IsLegacy = true
	receipt.ExecutionResources.IsLegacy = true
	return receipt, nil
}

func (s *StarknetRPC) GetNonce(id rpc.BlockID, address felt.Felt) (*felt.Felt, *jsonrpc.Error) {
	nonceReq := &cairo.NonceRequest{BlockID: cairo.NewFromJunoBlockID(id), Addr: &address}
	resp, jsonErr := s.adaptChainRead(nonceReq, "GetNonce")
	if jsonErr != nil {
		return nil, jsonErr
	}
	nr := resp.DataInterface.(*cairo.NonceResponse)
	return nr.Nonce, nr.Err
}

func (s *StarknetRPC) GetStorage(address, key felt.Felt, id rpc.BlockID) (*felt.Felt, *jsonrpc.Error) {
	storageReq := &cairo.StorageRequest{
		BlockID: cairo.NewFromJunoBlockID(id),
		Addr:    &address,
		Key:     &key,
	}
	resp, jsonErr := s.adaptChainRead(storageReq, "GetStorage")
	if jsonErr != nil {
		return nil, jsonErr
	}
	sr := resp.DataInterface.(*cairo.StorageResponse)
	return sr.Value, sr.Err
}

func (s *StarknetRPC) GetClass(id rpc.BlockID, classHash felt.Felt) (*rpc.Class, *jsonrpc.Error) {
	classReq := &cairo.ClassRequest{BlockID: cairo.NewFromJunoBlockID(id), ClassHash: &classHash}
	resp, jsonErr := s.adaptChainRead(classReq, "GetClass")
	if jsonErr != nil {
		return nil, jsonErr
	}
	cr := resp.DataInterface.(*cairo.ClassResponse)
	return cr.Class, cr.Err
}

func (s *StarknetRPC) GetClassAt(id rpc.BlockID, address felt.Felt) (*rpc.Class, *jsonrpc.Error) {
	classAtReq := &cairo.ClassAtRequest{BlockID: cairo.NewFromJunoBlockID(id), Addr: &address}
	resp, jsonErr := s.adaptChainRead(classAtReq, "GetClassAt")
	if jsonErr != nil {
		return nil, jsonErr
	}
	cr := resp.DataInterface.(*cairo.ClassResponse)
	return cr.Class, cr.Err
}

func (s *StarknetRPC) GetClassHashAt(id rpc.BlockID, address felt.Felt) (*felt.Felt, *jsonrpc.Error) {
	classHashAtReq := &cairo.ClassHashAtRequest{BlockID: cairo.NewFromJunoBlockID(id), Addr: &address}
	resp, jsonErr := s.adaptChainRead(classHashAtReq, "GetClassHashAt")
	if jsonErr != nil {
		return nil, jsonErr
	}
	cr := resp.DataInterface.(*cairo.ClassHashAtResponse)
	return cr.ClassHash, cr.Err
}

func (s *StarknetRPC) adaptChainRead(req any, funcName string) (*yucontext.ResponseData, *jsonrpc.Error) {
	byt, err := json.Marshal(req)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidJSON, err)
	}
	rdCall := &common.RdCall{
		TripodName: CairoTripod,
		FuncName:   funcName,
		Params:     string(byt),
	}
	resp, err := s.chain.HandleRead(rdCall)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err)
	}
	return resp, nil
}
