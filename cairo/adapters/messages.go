package adapters

import (
	"errors"
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/ethereum/go-ethereum/common"

	"math/big"
)

type MessageL1ToL2 struct {
	// The address of the L1 contract sending the message.
	From common.Address `json:"from_address" validate:"required"`
	// The address of the L1 contract sending the message.
	To       felt.Felt `json:"to_address" validate:"required"`
	Nonce    felt.Felt `json:"nonce" validate:"required"`
	Selector felt.Felt `json:"entry_point_selector" validate:"required"`
	// The payload of the message.
	Payload []felt.Felt `json:"payload" validate:"required"`
}

func (m *MessageL1ToL2) EncodeTo() ([]*big.Int, error) {
	var result []*big.Int

	// From
	result = append(result, new(big.Int).SetBytes(m.From.Bytes()))

	// To
	if m.To.IsZero() {
		return nil, errors.New("To field is zero (invalid)")
	}
	result = append(result, m.To.BigInt(new(big.Int)))

	// Nonce
	result = append(result, m.Nonce.BigInt(new(big.Int)))

	// Selector
	if m.Selector.IsZero() {
		return nil, errors.New("Selector field is zero (invalid)")
	}
	result = append(result, m.Selector.BigInt(new(big.Int)))

	payloadSize := big.NewInt(int64(len(m.Payload)))
	result = append(result, payloadSize)
	// Payload
	for _, p := range m.Payload {
		if !p.IsZero() {
			result = append(result, p.BigInt(new(big.Int)))
		}
	}

	return result, nil
}

func (m *MessageL1ToL2) SizeInFelts() int {
	size := 0
	size += sizeOfCommonAddress(m.From)
	size += sizeOfFelt(m.To)
	size += sizeOfFelt(m.Selector)
	// for payload length field
	size += 1
	for _, p := range m.Payload {
		size += sizeOfFelt(p)
	}
	return size
}

func sizeOfCommonAddress(addr common.Address) int {
	return 1
}

func sizeOfFelt(f felt.Felt) int {
	return 1
}

// MessageL2ToL1 L2ToL1Message
type MessageL2ToL1 struct {
	From *felt.Felt     `json:"from_address,omitempty"`
	To   common.Address `json:"to_address"`

	Payload []*felt.Felt `json:"payload"`
}

func (m *MessageL2ToL1) EncodeTo() ([]*big.Int, error) {
	var result []*big.Int

	// From
	if m.From != nil {
		result = append(result, m.From.BigInt(new(big.Int)))
	} else {
		return nil, errors.New("From field is nil")
	}

	// To
	result = append(result, new(big.Int).SetBytes(m.To.Bytes()))
	fmt.Println("To:", new(big.Int).SetBytes(m.To.Bytes()))

	payloadSize := big.NewInt(int64(len(m.Payload)))
	result = append(result, payloadSize)
	// Payload
	for _, p := range m.Payload {
		if p != nil {
			result = append(result, p.BigInt(new(big.Int)))
		}
	}

	return result, nil
}

func (m *MessageL2ToL1) SizeInFelts() int {
	size := 0

	if m.From != nil {
		size += sizeOfFelt(*m.From)
	}

	size += sizeOfCommonAddress(m.To)

	// for payload length field
	size += 1

	for _, p := range m.Payload {
		size += sizeOfFelt(*p)
	}

	return size
}
