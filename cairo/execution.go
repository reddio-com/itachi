package cairo

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
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

func (c *Cairo) adaptBroadcastedTransaction(bcTxn *rpc.BroadcastedTransaction) (core.Transaction, core.Class, *felt.Felt, error) {

}
