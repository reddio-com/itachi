package starknetrpc

/**
* Chain Handlers
* Block Handlers
* Transaction Handlers
* State Handlers
 */
import (
	"context"
	"encoding/json"
	"errors"
	"itachi/cairo"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/starknet.go/hash"
	sdk "github.com/NethermindEth/starknet.go/rpc"
	"github.com/yu-org/yu/common"
	yucore "github.com/yu-org/yu/core"
	yucontext "github.com/yu-org/yu/core/context"
)

func (s *StarknetRPC) SpecVersionV0_7() (string, *jsonrpc.Error) {
	return "0.7.0", nil
}
func (s *StarknetRPC) SpecVersionV0_6() (string, *jsonrpc.Error) {
	return "0.6.0", nil
}
func (s *StarknetRPC) SpecVersionV0_5() (string, *jsonrpc.Error) {
	return "0.5.0", nil
}

////////////////////////////////////////
///		Chain Handlers			////////
////////////////////////////////////////

// GetNetwork returns the network of the chain.
// FIXME: Implement full logic for GetChainID, including error handling and dynamic chain ID retrieval.
func (s *StarknetRPC) GetChainID() (*felt.Felt, *jsonrpc.Error) {
	return new(felt.Felt).SetBytes([]byte("SN_ITACHI")), nil
}

////////////////////////////////////////
////		Block Handlers		////////
////////////////////////////////////////

func (s *StarknetRPC) GetBlockWithTxHashes(id rpc.BlockID) (*cairo.BlockWithTxHashes, *jsonrpc.Error) {
	req := &cairo.BlockWithTxHashesRequest{BlockID: cairo.NewFromJunoBlockID(id)}
	resp, jsonErr := s.adaptChainRead(req, "GetBlockWithTxHashes")
	if jsonErr != nil {
		return nil, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockWithTxHashesResponse)
	return res.BlockWithTxHashes, res.Err
}

func (s *StarknetRPC) GetBlockWithTxs(id rpc.BlockID) (*cairo.BlockWithTxs, *jsonrpc.Error) {
	req := &cairo.BlockWithTxsRequest{BlockID: cairo.NewFromJunoBlockID(id)}
	resp, jsonErr := s.adaptChainRead(req, "GetBlockWithTxs")
	if jsonErr != nil {
		return nil, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockWithTxsResponse)
	return res.BlockWithTxs, res.Err
}

// BlockHashAndNumber returns the block number of the latest Finalized block.
func (s *StarknetRPC) GetBlockNumber() (uint64, *jsonrpc.Error) {
	req := ""
	resp, jsonErr := s.adaptChainRead(req, "GetBlockNumber")
	if jsonErr != nil {
		return 0, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockNumberResponse)

	return res.BlockNumber, nil
}

// BlockHashAndNumber returns the block number and hash of the latest Finalized block.
func (s *StarknetRPC) GetBlockHashAndNumber() (*rpc.BlockHashAndNumber, *jsonrpc.Error) {
	req := ""
	resp, jsonErr := s.adaptChainRead(req, "GetBlockHashAndNumber")
	if jsonErr != nil {
		return nil, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockHashAndNumberResponse)
	return res.BlockHashAndNumber, nil
}

// GetBlockTransactionCount returns the number of transactions in the block with the given ID.
func (s *StarknetRPC) GetBlockTransactionCount(id rpc.BlockID) (uint64, *jsonrpc.Error) {
	req := &cairo.BlockTransactionCountRequest{BlockID: cairo.NewFromJunoBlockID(id)}
	resp, jsonErr := s.adaptChainRead(req, "GetBlockTransactionCount")
	if jsonErr != nil {
		return 0, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockTransactionCountResponse)
	return res.TxsNumber, nil
}

////////////////////////////////////////
////		Transaction Handlers	////
////////////////////////////////////////

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

