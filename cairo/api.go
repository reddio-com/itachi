package cairo

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/encoder"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/types"
)

func NewFromJunoBlockID(id rpc.BlockID) BlockID {
	return BlockID{
		Pending: id.Pending,
		Latest:  id.Latest,
		Hash:    id.Hash,
		Number:  id.Number,
	}
}

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
		ctx.Json(http.StatusBadRequest, &TransactionResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	signedTx, err := c.TxDB.GetTxn(tq.Hash.Bytes())
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &TransactionResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	txReq := new(TxRequest)
	err = signedTx.BindJson(txReq)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &TransactionResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	ctx.JsonOk(&TransactionResponse{Tx: &txReq.Tx.Transaction})
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
		ctx.Json(http.StatusBadRequest, &ReceiptResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}

	starkReceipt, err := c.getReceipt(rq.Hash)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &ReceiptResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	ctx.JsonOk(&ReceiptResponse{Receipt: starkReceipt})
}

func (c *Cairo) getReceipt(hash felt.Felt) (*rpc.TransactionReceipt, error) {
	receipt, err := c.TxDB.GetReceipt(hash.Bytes())
	if err != nil {
		return nil, err
	}
	if receipt == nil {
		return nil, errors.New("no receipt found")
	}
	starkReceipt := new(rpc.TransactionReceipt)
	err = encoder.Unmarshal(receipt.Extra, starkReceipt)
	return starkReceipt, err
}

type BlockWithTxHashesRequest struct {
	BlockID BlockID `json:"block_id"`
}

type BlockWithTxHashesResponse struct {
	BlockWithTxHashes *BlockWithTxHashes `json:"block_with_tx_hashes"`
	Err               *jsonrpc.Error     `json:"err"`
}

func (c *Cairo) GetBlockWithTxHashes(ctx *context.ReadContext) {
	var br BlockWithTxHashesRequest
	err := ctx.BindJson(&br)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &BlockWithTxHashesResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	var compactBlock *types.CompactBlock
	compactBlock, err = c.getYuCompactBlock(br.BlockID)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockWithTxHashesResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	status, err := c.blockStatus(br.BlockID)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockWithTxsResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	txHashes := make([]*felt.Felt, 0)
	for _, txHash := range compactBlock.TxnsHashes {
		txHashes = append(txHashes, new(felt.Felt).SetBytes(txHash.Bytes()))
	}
	ctx.JsonOk(&BlockWithTxHashesResponse{BlockWithTxHashes: &BlockWithTxHashes{
		Status:      status,
		BlockHeader: c.adaptStarkBlockHeader(compactBlock),
		TxnHashes:   txHashes,
	}})
}

type BlockWithTxsRequest struct {
	BlockID BlockID `json:"block_id"`
}

