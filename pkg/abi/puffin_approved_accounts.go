// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// PuffinApprovedAccountsMetaData contains all meta data concerning the PuffinApprovedAccounts contract.
var PuffinApprovedAccountsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PuffinApprovedAccountsABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinApprovedAccountsMetaData.ABI instead.
var PuffinApprovedAccountsABI = PuffinApprovedAccountsMetaData.ABI

// PuffinApprovedAccounts is an auto generated Go binding around an Ethereum contract.
type PuffinApprovedAccounts struct {
	PuffinApprovedAccountsCaller     // Read-only binding to the contract
	PuffinApprovedAccountsTransactor // Write-only binding to the contract
	PuffinApprovedAccountsFilterer   // log filterer for contract events
}

// PuffinApprovedAccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinApprovedAccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinApprovedAccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinApprovedAccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinApprovedAccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinApprovedAccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinApprovedAccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinApprovedAccountsSession struct {
	Contract     *PuffinApprovedAccounts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PuffinApprovedAccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinApprovedAccountsCallerSession struct {
	Contract *PuffinApprovedAccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// PuffinApprovedAccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinApprovedAccountsTransactorSession struct {
	Contract     *PuffinApprovedAccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PuffinApprovedAccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinApprovedAccountsRaw struct {
	Contract *PuffinApprovedAccounts // Generic contract binding to access the raw methods on
}

// PuffinApprovedAccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinApprovedAccountsCallerRaw struct {
	Contract *PuffinApprovedAccountsCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinApprovedAccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinApprovedAccountsTransactorRaw struct {
	Contract *PuffinApprovedAccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinApprovedAccounts creates a new instance of PuffinApprovedAccounts, bound to a specific deployed contract.
func NewPuffinApprovedAccounts(address common.Address, backend bind.ContractBackend) (*PuffinApprovedAccounts, error) {
	contract, err := bindPuffinApprovedAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovedAccounts{PuffinApprovedAccountsCaller: PuffinApprovedAccountsCaller{contract: contract}, PuffinApprovedAccountsTransactor: PuffinApprovedAccountsTransactor{contract: contract}, PuffinApprovedAccountsFilterer: PuffinApprovedAccountsFilterer{contract: contract}}, nil
}

// NewPuffinApprovedAccountsCaller creates a new read-only instance of PuffinApprovedAccounts, bound to a specific deployed contract.
func NewPuffinApprovedAccountsCaller(address common.Address, caller bind.ContractCaller) (*PuffinApprovedAccountsCaller, error) {
	contract, err := bindPuffinApprovedAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovedAccountsCaller{contract: contract}, nil
}

// NewPuffinApprovedAccountsTransactor creates a new write-only instance of PuffinApprovedAccounts, bound to a specific deployed contract.
func NewPuffinApprovedAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinApprovedAccountsTransactor, error) {
	contract, err := bindPuffinApprovedAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovedAccountsTransactor{contract: contract}, nil
}

// NewPuffinApprovedAccountsFilterer creates a new log filterer instance of PuffinApprovedAccounts, bound to a specific deployed contract.
func NewPuffinApprovedAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinApprovedAccountsFilterer, error) {
	contract, err := bindPuffinApprovedAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovedAccountsFilterer{contract: contract}, nil
}

// bindPuffinApprovedAccounts binds a generic wrapper to an already deployed contract.
func bindPuffinApprovedAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinApprovedAccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinApprovedAccounts *PuffinApprovedAccountsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinApprovedAccounts.Contract.PuffinApprovedAccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinApprovedAccounts *PuffinApprovedAccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.PuffinApprovedAccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinApprovedAccounts *PuffinApprovedAccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.PuffinApprovedAccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinApprovedAccounts *PuffinApprovedAccountsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinApprovedAccounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.contract.Transact(opts, method, params...)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsCaller) IsApproved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinApprovedAccounts.contract.Call(opts, &out, "isApproved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsSession) IsApproved(arg0 common.Address) (bool, error) {
	return _PuffinApprovedAccounts.Contract.IsApproved(&_PuffinApprovedAccounts.CallOpts, arg0)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsCallerSession) IsApproved(arg0 common.Address) (bool, error) {
	return _PuffinApprovedAccounts.Contract.IsApproved(&_PuffinApprovedAccounts.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinApprovedAccounts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsSession) Owner() (common.Address, error) {
	return _PuffinApprovedAccounts.Contract.Owner(&_PuffinApprovedAccounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsCallerSession) Owner() (common.Address, error) {
	return _PuffinApprovedAccounts.Contract.Owner(&_PuffinApprovedAccounts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactor) Approve(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.contract.Transact(opts, "approve", user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsSession) Approve(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.Approve(&_PuffinApprovedAccounts.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactorSession) Approve(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.Approve(&_PuffinApprovedAccounts.TransactOpts, user)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address user) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactor) Remove(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.contract.Transact(opts, "remove", user)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address user) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsSession) Remove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.Remove(&_PuffinApprovedAccounts.TransactOpts, user)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address user) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactorSession) Remove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.Remove(&_PuffinApprovedAccounts.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.RenounceOwnership(&_PuffinApprovedAccounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.RenounceOwnership(&_PuffinApprovedAccounts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.TransferOwnership(&_PuffinApprovedAccounts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinApprovedAccounts *PuffinApprovedAccountsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinApprovedAccounts.Contract.TransferOwnership(&_PuffinApprovedAccounts.TransactOpts, newOwner)
}

// PuffinApprovedAccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinApprovedAccounts contract.
type PuffinApprovedAccountsOwnershipTransferredIterator struct {
	Event *PuffinApprovedAccountsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PuffinApprovedAccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinApprovedAccountsOwnershipTransferred)
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
		it.Event = new(PuffinApprovedAccountsOwnershipTransferred)
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
func (it *PuffinApprovedAccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinApprovedAccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinApprovedAccountsOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinApprovedAccounts contract.
type PuffinApprovedAccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinApprovedAccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinApprovedAccounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovedAccountsOwnershipTransferredIterator{contract: _PuffinApprovedAccounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinApprovedAccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinApprovedAccounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinApprovedAccountsOwnershipTransferred)
				if err := _PuffinApprovedAccounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinApprovedAccounts *PuffinApprovedAccountsFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinApprovedAccountsOwnershipTransferred, error) {
	event := new(PuffinApprovedAccountsOwnershipTransferred)
	if err := _PuffinApprovedAccounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
