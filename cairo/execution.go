package cairo

import (
	"github.com/yu-org/yu/core/startup"
	"github.com/yu-org/yu/core/types"
)

func (c *Cairo) TxnExecute(block *types.Block) error {
	for _, txn := range block.Txns {
		wrCall := txn.Raw.WrCall
		wr, err := startup.Land.GetWriting(wrCall.TripodName, wrCall.FuncName)
		if err != nil {
			return err
		}

	}
}
