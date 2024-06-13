package ethrpc

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	yucommon "github.com/yu-org/yu/common"
	yucore "github.com/yu-org/yu/core"
	"github.com/yu-org/yu/core/kernel"
	"itachi/evm"
	"math/big"
	"time"
)

type EthAPIBackend struct {
	allowUnprotectedTxs bool
	ethChainCfg         *params.ChainConfig
	chain               *kernel.Kernel
}

func (e *EthAPIBackend) SyncProgress() ethereum.SyncProgress {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) FeeHistory(ctx context.Context, blockCount uint64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, []*big.Int, []float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) BlobBaseFee(ctx context.Context) *big.Int {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) ChainDb() ethdb.Database {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) AccountManager() *accounts.Manager {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) ExtRPCEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) RPCGasCap() uint64 {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) RPCEVMTimeout() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) RPCTxFeeCap() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) UnprotectedAllowed() bool {
	return e.allowUnprotectedTxs
}

func (e *EthAPIBackend) SetHead(number uint64) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) CurrentHeader() *types.Header {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) CurrentBlock() *types.Header {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) Pending() (*types.Block, types.Receipts, *state.StateDB) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetEVM(ctx context.Context, msg *core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config, blockCtx *vm.BlockContext) *vm.EVM {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	block, err := e.chain.Chain.GetEndBlock()
	if err != nil {
		return err
	}
	signer := types.MakeSigner(e.ethChainCfg, new(big.Int).SetUint64(uint64(block.Height)), block.Timestamp)

	sender, err := types.Sender(signer, signedTx)
	if err != nil {
		return err
	}
	// TODO: should judge if is CreateTxrequest
	txReq := &evm.TxRequest{
		Input:    signedTx.Data(),
		Origin:   sender,
		GasLimit: signedTx.Gas(),
		GasPrice: signedTx.GasPrice(),
		Value:    signedTx.Value(),
	}
	byt, err := json.Marshal(txReq)
	if err != nil {
		return err
	}
	signedWrCall := &yucore.SignedWrCall{
		Call: &yucommon.WrCall{
			TripodName: SolidityTripod,
			FuncName:   "ExecuteTxn",
			Params:     string(byt),
		},
	}
	return e.chain.HandleTxn(signedWrCall)
}

func (e *EthAPIBackend) GetTransaction(ctx context.Context, txHash common.Hash) (bool, *types.Transaction, common.Hash, uint64, uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetPoolTransactions() (types.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) Stats() (pending int, queued int) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) TxPoolContent() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) TxPoolContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SubscribeNewTxsEvent(events chan<- core.NewTxsEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) ChainConfig() *params.ChainConfig {
	return e.ethChainCfg
}

func (e *EthAPIBackend) Engine() consensus.Engine {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetBody(ctx context.Context, hash common.Hash, number rpc.BlockNumber) (*types.Body, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) GetLogs(ctx context.Context, blockHash common.Hash, number uint64) ([][]*types.Log, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) BloomStatus() (uint64, uint64) {
	//TODO implement me
	panic("implement me")
}

func (e *EthAPIBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	//TODO implement me
	panic("implement me")
}
