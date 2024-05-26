package starknetrpc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"

	// "github.com/NethermindEth/juno/db"

	// "github.com/NethermindEth/juno/blockchain"
	"itachi/cairo"

	"github.com/NethermindEth/starknet.go/hash"
	sdk "github.com/NethermindEth/starknet.go/rpc"
	"github.com/yu-org/yu/common"
	yucore "github.com/yu-org/yu/core"
	yucontext "github.com/yu-org/yu/core/context"
)

func (s *StarknetRPC) GetChainID() (*felt.Felt, *jsonrpc.Error) {
	return s.network.ChainID(), nil
}

func (s *StarknetRPC) BlockNumber() (uint64, *jsonrpc.Error) {
	// num, _ := s.bcReader.Height()
	// return num, nil
	resp, err := s.adaptChainRead(nil, "LatestBlock")
	if err != nil {
		return 0, rpc.ErrNoBlock
	}
	res := resp.DataInterface.(*cairo.LatestBlockResponse)
	return uint64(res.Block.Height), res.Err
}

func (s *StarknetRPC) BlockHashAndNumber() (*rpc.BlockHashAndNumber, *jsonrpc.Error) {
	// block, err := s.bcReader.Head()
	// if err != nil {
	// 	return nil, rpc.ErrBlockNotFound
	// }
	// return &rpc.BlockHashAndNumber{Number: block.Number, Hash: block.Hash}, nil
	resp, err := s.adaptChainRead(nil, "LatestBlock")
	if err != nil {
		return nil, rpc.ErrNoBlock
	}
	res := resp.DataInterface.(*cairo.LatestBlockResponse)
	return &rpc.BlockHashAndNumber{Number: uint64(res.Block.Height), Hash: new(felt.Felt).SetBytes(res.Block.Hash.Bytes())}, nil
}

// func (s *StarknetRPC) GetStateUpdate(id rpc.BlockID) (*rpc.StateUpdate, *jsonrpc.Error) {
// 	var update *core.StateUpdate
// 	var err error
// 	if id.Latest {
// 		if height, heightErr := s.bcReader.Height(); heightErr != nil {
// 			err = heightErr
// 		} else {
// 			update, err = s.bcReader.StateUpdateByNumber(height)
// 		}
// 	} else if id.Pending {
// 		var pending blockchain.Pending
// 		pending, err = s.bcReader.Pending()
// 		if err == nil {
// 			update = pending.StateUpdate
// 		}
// 	} else if id.Hash != nil {
// 		update, err = s.bcReader.StateUpdateByHash(id.Hash)
// 	} else {
// 		update, err = s.bcReader.StateUpdateByNumber(id.Number)
// 	}
// 	if err != nil {
// 		if errors.Is(err, db.ErrKeyNotFound) {
// 			return nil, rpc.ErrBlockNotFound
// 		}
// 		return nil, rpc.ErrInternal.CloneWithData(err)
// 	}

// 	nonces := make([]rpc.Nonce, 0, len(update.StateDiff.Nonces))
// 	for addr, nonce := range update.StateDiff.Nonces {
// 		nonces = append(nonces, rpc.Nonce{ContractAddress: addr, Nonce: *nonce})
// 	}

// 	storageDiffs := make([]rpc.StorageDiff, 0, len(update.StateDiff.StorageDiffs))
// 	for addr, diffs := range update.StateDiff.StorageDiffs {
// 		entries := make([]rpc.Entry, 0, len(diffs))
// 		for key, value := range diffs {
// 			entries = append(entries, rpc.Entry{
// 				Key:   key,
// 				Value: *value,
// 			})
// 		}

// 		storageDiffs = append(storageDiffs, rpc.StorageDiff{
// 			Address:        addr,
// 			StorageEntries: entries,
// 		})
// 	}

// 	deployedContracts := make([]rpc.DeployedContract, 0, len(update.StateDiff.DeployedContracts))
// 	for addr, classHash := range update.StateDiff.DeployedContracts {
// 		deployedContracts = append(deployedContracts, rpc.DeployedContract{
// 			Address:   addr,
// 			ClassHash: *classHash,
// 		})
// 	}

// 	declaredClasses := make([]rpc.DeclaredClass, 0, len(update.StateDiff.DeclaredV1Classes))
// 	for classHash, compiledClassHash := range update.StateDiff.DeclaredV1Classes {
// 		declaredClasses = append(declaredClasses, rpc.DeclaredClass{
// 			ClassHash:         classHash,
// 			CompiledClassHash: *compiledClassHash,
// 		})
// 	}

