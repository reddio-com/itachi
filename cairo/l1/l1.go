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

type L1 struct {
	itachi      *kernel.Kernel
	ethL1       *EthSubscriber
	starknetRPC *starknetrpc.StarknetRPC
}

func NewL1(itachi *kernel.Kernel, cfg *config.Config, s *starknetrpc.StarknetRPC) (*L1, error) {
	ethL1, err := NewEthSubscriber(cfg.EthClientAddress, common.HexToAddress(cfg.EthContractAddress))
	if err != nil {
		return nil, err
	}
	return &L1{
		itachi:      itachi,
		ethL1:       ethL1,
		starknetRPC: s,
	}, nil
}

func StartupL1(itachi *kernel.Kernel, cfg *config.Config, s *starknetrpc.StarknetRPC) {
	if cfg.EnableL1 {
		l1, err := NewL1(itachi, cfg, s)
		if err != nil {
			logrus.Fatal("init L1 client failed: ", err)
		}
		err = l1.Run(context.Background())
		if err != nil {
			logrus.Fatal("l1 client run failed: ", err)
		}
	}
}

func (l *L1) Run(ctx context.Context) error {
	msgChan := make(chan *contract.StarknetLogMessageToL2)
	sub, err := l.ethL1.WatchLogMessageToL2(ctx, msgChan, nil, nil, nil)
	if err != nil {
		return err
	}

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
				response, jsonRpcErr := l.starknetRPC.AddTransaction(*broadcastedTxn)
				if jsonRpcErr != nil {
					logrus.Errorf("Error adding transaction: %v", err)
				} else {
					logrus.Infof("L1 Transaction added: %v", response)
				}
			case subErr := <-sub.Err():
				logrus.Errorf("L1 update subscription failed: %v, Resubscribing...", subErr)
				sub.Unsubscribe()

				sub, err = l.ethL1.WatchLogMessageToL2(ctx, msgChan, nil, nil, nil)
				if err != nil {
					logrus.Errorf("Resubscribe failed: %v", err)
				}
			case <-ctx.Done():
				sub.Unsubscribe()
				return
			}
		}
	}()
	return nil
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
			Version:         new(felt.Felt).SetUint64(0),
		},
		PaidFeeOnL1: new(felt.Felt).SetBigInt(event.Fee),
	}, nil
}
