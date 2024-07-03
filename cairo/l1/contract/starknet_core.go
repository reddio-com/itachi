// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StarknetCoreMetaData contains all meta data concerning the StarknetCore contract.
var StarknetCoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldConfigHash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newConfigHash\",\"type\":\"uint256\"}],\"name\":\"ConfigHashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"}],\"name\":\"ConsumedMessageToL1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ConsumedMessageToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"}],\"name\":\"LogMessageToL1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogMessageToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptedGovernor\",\"type\":\"address\"}],\"name\":\"LogNewGovernorAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominatedGovernor\",\"type\":\"address\"}],\"name\":\"LogNominatedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogNominationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedGovernor\",\"type\":\"address\"}],\"name\":\"LogRemovedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"blockNumber\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHash\",\"type\":\"uint256\"}],\"name\":\"LogStateUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MessageToL2Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MessageToL2CancellationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProgramHash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProgramHash\",\"type\":\"uint256\"}],\"name\":\"ProgramHashChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"cancelL1ToL2Message\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"}],\"name\":\"consumeMessageFromL2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxL1MsgFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identify\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"l1ToL2MessageCancellations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1ToL2MessageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"l1ToL2Messages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"l2ToL1Messages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCancellationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"programHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"}],\"name\":\"sendMessageToL2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newConfigHash\",\"type\":\"uint256\"}],\"name\":\"setConfigHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayInSeconds\",\"type\":\"uint256\"}],\"name\":\"setMessageCancellationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProgramHash\",\"type\":\"uint256\"}],\"name\":\"setProgramHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"starknetAcceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"starknetCancelNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"starknetIsGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"starknetNominateNewGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governorForRemoval\",\"type\":\"address\"}],\"name\":\"starknetRemoveGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"startL1ToL2MessageCancellation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateBlockHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateBlockNumber\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"programOutput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"onchainDataHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"onchainDataSize\",\"type\":\"uint256\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"programOutput\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"kzgProof\",\"type\":\"bytes\"}],\"name\":\"updateStateKzgDA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StarknetCoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StarknetCoreMetaData.ABI instead.
var StarknetCoreABI = StarknetCoreMetaData.ABI

// StarknetCore is an auto generated Go binding around an Ethereum contract.
type StarknetCore struct {
	StarknetCoreCaller     // Read-only binding to the contract
	StarknetCoreTransactor // Write-only binding to the contract
	StarknetCoreFilterer   // Log filterer for contract events
}

// StarknetCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StarknetCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarknetCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StarknetCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarknetCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StarknetCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarknetCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StarknetCoreSession struct {
	Contract     *StarknetCore     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StarknetCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StarknetCoreCallerSession struct {
	Contract *StarknetCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StarknetCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StarknetCoreTransactorSession struct {
	Contract     *StarknetCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StarknetCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StarknetCoreRaw struct {
	Contract *StarknetCore // Generic contract binding to access the raw methods on
}

// StarknetCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StarknetCoreCallerRaw struct {
	Contract *StarknetCoreCaller // Generic read-only contract binding to access the raw methods on
}

// StarknetCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StarknetCoreTransactorRaw struct {
	Contract *StarknetCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStarknetCore creates a new instance of StarknetCore, bound to a specific deployed contract.
func NewStarknetCore(address common.Address, backend bind.ContractBackend) (*StarknetCore, error) {
	contract, err := bindStarknetCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StarknetCore{StarknetCoreCaller: StarknetCoreCaller{contract: contract}, StarknetCoreTransactor: StarknetCoreTransactor{contract: contract}, StarknetCoreFilterer: StarknetCoreFilterer{contract: contract}}, nil
}

// NewStarknetCoreCaller creates a new read-only instance of StarknetCore, bound to a specific deployed contract.
func NewStarknetCoreCaller(address common.Address, caller bind.ContractCaller) (*StarknetCoreCaller, error) {
	contract, err := bindStarknetCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreCaller{contract: contract}, nil
}

// NewStarknetCoreTransactor creates a new write-only instance of StarknetCore, bound to a specific deployed contract.
func NewStarknetCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StarknetCoreTransactor, error) {
	contract, err := bindStarknetCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreTransactor{contract: contract}, nil
}

// NewStarknetCoreFilterer creates a new log filterer instance of StarknetCore, bound to a specific deployed contract.
func NewStarknetCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StarknetCoreFilterer, error) {
	contract, err := bindStarknetCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreFilterer{contract: contract}, nil
}

// bindStarknetCore binds a generic wrapper to an already deployed contract.
func bindStarknetCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StarknetCoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarknetCore *StarknetCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarknetCore.Contract.StarknetCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarknetCore *StarknetCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarknetCore *StarknetCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarknetCore *StarknetCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarknetCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarknetCore *StarknetCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarknetCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarknetCore *StarknetCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarknetCore.Contract.contract.Transact(opts, method, params...)
}

// ConfigHash is a free data retrieval call binding the contract method 0xe1f1176d.
//
// Solidity: function configHash() view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) ConfigHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "configHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfigHash is a free data retrieval call binding the contract method 0xe1f1176d.
//
// Solidity: function configHash() view returns(uint256)
func (_StarknetCore *StarknetCoreSession) ConfigHash() (*big.Int, error) {
	return _StarknetCore.Contract.ConfigHash(&_StarknetCore.CallOpts)
}

// ConfigHash is a free data retrieval call binding the contract method 0xe1f1176d.
//
// Solidity: function configHash() view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) ConfigHash() (*big.Int, error) {
	return _StarknetCore.Contract.ConfigHash(&_StarknetCore.CallOpts)
}

// GetMaxL1MsgFee is a free data retrieval call binding the contract method 0x54eccba4.
//
// Solidity: function getMaxL1MsgFee() pure returns(uint256)
func (_StarknetCore *StarknetCoreCaller) GetMaxL1MsgFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "getMaxL1MsgFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxL1MsgFee is a free data retrieval call binding the contract method 0x54eccba4.
//
// Solidity: function getMaxL1MsgFee() pure returns(uint256)
func (_StarknetCore *StarknetCoreSession) GetMaxL1MsgFee() (*big.Int, error) {
	return _StarknetCore.Contract.GetMaxL1MsgFee(&_StarknetCore.CallOpts)
}

