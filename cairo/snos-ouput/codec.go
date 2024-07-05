package snos_ouput

import (
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"itachi/cairo/adapters"
	"math/big"
)

type StarknetOsOutput struct {

	// The state commitment before this block.
	PrevStateRoot *felt.Felt `json:"old_root"`
	// The state commitment after this block.
	NewStateRoot *felt.Felt `json:"new_root"`
	// The number (height) of this block.
	BlockNumber *felt.Felt `json:"block_number"`
	// The hash of this block.
	BlockHash *felt.Felt `json:"block_hash"`
	// The Starknet chain config hash
	ConfigHash *felt.Felt `json:"config_hash"`

	// KZG_DA
	KzgDA *felt.Felt `json:"kzgDA"`

	MessagesToL1 []*adapters.MessageL2ToL1 `json:"messages_to_l1"`
	MessagesToL2 []*adapters.MessageL1ToL2 `json:"messages_to_l2"`
}

type SnosCodec interface {
	SizeInFelts() int
	EncodeToFelts() []*felt.Felt
	Decode(input []*felt.Felt) error
}

func CalculateFeltSizeInFelts(f *felt.Felt) int {

	return 1
}

func (o *StarknetOsOutput) SizeInFelts() int {
	size := 0
	size += CalculateFeltSizeInFelts(o.PrevStateRoot)
	size += CalculateFeltSizeInFelts(o.NewStateRoot)
	size += CalculateFeltSizeInFelts(o.BlockNumber)
	size += CalculateFeltSizeInFelts(o.BlockHash)
	size += CalculateFeltSizeInFelts(o.ConfigHash)
	size += CalculateFeltSizeInFelts(o.KzgDA)
	// for messagesToL1 length field
	size += 1

	for _, msg := range o.MessagesToL1 {
		size += msg.SizeInFelts()
	}
	//for _, msg := range o.MessagesToL2 {
	//	size += msg.SizeInFelts()
	//}
	return size
}

func (output *StarknetOsOutput) EncodeTo() ([]*big.Int, error) {
	var result []*big.Int

	// Convert type Felt to *big.Int
	result = append(result, output.PrevStateRoot.BigInt(new(big.Int)))
	//print tempBigInt
	fmt.Println("PrevStateRoot as *big.Int:", result[len(result)-1].String())

	tempBigInt := new(big.Int)
	tempBigInt.SetUint64(0)

	result = append(result, output.NewStateRoot.BigInt(new(big.Int)))

	result = append(result, output.BlockNumber.BigInt(new(big.Int)))

	result = append(result, output.BlockHash.BigInt(new(big.Int)))

	result = append(result, output.ConfigHash.BigInt(new(big.Int)))

	// set KzgDA 0
	result = append(result, big.NewInt(0))

	if len(output.MessagesToL1) == 0 {
		result = append(result, big.NewInt(int64(0)))
	} else {
		for _, msg := range output.MessagesToL1 {
			msgBigInts, err := msg.EncodeTo()
			if err != nil {
				return nil, err
			}

			messagesToL1Size := big.NewInt(int64(msg.SizeInFelts()))
			result = append(result, messagesToL1Size)

			result = append(result, msgBigInts...)
		}
	}

	if len(output.MessagesToL2) == 0 {
		result = append(result, big.NewInt(int64(0)))
	} else {
		for _, msg := range output.MessagesToL2 {
			msgBigInts, err := msg.EncodeTo()
			if err != nil {
				return nil, err
			}

			messagesToL2Size := big.NewInt(int64(msg.SizeInFelts()))
			result = append(result, messagesToL2Size)

			result = append(result, msgBigInts...)
		}
	}

	return result, nil

}