// TransactionByBlockIDAndIndex returns the details of a transaction identified by the given
// BlockID and index.
//
// It follows the specification defined here:
// https://github.com/starkware-libs/starknet-specs/blob/master/api/starknet_api_openrpc.json#L184
func (s *StarknetRPC) GetTransactionByBlockIDAndIndex(id rpc.BlockID, txIndex int) (*rpc.Transaction, *jsonrpc.Error) {
	req := &cairo.TransactionByBlockIDAndIndexRequest{BlockID: cairo.NewFromJunoBlockID(id), TxIndex: txIndex}
	resp, jsonErr := s.adaptChainRead(req, "GetTransactionByBlockIDAndIndex")
	if jsonErr != nil {
		return nil, jsonErr
	}
	tr := resp.DataInterface.(*cairo.TransactionByBlockIDAndIndexResponse)
	return tr.Tx, tr.Err
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

func (s *StarknetRPC) EstimateFee(broadcastedTxns []rpc.BroadcastedTransaction,
	simulationFlags []cairo.SimulationFlag, id rpc.BlockID,
) ([]cairo.FeeEstimate, *jsonrpc.Error) {
	result, err := s.simulateTransactions(id, broadcastedTxns, append(simulationFlags, cairo.SkipFeeChargeFlag), false, true)
	if err != nil {
		return nil, err
	}

	return utils.Map(result, func(tx cairo.SimulatedTransaction) cairo.FeeEstimate {
		return tx.FeeEstimation
	}), nil
}

func (s *StarknetRPC) LegacyEstimateFee(broadcastedTxns []rpc.BroadcastedTransaction, id rpc.BlockID) ([]cairo.FeeEstimate, *jsonrpc.Error) {
	result, err := s.simulateTransactions(id, broadcastedTxns, []cairo.SimulationFlag{cairo.SkipFeeChargeFlag}, true, true)
	if err != nil && err.Code == rpc.ErrTransactionExecutionError.Code {
		return nil, makeContractError(errors.New(err.Data.(rpc.TransactionExecutionErrorData).ExecutionError))
	}

	return utils.Map(result, func(tx cairo.SimulatedTransaction) cairo.FeeEstimate {
		return tx.FeeEstimation
	}), nil
}
func (s *StarknetRPC) SimulateTransactions(
	id rpc.BlockID,
	transactions []rpc.BroadcastedTransaction,
	simulationFlags []cairo.SimulationFlag,
) ([]cairo.SimulatedTransaction, *jsonrpc.Error) {
	return s.simulateTransactions(id, transactions, simulationFlags, false, false)
}

func (s *StarknetRPC) LegacySimulateTransactions(
	id rpc.BlockID,
	transactions []rpc.BroadcastedTransaction,
	simulationFlags []cairo.SimulationFlag,
) ([]cairo.SimulatedTransaction, *jsonrpc.Error) {
	simu, err := s.simulateTransactions(id, transactions, simulationFlags, true, true)
	if err.Code == rpc.ErrTransactionExecutionError.Code {
		return nil, makeContractError(errors.New(err.Data.(rpc.TransactionExecutionErrorData).ExecutionError))
	}
	return simu, err
}
func (s *StarknetRPC) EstimateMessageFee(msg rpc.MsgFromL1, id rpc.BlockID) (*cairo.FeeEstimate, *jsonrpc.Error) { //nolint:gocritic
	calldata := make([]*felt.Felt, 0, len(msg.Payload)+1)
	// The order of the calldata parameters matters. msg.From must be prepended.
	calldata = append(calldata, new(felt.Felt).SetBytes(msg.From.Bytes()))
	for payloadIdx := range msg.Payload {
		calldata = append(calldata, &msg.Payload[payloadIdx])
	}
	tx := rpc.BroadcastedTransaction{
		Transaction: rpc.Transaction{
			Type:               rpc.TxnL1Handler,
			ContractAddress:    &msg.To,
			EntryPointSelector: &msg.Selector,
			CallData:           &calldata,
			Version:            &felt.Zero, // Needed for transaction hash calculation.
			Nonce:              &felt.Zero, // Needed for transaction hash calculation.
		},
		// Needed to marshal to blockifier type.
		// Must be greater than zero to successfully execute transaction.
		PaidFeeOnL1: new(felt.Felt).SetUint64(1),
	}
	estimates, rpcErr := s.EstimateFee([]rpc.BroadcastedTransaction{tx}, nil, id)
	if rpcErr != nil {
		if rpcErr.Code == rpc.ErrTransactionExecutionError.Code {
			data := rpcErr.Data.(rpc.TransactionExecutionErrorData)
			return nil, makeContractError(errors.New(data.ExecutionError))
		}
		return nil, rpcErr
	}
	return &estimates[0], nil
}

func (s *StarknetRPC) simulateTransactions(
	id rpc.BlockID,
	transactions []rpc.BroadcastedTransaction,
	simulationFlags []cairo.SimulationFlag,
	legacyJson, errOnRevert bool,
) ([]cairo.SimulatedTransaction, *jsonrpc.Error) {
	simuReq := &cairo.SimulateRequest{
		BlockID:         cairo.NewFromJunoBlockID(id),
		Txs:             transactions,
		SimulationFlags: simulationFlags,
		LegacyJson:      legacyJson,
		ErrOnRevert:     errOnRevert,
	}
	resp, jsonErr := s.adaptChainRead(simuReq, "SimulateTransactions")
	if jsonErr != nil {
		return nil, jsonErr
	}
	sr := resp.DataInterface.(*cairo.SimulateResponse)
	return sr.Txs, sr.Err
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

////////////////////////////////////////
////		State Handlers		    ////
////////////////////////////////////////

// StateUpdate returns the state update identified by the given BlockID.
//
// It follows the specification defined here:
// https://github.com/starkware-libs/starknet-specs/blob/master/api/starknet_api_openrpc.json#L77
func (s *StarknetRPC) GetStateUpdate(id rpc.BlockID) (*cairo.StateUpdate, *jsonrpc.Error) {
	req := &cairo.StateUpdateRequest{BlockID: cairo.NewFromJunoBlockID(id)}
	resp, jsonErr := s.adaptChainRead(req, "GetStateUpdate")
	if jsonErr != nil {
		return nil, jsonErr
	}
	sr := resp.DataInterface.(*cairo.StateUpdateResponse)

	return sr.StateUpdate, sr.Err

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

////////////////////////////////////////
////		Contract Handlers		////
////////////////////////////////////////

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

func makeContractError(err error) *jsonrpc.Error {
	return rpc.ErrContractError.CloneWithData(rpc.ContractErrorData{
		RevertError: err.Error(),
	})
}
