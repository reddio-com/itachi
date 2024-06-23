package starknetrpc

import (
	"fmt"
	"itachi/cairo"
	"itachi/cairo/config"
	"math/big"
	"testing"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/joho/godotenv"
	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/require"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
)

const CONFIG_PATH string = "../../conf/cairo_cfg.toml"
const DEFAULT_BASE int = 4

var (
	// set the environment for the test, default: mock
	testEnv = "mock"

	// testConfigurations are predefined test configurations
	// Mainnet =  1
	// Goerli = 2
	// Goerli2 = 3
	// Integration = 4
	// Sepolia = 5
	// SepoliaIntegration =6
	testConfigurations = map[string]testConfiguration{
		// Requires a Mainnet Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"Mainnet": {
			base: 1,
		},
		// Requires a Goerli Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"Goerli": {
			base: 2,
		},
		// Requires a Goerli2 Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"Goerli2": {
			base: 3, // Update this base URL as needed
		},
		// Requires an Integration Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"Integration": {
			base: 4, // Update this base URL as needed
		},
		// Requires a Sepolia Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"Sepolia": {
			base: 5, // Update this base URL as needed
		},
		// Requires a SepoliaIntegration Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"SepoliaIntegration": {
			base: 6, // Update this base URL as needed
		},
		// Used with a mock as a standard configuration, see `mock_test.go`
		"mock": {},
	}
)

////////////////////////////////////////
////		Type Struct			////////
////////////////////////////////////////

// testConfiguration is a type that is used to configure tests
type testConfiguration struct {
	starknetRpc *StarknetRPC
	base        int
}

////////////////////////////////////////
////		Functions			////////
////////////////////////////////////////

////////////////////////////////////////
////		UtilFunctions		////////
////////////////////////////////////////

// Uint64ToFelt generates a new *felt.Felt from a given uint64 number.
//
// Parameters:
// - num: the uint64 number to convert to a *felt.Felt
// Returns:
// - *felt.Felt: a *felt.Felt
func Uint64ToFelt(num uint64) *felt.Felt {
	return new(felt.Felt).SetUint64(num)
}

// HexToFelt converts a hexadecimal string to a *felt.Felt object.
//
// Parameters:
// - hex: the input hexadecimal string to be converted.
// Returns:
// - *felt.Felt: a *felt.Felt object
// - error: if conversion fails
func HexToFelt(hex string) (*felt.Felt, error) {
	return new(felt.Felt).SetString(hex)
}

// HexArrToFelt converts an array of hexadecimal strings to an array of felt objects.
//
// The function iterates over each element in the hexArr array and calls the HexToFelt function to convert each hexadecimal value to a felt object.
// If any error occurs during the conversion, the function will return nil and the corresponding error.
// Otherwise, it appends the converted felt object to the feltArr array.
// Finally, the function returns the feltArr array containing all the converted felt objects.
//
// Parameters:
// - hexArr: an array of strings representing hexadecimal values
// Returns:
// - []*felt.Felt: an array of *felt.Felt objects, or nil if there was
// - error: an error if any
func HexArrToFelt(hexArr []string) ([]*felt.Felt, error) {

	feltArr := make([]*felt.Felt, len(hexArr))
	for i, e := range hexArr {
		felt, err := HexToFelt(e)
		if err != nil {
			return nil, err
		}
		feltArr[i] = felt
	}
	return feltArr, nil

}

// FeltToBigInt converts a Felt value to a *big.Int.
//
// Parameters:
// - f: the Felt value to convert
// Returns:
// - *big.Int: the converted value
func FeltToBigInt(f *felt.Felt) *big.Int {
	tmp := f.Bytes()
	return new(big.Int).SetBytes(tmp[:])
}

// BigIntToFelt converts a big integer to a felt.Felt.
//
// Parameters:
// - big: the big integer to convert
// Returns:
// - *felt.Felt: the converted value
func BigIntToFelt(big *big.Int) *felt.Felt {
	return new(felt.Felt).SetBytes(big.Bytes())
}

// FeltArrToBigIntArr converts an array of Felt objects to an array of big.Int objects.
//
// Parameters:
// - f: the array of Felt objects to convert
// Returns:
// - []*big.Int: the array of big.Int objects
func FeltArrToBigIntArr(f []*felt.Felt) []*big.Int {
	var bigArr []*big.Int
	for _, felt := range f {
		bigArr = append(bigArr, FeltToBigInt(felt))
	}
	return bigArr
}

// initStarknetRPC initializes and returns a new StarknetRPC instance.
func initStarknetRPC(base int) (*StarknetRPC, error) {
	startup.InitDefaultKernelConfig()
	poaCfg := poa.DefaultCfg(0)
	crCfg := config.LoadCairoCfg(CONFIG_PATH)
	crCfg.Network = base
	poaTri := poa.NewPoa(poaCfg)
	cairoTri := cairo.NewCairo(crCfg)
	chain := startup.InitDefaultKernel(
		poaTri, cairoTri,
	)
	starknetRpc, err := NewStarknetRPC(chain, crCfg)
	if err != nil {
		return nil, err
	}

	return starknetRpc, nil
}

func beforeEach(t *testing.T) *testConfiguration {
	t.Helper()
	_ = godotenv.Load(fmt.Sprintf(".env.%s", testEnv), ".env")
	testConfig, ok := testConfigurations[testEnv]
	if !ok {
		t.Fatal("env supports mock, testnet, mainnet, devnet, integration")
	}
	base := DEFAULT_BASE
	testConfig.base = base

	starknetRpc, err := initStarknetRPC(testConfig.base)
	if err != nil {
		t.Fatal("connect should succeed, instead:", err)
	}
	testConfig.starknetRpc = starknetRpc

	return &testConfig
}

