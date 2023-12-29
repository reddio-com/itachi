package cairo

import (
	"encoding/json"
	"errors"
	"github.com/NethermindEth/juno/adapters/sn2core"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/starknet"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/jinzhu/copier"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/startup"
	"github.com/yu-org/yu/core/types"
)

func (c *Cairo) TxnExecute(block *types.Block) error {
	var (
		starknetTxns []core.Transaction
		classes      []core.Class
		paidFeesOnL1 []*felt.Felt
	)
	for _, txn := range block.Txns {
		wrCall := txn.Raw.WrCall
		ctx, err := context.NewWriteContext(txn, block)
		if err != nil {
			return err
		}
		_, err = startup.Land.GetWriting(wrCall.TripodName, wrCall.FuncName)
		if err != nil {
			return err
		}
		btxn := new(rpc.BroadcastedTransaction)
		err = ctx.BindJson(btxn)
		if err != nil {
			return err
		}
		tx, class, paidFeeOnL1, err := c.adaptBroadcastedTransaction(btxn)
		if err != nil {
			return err
		}
		starknetTxns = append(starknetTxns, tx)
		classes = append(classes, class)
		paidFeesOnL1 = append(paidFeesOnL1, paidFeeOnL1)
	}
	blockNumber := uint64(block.Height)
	blockTimestamp := block.Timestamp

	actualFee, traces, err := c.execute(starknetTxns, classes, blockNumber, blockTimestamp, paidFeesOnL1)
	if err != nil {
		return err
	}
}

func (c *Cairo) call(
	contractAddr, classHash, selector *felt.Felt,
	calldata []felt.Felt,
	blockNumber, blockTimestamp uint64,
) ([]*felt.Felt, error) {
	return c.cairoVM.Call(contractAddr, classHash, selector, calldata, blockNumber, blockTimestamp, c.state, c.network)
}

func (c *Cairo) execute(
	txns []core.Transaction, declaredClasses []core.Class,
	blockNumber, blockTimestamp uint64, paidFeesOnL1 []*felt.Felt,
	gasPriceWEI, gasPriceSTRK *felt.Felt, legacyTraceJSON bool,
) ([]*felt.Felt, []vm.TransactionTrace, error) {
	return c.cairoVM.Execute(txns, declaredClasses, blockNumber, blockTimestamp, c.sequencerAddr,
		c.state, c.network, paidFeesOnL1, c.cfg.SkipChargeFee, c.cfg.SkipValidate, c.cfg.ErrOnRevert, gasPriceWEI, gasPriceSTRK, legacyTraceJSON)
}

func (c *Cairo) adaptBroadcastedTransaction(bcTxn *rpc.BroadcastedTransaction) (core.Transaction, core.Class, *felt.Felt, error) {
	var feederTxn starknet.Transaction
	if err := copier.Copy(&feederTxn, bcTxn.Transaction); err != nil {
		return nil, nil, nil, err
	}

	txn, err := sn2core.AdaptTransaction(&feederTxn)
	if err != nil {
		return nil, nil, nil, err
	}

	var declaredClass core.Class
	if len(bcTxn.ContractClass) != 0 {
		declaredClass, err = adaptDeclaredClass(bcTxn.ContractClass)
		if err != nil {
			return nil, nil, nil, err
		}
	} else if bcTxn.Type == rpc.TxnDeclare {
		return nil, nil, nil, errors.New("declare without a class definition")
	}

	if t, ok := txn.(*core.DeclareTransaction); ok {
		t.ClassHash, err = declaredClass.Hash()
		if err != nil {
			return nil, nil, nil, err
		}
	}

	txnHash, err := core.TransactionHash(txn, c.network)
	if err != nil {
		return nil, nil, nil, err
	}

	var paidFeeOnL1 *felt.Felt
	switch t := txn.(type) {
	case *core.DeclareTransaction:
		t.TransactionHash = txnHash
	case *core.InvokeTransaction:
		t.TransactionHash = txnHash
	case *core.DeployAccountTransaction:
		t.TransactionHash = txnHash
	case *core.L1HandlerTransaction:
		t.TransactionHash = txnHash
		paidFeeOnL1 = bcTxn.PaidFeeOnL1
	default:
		return nil, nil, nil, errors.New("unsupported transaction")
	}

	if txn.Hash() == nil {
		return nil, nil, nil, errors.New("deprecated transaction type")
	}
	return txn, declaredClass, paidFeeOnL1, nil
}

func adaptDeclaredClass(declaredClass json.RawMessage) (core.Class, error) {
	var feederClass starknet.ClassDefinition
	err := json.Unmarshal(declaredClass, &feederClass)
	if err != nil {
		return nil, err
	}

	switch {
	case feederClass.V1 != nil:
		compiledClass, cErr := starknet.Compile(feederClass.V1)
		if cErr != nil {
			return nil, cErr
		}
		return sn2core.AdaptCairo1Class(feederClass.V1, compiledClass)
	case feederClass.V0 != nil:
		// strip the quotes
		base64Program := string(feederClass.V0.Program[1 : len(feederClass.V0.Program)-1])
		feederClass.V0.Program, err = utils.Gzip64Decode(base64Program)
		if err != nil {
			return nil, err
		}

		return sn2core.AdaptCairo0Class(feederClass.V0)
	default:
		return nil, errors.New("empty class")
	}
}
