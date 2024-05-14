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

// StarknetMetaData contains all meta data concerning the Starknet contract.
var StarknetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"}],\"name\":\"ConsumedMessageToL1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ConsumedMessageToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"}],\"name\":\"LogMessageToL1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogMessageToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MessageToL2Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toAddress\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payload\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MessageToL2CancellationStarted\",\"type\":\"event\"}]",
}

// StarknetABI is the input ABI used to generate the binding from.
// Deprecated: Use StarknetMetaData.ABI instead.
var StarknetABI = StarknetMetaData.ABI

// Starknet is an auto generated Go binding around an Ethereum contract.
type Starknet struct {
	StarknetCaller     // Read-only binding to the contract
	StarknetTransactor // Write-only binding to the contract
	StarknetFilterer   // Log filterer for contract events
}

// StarknetCaller is an auto generated read-only Go binding around an Ethereum contract.
type StarknetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarknetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StarknetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarknetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StarknetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarknetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StarknetSession struct {
	Contract     *Starknet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StarknetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StarknetCallerSession struct {
	Contract *StarknetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StarknetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StarknetTransactorSession struct {
	Contract     *StarknetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StarknetRaw is an auto generated low-level Go binding around an Ethereum contract.
type StarknetRaw struct {
	Contract *Starknet // Generic contract binding to access the raw methods on
}

// StarknetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StarknetCallerRaw struct {
	Contract *StarknetCaller // Generic read-only contract binding to access the raw methods on
}

// StarknetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StarknetTransactorRaw struct {
	Contract *StarknetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStarknet creates a new instance of Starknet, bound to a specific deployed contract.
func NewStarknet(address common.Address, backend bind.ContractBackend) (*Starknet, error) {
	contract, err := bindStarknet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Starknet{StarknetCaller: StarknetCaller{contract: contract}, StarknetTransactor: StarknetTransactor{contract: contract}, StarknetFilterer: StarknetFilterer{contract: contract}}, nil
}

// NewStarknetCaller creates a new read-only instance of Starknet, bound to a specific deployed contract.
func NewStarknetCaller(address common.Address, caller bind.ContractCaller) (*StarknetCaller, error) {
	contract, err := bindStarknet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StarknetCaller{contract: contract}, nil
}

// NewStarknetTransactor creates a new write-only instance of Starknet, bound to a specific deployed contract.
func NewStarknetTransactor(address common.Address, transactor bind.ContractTransactor) (*StarknetTransactor, error) {
	contract, err := bindStarknet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StarknetTransactor{contract: contract}, nil
}

// NewStarknetFilterer creates a new log filterer instance of Starknet, bound to a specific deployed contract.
func NewStarknetFilterer(address common.Address, filterer bind.ContractFilterer) (*StarknetFilterer, error) {
	contract, err := bindStarknet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StarknetFilterer{contract: contract}, nil
}

// bindStarknet binds a generic wrapper to an already deployed contract.
func bindStarknet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StarknetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Starknet *StarknetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Starknet.Contract.StarknetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Starknet *StarknetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Starknet.Contract.StarknetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Starknet *StarknetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Starknet.Contract.StarknetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Starknet *StarknetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Starknet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Starknet *StarknetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Starknet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Starknet *StarknetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Starknet.Contract.contract.Transact(opts, method, params...)
}

// StarknetConsumedMessageToL1Iterator is returned from FilterConsumedMessageToL1 and is used to iterate over the raw logs and unpacked data for ConsumedMessageToL1 events raised by the Starknet contract.
type StarknetConsumedMessageToL1Iterator struct {
	Event *StarknetConsumedMessageToL1 // Event containing the contract specifics and raw log

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
func (it *StarknetConsumedMessageToL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetConsumedMessageToL1)
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
		it.Event = new(StarknetConsumedMessageToL1)
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
func (it *StarknetConsumedMessageToL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetConsumedMessageToL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetConsumedMessageToL1 represents a ConsumedMessageToL1 event raised by the Starknet contract.
type StarknetConsumedMessageToL1 struct {
	FromAddress *big.Int
	ToAddress   common.Address
	Payload     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsumedMessageToL1 is a free log retrieval operation binding the contract event 0x7a06c571aa77f34d9706c51e5d8122b5595aebeaa34233bfe866f22befb973b1.
//
// Solidity: event ConsumedMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_Starknet *StarknetFilterer) FilterConsumedMessageToL1(opts *bind.FilterOpts, fromAddress []*big.Int, toAddress []common.Address) (*StarknetConsumedMessageToL1Iterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Starknet.contract.FilterLogs(opts, "ConsumedMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &StarknetConsumedMessageToL1Iterator{contract: _Starknet.contract, event: "ConsumedMessageToL1", logs: logs, sub: sub}, nil
}

// WatchConsumedMessageToL1 is a free log subscription operation binding the contract event 0x7a06c571aa77f34d9706c51e5d8122b5595aebeaa34233bfe866f22befb973b1.
//
// Solidity: event ConsumedMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_Starknet *StarknetFilterer) WatchConsumedMessageToL1(opts *bind.WatchOpts, sink chan<- *StarknetConsumedMessageToL1, fromAddress []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Starknet.contract.WatchLogs(opts, "ConsumedMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetConsumedMessageToL1)
				if err := _Starknet.contract.UnpackLog(event, "ConsumedMessageToL1", log); err != nil {
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
func (_Starknet *StarknetFilterer) ParseConsumedMessageToL1(log types.Log) (*StarknetConsumedMessageToL1, error) {
	event := new(StarknetConsumedMessageToL1)
	if err := _Starknet.contract.UnpackLog(event, "ConsumedMessageToL1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetConsumedMessageToL2Iterator is returned from FilterConsumedMessageToL2 and is used to iterate over the raw logs and unpacked data for ConsumedMessageToL2 events raised by the Starknet contract.
type StarknetConsumedMessageToL2Iterator struct {
	Event *StarknetConsumedMessageToL2 // Event containing the contract specifics and raw log

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
func (it *StarknetConsumedMessageToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetConsumedMessageToL2)
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
		it.Event = new(StarknetConsumedMessageToL2)
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
func (it *StarknetConsumedMessageToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetConsumedMessageToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetConsumedMessageToL2 represents a ConsumedMessageToL2 event raised by the Starknet contract.
type StarknetConsumedMessageToL2 struct {
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
func (_Starknet *StarknetFilterer) FilterConsumedMessageToL2(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetConsumedMessageToL2Iterator, error) {

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

	logs, sub, err := _Starknet.contract.FilterLogs(opts, "ConsumedMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetConsumedMessageToL2Iterator{contract: _Starknet.contract, event: "ConsumedMessageToL2", logs: logs, sub: sub}, nil
}

// WatchConsumedMessageToL2 is a free log subscription operation binding the contract event 0x9592d37825c744e33fa80c469683bbd04d336241bb600b574758efd182abe26a.
//
// Solidity: event ConsumedMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_Starknet *StarknetFilterer) WatchConsumedMessageToL2(opts *bind.WatchOpts, sink chan<- *StarknetConsumedMessageToL2, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Starknet.contract.WatchLogs(opts, "ConsumedMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetConsumedMessageToL2)
				if err := _Starknet.contract.UnpackLog(event, "ConsumedMessageToL2", log); err != nil {
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
func (_Starknet *StarknetFilterer) ParseConsumedMessageToL2(log types.Log) (*StarknetConsumedMessageToL2, error) {
	event := new(StarknetConsumedMessageToL2)
	if err := _Starknet.contract.UnpackLog(event, "ConsumedMessageToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetLogMessageToL1Iterator is returned from FilterLogMessageToL1 and is used to iterate over the raw logs and unpacked data for LogMessageToL1 events raised by the Starknet contract.
type StarknetLogMessageToL1Iterator struct {
	Event *StarknetLogMessageToL1 // Event containing the contract specifics and raw log

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
func (it *StarknetLogMessageToL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetLogMessageToL1)
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
		it.Event = new(StarknetLogMessageToL1)
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
func (it *StarknetLogMessageToL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetLogMessageToL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetLogMessageToL1 represents a LogMessageToL1 event raised by the Starknet contract.
type StarknetLogMessageToL1 struct {
	FromAddress *big.Int
	ToAddress   common.Address
	Payload     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogMessageToL1 is a free log retrieval operation binding the contract event 0x4264ac208b5fde633ccdd42e0f12c3d6d443a4f3779bbf886925b94665b63a22.
//
// Solidity: event LogMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_Starknet *StarknetFilterer) FilterLogMessageToL1(opts *bind.FilterOpts, fromAddress []*big.Int, toAddress []common.Address) (*StarknetLogMessageToL1Iterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Starknet.contract.FilterLogs(opts, "LogMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &StarknetLogMessageToL1Iterator{contract: _Starknet.contract, event: "LogMessageToL1", logs: logs, sub: sub}, nil
}

// WatchLogMessageToL1 is a free log subscription operation binding the contract event 0x4264ac208b5fde633ccdd42e0f12c3d6d443a4f3779bbf886925b94665b63a22.
//
// Solidity: event LogMessageToL1(uint256 indexed fromAddress, address indexed toAddress, uint256[] payload)
func (_Starknet *StarknetFilterer) WatchLogMessageToL1(opts *bind.WatchOpts, sink chan<- *StarknetLogMessageToL1, fromAddress []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Starknet.contract.WatchLogs(opts, "LogMessageToL1", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetLogMessageToL1)
				if err := _Starknet.contract.UnpackLog(event, "LogMessageToL1", log); err != nil {
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
func (_Starknet *StarknetFilterer) ParseLogMessageToL1(log types.Log) (*StarknetLogMessageToL1, error) {
	event := new(StarknetLogMessageToL1)
	if err := _Starknet.contract.UnpackLog(event, "LogMessageToL1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetLogMessageToL2Iterator is returned from FilterLogMessageToL2 and is used to iterate over the raw logs and unpacked data for LogMessageToL2 events raised by the Starknet contract.
type StarknetLogMessageToL2Iterator struct {
	Event *StarknetLogMessageToL2 // Event containing the contract specifics and raw log

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
func (it *StarknetLogMessageToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetLogMessageToL2)
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
		it.Event = new(StarknetLogMessageToL2)
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
func (it *StarknetLogMessageToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetLogMessageToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetLogMessageToL2 represents a LogMessageToL2 event raised by the Starknet contract.
type StarknetLogMessageToL2 struct {
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
func (_Starknet *StarknetFilterer) FilterLogMessageToL2(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetLogMessageToL2Iterator, error) {

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

	logs, sub, err := _Starknet.contract.FilterLogs(opts, "LogMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetLogMessageToL2Iterator{contract: _Starknet.contract, event: "LogMessageToL2", logs: logs, sub: sub}, nil
}

// WatchLogMessageToL2 is a free log subscription operation binding the contract event 0xdb80dd488acf86d17c747445b0eabb5d57c541d3bd7b6b87af987858e5066b2b.
//
// Solidity: event LogMessageToL2(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce, uint256 fee)
func (_Starknet *StarknetFilterer) WatchLogMessageToL2(opts *bind.WatchOpts, sink chan<- *StarknetLogMessageToL2, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Starknet.contract.WatchLogs(opts, "LogMessageToL2", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetLogMessageToL2)
				if err := _Starknet.contract.UnpackLog(event, "LogMessageToL2", log); err != nil {
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
func (_Starknet *StarknetFilterer) ParseLogMessageToL2(log types.Log) (*StarknetLogMessageToL2, error) {
	event := new(StarknetLogMessageToL2)
	if err := _Starknet.contract.UnpackLog(event, "LogMessageToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetMessageToL2CanceledIterator is returned from FilterMessageToL2Canceled and is used to iterate over the raw logs and unpacked data for MessageToL2Canceled events raised by the Starknet contract.
type StarknetMessageToL2CanceledIterator struct {
	Event *StarknetMessageToL2Canceled // Event containing the contract specifics and raw log

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
func (it *StarknetMessageToL2CanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetMessageToL2Canceled)
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
		it.Event = new(StarknetMessageToL2Canceled)
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
func (it *StarknetMessageToL2CanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetMessageToL2CanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetMessageToL2Canceled represents a MessageToL2Canceled event raised by the Starknet contract.
type StarknetMessageToL2Canceled struct {
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
func (_Starknet *StarknetFilterer) FilterMessageToL2Canceled(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetMessageToL2CanceledIterator, error) {

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

	logs, sub, err := _Starknet.contract.FilterLogs(opts, "MessageToL2Canceled", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetMessageToL2CanceledIterator{contract: _Starknet.contract, event: "MessageToL2Canceled", logs: logs, sub: sub}, nil
}

// WatchMessageToL2Canceled is a free log subscription operation binding the contract event 0x8abd2ec2e0a10c82f5b60ea00455fa96c41fd144f225fcc52b8d83d94f803ed8.
//
// Solidity: event MessageToL2Canceled(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_Starknet *StarknetFilterer) WatchMessageToL2Canceled(opts *bind.WatchOpts, sink chan<- *StarknetMessageToL2Canceled, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Starknet.contract.WatchLogs(opts, "MessageToL2Canceled", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetMessageToL2Canceled)
				if err := _Starknet.contract.UnpackLog(event, "MessageToL2Canceled", log); err != nil {
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
func (_Starknet *StarknetFilterer) ParseMessageToL2Canceled(log types.Log) (*StarknetMessageToL2Canceled, error) {
	event := new(StarknetMessageToL2Canceled)
	if err := _Starknet.contract.UnpackLog(event, "MessageToL2Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarknetMessageToL2CancellationStartedIterator is returned from FilterMessageToL2CancellationStarted and is used to iterate over the raw logs and unpacked data for MessageToL2CancellationStarted events raised by the Starknet contract.
type StarknetMessageToL2CancellationStartedIterator struct {
	Event *StarknetMessageToL2CancellationStarted // Event containing the contract specifics and raw log

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
func (it *StarknetMessageToL2CancellationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarknetMessageToL2CancellationStarted)
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
		it.Event = new(StarknetMessageToL2CancellationStarted)
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
func (it *StarknetMessageToL2CancellationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarknetMessageToL2CancellationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarknetMessageToL2CancellationStarted represents a MessageToL2CancellationStarted event raised by the Starknet contract.
type StarknetMessageToL2CancellationStarted struct {
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
func (_Starknet *StarknetFilterer) FilterMessageToL2CancellationStarted(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (*StarknetMessageToL2CancellationStartedIterator, error) {

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

	logs, sub, err := _Starknet.contract.FilterLogs(opts, "MessageToL2CancellationStarted", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &StarknetMessageToL2CancellationStartedIterator{contract: _Starknet.contract, event: "MessageToL2CancellationStarted", logs: logs, sub: sub}, nil
}

// WatchMessageToL2CancellationStarted is a free log subscription operation binding the contract event 0x2e00dccd686fd6823ec7dc3e125582aa82881b6ff5f6b5a73856e1ea8338a3be.
//
// Solidity: event MessageToL2CancellationStarted(address indexed fromAddress, uint256 indexed toAddress, uint256 indexed selector, uint256[] payload, uint256 nonce)
func (_Starknet *StarknetFilterer) WatchMessageToL2CancellationStarted(opts *bind.WatchOpts, sink chan<- *StarknetMessageToL2CancellationStarted, fromAddress []common.Address, toAddress []*big.Int, selector []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Starknet.contract.WatchLogs(opts, "MessageToL2CancellationStarted", fromAddressRule, toAddressRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarknetMessageToL2CancellationStarted)
				if err := _Starknet.contract.UnpackLog(event, "MessageToL2CancellationStarted", log); err != nil {
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
func (_Starknet *StarknetFilterer) ParseMessageToL2CancellationStarted(log types.Log) (*StarknetMessageToL2CancellationStarted, error) {
	event := new(StarknetMessageToL2CancellationStarted)
	if err := _Starknet.contract.UnpackLog(event, "MessageToL2CancellationStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