// GetMaxL1MsgFee is a free data retrieval call binding the contract method 0x54eccba4.
//
// Solidity: function getMaxL1MsgFee() pure returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) GetMaxL1MsgFee() (*big.Int, error) {
	return _StarknetCore.Contract.GetMaxL1MsgFee(&_StarknetCore.CallOpts)
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() pure returns(string)
func (_StarknetCore *StarknetCoreCaller) Identify(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "identify")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() pure returns(string)
func (_StarknetCore *StarknetCoreSession) Identify() (string, error) {
	return _StarknetCore.Contract.Identify(&_StarknetCore.CallOpts)
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() pure returns(string)
func (_StarknetCore *StarknetCoreCallerSession) Identify() (string, error) {
	return _StarknetCore.Contract.Identify(&_StarknetCore.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_StarknetCore *StarknetCoreCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_StarknetCore *StarknetCoreSession) IsFinalized() (bool, error) {
	return _StarknetCore.Contract.IsFinalized(&_StarknetCore.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_StarknetCore *StarknetCoreCallerSession) IsFinalized() (bool, error) {
	return _StarknetCore.Contract.IsFinalized(&_StarknetCore.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_StarknetCore *StarknetCoreCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_StarknetCore *StarknetCoreSession) IsFrozen() (bool, error) {
	return _StarknetCore.Contract.IsFrozen(&_StarknetCore.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_StarknetCore *StarknetCoreCallerSession) IsFrozen() (bool, error) {
	return _StarknetCore.Contract.IsFrozen(&_StarknetCore.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address user) view returns(bool)
func (_StarknetCore *StarknetCoreCaller) IsOperator(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "isOperator", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address user) view returns(bool)
func (_StarknetCore *StarknetCoreSession) IsOperator(user common.Address) (bool, error) {
	return _StarknetCore.Contract.IsOperator(&_StarknetCore.CallOpts, user)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address user) view returns(bool)
func (_StarknetCore *StarknetCoreCallerSession) IsOperator(user common.Address) (bool, error) {
	return _StarknetCore.Contract.IsOperator(&_StarknetCore.CallOpts, user)
}

// L1ToL2MessageCancellations is a free data retrieval call binding the contract method 0x9be446bf.
//
// Solidity: function l1ToL2MessageCancellations(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) L1ToL2MessageCancellations(opts *bind.CallOpts, msgHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "l1ToL2MessageCancellations", msgHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1ToL2MessageCancellations is a free data retrieval call binding the contract method 0x9be446bf.
//
// Solidity: function l1ToL2MessageCancellations(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreSession) L1ToL2MessageCancellations(msgHash [32]byte) (*big.Int, error) {
	return _StarknetCore.Contract.L1ToL2MessageCancellations(&_StarknetCore.CallOpts, msgHash)
}

// L1ToL2MessageCancellations is a free data retrieval call binding the contract method 0x9be446bf.
//
// Solidity: function l1ToL2MessageCancellations(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) L1ToL2MessageCancellations(msgHash [32]byte) (*big.Int, error) {
	return _StarknetCore.Contract.L1ToL2MessageCancellations(&_StarknetCore.CallOpts, msgHash)
}

// L1ToL2MessageNonce is a free data retrieval call binding the contract method 0x018cccdf.
//
// Solidity: function l1ToL2MessageNonce() view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) L1ToL2MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "l1ToL2MessageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1ToL2MessageNonce is a free data retrieval call binding the contract method 0x018cccdf.
//
// Solidity: function l1ToL2MessageNonce() view returns(uint256)
func (_StarknetCore *StarknetCoreSession) L1ToL2MessageNonce() (*big.Int, error) {
	return _StarknetCore.Contract.L1ToL2MessageNonce(&_StarknetCore.CallOpts)
}

// L1ToL2MessageNonce is a free data retrieval call binding the contract method 0x018cccdf.
//
// Solidity: function l1ToL2MessageNonce() view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) L1ToL2MessageNonce() (*big.Int, error) {
	return _StarknetCore.Contract.L1ToL2MessageNonce(&_StarknetCore.CallOpts)
}

// L1ToL2Messages is a free data retrieval call binding the contract method 0x77c7d7a9.
//
// Solidity: function l1ToL2Messages(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) L1ToL2Messages(opts *bind.CallOpts, msgHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "l1ToL2Messages", msgHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1ToL2Messages is a free data retrieval call binding the contract method 0x77c7d7a9.
//
// Solidity: function l1ToL2Messages(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreSession) L1ToL2Messages(msgHash [32]byte) (*big.Int, error) {
	return _StarknetCore.Contract.L1ToL2Messages(&_StarknetCore.CallOpts, msgHash)
}

// L1ToL2Messages is a free data retrieval call binding the contract method 0x77c7d7a9.
//
// Solidity: function l1ToL2Messages(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) L1ToL2Messages(msgHash [32]byte) (*big.Int, error) {
	return _StarknetCore.Contract.L1ToL2Messages(&_StarknetCore.CallOpts, msgHash)
}

// L2ToL1Messages is a free data retrieval call binding the contract method 0xa46efaf3.
//
// Solidity: function l2ToL1Messages(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) L2ToL1Messages(opts *bind.CallOpts, msgHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "l2ToL1Messages", msgHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Messages is a free data retrieval call binding the contract method 0xa46efaf3.
//
// Solidity: function l2ToL1Messages(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreSession) L2ToL1Messages(msgHash [32]byte) (*big.Int, error) {
	return _StarknetCore.Contract.L2ToL1Messages(&_StarknetCore.CallOpts, msgHash)
}

// L2ToL1Messages is a free data retrieval call binding the contract method 0xa46efaf3.
//
// Solidity: function l2ToL1Messages(bytes32 msgHash) view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) L2ToL1Messages(msgHash [32]byte) (*big.Int, error) {
	return _StarknetCore.Contract.L2ToL1Messages(&_StarknetCore.CallOpts, msgHash)
}

// MessageCancellationDelay is a free data retrieval call binding the contract method 0x8303bd8a.
//
// Solidity: function messageCancellationDelay() view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) MessageCancellationDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "messageCancellationDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCancellationDelay is a free data retrieval call binding the contract method 0x8303bd8a.
//
// Solidity: function messageCancellationDelay() view returns(uint256)
func (_StarknetCore *StarknetCoreSession) MessageCancellationDelay() (*big.Int, error) {
	return _StarknetCore.Contract.MessageCancellationDelay(&_StarknetCore.CallOpts)
}

// MessageCancellationDelay is a free data retrieval call binding the contract method 0x8303bd8a.
//
// Solidity: function messageCancellationDelay() view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) MessageCancellationDelay() (*big.Int, error) {
	return _StarknetCore.Contract.MessageCancellationDelay(&_StarknetCore.CallOpts)
}

// ProgramHash is a free data retrieval call binding the contract method 0x8a9bf090.
//
// Solidity: function programHash() view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) ProgramHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "programHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProgramHash is a free data retrieval call binding the contract method 0x8a9bf090.
//
// Solidity: function programHash() view returns(uint256)
func (_StarknetCore *StarknetCoreSession) ProgramHash() (*big.Int, error) {
	return _StarknetCore.Contract.ProgramHash(&_StarknetCore.CallOpts)
}

// ProgramHash is a free data retrieval call binding the contract method 0x8a9bf090.
//
// Solidity: function programHash() view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) ProgramHash() (*big.Int, error) {
	return _StarknetCore.Contract.ProgramHash(&_StarknetCore.CallOpts)
}

// StarknetIsGovernor is a free data retrieval call binding the contract method 0x01a01590.
//
// Solidity: function starknetIsGovernor(address user) view returns(bool)
func (_StarknetCore *StarknetCoreCaller) StarknetIsGovernor(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "starknetIsGovernor", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StarknetIsGovernor is a free data retrieval call binding the contract method 0x01a01590.
//
// Solidity: function starknetIsGovernor(address user) view returns(bool)
func (_StarknetCore *StarknetCoreSession) StarknetIsGovernor(user common.Address) (bool, error) {
	return _StarknetCore.Contract.StarknetIsGovernor(&_StarknetCore.CallOpts, user)
}

// StarknetIsGovernor is a free data retrieval call binding the contract method 0x01a01590.
//
// Solidity: function starknetIsGovernor(address user) view returns(bool)
func (_StarknetCore *StarknetCoreCallerSession) StarknetIsGovernor(user common.Address) (bool, error) {
	return _StarknetCore.Contract.StarknetIsGovernor(&_StarknetCore.CallOpts, user)
}

// StateBlockHash is a free data retrieval call binding the contract method 0x382d83e3.
//
// Solidity: function stateBlockHash() view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) StateBlockHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "stateBlockHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateBlockHash is a free data retrieval call binding the contract method 0x382d83e3.
//
// Solidity: function stateBlockHash() view returns(uint256)
func (_StarknetCore *StarknetCoreSession) StateBlockHash() (*big.Int, error) {
	return _StarknetCore.Contract.StateBlockHash(&_StarknetCore.CallOpts)
}

// StateBlockHash is a free data retrieval call binding the contract method 0x382d83e3.
//
// Solidity: function stateBlockHash() view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) StateBlockHash() (*big.Int, error) {
	return _StarknetCore.Contract.StateBlockHash(&_StarknetCore.CallOpts)
}

// StateBlockNumber is a free data retrieval call binding the contract method 0x35befa5d.
//
// Solidity: function stateBlockNumber() view returns(int256)
func (_StarknetCore *StarknetCoreCaller) StateBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "stateBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateBlockNumber is a free data retrieval call binding the contract method 0x35befa5d.
//
// Solidity: function stateBlockNumber() view returns(int256)
func (_StarknetCore *StarknetCoreSession) StateBlockNumber() (*big.Int, error) {
	return _StarknetCore.Contract.StateBlockNumber(&_StarknetCore.CallOpts)
}

