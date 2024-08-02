package cairo

import (
	"encoding/json"
	"fmt"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/vm"
)

type CallRequest struct {
	ContractAddr *felt.Felt  `json:"contract_addr"`
	Selector     *felt.Felt  `json:"selector"`
	Calldata     []felt.Felt `json:"calldata"`
	BlockID      BlockID     `json:"block_id"`
}

type CallResponse struct {
	ReturnData []*felt.Felt   `json:"return_data"`
	Err        *jsonrpc.Error `json:"err"`
}

type TxRequest struct {
	Tx              *rpc.BroadcastedTransaction `json:"tx"`
	GasPriceWEI     *felt.Felt                  `json:"gas_price_wei"`
	GasPriceSTRK    *felt.Felt                  `json:"gas_price_strk"`
	LegacyTraceJson bool                        `json:"legacy_trace_json"`
}

////////////////////////////////////////
///		RPC Types 		        ////////
////////////////////////////////////////

/*************************
	Block Types
**************************/

type BlockID struct {
	Pending bool       `json:"pending"`
	Latest  bool       `json:"latest"`
	Hash    *felt.Felt `json:"hash"`
	Number  uint64     `json:"number"`
}

type BlockStatus uint8

const (
	BlockPending BlockStatus = iota
	BlockAcceptedL2
	BlockAcceptedL1
	BlockRejected
)

func (s BlockStatus) MarshalText() ([]byte, error) {
	switch s {
	case BlockPending:
		return []byte("PENDING"), nil
	case BlockAcceptedL2:
		return []byte("ACCEPTED_ON_L2"), nil
	case BlockAcceptedL1:
		return []byte("ACCEPTED_ON_L1"), nil
	case BlockRejected:
		return []byte("REJECTED"), nil
	default:
		return nil, fmt.Errorf("unknown block status %v", s)
	}
}

type ResourcePrice struct {
	InFri *felt.Felt `json:"price_in_fri"`
	InWei *felt.Felt `json:"price_in_wei"`
}
type L1DAMode uint8

const (
	Blob L1DAMode = iota
	Calldata
)

func (l L1DAMode) MarshalText() ([]byte, error) {
	switch l {
	case Blob:
		return []byte("BLOB"), nil
	case Calldata:
		return []byte("CALLDATA"), nil
	default:
		return nil, fmt.Errorf("unknown L1DAMode value = %v", l)
	}
}

// https://github.com/starkware-libs/starknet-specs/blob/a789ccc3432c57777beceaa53a34a7ae2f25fda0/api/starknet_api_openrpc.json#L1072
type BlockHeader struct {
	Hash             *felt.Felt     `json:"block_hash,omitempty"`
	ParentHash       *felt.Felt     `json:"parent_hash"`
	Number           *uint64        `json:"block_number,omitempty"`
	NewRoot          *felt.Felt     `json:"new_root,omitempty"`
	Timestamp        uint64         `json:"timestamp"`
	SequencerAddress *felt.Felt     `json:"sequencer_address,omitempty"`
	L1GasPrice       *ResourcePrice `json:"l1_gas_price"`
	L1DataGasPrice   *ResourcePrice `json:"l1_data_gas_price,omitempty"`
	L1DAMode         *L1DAMode      `json:"l1_da_mode,omitempty"`
	StarknetVersion  string         `json:"starknet_version"`
}

// https://github.com/starkware-libs/starknet-specs/blob/a789ccc3432c57777beceaa53a34a7ae2f25fda0/api/starknet_api_openrpc.json#L1131
type BlockWithTxs struct {
	Status BlockStatus `json:"status,omitempty"`
	BlockHeader
	Transactions []*rpc.Transaction `json:"transactions"`
}

// https://github.com/starkware-libs/starknet-specs/blob/a789ccc3432c57777beceaa53a34a7ae2f25fda0/api/starknet_api_openrpc.json#L1109
type BlockWithTxHashes struct {
	Status BlockStatus `json:"status,omitempty"`
	BlockHeader
	TxnHashes []*felt.Felt `json:"transactions"`
}

/*************************
	StateUpdate Types
**************************/

// https://github.com/starkware-libs/starknet-specs/blob/8016dd08ed7cd220168db16f24c8a6827ab88317/api/starknet_api_openrpc.json#L909
type StateUpdate struct {
	BlockHash *felt.Felt `json:"block_hash,omitempty"`
	NewRoot   *felt.Felt `json:"new_root,omitempty"`
	OldRoot   *felt.Felt `json:"old_root"`
	StateDiff *StateDiff `json:"state_diff"`
}

