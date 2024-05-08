package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/yu-org/yu/core/context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func NewEnv(cfg *Config) *vm.EVM {
	txContext := vm.TxContext{
		Origin:     cfg.Origin,
		GasPrice:   cfg.GasPrice,
		BlobHashes: cfg.BlobHashes,
		BlobFeeCap: cfg.BlobFeeCap,
	}
	blockContext := vm.BlockContext{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     cfg.GetHashFn,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    cfg.GasLimit,
		BaseFee:     cfg.BaseFee,
		BlobBaseFee: cfg.BlobBaseFee,
		Random:      cfg.Random,
	}

	return vm.NewEVM(blockContext, txContext, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}

type Config struct {
	ChainConfig *params.ChainConfig
	Difficulty  *big.Int
	Origin      common.Address
	Coinbase    common.Address
	BlockNumber *big.Int
	Time        uint64
	GasLimit    uint64
	GasPrice    *big.Int
	Value       *big.Int
	Debug       bool
	EVMConfig   vm.Config
	BaseFee     *big.Int
	BlobBaseFee *big.Int
	BlobHashes  []common.Hash
	BlobFeeCap  *big.Int
	Random      *common.Hash

	State     *state.StateDB
	GetHashFn func(n uint64) common.Hash
}

// sets defaults on the config
func setDefaults(cfg *Config) {
	if cfg.ChainConfig == nil {
		cfg.ChainConfig = &params.ChainConfig{
			ChainID:             big.NewInt(1),
			HomesteadBlock:      new(big.Int),
			DAOForkBlock:        new(big.Int),
			DAOForkSupport:      false,
			EIP150Block:         new(big.Int),
			EIP155Block:         new(big.Int),
			EIP158Block:         new(big.Int),
			ByzantiumBlock:      new(big.Int),
			ConstantinopleBlock: new(big.Int),
			PetersburgBlock:     new(big.Int),
			IstanbulBlock:       new(big.Int),
			MuirGlacierBlock:    new(big.Int),
			BerlinBlock:         new(big.Int),
			LondonBlock:         new(big.Int),
		}
	}

	if cfg.Difficulty == nil {
		cfg.Difficulty = new(big.Int)
	}
	if cfg.GasLimit == 0 {
		cfg.GasLimit = math.MaxUint64
	}
	if cfg.GasPrice == nil {
		cfg.GasPrice = new(big.Int)
	}
	if cfg.Value == nil {
		cfg.Value = new(big.Int)
	}
	if cfg.BlockNumber == nil {
		cfg.BlockNumber = new(big.Int)
	}
	if cfg.GetHashFn == nil {
		cfg.GetHashFn = func(n uint64) common.Hash {
			return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
		}
	}
	if cfg.BaseFee == nil {
		cfg.BaseFee = big.NewInt(params.InitialBaseFee)
	}
	if cfg.BlobBaseFee == nil {
		cfg.BlobBaseFee = big.NewInt(params.BlobTxMinBlobGasprice)
	}
}

func (cfg *Config) ExecuteTxn(ctx *context.WriteContext) error {
	txReq := new(TxRequest)
	err := ctx.BindJson(txReq)
	if err != nil {
		return err
	}

	tx, class, paidFeeOnL1, err := e.adaptBroadcastedTransaction(txReq.Tx)
	if err != nil {
		return err
	}

	var (
		starknetTxns = make([]core.Transaction, 0)
		classes      = make([]core.Class, 0)
		paidFeesOnL1 = make([]*felt.Felt, 0)
	)
	if tx != nil {
		starknetTxns = append(starknetTxns, tx)
	}
	if class != nil {
		classes = append(classes, class)
	}
	if paidFeeOnL1 != nil {
		paidFeesOnL1 = append(paidFeesOnL1, paidFeeOnL1)
	}

	blockNumber := uint64(ctx.Block.Height)
	blockTimestamp := ctx.Block.Timestamp

	// FIXME: GasPriceWEI, GasPriceSTRK should be filled.
	gasPrice := new(felt.Felt).SetUint64(1)
	vmenv = NewEnv(cfg)
	actualFees, traces, err := vmenv.execute(
		starknetTxns, classes, blockNumber, blockTimestamp,
		paidFeesOnL1, gasPrice, gasPrice, txReq.LegacyTraceJson,
	)

	var starkReceipt *rpc.TransactionReceipt
	if len(traces) > 0 && len(actualFees) > 0 {
		starkReceipt = makeStarkReceipt(traces[0], ctx.Block, tx, actualFees[0])
	}
	if err != nil {
		// fmt.Printf("execute txn(%s) error: %v \n", tx.Hash(), err)
		starkReceipt = makeErrStarkReceipt(ctx.Block, tx, err)
	}

	//spew.Dump(starkReceipt)

	receiptByt, err := encoder.Marshal(starkReceipt)
	if err != nil {
		return err
	}
	ctx.EmitExtra(receiptByt)
	return nil
}