// StateBlockNumber is a free data retrieval call binding the contract method 0x35befa5d.
//
// Solidity: function stateBlockNumber() view returns(int256)
func (_StarknetCore *StarknetCoreCallerSession) StateBlockNumber() (*big.Int, error) {
	return _StarknetCore.Contract.StateBlockNumber(&_StarknetCore.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(uint256)
func (_StarknetCore *StarknetCoreCaller) StateRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarknetCore.contract.Call(opts, &out, "stateRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(uint256)
func (_StarknetCore *StarknetCoreSession) StateRoot() (*big.Int, error) {
	return _StarknetCore.Contract.StateRoot(&_StarknetCore.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(uint256)
func (_StarknetCore *StarknetCoreCallerSession) StateRoot() (*big.Int, error) {
	return _StarknetCore.Contract.StateRoot(&_StarknetCore.CallOpts)
}

// CancelL1ToL2Message is a paid mutator transaction binding the contract method 0x6170ff1b.
//
// Solidity: function cancelL1ToL2Message(uint256 toAddress, uint256 selector, uint256[] payload, uint256 nonce) returns(bytes32)
func (_StarknetCore *StarknetCoreTransactor) CancelL1ToL2Message(opts *bind.TransactOpts, toAddress *big.Int, selector *big.Int, payload []*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "cancelL1ToL2Message", toAddress, selector, payload, nonce)
}

// CancelL1ToL2Message is a paid mutator transaction binding the contract method 0x6170ff1b.
//
// Solidity: function cancelL1ToL2Message(uint256 toAddress, uint256 selector, uint256[] payload, uint256 nonce) returns(bytes32)
func (_StarknetCore *StarknetCoreSession) CancelL1ToL2Message(toAddress *big.Int, selector *big.Int, payload []*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.CancelL1ToL2Message(&_StarknetCore.TransactOpts, toAddress, selector, payload, nonce)
}

// CancelL1ToL2Message is a paid mutator transaction binding the contract method 0x6170ff1b.
//
// Solidity: function cancelL1ToL2Message(uint256 toAddress, uint256 selector, uint256[] payload, uint256 nonce) returns(bytes32)
func (_StarknetCore *StarknetCoreTransactorSession) CancelL1ToL2Message(toAddress *big.Int, selector *big.Int, payload []*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.CancelL1ToL2Message(&_StarknetCore.TransactOpts, toAddress, selector, payload, nonce)
}

// ConsumeMessageFromL2 is a paid mutator transaction binding the contract method 0x2c9dd5c0.
//
// Solidity: function consumeMessageFromL2(uint256 fromAddress, uint256[] payload) returns(bytes32)
func (_StarknetCore *StarknetCoreTransactor) ConsumeMessageFromL2(opts *bind.TransactOpts, fromAddress *big.Int, payload []*big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "consumeMessageFromL2", fromAddress, payload)
}

// ConsumeMessageFromL2 is a paid mutator transaction binding the contract method 0x2c9dd5c0.
//
// Solidity: function consumeMessageFromL2(uint256 fromAddress, uint256[] payload) returns(bytes32)
func (_StarknetCore *StarknetCoreSession) ConsumeMessageFromL2(fromAddress *big.Int, payload []*big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.ConsumeMessageFromL2(&_StarknetCore.TransactOpts, fromAddress, payload)
}

// ConsumeMessageFromL2 is a paid mutator transaction binding the contract method 0x2c9dd5c0.
//
// Solidity: function consumeMessageFromL2(uint256 fromAddress, uint256[] payload) returns(bytes32)
func (_StarknetCore *StarknetCoreTransactorSession) ConsumeMessageFromL2(fromAddress *big.Int, payload []*big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.ConsumeMessageFromL2(&_StarknetCore.TransactOpts, fromAddress, payload)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_StarknetCore *StarknetCoreTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_StarknetCore *StarknetCoreSession) Finalize() (*types.Transaction, error) {
	return _StarknetCore.Contract.Finalize(&_StarknetCore.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_StarknetCore *StarknetCoreTransactorSession) Finalize() (*types.Transaction, error) {
	return _StarknetCore.Contract.Finalize(&_StarknetCore.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_StarknetCore *StarknetCoreTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_StarknetCore *StarknetCoreSession) Initialize(data []byte) (*types.Transaction, error) {
	return _StarknetCore.Contract.Initialize(&_StarknetCore.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_StarknetCore *StarknetCoreTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _StarknetCore.Contract.Initialize(&_StarknetCore.TransactOpts, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_StarknetCore *StarknetCoreTransactor) RegisterOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "registerOperator", newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_StarknetCore *StarknetCoreSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.RegisterOperator(&_StarknetCore.TransactOpts, newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_StarknetCore *StarknetCoreTransactorSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.RegisterOperator(&_StarknetCore.TransactOpts, newOperator)
}

// SendMessageToL2 is a paid mutator transaction binding the contract method 0x3e3aa6c5.
//
// Solidity: function sendMessageToL2(uint256 toAddress, uint256 selector, uint256[] payload) payable returns(bytes32, uint256)
func (_StarknetCore *StarknetCoreTransactor) SendMessageToL2(opts *bind.TransactOpts, toAddress *big.Int, selector *big.Int, payload []*big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "sendMessageToL2", toAddress, selector, payload)
}

// SendMessageToL2 is a paid mutator transaction binding the contract method 0x3e3aa6c5.
//
// Solidity: function sendMessageToL2(uint256 toAddress, uint256 selector, uint256[] payload) payable returns(bytes32, uint256)
func (_StarknetCore *StarknetCoreSession) SendMessageToL2(toAddress *big.Int, selector *big.Int, payload []*big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SendMessageToL2(&_StarknetCore.TransactOpts, toAddress, selector, payload)
}

// SendMessageToL2 is a paid mutator transaction binding the contract method 0x3e3aa6c5.
//
// Solidity: function sendMessageToL2(uint256 toAddress, uint256 selector, uint256[] payload) payable returns(bytes32, uint256)
func (_StarknetCore *StarknetCoreTransactorSession) SendMessageToL2(toAddress *big.Int, selector *big.Int, payload []*big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SendMessageToL2(&_StarknetCore.TransactOpts, toAddress, selector, payload)
}

// SetConfigHash is a paid mutator transaction binding the contract method 0x3d07b336.
//
// Solidity: function setConfigHash(uint256 newConfigHash) returns()
func (_StarknetCore *StarknetCoreTransactor) SetConfigHash(opts *bind.TransactOpts, newConfigHash *big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "setConfigHash", newConfigHash)
}

// SetConfigHash is a paid mutator transaction binding the contract method 0x3d07b336.
//
// Solidity: function setConfigHash(uint256 newConfigHash) returns()
func (_StarknetCore *StarknetCoreSession) SetConfigHash(newConfigHash *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SetConfigHash(&_StarknetCore.TransactOpts, newConfigHash)
}

// SetConfigHash is a paid mutator transaction binding the contract method 0x3d07b336.
//
// Solidity: function setConfigHash(uint256 newConfigHash) returns()
func (_StarknetCore *StarknetCoreTransactorSession) SetConfigHash(newConfigHash *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SetConfigHash(&_StarknetCore.TransactOpts, newConfigHash)
}

// SetMessageCancellationDelay is a paid mutator transaction binding the contract method 0xc99d397f.
//
// Solidity: function setMessageCancellationDelay(uint256 delayInSeconds) returns()
func (_StarknetCore *StarknetCoreTransactor) SetMessageCancellationDelay(opts *bind.TransactOpts, delayInSeconds *big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "setMessageCancellationDelay", delayInSeconds)
}

// SetMessageCancellationDelay is a paid mutator transaction binding the contract method 0xc99d397f.
//
// Solidity: function setMessageCancellationDelay(uint256 delayInSeconds) returns()
func (_StarknetCore *StarknetCoreSession) SetMessageCancellationDelay(delayInSeconds *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SetMessageCancellationDelay(&_StarknetCore.TransactOpts, delayInSeconds)
}

// SetMessageCancellationDelay is a paid mutator transaction binding the contract method 0xc99d397f.
//
// Solidity: function setMessageCancellationDelay(uint256 delayInSeconds) returns()
func (_StarknetCore *StarknetCoreTransactorSession) SetMessageCancellationDelay(delayInSeconds *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SetMessageCancellationDelay(&_StarknetCore.TransactOpts, delayInSeconds)
}

// SetProgramHash is a paid mutator transaction binding the contract method 0xe87e7332.
//
// Solidity: function setProgramHash(uint256 newProgramHash) returns()
func (_StarknetCore *StarknetCoreTransactor) SetProgramHash(opts *bind.TransactOpts, newProgramHash *big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "setProgramHash", newProgramHash)
}

// SetProgramHash is a paid mutator transaction binding the contract method 0xe87e7332.
//
// Solidity: function setProgramHash(uint256 newProgramHash) returns()
func (_StarknetCore *StarknetCoreSession) SetProgramHash(newProgramHash *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SetProgramHash(&_StarknetCore.TransactOpts, newProgramHash)
}

// SetProgramHash is a paid mutator transaction binding the contract method 0xe87e7332.
//
// Solidity: function setProgramHash(uint256 newProgramHash) returns()
func (_StarknetCore *StarknetCoreTransactorSession) SetProgramHash(newProgramHash *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.SetProgramHash(&_StarknetCore.TransactOpts, newProgramHash)
}

// StarknetAcceptGovernance is a paid mutator transaction binding the contract method 0x946be3ed.
//
// Solidity: function starknetAcceptGovernance() returns()
func (_StarknetCore *StarknetCoreTransactor) StarknetAcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "starknetAcceptGovernance")
}

// StarknetAcceptGovernance is a paid mutator transaction binding the contract method 0x946be3ed.
//
// Solidity: function starknetAcceptGovernance() returns()
func (_StarknetCore *StarknetCoreSession) StarknetAcceptGovernance() (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetAcceptGovernance(&_StarknetCore.TransactOpts)
}

// StarknetAcceptGovernance is a paid mutator transaction binding the contract method 0x946be3ed.
//
// Solidity: function starknetAcceptGovernance() returns()
func (_StarknetCore *StarknetCoreTransactorSession) StarknetAcceptGovernance() (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetAcceptGovernance(&_StarknetCore.TransactOpts)
}

// StarknetCancelNomination is a paid mutator transaction binding the contract method 0xe37fec25.
//
// Solidity: function starknetCancelNomination() returns()
func (_StarknetCore *StarknetCoreTransactor) StarknetCancelNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "starknetCancelNomination")
}

