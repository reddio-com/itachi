package cairo

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"github.com/sirupsen/logrus"
)

type MockCairoVM struct {
}

func (m *MockCairoVM) Call(
	contractAddr, classHash, selector *felt.Felt,
	calldata []felt.Felt,
	blockNumber, blockTimestamp uint64,
	state core.StateReader, network utils.Network,
) ([]*felt.Felt, error) {
	logrus.Info("Mock CairoVM.Call() here!")
	return []*felt.Felt{&felt.Zero}, nil
}

func (m *MockCairoVM) Execute(
	txns []core.Transaction, declaredClasses []core.Class,
	blockNumber, blockTimestamp uint64,
	sequencerAddress *felt.Felt, state core.StateReader,
	network utils.Network, paidFeesOnL1 []*felt.Felt,
	skipChargeFee, skipValidate, errOnRevert bool,
	gasPriceWEI *felt.Felt, gasPriceSTRK *felt.Felt,
	legacyTraceJSON bool,
) ([]*felt.Felt, []vm.TransactionTrace, error) {
	logrus.Info("Mock CairoVM.Execute() here!")
	return []*felt.Felt{&felt.Zero},
		[]vm.TransactionTrace{
			{Type: vm.TxnDeployAccount},
			{Type: vm.TxnDeclare},
			{Type: vm.TxnInvoke},
		}, nil
}