////////////////////////////////////////
////		TestFunctions		////////
////////////////////////////////////////

func TestGetChainId(t *testing.T) {
	//for _, n := range []utils.Network{utils.Mainnet, utils.Goerli, utils.Goerli2, utils.Integration} {
	for _, n := range []utils.Network{utils.Integration} {
		testConfig := beforeEach(t)
		t.Run(n.String(), func(t *testing.T) {
			t.Log("n.ChainID() = ", n.ChainID())

			cID, jsonErr := testConfig.starknetRpc.GetChainID()
			require.Nil(t, jsonErr)
			t.Log("cID = ", cID)
			//assert
			assert.Equal(t, n.ChainID(), cID)
		})
	}
}

// TODO: Implement the following tests
func TestGetBlockWithTxHashes(t *testing.T) {
	errTests := map[string]rpc.BlockID{
		"latest":  {Latest: true},
		"pending": {Pending: true},
		"hash":    {Hash: new(felt.Felt).SetUint64(1)},
		"number":  {Number: 1},
	}
	testConfig := beforeEach(t)

	//test for error
	for description, id := range errTests {
		t.Run(description, func(t *testing.T) {
			starknetRpc := testConfig.starknetRpc
			block, rpcErr := starknetRpc.GetBlockWithTxHashes(id)
			assert.Nil(t, block)
			assert.Equal(t, jsonrpc.InternalError, rpcErr.Code)
		})
	}
	//test Latest
	t.Run("latest", func(t *testing.T) {
		t.Helper()
		starknetRpc := testConfig.starknetRpc
		id := rpc.BlockID{Latest: true}
		block, rpcErr := starknetRpc.GetBlockWithTxHashes(id)
		assert.Nil(t, rpcErr)
		t.Log("block = ", block)
		assert.NotNil(t, block)
	})
	//test pending

}

func TestClass(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		BlockID                       rpc.BlockID
		ClassHash                     *felt.Felt
		ExpectedProgram               string
		ExpectedEntryPointConstructor core.SierraEntryPoint
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				BlockID:         rpc.BlockID{Pending: true},
				ClassHash:       TransferHexToFelt("0xdeadbeef"),
				ExpectedProgram: "H4sIAAAAAAAA",
			},
		},
		"testnet": {
			// v0 class
			{
				BlockID:         rpc.BlockID{Latest: true},
				ClassHash:       TransferHexToFelt("0x036c7e49a16f8fc760a6fbdf71dde543d98be1fee2eda5daff59a0eeae066ed9"),
				ExpectedProgram: "H4sIAAAAAAAA",
			},
			// v2 classes
			{
				BlockID:                       rpc.BlockID{Latest: true},
				ClassHash:                     TransferHexToFelt("0x00816dd0297efc55dc1e7559020a3a825e81ef734b558f03c83325d4da7e6253"),
				ExpectedProgram:               TransferHexToFelt("0x576402000a0028a9c00a010").String(),
				ExpectedEntryPointConstructor: core.SierraEntryPoint{Index: 34, Selector: TransferHexToFelt("0x28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194")},
			},
			{
				BlockID:                       rpc.BlockID{Latest: true},
				ClassHash:                     TransferHexToFelt("0x01f372292df22d28f2d4c5798734421afe9596e6a566b8bc9b7b50e26521b855"),
				ExpectedProgram:               TransferHexToFelt("0xe70d09071117174f17170d4fe60d09071117").String(),
				ExpectedEntryPointConstructor: core.SierraEntryPoint{Index: 2, Selector: TransferHexToFelt("0x28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194")},
			},
		},
		"mainnet": {
			// v2 class
			{
				BlockID:                       rpc.BlockID{Latest: true},
				ClassHash:                     TransferHexToFelt("0x029927c8af6bccf3f6fda035981e765a7bdbf18a2dc0d630494f8758aa908e2b"),
				ExpectedProgram:               TransferHexToFelt("0x9fa00900700e00712e12500712e").String(),
				ExpectedEntryPointConstructor: core.SierraEntryPoint{Index: 32, Selector: TransferHexToFelt("0x28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194")},
			},
		},
	}[testEnv]

	for _, test := range testSet {
		require := require.New(t)
		resp, err := testConfig.starknetRpc.GetClass(test.BlockID, *test.ClassHash)
		if err != nil {
			require.Nil(resp)
			require.Equal(jsonrpc.InternalError, err.Code)
		}

		t.Log("resp = ", resp)
	}
}

// TestHexToFelt generates a felt.Felt from a hexadecimal string.
//
// Parameters:
// - t: the testing.TB object for test logging and reporting
// - hex: the hexadecimal string to convert to a felt.Felt
// Returns:
// - *felt.Felt: the generated felt.Felt object
func TransferHexToFelt(hex string) *felt.Felt {
	f, _ := HexToFelt(hex)
	return f
}

// TestHexArrToFelt generates a slice of *felt.Felt from a slice of strings representing hexadecimal values.
//
// Parameters:
// - t: A testing.TB interface used for test logging and error reporting
// - hexArr: A slice of strings representing hexadecimal values
// Returns:
// - []*felt.Felt: a slice of *felt.Felt
func TransferHexArrToFelt(hexArr []string) []*felt.Felt {
	feltArr, _ := HexArrToFelt(hexArr)
	return feltArr
}
