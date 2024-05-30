package l1

import (
	"context"
	"itachi/cairo/config"
	"itachi/cairo/l1/contract"
	"itachi/cairo/starknetrpc"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/core/kernel"
)

const CairoTripod = "cairo"

type L1 struct {
	itachi *kernel.Kernel
	ethL1  *EthSubscriber
}

func NewL1(itachi *kernel.Kernel, cfg *config.Config) (*L1, error) {
	ethL1, err := NewEthSubscriber(cfg.EthClientAddress, common.HexToAddress(cfg.EthContractAddress))
	if err != nil {
		return nil, err
	}
	return &L1{
		itachi: itachi,
		ethL1:  ethL1,
	}, nil
}

func (l *L1) Run(ctx context.Context, s *starknetrpc.StarknetRPC) {
	msgChan := make(chan *contract.StarknetLogMessageToL2)
	l.ethL1.WatchLogMessageToL2(ctx, msgChan, nil, nil, nil)

	// Listen for msgChan
	go func() {
		for {
			select {
			case msg := <-msgChan:
				l1Txn := parseEventToL1Txn(msg)
				broadcastedTxn := convertL1TxnToBroadcastedTxn(l1Txn)
				response, err := s.AddTransaction(*broadcastedTxn)
				if err != nil {
					logrus.Errorf("Error adding transaction: %v", err)
				} else {
					logrus.Infof("Transaction added: %v", response)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

}

func parseEventToL1Txn(event *contract.StarknetLogMessageToL2) *core.L1HandlerTransaction {
	callData := make([]*felt.Felt, 0)
	callData = append(callData, new(felt.Felt).SetBigInt(event.FromAddress.Big()))
	for _, payload := range event.Payload {
		data := new(felt.Felt).SetBigInt(payload)
		callData = append(callData, data)
	}
	return &core.L1HandlerTransaction{
		ContractAddress:    new(felt.Felt).SetBigInt(event.ToAddress),
		EntryPointSelector: new(felt.Felt).SetBigInt(event.Selector),
		Nonce:              new(felt.Felt).SetBigInt(event.Nonce),
		CallData:           callData,
		Version:            new(core.TransactionVersion),
	}
}

func convertL1TxnToBroadcastedTxn(l1Txn *core.L1HandlerTransaction) *rpc.BroadcastedTransaction {
	return &rpc.BroadcastedTransaction{
		Transaction: rpc.Transaction{
			Type:            rpc.TxnL1Handler,
			ContractAddress: l1Txn.ContractAddress,
			Nonce:           l1Txn.Nonce,
			CallData:        &l1Txn.CallData,
		},
		PaidFeeOnL1: nil,
	}
}