type BlockWithTxsResponse struct {
	BlockWithTxs *BlockWithTxs  `json:"block_with_txs"`
	Err          *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetBlockWithTxs(ctx *context.ReadContext) {
	var br BlockWithTxsRequest
	err := ctx.BindJson(&br)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &BlockWithTxsResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	var compactBlock *types.CompactBlock
	compactBlock, err = c.getYuCompactBlock(br.BlockID)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockWithTxsResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	starkTxs, err := c.getTransactionsByCompactBlock(compactBlock)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockWithTxsResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	status, err := c.blockStatus(br.BlockID)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockWithTxsResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	blockWithTxs := &BlockWithTxs{
		Status:       status,
		BlockHeader:  c.adaptStarkBlockHeader(compactBlock),
		Transactions: starkTxs,
	}

	ctx.JsonOk(&BlockWithTxsResponse{BlockWithTxs: blockWithTxs})
}

func (c *Cairo) getTransactionsByCompactBlock(compactBlock *types.CompactBlock) ([]*rpc.Transaction, error) {
	starkTxs := make([]*rpc.Transaction, 0)
	for _, txHash := range compactBlock.TxnsHashes {
		var yuTxn *types.SignedTxn
		yuTxn, err := c.TxDB.GetTxn(txHash)
		if err != nil {
			return nil, err
		}
		txReq := new(TxRequest)
		err = yuTxn.BindJson(txReq)
		if err != nil {
			return nil, err
		}
		starkCoreTx, _, _, err := c.adaptBroadcastedTransaction(txReq.Tx)
		if err != nil {
			return nil, err
		}
		starkTx := rpc.AdaptTransaction(starkCoreTx)
		starkTxs = append(starkTxs, starkTx)
	}
	return starkTxs, nil
}

// FIXME: the following functions are not implemented yet
func (c *Cairo) blockStatus(id BlockID) (BlockStatus, error) {
	status := BlockAcceptedL2
	if id.Pending {
		status = BlockPending
	}
	return status, nil
}

//	func (c *Cairo) getYuBlock(id BlockID) (*types.Block, error) {
//		switch {
//		case id.Latest || id.Pending:
//			return c.Chain.GetEndBlock()
//		case id.Hash != nil:
//			return c.Chain.GetBlock(id.Hash.Bytes())
//		case id.Number != 0:
//			return c.Chain.GetBlockByHeight(common.BlockNum(id.Number))
//		default:
//			return c.Chain.GetBlockByHeight(common.BlockNum(id.Number))
//		}
//	}
func (c *Cairo) getYuCompactBlock(id BlockID) (*types.CompactBlock, error) {
	switch {
	case id.Latest || id.Pending:
		return c.Chain.GetEndCompactBlock()
	case id.Hash != nil:
		return c.Chain.GetCompactBlock(id.Hash.Bytes())
	case id.Number != 0:
		return c.Chain.GetCompactBlockByHeight(common.BlockNum(id.Number))
	default:
		return c.Chain.GetCompactBlockByHeight(common.BlockNum(id.Number))
	}
}
func (c *Cairo) adaptStarkBlockHeader(yuBlock *types.CompactBlock) BlockHeader {
	var blockNumber *uint64
	// if header.Hash == nil it's a pending block
	if yuBlock.Header.Hash.Bytes() != nil {
		height := uint64(yuBlock.Header.Height)
		blockNumber = &height
	}
	sequencerAddress := c.sequencerAddr
	if sequencerAddress == nil {
		sequencerAddress = &felt.Zero
	}

	// FIXME: L1DAMode just for test, after the feeder function is implemented, it should be replaced
	L1DAMode := Calldata
	return BlockHeader{
		Hash:       new(felt.Felt).SetBytes(yuBlock.Header.Hash.Bytes()),
		ParentHash: new(felt.Felt).SetBytes(yuBlock.PrevHash.Bytes()),
		Number:     blockNumber,
		// FIXME
		NewRoot:          new(felt.Felt).SetBytes(yuBlock.StateRoot.Bytes()),
		Timestamp:        yuBlock.Timestamp,
		SequencerAddress: sequencerAddress,
		// FIXME: L1GasPrice, StarknetVersion，L1DataGasPrice，L1DAMode just for test, after the feeder function is implemented, it should be replaced
		L1GasPrice: &ResourcePrice{
			InFri: new(felt.Felt).SetUint64(52953095),
			InWei: new(felt.Felt).SetUint64(52953096),
		},
		L1DataGasPrice: &ResourcePrice{
			InFri: new(felt.Felt).SetUint64(52953097),
			InWei: new(felt.Felt).SetUint64(52953098),
		},
		L1DAMode:        &L1DAMode,
		StarknetVersion: "0.13.0",
	}
}

type BlockNumberResponse struct {
	BlockNumber uint64         `json:"block_number"`
	Err         *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetBlockNumber(ctx *context.ReadContext) {
	var err error
	var compactBlock *types.CompactBlock
	compactBlock, err = c.Chain.LastFinalizedCompact()
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockNumberResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	ctx.JsonOk(&BlockNumberResponse{BlockNumber: uint64(compactBlock.Height)})
}

type BlockHashAndNumberResponse struct {
	BlockHashAndNumber *rpc.BlockHashAndNumber `json:"block_hash_and_number"`
	Err                *jsonrpc.Error          `json:"err"`
}

func (c *Cairo) GetBlockHashAndNumber(ctx *context.ReadContext) {
	var err error
	var compactBlock *types.CompactBlock
	compactBlock, err = c.Chain.LastFinalizedCompact()
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockNumberResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	feltHash := new(felt.Felt).SetBytes(compactBlock.Hash.Bytes())
	blockHashAndNumber := &rpc.BlockHashAndNumber{Hash: feltHash, Number: uint64(compactBlock.Height)}
	ctx.JsonOk(&BlockHashAndNumberResponse{BlockHashAndNumber: blockHashAndNumber})
}

type BlockTransactionCountRequest struct {
	BlockID BlockID `json:"block_id"`
}

type BlockTransactionCountResponse struct {
	TxsNumber uint64         `json:"TxsNumber"`
	Err       *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetBlockTransactionCount(ctx *context.ReadContext) {
	//get Block by blockID
	var br BlockTransactionCountRequest
	err := ctx.BindJson(&br)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &BlockTransactionCountResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}

	var compactBlock *types.CompactBlock
	compactBlock, err = c.getYuCompactBlock(br.BlockID)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &BlockTransactionCountResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	//get the number of transactions in the block
	var txnsNumber uint64 = uint64(len(compactBlock.TxnsHashes))
	ctx.JsonOk(&BlockTransactionCountResponse{TxsNumber: txnsNumber})
}

type TransactionByBlockIDAndIndexRequest struct {
	BlockID BlockID `json:"block_id"`
	TxIndex int     `json:"index"`
}

type TransactionByBlockIDAndIndexResponse struct {
	Tx  *rpc.Transaction `json:"tx"`
	Err *jsonrpc.Error   `json:"err"`
}

func (c *Cairo) GetTransactionByBlockIDAndIndex(ctx *context.ReadContext) {
	var tq TransactionByBlockIDAndIndexRequest
	err := ctx.BindJson(&tq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &TransactionResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	//check if the index is valid
	if tq.TxIndex < 0 {
		ctx.Json(http.StatusBadRequest, &TransactionByBlockIDAndIndexResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, "index must be a non-negative integer")})
		return
	}
	//check if the block is pending
	compactBlock, err := c.getYuCompactBlock(tq.BlockID)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &TransactionByBlockIDAndIndexResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	if uint64(tq.TxIndex) >= uint64(len(compactBlock.TxnsHashes)) {
		ctx.Json(http.StatusInternalServerError, &TransactionByBlockIDAndIndexResponse{Err: jsonrpc.Err(jsonrpc.InternalError, "index out of range")})
		return
	}
	//get the transaction by index
	txHash := compactBlock.TxnsHashes[tq.TxIndex]
	signedTx, err := c.TxDB.GetTxn(common.Hash(txHash.Bytes()))
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &TransactionByBlockIDAndIndexResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	txReq := new(TxRequest)
	err = signedTx.BindJson(txReq)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &TransactionByBlockIDAndIndexResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	ctx.JsonOk(&TransactionByBlockIDAndIndexResponse{Tx: &txReq.Tx.Transaction})
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
		ctx.Json(http.StatusBadRequest, &TransactionStatusResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	stxn, _ := c.Pool.GetTxn(tr.Hash.Bytes()) // will not return error here
	if stxn != nil {
		ctx.JsonOk(&TransactionStatusResponse{Status: &rpc.TransactionStatus{Finality: rpc.TxnStatusReceived}})
		return
	}

	starkReceipt, err := c.getReceipt(tr.Hash)
	if err != nil {
		// TODO: when ErrTxnHashNotFound, should fetch from ETH L1
		ctx.Json(http.StatusInternalServerError, &TransactionStatusResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	ctx.JsonOk(&TransactionStatusResponse{Status: &rpc.TransactionStatus{
		Finality:  rpc.TxnStatus(starkReceipt.FinalityStatus),
		Execution: starkReceipt.ExecutionStatus,
	}})
}

type StateUpdateRequest struct {
	BlockID BlockID `json:"block_id"`
}

type StateUpdateResponse struct {
	StateUpdate *StateUpdate   `json:"state_update"`
	Err         *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetStateUpdate(ctx *context.ReadContext) {
	var tr StateUpdateRequest
	err := ctx.BindJson(&tr)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &TransactionStatusResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	update, err := c.cairoState.GetStateUpdateByNumber(tr.BlockID.Number)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &StateUpdateResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	nonces := make([]Nonce, 0, len(update.StateDiff.Nonces))
	for addr, nonce := range update.StateDiff.Nonces {
		nonces = append(nonces, Nonce{ContractAddress: addr, Nonce: *nonce})
	}

	storageDiffs := make([]StorageDiff, 0, len(update.StateDiff.StorageDiffs))
	for addr, diffs := range update.StateDiff.StorageDiffs {
		entries := make([]Entry, 0, len(diffs))
		for key, value := range diffs {
			entries = append(entries, Entry{
				Key:   key,
				Value: *value,
			})
		}

		storageDiffs = append(storageDiffs, StorageDiff{
			Address:        addr,
			StorageEntries: entries,
		})
	}

	deployedContracts := make([]DeployedContract, 0, len(update.StateDiff.DeployedContracts))
	for addr, classHash := range update.StateDiff.DeployedContracts {
		deployedContracts = append(deployedContracts, DeployedContract{
			Address:   addr,
			ClassHash: *classHash,
		})
	}

	declaredClasses := make([]DeclaredClass, 0, len(update.StateDiff.DeclaredV1Classes))
	for classHash, compiledClassHash := range update.StateDiff.DeclaredV1Classes {
		declaredClasses = append(declaredClasses, DeclaredClass{
			ClassHash:         classHash,
			CompiledClassHash: *compiledClassHash,
		})
	}

	replacedClasses := make([]ReplacedClass, 0, len(update.StateDiff.ReplacedClasses))
	for addr, classHash := range update.StateDiff.ReplacedClasses {
		replacedClasses = append(replacedClasses, ReplacedClass{
			ClassHash:       *classHash,
			ContractAddress: addr,
		})
	}

	ctx.JsonOk(&StateUpdateResponse{StateUpdate: &StateUpdate{
		BlockHash: update.BlockHash,
		OldRoot:   update.OldRoot,
		NewRoot:   update.NewRoot,
		StateDiff: &StateDiff{
			DeprecatedDeclaredClasses: update.StateDiff.DeclaredV0Classes,
			DeclaredClasses:           declaredClasses,
			ReplacedClasses:           replacedClasses,
			Nonces:                    nonces,
			StorageDiffs:              storageDiffs,
			DeployedContracts:         deployedContracts,
		},
	}})
}

type NonceRequest struct {
	BlockID BlockID    `json:"block_id"`
	Addr    *felt.Felt `json:"addr"`
}

type NonceResponse struct {
	Nonce *felt.Felt     `json:"nonce"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetNonce(ctx *context.ReadContext) {
	var nq NonceRequest
	err := ctx.BindJson(&nq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &NonceResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}

	var nonce *felt.Felt
	switch {
	case nq.BlockID.Latest || nq.BlockID.Pending:
		nonce, err = c.cairoState.ContractNonce(nq.Addr)
		fmt.Print("err: ", err)
	default:
		nonce, err = c.cairoState.ContractNonceAt(nq.Addr, nq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &NonceResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	ctx.JsonOk(&NonceResponse{Nonce: nonce})
}

type ClassRequest struct {
	BlockID   BlockID    `json:"block_id"`
	ClassHash *felt.Felt `json:"class_hash"`
}

type ClassAtRequest struct {
	BlockID BlockID    `json:"block_id"`
	Addr    *felt.Felt `json:"addr"`
}

type ClassResponse struct {
	Class *rpc.Class     `json:"class"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetClass(ctx *context.ReadContext) {
	var cq ClassRequest
	err := ctx.BindJson(&cq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &ClassResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err)})
		return
	}

	c.getClass(ctx, &cq.BlockID, cq.ClassHash)
}

func (c *Cairo) GetClassAt(ctx *context.ReadContext) {
	var cq ClassAtRequest
	err := ctx.BindJson(&cq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &ClassResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	var classHash *felt.Felt
	switch {
	case cq.BlockID.Latest || cq.BlockID.Pending:
		classHash, err = c.cairoState.ContractClassHash(cq.Addr)
	default:
		classHash, err = c.cairoState.ContractClassHashAt(cq.Addr, cq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &ClassResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	c.getClass(ctx, &cq.BlockID, classHash)
}

func (c *Cairo) getClass(ctx *context.ReadContext, blockID *BlockID, classHash *felt.Felt) {
	class, err := c.cairoState.Class(classHash)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &ClassResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	if !blockID.Latest {
		if blockID.Number < class.At {
			ctx.Json(http.StatusBadRequest, &ClassResponse{Err: rpc.ErrClassHashNotFound})
			return
		}
	}
	rpcClass := declaredClassToClass(class)
	if rpcClass != nil {
		ctx.JsonOk(&ClassResponse{Class: rpcClass})
	} else {
		ctx.Json(http.StatusBadRequest, &ClassResponse{Err: rpc.ErrClassHashNotFound})
	}
}

type ClassHashAtRequest struct {
	BlockID BlockID    `json:"block_id"`
	Addr    *felt.Felt `json:"addr"`
}

type ClassHashAtResponse struct {
	ClassHash *felt.Felt     `json:"class_hash"`
	Err       *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetClassHashAt(ctx *context.ReadContext) {
	var cq ClassHashAtRequest
	err := ctx.BindJson(&cq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &ClassHashAtResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}

	var classHash *felt.Felt
	switch {
	case cq.BlockID.Latest || cq.BlockID.Pending:
		classHash, err = c.cairoState.ContractClassHash(cq.Addr)
	default:
		classHash, err = c.cairoState.ContractClassHashAt(cq.Addr, cq.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &ClassHashAtResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	ctx.JsonOk(&ClassHashAtResponse{ClassHash: classHash})
}

type StorageRequest struct {
	BlockID BlockID    `json:"block_id"`
	Addr    *felt.Felt `json:"addr"`
	Key     *felt.Felt `json:"key"`
}

type StorageResponse struct {
	Value *felt.Felt     `json:"value"`
	Err   *jsonrpc.Error `json:"err"`
}

func (c *Cairo) GetStorage(ctx *context.ReadContext) {
	var sr StorageRequest
	err := ctx.BindJson(&sr)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &StorageResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}

	var value *felt.Felt
	switch {
	case sr.BlockID.Latest || sr.BlockID.Pending:
		value, err = c.cairoState.ContractStorage(sr.Addr, sr.Key)
	default:
		value, err = c.cairoState.ContractStorageAt(sr.Addr, sr.Key, sr.BlockID.Number)
	}
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &StorageResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}
	ctx.JsonOk(&StorageResponse{Value: value})
}

type SimulateRequest struct {
	BlockID         BlockID                      `json:"block_id"`
	Txs             []rpc.BroadcastedTransaction `json:"txs"`
	SimulationFlags []SimulationFlag             `json:"simulation_flags"`
	LegacyJson      bool                         `json:"legacy_json"`
	ErrOnRevert     bool                         `json:"err_on_revert"`
}

type SimulateResponse struct {
	Txs []SimulatedTransaction `json:"txs"`
	Err *jsonrpc.Error         `json:"err"`
}

func (c *Cairo) SimulateTransactions(ctx *context.ReadContext) {
	var (
		gasPrice     = new(felt.Felt).SetUint64(1)
		gasPriceSTRK = new(felt.Felt).SetUint64(1)
	)

	var sq SimulateRequest
	err := ctx.BindJson(&sq)
	if err != nil {
		ctx.Json(http.StatusBadRequest, &SimulateResponse{Err: jsonrpc.Err(jsonrpc.InvalidJSON, err.Error())})
		return
	}
	skipFeeCharge := slices.Contains(sq.SimulationFlags, SkipFeeChargeFlag)
	skipValidate := slices.Contains(sq.SimulationFlags, SkipValidateFlag)
	// TODO: try get more BlockID
	block, err := c.GetCurrentBlock()
	if err != nil {
		ctx.Json(http.StatusInternalServerError, &SimulateResponse{Err: jsonrpc.Err(jsonrpc.InternalError, err.Error())})
		return
	}

	var txns []core.Transaction
	var classes []core.Class

	paidFeesOnL1 := make([]*felt.Felt, 0)
	for idx := range sq.Txs {
		txn, declaredClass, paidFeeOnL1, aErr := AdaptBroadcastedTransaction(&sq.Txs[idx], c.network)
		if aErr != nil {
			ctx.Json(http.StatusInternalServerError, &SimulateResponse{Err: jsonrpc.Err(jsonrpc.InvalidParams, aErr.Error())})
			return
		}

		if paidFeeOnL1 != nil {
			paidFeesOnL1 = append(paidFeesOnL1, paidFeeOnL1)
		}

		txns = append(txns, txn)
		if declaredClass != nil {
			classes = append(classes, declaredClass)
		}
	}
	if c.sequencerAddr == nil {
		c.sequencerAddr = core.NetworkBlockHashMetaInfo(c.network).FallBackSequencerAddress
	}
	fees, traces, err := c.cairoVM.Execute(
		txns, classes, uint64(block.Height),
		block.Timestamp, c.sequencerAddr,
		c.cairoState.State, c.network, paidFeesOnL1,
		skipFeeCharge, skipValidate, sq.ErrOnRevert,
		gasPrice, gasPriceSTRK, sq.LegacyJson,
	)
	if err != nil {
		if errors.Is(err, utils.ErrResourceBusy) {
			ctx.Json(http.StatusInternalServerError, &SimulateResponse{Err: rpc.ErrInternal.CloneWithData(err.Error())})
			return
		}
		var txnExecutionError vm.TransactionExecutionError
		if errors.As(err, &txnExecutionError) {
			ctx.Json(http.StatusInternalServerError, &SimulateResponse{Err: makeTransactionExecutionError(&txnExecutionError)})
			return
		}
		ctx.Json(http.StatusInternalServerError, &SimulateResponse{Err: rpc.ErrUnexpectedError.CloneWithData(err.Error())})
		return
	}

	var result []SimulatedTransaction
	for i, overallFee := range fees {
		feeUnit := feeUnit(txns[i])

		if feeUnit == rpc.FRI {
			if gasPrice = gasPriceSTRK; gasPrice == nil {
				gasPrice = &felt.Zero
			}
		}

		estimate := FeeEstimate{
			GasConsumed:     new(felt.Felt).Div(overallFee, gasPrice),
			GasPrice:        gasPrice,
			DataGasConsumed: &felt.Zero, //FixMe
			DataGasPrice:    &felt.Zero, //FixMe
			OverallFee:      overallFee,
		}

		if !sq.LegacyJson {
			estimate.Unit = utils.Ptr(feeUnit)
		}
		result = append(result, SimulatedTransaction{
			TransactionTrace: &traces[i],
			FeeEstimation:    estimate,
		})
	}

	ctx.JsonOk(&SimulateResponse{Txs: result})
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

func makeTransactionExecutionError(err *vm.TransactionExecutionError) *jsonrpc.Error {
	return rpc.ErrTransactionExecutionError.CloneWithData(rpc.TransactionExecutionErrorData{
		TransactionIndex: err.Index,
		ExecutionError:   err.Cause.Error(),
	})
}
