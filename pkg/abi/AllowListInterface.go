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

// AllowListInterfaceMetaData contains all meta data concerning the AllowListInterface contract.
var AllowListInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readAllowList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AllowListInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use AllowListInterfaceMetaData.ABI instead.
var AllowListInterfaceABI = AllowListInterfaceMetaData.ABI

// AllowListInterface is an auto generated Go binding around an Ethereum contract.
type AllowListInterface struct {
	AllowListInterfaceCaller     // Read-only binding to the contract
	AllowListInterfaceTransactor // Write-only binding to the contract
	AllowListInterfaceFilterer   // Log filterer for contract events
}

// AllowListInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllowListInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowListInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllowListInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowListInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllowListInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowListInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllowListInterfaceSession struct {
	Contract     *AllowListInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AllowListInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllowListInterfaceCallerSession struct {
	Contract *AllowListInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AllowListInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllowListInterfaceTransactorSession struct {
	Contract     *AllowListInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AllowListInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllowListInterfaceRaw struct {
	Contract *AllowListInterface // Generic contract binding to access the raw methods on
}

// AllowListInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllowListInterfaceCallerRaw struct {
	Contract *AllowListInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// AllowListInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllowListInterfaceTransactorRaw struct {
	Contract *AllowListInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllowListInterface creates a new instance of AllowListInterface, bound to a specific deployed contract.
func NewAllowListInterface(address common.Address, backend bind.ContractBackend) (*AllowListInterface, error) {
	contract, err := bindAllowListInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllowListInterface{AllowListInterfaceCaller: AllowListInterfaceCaller{contract: contract}, AllowListInterfaceTransactor: AllowListInterfaceTransactor{contract: contract}, AllowListInterfaceFilterer: AllowListInterfaceFilterer{contract: contract}}, nil
}

// NewAllowListInterfaceCaller creates a new read-only instance of AllowListInterface, bound to a specific deployed contract.
func NewAllowListInterfaceCaller(address common.Address, caller bind.ContractCaller) (*AllowListInterfaceCaller, error) {
	contract, err := bindAllowListInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllowListInterfaceCaller{contract: contract}, nil
}

// NewAllowListInterfaceTransactor creates a new write-only instance of AllowListInterface, bound to a specific deployed contract.
func NewAllowListInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AllowListInterfaceTransactor, error) {
	contract, err := bindAllowListInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllowListInterfaceTransactor{contract: contract}, nil
}

// NewAllowListInterfaceFilterer creates a new log filterer instance of AllowListInterface, bound to a specific deployed contract.
func NewAllowListInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AllowListInterfaceFilterer, error) {
	contract, err := bindAllowListInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllowListInterfaceFilterer{contract: contract}, nil
}

// bindAllowListInterface binds a generic wrapper to an already deployed contract.
func bindAllowListInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AllowListInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowListInterface *AllowListInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllowListInterface.Contract.AllowListInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowListInterface *AllowListInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowListInterface.Contract.AllowListInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowListInterface *AllowListInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowListInterface.Contract.AllowListInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowListInterface *AllowListInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllowListInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowListInterface *AllowListInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowListInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowListInterface *AllowListInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowListInterface.Contract.contract.Transact(opts, method, params...)
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256)
func (_AllowListInterface *AllowListInterfaceCaller) ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AllowListInterface.contract.Call(opts, &out, "readAllowList", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256)
func (_AllowListInterface *AllowListInterfaceSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	return _AllowListInterface.Contract.ReadAllowList(&_AllowListInterface.CallOpts, addr)
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256)
func (_AllowListInterface *AllowListInterfaceCallerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	return _AllowListInterface.Contract.ReadAllowList(&_AllowListInterface.CallOpts, addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_AllowListInterface *AllowListInterfaceTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.contract.Transact(opts, "setAdmin", addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_AllowListInterface *AllowListInterfaceSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.Contract.SetAdmin(&_AllowListInterface.TransactOpts, addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_AllowListInterface *AllowListInterfaceTransactorSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.Contract.SetAdmin(&_AllowListInterface.TransactOpts, addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_AllowListInterface *AllowListInterfaceTransactor) SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.contract.Transact(opts, "setEnabled", addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_AllowListInterface *AllowListInterfaceSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.Contract.SetEnabled(&_AllowListInterface.TransactOpts, addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_AllowListInterface *AllowListInterfaceTransactorSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.Contract.SetEnabled(&_AllowListInterface.TransactOpts, addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_AllowListInterface *AllowListInterfaceTransactor) SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.contract.Transact(opts, "setNone", addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_AllowListInterface *AllowListInterfaceSession) SetNone(addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.Contract.SetNone(&_AllowListInterface.TransactOpts, addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_AllowListInterface *AllowListInterfaceTransactorSession) SetNone(addr common.Address) (*types.Transaction, error) {
	return _AllowListInterface.Contract.SetNone(&_AllowListInterface.TransactOpts, addr)
}