// 	replacedClasses := make([]rpc.ReplacedClass, 0, len(update.StateDiff.ReplacedClasses))
// 	for addr, classHash := range update.StateDiff.ReplacedClasses {
// 		replacedClasses = append(replacedClasses, rpc.ReplacedClass{
// 			ClassHash:       *classHash,
// 			ContractAddress: addr,
// 		})
// 	}

// 	return &rpc.StateUpdate{
// 		BlockHash: update.BlockHash,
// 		OldRoot:   update.OldRoot,
// 		NewRoot:   update.NewRoot,
// 		StateDiff: &rpc.StateDiff{
// 			DeprecatedDeclaredClasses: update.StateDiff.DeclaredV0Classes,
// 			DeclaredClasses:           declaredClasses,
// 			ReplacedClasses:           replacedClasses,
// 			Nonces:                    nonces,
// 			StorageDiffs:              storageDiffs,
// 			DeployedContracts:         deployedContracts,
// 		},
// 	}, nil
// }

// func (s *StarknetRPC) Syncing() (*rpc.Sync, *jsonrpc.Error) {
// 	return nil, nil
// }

// func (s *StarknetRPC) TraceTransaction(ctx context.Context, hash felt.Felt) (*vm.TransactionTrace, *jsonrpc.Error) {
// 	receipt, err := s.GetReceiptByHash(hash)
// 	if err != nil {
// 		return nil, err
// 	}
// 	blockhash := receipt.BlockHash
// 	if blockhash == nil {

// 	}
// }

// func (s *StarknetRPC) TraceBlockTransactions() () {}

// func (s *StarknetRPC) traceTransaction(ctx context.Context, blockHash *felt.Felt) ([]rpc.TracedBlockTransaction, *jsonrpc.Error) {
// 	// not pending
// 	if blockHash != nil {

// 	}
// }

func (s *StarknetRPC) GetBlockWithTxHashes(id rpc.BlockID) (*rpc.BlockWithTxHashes, *jsonrpc.Error) {
	req := &cairo.BlockWithTxHashesRequest{BlockID: cairo.NewFromJunoBlockID(id)}
	resp, jsonErr := s.adaptChainRead(req, "GetBlockWithTxHashes")
	if jsonErr != nil {
		return nil, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockWithTxHashesResponse)
	return res.BlockWithTxHashes, res.Err
}