// StarknetCancelNomination is a paid mutator transaction binding the contract method 0xe37fec25.
//
// Solidity: function starknetCancelNomination() returns()
func (_StarknetCore *StarknetCoreSession) StarknetCancelNomination() (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetCancelNomination(&_StarknetCore.TransactOpts)
}

// StarknetCancelNomination is a paid mutator transaction binding the contract method 0xe37fec25.
//
// Solidity: function starknetCancelNomination() returns()
func (_StarknetCore *StarknetCoreTransactorSession) StarknetCancelNomination() (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetCancelNomination(&_StarknetCore.TransactOpts)
}

// StarknetNominateNewGovernor is a paid mutator transaction binding the contract method 0x91a66a26.
//
// Solidity: function starknetNominateNewGovernor(address newGovernor) returns()
func (_StarknetCore *StarknetCoreTransactor) StarknetNominateNewGovernor(opts *bind.TransactOpts, newGovernor common.Address) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "starknetNominateNewGovernor", newGovernor)
}

// StarknetNominateNewGovernor is a paid mutator transaction binding the contract method 0x91a66a26.
//
// Solidity: function starknetNominateNewGovernor(address newGovernor) returns()
func (_StarknetCore *StarknetCoreSession) StarknetNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetNominateNewGovernor(&_StarknetCore.TransactOpts, newGovernor)
}

// StarknetNominateNewGovernor is a paid mutator transaction binding the contract method 0x91a66a26.
//
// Solidity: function starknetNominateNewGovernor(address newGovernor) returns()
func (_StarknetCore *StarknetCoreTransactorSession) StarknetNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetNominateNewGovernor(&_StarknetCore.TransactOpts, newGovernor)
}

// StarknetRemoveGovernor is a paid mutator transaction binding the contract method 0x84f921cd.
//
// Solidity: function starknetRemoveGovernor(address governorForRemoval) returns()
func (_StarknetCore *StarknetCoreTransactor) StarknetRemoveGovernor(opts *bind.TransactOpts, governorForRemoval common.Address) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "starknetRemoveGovernor", governorForRemoval)
}

// StarknetRemoveGovernor is a paid mutator transaction binding the contract method 0x84f921cd.
//
// Solidity: function starknetRemoveGovernor(address governorForRemoval) returns()
func (_StarknetCore *StarknetCoreSession) StarknetRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetRemoveGovernor(&_StarknetCore.TransactOpts, governorForRemoval)
}

// StarknetRemoveGovernor is a paid mutator transaction binding the contract method 0x84f921cd.
//
// Solidity: function starknetRemoveGovernor(address governorForRemoval) returns()
func (_StarknetCore *StarknetCoreTransactorSession) StarknetRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.StarknetRemoveGovernor(&_StarknetCore.TransactOpts, governorForRemoval)
}

// StartL1ToL2MessageCancellation is a paid mutator transaction binding the contract method 0x7a98660b.
//
// Solidity: function startL1ToL2MessageCancellation(uint256 toAddress, uint256 selector, uint256[] payload, uint256 nonce) returns(bytes32)
func (_StarknetCore *StarknetCoreTransactor) StartL1ToL2MessageCancellation(opts *bind.TransactOpts, toAddress *big.Int, selector *big.Int, payload []*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "startL1ToL2MessageCancellation", toAddress, selector, payload, nonce)
}

// StartL1ToL2MessageCancellation is a paid mutator transaction binding the contract method 0x7a98660b.
//
// Solidity: function startL1ToL2MessageCancellation(uint256 toAddress, uint256 selector, uint256[] payload, uint256 nonce) returns(bytes32)
func (_StarknetCore *StarknetCoreSession) StartL1ToL2MessageCancellation(toAddress *big.Int, selector *big.Int, payload []*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.StartL1ToL2MessageCancellation(&_StarknetCore.TransactOpts, toAddress, selector, payload, nonce)
}

// StartL1ToL2MessageCancellation is a paid mutator transaction binding the contract method 0x7a98660b.
//
// Solidity: function startL1ToL2MessageCancellation(uint256 toAddress, uint256 selector, uint256[] payload, uint256 nonce) returns(bytes32)
func (_StarknetCore *StarknetCoreTransactorSession) StartL1ToL2MessageCancellation(toAddress *big.Int, selector *big.Int, payload []*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.StartL1ToL2MessageCancellation(&_StarknetCore.TransactOpts, toAddress, selector, payload, nonce)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_StarknetCore *StarknetCoreTransactor) UnregisterOperator(opts *bind.TransactOpts, removedOperator common.Address) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "unregisterOperator", removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_StarknetCore *StarknetCoreSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.UnregisterOperator(&_StarknetCore.TransactOpts, removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_StarknetCore *StarknetCoreTransactorSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _StarknetCore.Contract.UnregisterOperator(&_StarknetCore.TransactOpts, removedOperator)
}

// UpdateState is a paid mutator transaction binding the contract method 0x77552641.
//
// Solidity: function updateState(uint256[] programOutput, uint256 onchainDataHash, uint256 onchainDataSize) returns()
func (_StarknetCore *StarknetCoreTransactor) UpdateState(opts *bind.TransactOpts, programOutput []*big.Int, onchainDataHash *big.Int, onchainDataSize *big.Int) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "updateState", programOutput, onchainDataHash, onchainDataSize)
}

// UpdateState is a paid mutator transaction binding the contract method 0x77552641.
//
// Solidity: function updateState(uint256[] programOutput, uint256 onchainDataHash, uint256 onchainDataSize) returns()
func (_StarknetCore *StarknetCoreSession) UpdateState(programOutput []*big.Int, onchainDataHash *big.Int, onchainDataSize *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.UpdateState(&_StarknetCore.TransactOpts, programOutput, onchainDataHash, onchainDataSize)
}

// UpdateState is a paid mutator transaction binding the contract method 0x77552641.
//
// Solidity: function updateState(uint256[] programOutput, uint256 onchainDataHash, uint256 onchainDataSize) returns()
func (_StarknetCore *StarknetCoreTransactorSession) UpdateState(programOutput []*big.Int, onchainDataHash *big.Int, onchainDataSize *big.Int) (*types.Transaction, error) {
	return _StarknetCore.Contract.UpdateState(&_StarknetCore.TransactOpts, programOutput, onchainDataHash, onchainDataSize)
}

