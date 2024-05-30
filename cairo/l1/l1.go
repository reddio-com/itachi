package l1

import (
	"context"
	"fmt"
	"itachi/cairo/config"
	"itachi/cairo/l1/contract"
	"itachi/cairo/starknetrpc"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/core/kernel"
)

const CairoTripod = "cairo"

type L1 struct {
	itachi      *kernel.Kernel
	ethL1       *EthSubscriber
	starknetrpc *starknetrpc.StarknetRPC
}

func NewL1(itachi *kernel.Kernel, cfg *config.Config, s *starknetrpc.StarknetRPC) (*L1, error) {
	ethL1, err := NewEthSubscriber(cfg.EthClientAddress, common.HexToAddress(cfg.EthContractAddress))
	if err != nil {
		return nil, err
	}
	return &L1{
		itachi:      itachi,
		ethL1:       ethL1,
		starknetrpc: s,
	}, nil
}

func (l *L1) Run(ctx context.Context) {
	msgChan := make(chan *contract.StarknetLogMessageToL2)
	l.ethL1.WatchLogMessageToL2(ctx, msgChan, nil, nil, nil)

	// Listen for msgChan
	go func() {
		for {
			select {
			case msg := <-msgChan:
				broadcastedTxn, err := convertL1TxnToBroadcastedTxn(msg)
				if err != nil {
					logrus.Errorf("Error converting L1 txn to broadcasted txn: %v", err)
					continue
				}
				response, jsonRpcErr := l.starknetrpc.AddTransaction(*broadcastedTxn)
				if jsonRpcErr != nil {
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

func convertL1TxnToBroadcastedTxn(event *contract.StarknetLogMessageToL2) (*rpc.BroadcastedTransaction, error) {
	callData := make([]*felt.Felt, 0)
	callData = append(callData, new(felt.Felt).SetBigInt(event.FromAddress.Big()))
	for _, payload := range event.Payload {
		data := new(felt.Felt).SetBigInt(payload)
		callData = append(callData, data)
	}

	maxU128 := new(big.Int).SetUint64(1<<64 - 1)

	// Check if fee exceeds u128 max value
	if event.Fee.Cmp(maxU128) > 0 {
		return nil, fmt.Errorf("fee exceeds u128 max value")
	}

	return &rpc.BroadcastedTransaction{
		Transaction: rpc.Transaction{
			Type:            rpc.TxnL1Handler,
			ContractAddress: new(felt.Felt).SetBigInt(event.ToAddress),
			Nonce:           new(felt.Felt).SetBigInt(event.Nonce),
			CallData:        &callData,
		},
		PaidFeeOnL1: new(felt.Felt).SetBigInt(event.Fee),
	}, nil
}
