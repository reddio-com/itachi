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
)

//func (c *Cairo) TxnExecute(block *types.Block) error {
//
//	receipts := make(map[common.Hash]*types.Receipt)
//
//	for _, txn := range block.Txns {
//		wrCall := txn.Raw.WrCall
//		ctx, err := context.NewWriteContext(txn, block)
//		if err != nil {
//			return err
//		}
//		wr, err := startup.Land.GetWriting(wrCall.TripodName, wrCall.FuncName)
//		if err != nil {
//			return err
//		}
//		err = wr(ctx)
//		rcpt := types.NewReceipt(ctx.Events, err, ctx.Extra)
//		rcpt.FillMetadata(block, txn, ctx.LeiCost)
//		receipts[txn.TxnHash] = rcpt
//	}
//	blockNumber := uint64(block.Height)
//
//	// commit cairoState
//	err := c.cairoState.Commit(blockNumber)
//	if err != nil {
//		return err
//	}
//	return c.TxDB.SetReceipts(receipts)
//}

func (c *Cairo) execute(
	txns []core.Transaction, declaredClasses []core.Class,
	blockNumber, blockTimestamp uint64, paidFeesOnL1 []*felt.Felt,
	gasPriceWEI, gasPriceSTRK *felt.Felt, legacyTraceJSON bool,
) ([]*felt.Felt, []vm.TransactionTrace, error) {
	return c.cairoVM.Execute(
		txns, declaredClasses, blockNumber, blockTimestamp, c.sequencerAddr,
		c.cairoState.PendingStateWriter, c.network, paidFeesOnL1, c.cfg.SkipChargeFee, c.cfg.SkipValidate,
		c.cfg.ErrOnRevert, gasPriceWEI, gasPriceSTRK, legacyTraceJSON,
	)
}

func AdaptBroadcastedTransaction(bcTxn *rpc.BroadcastedTransaction, network utils.Network) (core.Transaction, core.Class, *felt.Felt, error) {
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

	txnHash, err := core.TransactionHash(txn, network)
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

func (c *Cairo) adaptBroadcastedTransaction(bcTxn *rpc.BroadcastedTransaction) (core.Transaction, core.Class, *felt.Felt, error) {
	return AdaptBroadcastedTransaction(bcTxn, c.network)
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