// UpdateStateKzgDA is a paid mutator transaction binding the contract method 0xb72d42a1.
//
// Solidity: function updateStateKzgDA(uint256[] programOutput, bytes kzgProof) returns()
func (_StarknetCore *StarknetCoreTransactor) UpdateStateKzgDA(opts *bind.TransactOpts, programOutput []*big.Int, kzgProof []byte) (*types.Transaction, error) {
	return _StarknetCore.contract.Transact(opts, "updateStateKzgDA", programOutput, kzgProof)
}

// UpdateStateKzgDA is a paid mutator transaction binding the contract method 0xb72d42a1.
//
// Solidity: function updateStateKzgDA(uint256[] programOutput, bytes kzgProof) returns()
func (_StarknetCore *StarknetCoreSession) UpdateStateKzgDA(programOutput []*big.Int, kzgProof []byte) (*types.Transaction, error) {
	return _StarknetCore.Contract.UpdateStateKzgDA(&_StarknetCore.TransactOpts, programOutput, kzgProof)
}

// UpdateStateKzgDA is a paid mutator transaction binding the contract method 0xb72d42a1.
//
// Solidity: function updateStateKzgDA(uint256[] programOutput, bytes kzgProof) returns()
func (_StarknetCore *StarknetCoreTransactorSession) UpdateStateKzgDA(programOutput []*big.Int, kzgProof []byte) (*types.Transaction, error) {
	return _StarknetCore.Contract.UpdateStateKzgDA(&_StarknetCore.TransactOpts, programOutput, kzgProof)
}

