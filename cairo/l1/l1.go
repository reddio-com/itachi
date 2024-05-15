package l1

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"itachi/cairo/l1/contract"
)

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