type StateDiff struct {
	StorageDiffs              []StorageDiff      `json:"storage_diffs"`
	Nonces                    []Nonce            `json:"nonces"`
	DeployedContracts         []DeployedContract `json:"deployed_contracts"`
	DeprecatedDeclaredClasses []*felt.Felt       `json:"deprecated_declared_classes"`
	DeclaredClasses           []DeclaredClass    `json:"declared_classes"`
	ReplacedClasses           []ReplacedClass    `json:"replaced_classes"`
}

type Nonce struct {
	ContractAddress felt.Felt `json:"contract_address"`
	Nonce           felt.Felt `json:"nonce"`
}

type StorageDiff struct {
	Address        felt.Felt `json:"address"`
	StorageEntries []Entry   `json:"storage_entries"`
}

type Entry struct {
	Key   felt.Felt `json:"key"`
	Value felt.Felt `json:"value"`
}

type DeployedContract struct {
	Address   felt.Felt `json:"address"`
	ClassHash felt.Felt `json:"class_hash"`
}

type ReplacedClass struct {
	ContractAddress felt.Felt `json:"contract_address"`
	ClassHash       felt.Felt `json:"class_hash"`
}

type DeclaredClass struct {
	ClassHash         felt.Felt `json:"class_hash"`
	CompiledClassHash felt.Felt `json:"compiled_class_hash"`
}

/*************************
	FeeUnit Types
**************************/

type FeeUnit byte

type FeeEstimate struct {
	GasConsumed     *felt.Felt   `json:"gas_consumed"`
	GasPrice        *felt.Felt   `json:"gas_price"`
	DataGasConsumed *felt.Felt   `json:"data_gas_consumed"`
	DataGasPrice    *felt.Felt   `json:"data_gas_price"`
	OverallFee      *felt.Felt   `json:"overall_fee"`
	Unit            *rpc.FeeUnit `json:"unit,omitempty"`
	// pre 13.1 response
	v0_6Response bool
}

func (f FeeEstimate) MarshalJSON() ([]byte, error) {
	if f.v0_6Response {
		return json.Marshal(struct {
			GasConsumed *felt.Felt   `json:"gas_consumed"`
			GasPrice    *felt.Felt   `json:"gas_price"`
			OverallFee  *felt.Felt   `json:"overall_fee"`
			Unit        *rpc.FeeUnit `json:"unit,omitempty"`
		}{
			GasConsumed: f.GasConsumed,
			GasPrice:    f.GasPrice,
			OverallFee:  f.OverallFee,
			Unit:        f.Unit,
		})
	} else {
		type alias FeeEstimate // avoid infinite recursion
		return json.Marshal(alias(f))
	}
}

/*
************************
	Simulation Types
*************************
*/

type SimulationFlag int

const (
	SkipValidateFlag SimulationFlag = iota + 1
	SkipFeeChargeFlag
)

func (s *SimulationFlag) UnmarshalJSON(bytes []byte) (err error) {
	var raw interface{}
	if err = json.Unmarshal(bytes, &raw); err != nil {
		return
	}

	switch v := raw.(type) {
	case string:
		switch v {
		case "SKIP_VALIDATE":
			*s = SkipValidateFlag
		case "SKIP_FEE_CHARGE":
			*s = SkipFeeChargeFlag
		default:
			err = fmt.Errorf("unknown simulation flag %q", v)
		}
	case float64: // JSON numbers are unmarshaled into float64
		switch int(v) {
		case 1:
			*s = SkipValidateFlag
		case 2:
			*s = SkipFeeChargeFlag
		default:
			err = fmt.Errorf("unknown simulation flag code %d", int(v))
		}
	default:
		err = fmt.Errorf("unexpected type for simulation flag: %T", v)
	}

	return
}

type SimulatedTransaction struct {
	TransactionTrace *vm.TransactionTrace `json:"transaction_trace,omitempty"`
	FeeEstimation    FeeEstimate          `json:"fee_estimation,omitempty"`
}

type TracedBlockTransaction struct {
	TraceRoot       *vm.TransactionTrace `json:"trace_root,omitempty"`
	TransactionHash *felt.Felt           `json:"transaction_hash,omitempty"`
}