// StarknetCoreConfigHashChangedIterator is returned from FilterConfigHashChanged and is used to iterate over the raw logs and unpacked data for ConfigHashChanged events raised by the StarknetCore contract.
type StarknetCoreConfigHashChangedIterator struct {
	Event *StarknetCoreConfigHashChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreConfigHashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreConfigHashChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreConfigHashChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreConfigHashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreConfigHashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreConfigHashChanged represents a ConfigHashChanged event raised by the StarknetCore contract.
type StarknetCoreConfigHashChanged struct {
	ChangedBy     common.Address
	OldConfigHash *big.Int
	NewConfigHash *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfigHashChanged is a free log retrieval operation binding the contract event 0x393c6beb5756a944b2967f15f31ff671e312e945d7a84fd3bdcfd6b408b2dc79.
//
// Solidity: event ConfigHashChanged(address indexed changedBy, uint256 oldConfigHash, uint256 newConfigHash)
func (_StarknetCore *StarknetCoreFilterer) FilterConfigHashChanged(opts *bind.FilterOpts, changedBy []common.Address) (*StarknetCoreConfigHashChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "ConfigHashChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreConfigHashChangedIterator{contract: _StarknetCore.contract, event: "ConfigHashChanged", logs: logs, sub: sub}, nil
}

// WatchConfigHashChanged is a free log subscription operation binding the contract event 0x393c6beb5756a944b2967f15f31ff671e312e945d7a84fd3bdcfd6b408b2dc79.
//
// Solidity: event ConfigHashChanged(address indexed changedBy, uint256 oldConfigHash, uint256 newConfigHash)
func (_StarknetCore *StarknetCoreFilterer) WatchConfigHashChanged(opts *bind.WatchOpts, sink chan<- *StarknetCoreConfigHashChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "ConfigHashChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreConfigHashChanged)
				if err := _StarknetCore.contract.UnpackLog(event, "ConfigHashChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigHashChanged is a log parse operation binding the contract event 0x393c6beb5756a944b2967f15f31ff671e312e945d7a84fd3bdcfd6b408b2dc79.
//
// Solidity: event ConfigHashChanged(address indexed changedBy, uint256 oldConfigHash, uint256 newConfigHash)
func (_StarknetCore *StarknetCoreFilterer) ParseConfigHashChanged(log types.Log) (*StarknetCoreConfigHashChanged, error) {
	event := new(StarknetCoreConfigHashChanged)
	if err := _StarknetCore.contract.UnpackLog(event, "ConfigHashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreConsumedMessageToL1Iterator is returned from FilterConsumedMessageToL1 and is used to iterate over the raw logs and unpacked data for ConsumedMessageToL1 events raised by the StarknetCore contract.
type StarknetCoreConsumedMessageToL1Iterator struct {
	Event *StarknetCoreConsumedMessageToL1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreConsumedMessageToL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreConsumedMessageToL1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreConsumedMessageToL1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreConsumedMessageToL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreConsumedMessageToL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreConsumedMessageToL1 represents a ConsumedMessageToL1 event raised by the StarknetCore contract.
type StarknetCoreConsumedMessageToL1 struct {
	FromAddress *big.Int
	ToAddress   common.Address
	Payload     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsumedMessageToL1 is a free log retrieval operation binding the contract event 0x7a06c571aa77f34d9706c51e5d8122b5595aebeaa34233bfe866f22befb973b1.
//
// Solidity: event ConsumedMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_StarknetCore *StarknetCoreFilterer) FilterConsumedMessageToL1(opts *bind.FilterOpts, fromAddress []*big.Int, toAddress []common.Address) (*StarknetCoreConsumedMessageToL1Iterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "ConsumedMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreConsumedMessageToL1Iterator{contract: _StarknetCore.contract, event: "ConsumedMessageToL1", logs: logs, sub: sub}, nil
}

// WatchConsumedMessageToL1 is a free log subscription operation binding the contract event 0x7a06c571aa77f34d9706c51e5d8122b5595aebeaa34233bfe866f22befb973b1.
//
// Solidity: event ConsumedMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_StarknetCore *StarknetCoreFilterer) WatchConsumedMessageToL1(opts *bind.WatchOpts, sink chan<- *StarknetCoreConsumedMessageToL1, fromAddress []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "ConsumedMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreConsumedMessageToL1)
				if err := _StarknetCore.contract.UnpackLog(event, "ConsumedMessageToL1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsumedMessageToL1 is a log parse operation binding the contract event 0x7a06c571aa77f34d9706c51e5d8122b5595aebeaa34233bfe866f22befb973b1.
//
// Solidity: event ConsumedMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_StarknetCore *StarknetCoreFilterer) ParseConsumedMessageToL1(log types.Log) (*StarknetCoreConsumedMessageToL1, error) {
	event := new(StarknetCoreConsumedMessageToL1)
	if err := _StarknetCore.contract.UnpackLog(event, "ConsumedMessageToL1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreConsumedMessageToL2Iterator is returned from FilterConsumedMessageToL2 and is used to iterate over the raw logs and unpacked data for ConsumedMessageToL2 events raised by the StarknetCore contract.
type StarknetCoreConsumedMessageToL2Iterator struct {
	Event *StarknetCoreConsumedMessageToL2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreConsumedMessageToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreConsumedMessageToL2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreConsumedMessageToL2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreConsumedMessageToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreConsumedMessageToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreConsumedMessageToL2 represents a ConsumedMessageToL2 event raised by the StarknetCore contract.
type StarknetCoreConsumedMessageToL2 struct {
	FromAddress common.Address
	ToAddress   *big.Int
	Selector    *big.Int
	Payload     []*big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsumedMessageToL2 is a free log retrieval operation binding the contract event 0x9592d37825c744e33fa80c469683bbd04d336241bb600b574758efd182abe26a.
//
// Solidity: event ConsumedMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) FilterConsumedMessageToL2(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetCoreConsumedMessageToL2Iterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "ConsumedMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreConsumedMessageToL2Iterator{contract: _StarknetCore.contract, event: "ConsumedMessageToL2", logs: logs, sub: sub}, nil
}

// WatchConsumedMessageToL2 is a free log subscription operation binding the contract event 0x9592d37825c744e33fa80c469683bbd04d336241bb600b574758efd182abe26a.
//
// Solidity: event ConsumedMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) WatchConsumedMessageToL2(opts *bind.WatchOpts, sink chan<- *StarknetCoreConsumedMessageToL2, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "ConsumedMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreConsumedMessageToL2)
				if err := _StarknetCore.contract.UnpackLog(event, "ConsumedMessageToL2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsumedMessageToL2 is a log parse operation binding the contract event 0x9592d37825c744e33fa80c469683bbd04d336241bb600b574758efd182abe26a.
//
// Solidity: event ConsumedMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) ParseConsumedMessageToL2(log types.Log) (*StarknetCoreConsumedMessageToL2, error) {
	event := new(StarknetCoreConsumedMessageToL2)
	if err := _StarknetCore.contract.UnpackLog(event, "ConsumedMessageToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the StarknetCore contract.
type StarknetCoreFinalizedIterator struct {
	Event *StarknetCoreFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreFinalized represents a Finalized event raised by the StarknetCore contract.
type StarknetCoreFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_StarknetCore *StarknetCoreFilterer) FilterFinalized(opts *bind.FilterOpts) (*StarknetCoreFinalizedIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreFinalizedIterator{contract: _StarknetCore.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_StarknetCore *StarknetCoreFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *StarknetCoreFinalized) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreFinalized)
				if err := _StarknetCore.contract.UnpackLog(event, "Finalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalized is a log parse operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_StarknetCore *StarknetCoreFilterer) ParseFinalized(log types.Log) (*StarknetCoreFinalized, error) {
	event := new(StarknetCoreFinalized)
	if err := _StarknetCore.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogMessageToL1Iterator is returned from FilterLogMessageToL1 and is used to iterate over the raw logs and unpacked data for LogMessageToL1 events raised by the StarknetCore contract.
type StarknetCoreLogMessageToL1Iterator struct {
	Event *StarknetCoreLogMessageToL1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogMessageToL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogMessageToL1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogMessageToL1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogMessageToL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogMessageToL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogMessageToL1 represents a LogMessageToL1 event raised by the StarknetCore contract.
type StarknetCoreLogMessageToL1 struct {
	FromAddress *big.Int
	ToAddress   common.Address
	Payload     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogMessageToL1 is a free log retrieval operation binding the contract event 0x4264ac208b5fde633ccdd42e0f12c3d6d443a4f3779bbf886925b94665b63a22.
//
// Solidity: event LogMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_StarknetCore *StarknetCoreFilterer) FilterLogMessageToL1(opts *bind.FilterOpts, fromAddress []*big.Int, toAddress []common.Address) (*StarknetCoreLogMessageToL1Iterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogMessageToL1Iterator{contract: _StarknetCore.contract, event: "LogMessageToL1", logs: logs, sub: sub}, nil
}

// WatchLogMessageToL1 is a free log subscription operation binding the contract event 0x4264ac208b5fde633ccdd42e0f12c3d6d443a4f3779bbf886925b94665b63a22.
//
// Solidity: event LogMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_StarknetCore *StarknetCoreFilterer) WatchLogMessageToL1(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogMessageToL1, fromAddress []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogMessageToL1)
				if err := _StarknetCore.contract.UnpackLog(event, "LogMessageToL1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMessageToL1 is a log parse operation binding the contract event 0x4264ac208b5fde633ccdd42e0f12c3d6d443a4f3779bbf886925b94665b63a22.
//
// Solidity: event LogMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_StarknetCore *StarknetCoreFilterer) ParseLogMessageToL1(log types.Log) (*StarknetCoreLogMessageToL1, error) {
	event := new(StarknetCoreLogMessageToL1)
	if err := _StarknetCore.contract.UnpackLog(event, "LogMessageToL1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogMessageToL2Iterator is returned from FilterLogMessageToL2 and is used to iterate over the raw logs and unpacked data for LogMessageToL2 events raised by the StarknetCore contract.
type StarknetCoreLogMessageToL2Iterator struct {
	Event *StarknetCoreLogMessageToL2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogMessageToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogMessageToL2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogMessageToL2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogMessageToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogMessageToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogMessageToL2 represents a LogMessageToL2 event raised by the StarknetCore contract.
type StarknetCoreLogMessageToL2 struct {
	FromAddress common.Address
	ToAddress   *big.Int
	Selector    *big.Int
	Payload     []*big.Int
	Nonce       *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogMessageToL2 is a free log retrieval operation binding the contract event 0xdb80dd488acf86d17c747445b0eabb5d57c541d3bd7b6b87af987858e5066b2b.
//
// Solidity: event LogMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce, uint256 fee)
func (_StarknetCore *StarknetCoreFilterer) FilterLogMessageToL2(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetCoreLogMessageToL2Iterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogMessageToL2Iterator{contract: _StarknetCore.contract, event: "LogMessageToL2", logs: logs, sub: sub}, nil
}

// WatchLogMessageToL2 is a free log subscription operation binding the contract event 0xdb80dd488acf86d17c747445b0eabb5d57c541d3bd7b6b87af987858e5066b2b.
//
// Solidity: event LogMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce, uint256 fee)
func (_StarknetCore *StarknetCoreFilterer) WatchLogMessageToL2(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogMessageToL2, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogMessageToL2)
				if err := _StarknetCore.contract.UnpackLog(event, "LogMessageToL2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMessageToL2 is a log parse operation binding the contract event 0xdb80dd488acf86d17c747445b0eabb5d57c541d3bd7b6b87af987858e5066b2b.
//
// Solidity: event LogMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce, uint256 fee)
func (_StarknetCore *StarknetCoreFilterer) ParseLogMessageToL2(log types.Log) (*StarknetCoreLogMessageToL2, error) {
	event := new(StarknetCoreLogMessageToL2)
	if err := _StarknetCore.contract.UnpackLog(event, "LogMessageToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogNewGovernorAcceptedIterator is returned from FilterLogNewGovernorAccepted and is used to iterate over the raw logs and unpacked data for LogNewGovernorAccepted events raised by the StarknetCore contract.
type StarknetCoreLogNewGovernorAcceptedIterator struct {
	Event *StarknetCoreLogNewGovernorAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogNewGovernorAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogNewGovernorAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogNewGovernorAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogNewGovernorAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogNewGovernorAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogNewGovernorAccepted represents a LogNewGovernorAccepted event raised by the StarknetCore contract.
type StarknetCoreLogNewGovernorAccepted struct {
	AcceptedGovernor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewGovernorAccepted is a free log retrieval operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_StarknetCore *StarknetCoreFilterer) FilterLogNewGovernorAccepted(opts *bind.FilterOpts) (*StarknetCoreLogNewGovernorAcceptedIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogNewGovernorAcceptedIterator{contract: _StarknetCore.contract, event: "LogNewGovernorAccepted", logs: logs, sub: sub}, nil
}

// WatchLogNewGovernorAccepted is a free log subscription operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_StarknetCore *StarknetCoreFilterer) WatchLogNewGovernorAccepted(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogNewGovernorAccepted) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogNewGovernorAccepted)
				if err := _StarknetCore.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewGovernorAccepted is a log parse operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_StarknetCore *StarknetCoreFilterer) ParseLogNewGovernorAccepted(log types.Log) (*StarknetCoreLogNewGovernorAccepted, error) {
	event := new(StarknetCoreLogNewGovernorAccepted)
	if err := _StarknetCore.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogNominatedGovernorIterator is returned from FilterLogNominatedGovernor and is used to iterate over the raw logs and unpacked data for LogNominatedGovernor events raised by the StarknetCore contract.
type StarknetCoreLogNominatedGovernorIterator struct {
	Event *StarknetCoreLogNominatedGovernor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogNominatedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogNominatedGovernor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogNominatedGovernor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogNominatedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogNominatedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogNominatedGovernor represents a LogNominatedGovernor event raised by the StarknetCore contract.
type StarknetCoreLogNominatedGovernor struct {
	NominatedGovernor common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogNominatedGovernor is a free log retrieval operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_StarknetCore *StarknetCoreFilterer) FilterLogNominatedGovernor(opts *bind.FilterOpts) (*StarknetCoreLogNominatedGovernorIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogNominatedGovernorIterator{contract: _StarknetCore.contract, event: "LogNominatedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogNominatedGovernor is a free log subscription operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_StarknetCore *StarknetCoreFilterer) WatchLogNominatedGovernor(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogNominatedGovernor) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogNominatedGovernor)
				if err := _StarknetCore.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNominatedGovernor is a log parse operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_StarknetCore *StarknetCoreFilterer) ParseLogNominatedGovernor(log types.Log) (*StarknetCoreLogNominatedGovernor, error) {
	event := new(StarknetCoreLogNominatedGovernor)
	if err := _StarknetCore.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogNominationCancelledIterator is returned from FilterLogNominationCancelled and is used to iterate over the raw logs and unpacked data for LogNominationCancelled events raised by the StarknetCore contract.
type StarknetCoreLogNominationCancelledIterator struct {
	Event *StarknetCoreLogNominationCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogNominationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogNominationCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogNominationCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogNominationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogNominationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogNominationCancelled represents a LogNominationCancelled event raised by the StarknetCore contract.
type StarknetCoreLogNominationCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNominationCancelled is a free log retrieval operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_StarknetCore *StarknetCoreFilterer) FilterLogNominationCancelled(opts *bind.FilterOpts) (*StarknetCoreLogNominationCancelledIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogNominationCancelledIterator{contract: _StarknetCore.contract, event: "LogNominationCancelled", logs: logs, sub: sub}, nil
}

// WatchLogNominationCancelled is a free log subscription operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_StarknetCore *StarknetCoreFilterer) WatchLogNominationCancelled(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogNominationCancelled) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogNominationCancelled)
				if err := _StarknetCore.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNominationCancelled is a log parse operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_StarknetCore *StarknetCoreFilterer) ParseLogNominationCancelled(log types.Log) (*StarknetCoreLogNominationCancelled, error) {
	event := new(StarknetCoreLogNominationCancelled)
	if err := _StarknetCore.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogOperatorAddedIterator is returned from FilterLogOperatorAdded and is used to iterate over the raw logs and unpacked data for LogOperatorAdded events raised by the StarknetCore contract.
type StarknetCoreLogOperatorAddedIterator struct {
	Event *StarknetCoreLogOperatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogOperatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogOperatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogOperatorAdded represents a LogOperatorAdded event raised by the StarknetCore contract.
type StarknetCoreLogOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorAdded is a free log retrieval operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_StarknetCore *StarknetCoreFilterer) FilterLogOperatorAdded(opts *bind.FilterOpts) (*StarknetCoreLogOperatorAddedIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogOperatorAddedIterator{contract: _StarknetCore.contract, event: "LogOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogOperatorAdded is a free log subscription operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_StarknetCore *StarknetCoreFilterer) WatchLogOperatorAdded(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogOperatorAdded)
				if err := _StarknetCore.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogOperatorAdded is a log parse operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_StarknetCore *StarknetCoreFilterer) ParseLogOperatorAdded(log types.Log) (*StarknetCoreLogOperatorAdded, error) {
	event := new(StarknetCoreLogOperatorAdded)
	if err := _StarknetCore.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogOperatorRemovedIterator is returned from FilterLogOperatorRemoved and is used to iterate over the raw logs and unpacked data for LogOperatorRemoved events raised by the StarknetCore contract.
type StarknetCoreLogOperatorRemovedIterator struct {
	Event *StarknetCoreLogOperatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogOperatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogOperatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogOperatorRemoved represents a LogOperatorRemoved event raised by the StarknetCore contract.
type StarknetCoreLogOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorRemoved is a free log retrieval operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_StarknetCore *StarknetCoreFilterer) FilterLogOperatorRemoved(opts *bind.FilterOpts) (*StarknetCoreLogOperatorRemovedIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogOperatorRemovedIterator{contract: _StarknetCore.contract, event: "LogOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogOperatorRemoved is a free log subscription operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_StarknetCore *StarknetCoreFilterer) WatchLogOperatorRemoved(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogOperatorRemoved)
				if err := _StarknetCore.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogOperatorRemoved is a log parse operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_StarknetCore *StarknetCoreFilterer) ParseLogOperatorRemoved(log types.Log) (*StarknetCoreLogOperatorRemoved, error) {
	event := new(StarknetCoreLogOperatorRemoved)
	if err := _StarknetCore.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogRemovedGovernorIterator is returned from FilterLogRemovedGovernor and is used to iterate over the raw logs and unpacked data for LogRemovedGovernor events raised by the StarknetCore contract.
type StarknetCoreLogRemovedGovernorIterator struct {
	Event *StarknetCoreLogRemovedGovernor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogRemovedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogRemovedGovernor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogRemovedGovernor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogRemovedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogRemovedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogRemovedGovernor represents a LogRemovedGovernor event raised by the StarknetCore contract.
type StarknetCoreLogRemovedGovernor struct {
	RemovedGovernor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRemovedGovernor is a free log retrieval operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_StarknetCore *StarknetCoreFilterer) FilterLogRemovedGovernor(opts *bind.FilterOpts) (*StarknetCoreLogRemovedGovernorIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogRemovedGovernorIterator{contract: _StarknetCore.contract, event: "LogRemovedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogRemovedGovernor is a free log subscription operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_StarknetCore *StarknetCoreFilterer) WatchLogRemovedGovernor(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogRemovedGovernor) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogRemovedGovernor)
				if err := _StarknetCore.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRemovedGovernor is a log parse operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_StarknetCore *StarknetCoreFilterer) ParseLogRemovedGovernor(log types.Log) (*StarknetCoreLogRemovedGovernor, error) {
	event := new(StarknetCoreLogRemovedGovernor)
	if err := _StarknetCore.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the StarknetCore contract.
type StarknetCoreLogStateTransitionFactIterator struct {
	Event *StarknetCoreLogStateTransitionFact // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogStateTransitionFact)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogStateTransitionFact)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogStateTransitionFact represents a LogStateTransitionFact event raised by the StarknetCore contract.
type StarknetCoreLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_StarknetCore *StarknetCoreFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*StarknetCoreLogStateTransitionFactIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogStateTransitionFactIterator{contract: _StarknetCore.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_StarknetCore *StarknetCoreFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogStateTransitionFact)
				if err := _StarknetCore.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogStateTransitionFact is a log parse operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_StarknetCore *StarknetCoreFilterer) ParseLogStateTransitionFact(log types.Log) (*StarknetCoreLogStateTransitionFact, error) {
	event := new(StarknetCoreLogStateTransitionFact)
	if err := _StarknetCore.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreLogStateUpdateIterator is returned from FilterLogStateUpdate and is used to iterate over the raw logs and unpacked data for LogStateUpdate events raised by the StarknetCore contract.
type StarknetCoreLogStateUpdateIterator struct {
	Event *StarknetCoreLogStateUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreLogStateUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreLogStateUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreLogStateUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreLogStateUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreLogStateUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreLogStateUpdate represents a LogStateUpdate event raised by the StarknetCore contract.
type StarknetCoreLogStateUpdate struct {
	GlobalRoot  *big.Int
	BlockNumber *big.Int
	BlockHash   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogStateUpdate is a free log retrieval operation binding the contract event 0xd342ddf7a308dec111745b00315c14b7efb2bdae570a6856e088ed0c65a3576c.
//
// Solidity: event LogStateUpdate(uint256 globalRoot, int256 blockNumber, uint256 blockHash)
func (_StarknetCore *StarknetCoreFilterer) FilterLogStateUpdate(opts *bind.FilterOpts) (*StarknetCoreLogStateUpdateIterator, error) {

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "LogStateUpdate")
	if err != nil {
		return nil, err
	}
	return &StarknetCoreLogStateUpdateIterator{contract: _StarknetCore.contract, event: "LogStateUpdate", logs: logs, sub: sub}, nil
}

// WatchLogStateUpdate is a free log subscription operation binding the contract event 0xd342ddf7a308dec111745b00315c14b7efb2bdae570a6856e088ed0c65a3576c.
//
// Solidity: event LogStateUpdate(uint256 globalRoot, int256 blockNumber, uint256 blockHash)
func (_StarknetCore *StarknetCoreFilterer) WatchLogStateUpdate(opts *bind.WatchOpts, sink chan<- *StarknetCoreLogStateUpdate) (event.Subscription, error) {

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "LogStateUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreLogStateUpdate)
				if err := _StarknetCore.contract.UnpackLog(event, "LogStateUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogStateUpdate is a log parse operation binding the contract event 0xd342ddf7a308dec111745b00315c14b7efb2bdae570a6856e088ed0c65a3576c.
//
// Solidity: event LogStateUpdate(uint256 globalRoot, int256 blockNumber, uint256 blockHash)
func (_StarknetCore *StarknetCoreFilterer) ParseLogStateUpdate(log types.Log) (*StarknetCoreLogStateUpdate, error) {
	event := new(StarknetCoreLogStateUpdate)
	if err := _StarknetCore.contract.UnpackLog(event, "LogStateUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreMessageToL2CanceledIterator is returned from FilterMessageToL2Canceled and is used to iterate over the raw logs and unpacked data for MessageToL2Canceled events raised by the StarknetCore contract.
type StarknetCoreMessageToL2CanceledIterator struct {
	Event *StarknetCoreMessageToL2Canceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreMessageToL2CanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreMessageToL2Canceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreMessageToL2Canceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreMessageToL2CanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreMessageToL2CanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreMessageToL2Canceled represents a MessageToL2Canceled event raised by the StarknetCore contract.
type StarknetCoreMessageToL2Canceled struct {
	FromAddress common.Address
	ToAddress   *big.Int
	Selector    *big.Int
	Payload     []*big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageToL2Canceled is a free log retrieval operation binding the contract event 0x8abd2ec2e0a10c82f5b60ea00455fa96c41fd144f225fcc52b8d83d94f803ed8.
//
// Solidity: event MessageToL2Canceled(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) FilterMessageToL2Canceled(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetCoreMessageToL2CanceledIterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "MessageToL2Canceled", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreMessageToL2CanceledIterator{contract: _StarknetCore.contract, event: "MessageToL2Canceled", logs: logs, sub: sub}, nil
}

// WatchMessageToL2Canceled is a free log subscription operation binding the contract event 0x8abd2ec2e0a10c82f5b60ea00455fa96c41fd144f225fcc52b8d83d94f803ed8.
//
// Solidity: event MessageToL2Canceled(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) WatchMessageToL2Canceled(opts *bind.WatchOpts, sink chan<- *StarknetCoreMessageToL2Canceled, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "MessageToL2Canceled", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreMessageToL2Canceled)
				if err := _StarknetCore.contract.UnpackLog(event, "MessageToL2Canceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageToL2Canceled is a log parse operation binding the contract event 0x8abd2ec2e0a10c82f5b60ea00455fa96c41fd144f225fcc52b8d83d94f803ed8.
//
// Solidity: event MessageToL2Canceled(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) ParseMessageToL2Canceled(log types.Log) (*StarknetCoreMessageToL2Canceled, error) {
	event := new(StarknetCoreMessageToL2Canceled)
	if err := _StarknetCore.contract.UnpackLog(event, "MessageToL2Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreMessageToL2CancellationStartedIterator is returned from FilterMessageToL2CancellationStarted and is used to iterate over the raw logs and unpacked data for MessageToL2CancellationStarted events raised by the StarknetCore contract.
type StarknetCoreMessageToL2CancellationStartedIterator struct {
	Event *StarknetCoreMessageToL2CancellationStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreMessageToL2CancellationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreMessageToL2CancellationStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreMessageToL2CancellationStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreMessageToL2CancellationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreMessageToL2CancellationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreMessageToL2CancellationStarted represents a MessageToL2CancellationStarted event raised by the StarknetCore contract.
type StarknetCoreMessageToL2CancellationStarted struct {
	FromAddress common.Address
	ToAddress   *big.Int
	Selector    *big.Int
	Payload     []*big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageToL2CancellationStarted is a free log retrieval operation binding the contract event 0x2e00dccd686fd6823ec7dc3e125582aa82881b6ff5f6b5a73856e1ea8338a3be.
//
// Solidity: event MessageToL2CancellationStarted(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) FilterMessageToL2CancellationStarted(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetCoreMessageToL2CancellationStartedIterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "MessageToL2CancellationStarted", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreMessageToL2CancellationStartedIterator{contract: _StarknetCore.contract, event: "MessageToL2CancellationStarted", logs: logs, sub: sub}, nil
}

// WatchMessageToL2CancellationStarted is a free log subscription operation binding the contract event 0x2e00dccd686fd6823ec7dc3e125582aa82881b6ff5f6b5a73856e1ea8338a3be.
//
// Solidity: event MessageToL2CancellationStarted(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) WatchMessageToL2CancellationStarted(opts *bind.WatchOpts, sink chan<- *StarknetCoreMessageToL2CancellationStarted, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "MessageToL2CancellationStarted", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreMessageToL2CancellationStarted)
				if err := _StarknetCore.contract.UnpackLog(event, "MessageToL2CancellationStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageToL2CancellationStarted is a log parse operation binding the contract event 0x2e00dccd686fd6823ec7dc3e125582aa82881b6ff5f6b5a73856e1ea8338a3be.
//
// Solidity: event MessageToL2CancellationStarted(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_StarknetCore *StarknetCoreFilterer) ParseMessageToL2CancellationStarted(log types.Log) (*StarknetCoreMessageToL2CancellationStarted, error) {
	event := new(StarknetCoreMessageToL2CancellationStarted)
	if err := _StarknetCore.contract.UnpackLog(event, "MessageToL2CancellationStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetCoreProgramHashChangedIterator is returned from FilterProgramHashChanged and is used to iterate over the raw logs and unpacked data for ProgramHashChanged events raised by the StarknetCore contract.
type StarknetCoreProgramHashChangedIterator struct {
	Event *StarknetCoreProgramHashChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarknetCoreProgramHashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetCoreProgramHashChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarknetCoreProgramHashChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarknetCoreProgramHashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetCoreProgramHashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetCoreProgramHashChanged represents a ProgramHashChanged event raised by the StarknetCore contract.
type StarknetCoreProgramHashChanged struct {
	ChangedBy      common.Address
	OldProgramHash *big.Int
	NewProgramHash *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProgramHashChanged is a free log retrieval operation binding the contract event 0x600a61c1b32ac42fb2fe76e8fc7582a98106668fc16dcd85567cd3937363e49b.
//
// Solidity: event ProgramHashChanged(address indexed changedBy, uint256 oldProgramHash, uint256 newProgramHash)
func (_StarknetCore *StarknetCoreFilterer) FilterProgramHashChanged(opts *bind.FilterOpts, changedBy []common.Address) (*StarknetCoreProgramHashChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _StarknetCore.contract.FilterLogs(opts, "ProgramHashChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &StarknetCoreProgramHashChangedIterator{contract: _StarknetCore.contract, event: "ProgramHashChanged", logs: logs, sub: sub}, nil
}

// WatchProgramHashChanged is a free log subscription operation binding the contract event 0x600a61c1b32ac42fb2fe76e8fc7582a98106668fc16dcd85567cd3937363e49b.
//
// Solidity: event ProgramHashChanged(address indexed changedBy, uint256 oldProgramHash, uint256 newProgramHash)
func (_StarknetCore *StarknetCoreFilterer) WatchProgramHashChanged(opts *bind.WatchOpts, sink chan<- *StarknetCoreProgramHashChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _StarknetCore.contract.WatchLogs(opts, "ProgramHashChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetCoreProgramHashChanged)
				if err := _StarknetCore.contract.UnpackLog(event, "ProgramHashChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProgramHashChanged is a log parse operation binding the contract event 0x600a61c1b32ac42fb2fe76e8fc7582a98106668fc16dcd85567cd3937363e49b.
//
// Solidity: event ProgramHashChanged(address indexed changedBy, uint256 oldProgramHash, uint256 newProgramHash)
func (_StarknetCore *StarknetCoreFilterer) ParseProgramHashChanged(log types.Log) (*StarknetCoreProgramHashChanged, error) {
	event := new(StarknetCoreProgramHashChanged)
	if err := _StarknetCore.contract.UnpackLog(event, "ProgramHashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