func (s *StarknetRPC) GetBlockWithTxs(id rpc.BlockID) (*rpc.BlockWithTxs, *jsonrpc.Error) {
	req := &cairo.BlockWithTxsRequest{BlockID: cairo.NewFromJunoBlockID(id)}
	resp, jsonErr := s.adaptChainRead(req, "GetBlockWithTxs")
	if jsonErr != nil {
		return nil, jsonErr
	}
	res := resp.DataInterface.(*cairo.BlockWithTxsResponse)
	return res.BlockWithTxs, res.Err
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

func (s *StarknetRPC) EstimateMessageFee(msg rpc.MsgFromL1, id rpc.BlockID) (*rpc.FeeEstimate, *jsonrpc.Error) {
	feeEstimate, err := s.estimateMessageFee(msg, id, s.EstimateFee)
	if err != nil {
		return nil, err
	}
	return feeEstimate, nil
}

type estimateFeeHandler func(broadcastedTxns []rpc.BroadcastedTransaction,
	simulationFlags []rpc.SimulationFlag, id rpc.BlockID,
) ([]rpc.FeeEstimate, *jsonrpc.Error)

func (s *StarknetRPC) estimateMessageFee(
	msg rpc.MsgFromL1,
	id rpc.BlockID,
	f estimateFeeHandler,
) (*rpc.FeeEstimate, *jsonrpc.Error) {
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
	estimates, _ := f([]rpc.BroadcastedTransaction{tx}, nil, id)
	return &estimates[0], nil
}

func (s *StarknetRPC) EstimateFee(broadcastedTxns []rpc.BroadcastedTransaction,
	simulationFlags []rpc.SimulationFlag, id rpc.BlockID,
) ([]rpc.FeeEstimate, *jsonrpc.Error) {
	result, err := s.simulateTransactions(id, broadcastedTxns, append(simulationFlags, rpc.SkipFeeChargeFlag), false, true)
	if err != nil {
		return nil, err
	}
	return utils.Map(result, func(tx rpc.SimulatedTransaction) rpc.FeeEstimate {
		return tx.FeeEstimation
	}), nil
}

func (s *StarknetRPC) LegacyEstimateFee(broadcastedTxns []rpc.BroadcastedTransaction, id rpc.BlockID) ([]rpc.FeeEstimate, *jsonrpc.Error) {
	result, err := s.simulateTransactions(id, broadcastedTxns, []rpc.SimulationFlag{rpc.SkipFeeChargeFlag}, true, true)
	if err != nil && err.Code == rpc.ErrTransactionExecutionError.Code {
		return nil, makeContractError(errors.New(err.Data.(rpc.TransactionExecutionErrorData).ExecutionError))
	}

	return utils.Map(result, func(tx rpc.SimulatedTransaction) rpc.FeeEstimate {
		return tx.FeeEstimation
	}), nil
}

func (s *StarknetRPC) SimulateTransactions(
	id rpc.BlockID,
	transactions []rpc.BroadcastedTransaction,
	simulationFlags []rpc.SimulationFlag,
) ([]rpc.SimulatedTransaction, *jsonrpc.Error) {
	return s.simulateTransactions(id, transactions, simulationFlags, false, false)
}

func (s *StarknetRPC) LegacySimulateTransactions(
	id rpc.BlockID,
	transactions []rpc.BroadcastedTransaction,
	simulationFlags []rpc.SimulationFlag,
) ([]rpc.SimulatedTransaction, *jsonrpc.Error) {
	simu, err := s.simulateTransactions(id, transactions, simulationFlags, true, true)
	if err.Code == rpc.ErrTransactionExecutionError.Code {
		return nil, makeContractError(errors.New(err.Data.(rpc.TransactionExecutionErrorData).ExecutionError))
	}
	return simu, err
}

func (s *StarknetRPC) simulateTransactions(
	id rpc.BlockID,
	transactions []rpc.BroadcastedTransaction,
	simulationFlags []rpc.SimulationFlag,
	legacyJson, errOnRevert bool,
) ([]rpc.SimulatedTransaction, *jsonrpc.Error) {
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

func (s *StarknetRPC) GetTransactionByHash(hash felt.Felt) (*rpc.Transaction, *jsonrpc.Error) {
	txReq := &cairo.TransactionRequest{Hash: hash}
	resp, jsonErr := s.adaptChainRead(txReq, "GetTransaction")
	if jsonErr != nil {
		return nil, jsonErr
	}
	tr := resp.DataInterface.(*cairo.TransactionResponse)
	return tr.Tx, tr.Err
}

func (s *StarknetRPC) GetBlockTransactionCount(id rpc.BlockID) (uint64, *jsonrpc.Error) {
	blockWithTxs, err := s.GetBlockWithTxs(id)
	if err != nil {
		return 0, err
	}
	return uint64(len(blockWithTxs.Transactions)), err
}

func (s *StarknetRPC) GetTransactionByBlockIdAndIndex(id rpc.BlockID, txIndex int) (*rpc.Transaction, *jsonrpc.Error) {
	if txIndex < 0 {
		return nil, rpc.ErrInvalidTxIndex
	}

	blockWithTxHashes, err := s.GetBlockWithTxHashes(id)
	if err != nil {
		return nil, err
	}
	hash := *blockWithTxHashes.TxnHashes[txIndex] //不是很确定这里是不是直接可以用index

	txReq := &cairo.TransactionRequest{Hash: hash}
	resp, jsonErr := s.adaptChainRead(txReq, "GetTransaction")
	if jsonErr != nil {
		return nil, jsonErr
	}
	tr := resp.DataInterface.(*cairo.TransactionResponse)
	return tr.Tx, tr.Err

	// if id.Pending {
	// 	pending, err := s.bcReader.Pending()
	// 	if err != nil {
	// 		return nil, rpc.ErrBlockNotFound
	// 	}

	// 	if uint64(txIndex) > pending.Block.TransactionCount {
	// 		return nil, rpc.ErrInvalidTxIndex
	// 	}

	// 	hash := *pending.Block.Transactions[txIndex].Hash()
	// 	txReq := &cairo.TransactionRequest{Hash: hash}
	// 	resp, jsonErr := s.adaptChainRead(txReq, "GetTransaction")
	// 	if jsonErr != nil {
	// 		return nil, jsonErr
	// 	}
	// 	tr := resp.DataInterface.(*cairo.TransactionResponse)
	// 	return tr.Tx, tr.Err
	// }
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

func (s *StarknetRPC) SpecVersion() (string, *jsonrpc.Error) {
	return "0.6.0", nil
}

func (s *StarknetRPC) LegacySpecVersion() (string, *jsonrpc.Error) {
	return "0.5.1", nil
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
