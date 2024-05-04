package evm

import (
	"github.com/ethereum/go-ethereum/core/state"
)

type State struct {
	*state.StateDB
}
