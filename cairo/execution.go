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
	"github.com/yu-org/yu/core/result"
	"github.com/yu-org/yu/core/startup"
	"github.com/yu-org/yu/core/types"
)

func (c *Cairo) TxnExecute(block *types.Block) error {

	var results []*result.Result

	for _, txn := range block.Txns {
		wrCall := txn.Raw.WrCall
		ctx, err := context.NewWriteContext(txn, block)
		if err != nil {
			return err
		}
		wr, err := startup.Land.GetWriting(wrCall.TripodName, wrCall.FuncName)
		if err != nil {
			return err
		}
		//txReq := new(TxRequest)
		//err = ctx.BindJson(txReq)
		//if err != nil {
		//	return err
		//}
		//tx, class, paidFeeOnL1, err := c.adaptBroadcastedTransaction(txReq.Tx)
		//if err != nil {
		//	return err
		//}
		//starknetTxns = append(starknetTxns, tx)
		//classes = append(classes, class)
		//paidFeesOnL1 = append(paidFeesOnL1, paidFeeOnL1)
		err = wr(ctx)
		if err != nil {
			ctx.EmitError(err)
			continue
		}
		for _, event := range ctx.Events {
			results = append(results, result.NewEvent(event))
		}
		if ctx.Error != nil {
			results = append(results, result.NewError(ctx.Error))
		}
	}
	blockNumber := uint64(block.Height)
	//blockTimestamp := block.Timestamp
	//blockHash := block.Hash

	//pendingState := c.newPendingStateWriter()
	//_, traces, err := c.execute(
	//	pendingState, starknetTxns, classes, blockNumber, blockTimestamp,
	//	paidFeesOnL1, &felt.Zero, &felt.Zero, false,
	//)
	//if err != nil {
	//	return err
	//}

	// commit cairoState
	err := c.cairoState.Commit(blockNumber)
	if err != nil {
		return err
	}
	//stateDiff, newClasses := pendingState.StateDiffAndClasses()
	//err = c.cairoState.Update(blockNumber, stateDiff, newClasses)
	//if err != nil {
	//	return err
	//}

	// store events

	//for _, trace := range traces {
	//	if trace.ExecuteInvocation != nil {
	//		for _, event := range trace.ExecuteInvocation.Events {
	//			eventByt, terr := json.Marshal(event)
	//			if terr != nil {
	//				return terr
	//			}
	//			callerByt := trace.ExecuteInvocation.CallerAddress.Bytes()
	//			caller := common.BytesToAddress(callerByt[:])
	//			yuEvent := &result.Event{
	//				Caller:    &caller,
	//				BlockHash: blockHash,
	//				Height:    block.Height,
	//				Value:     eventByt,
	//			}
	//			results = append(results, result.NewEvent(yuEvent))
	//		}
	//	}
	//
	//}
	return c.TxDB.SetResults(results)
}

func (c *Cairo) execute(
	txns []core.Transaction, declaredClasses []core.Class,
	blockNumber, blockTimestamp uint64, paidFeesOnL1 []*felt.Felt,
	gasPriceWEI, gasPriceSTRK *felt.Felt, legacyTraceJSON bool,
) ([]*felt.Felt, []vm.TransactionTrace, error) {
	return c.cairoVM.Execute(
		txns, declaredClasses, blockNumber, blockTimestamp, c.sequencerAddr,
		c.cairoState, c.network, paidFeesOnL1, c.cfg.SkipChargeFee, c.cfg.SkipValidate,
		c.cfg.ErrOnRevert, gasPriceWEI, gasPriceSTRK, legacyTraceJSON,
	)
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
