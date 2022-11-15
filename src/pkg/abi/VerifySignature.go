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

// VerifySignatureMetaData contains all meta data concerning the VerifySignature contract.
var VerifySignatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedMessage\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"VerifyMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VerifySignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifySignatureMetaData.ABI instead.
var VerifySignatureABI = VerifySignatureMetaData.ABI

// VerifySignature is an auto generated Go binding around an Ethereum contract.
type VerifySignature struct {
	VerifySignatureCaller     // Read-only binding to the contract
	VerifySignatureTransactor // Write-only binding to the contract
	VerifySignatureFilterer   // Log filterer for contract events
}

// VerifySignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifySignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifySignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifySignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySignatureSession struct {
	Contract     *VerifySignature  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifySignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifySignatureCallerSession struct {
	Contract *VerifySignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VerifySignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifySignatureTransactorSession struct {
	Contract     *VerifySignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VerifySignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifySignatureRaw struct {
	Contract *VerifySignature // Generic contract binding to access the raw methods on
}

// VerifySignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifySignatureCallerRaw struct {
	Contract *VerifySignatureCaller // Generic read-only contract binding to access the raw methods on
}

// VerifySignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifySignatureTransactorRaw struct {
	Contract *VerifySignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifySignature creates a new instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignature(address common.Address, backend bind.ContractBackend) (*VerifySignature, error) {
	contract, err := bindVerifySignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifySignature{VerifySignatureCaller: VerifySignatureCaller{contract: contract}, VerifySignatureTransactor: VerifySignatureTransactor{contract: contract}, VerifySignatureFilterer: VerifySignatureFilterer{contract: contract}}, nil
}

// NewVerifySignatureCaller creates a new read-only instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureCaller(address common.Address, caller bind.ContractCaller) (*VerifySignatureCaller, error) {
	contract, err := bindVerifySignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureCaller{contract: contract}, nil
}

// NewVerifySignatureTransactor creates a new write-only instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifySignatureTransactor, error) {
	contract, err := bindVerifySignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureTransactor{contract: contract}, nil
}

// NewVerifySignatureFilterer creates a new log filterer instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifySignatureFilterer, error) {
	contract, err := bindVerifySignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureFilterer{contract: contract}, nil
}

// bindVerifySignature binds a generic wrapper to an already deployed contract.
func bindVerifySignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifySignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifySignature *VerifySignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifySignature.Contract.VerifySignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifySignature *VerifySignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifySignature.Contract.VerifySignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifySignature *VerifySignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifySignature.Contract.VerifySignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifySignature *VerifySignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifySignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifySignature *VerifySignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifySignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifySignature *VerifySignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifySignature.Contract.contract.Transact(opts, method, params...)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x659934c1.
//
// Solidity: function VerifyMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_VerifySignature *VerifySignatureCaller) VerifyMessage(opts *bind.CallOpts, _hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VerifySignature.contract.Call(opts, &out, "VerifyMessage", _hashedMessage, _v, _r, _s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x659934c1.
//
// Solidity: function VerifyMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_VerifySignature *VerifySignatureSession) VerifyMessage(_hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _VerifySignature.Contract.VerifyMessage(&_VerifySignature.CallOpts, _hashedMessage, _v, _r, _s)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x659934c1.
//
// Solidity: function VerifyMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_VerifySignature *VerifySignatureCallerSession) VerifyMessage(_hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _VerifySignature.Contract.VerifyMessage(&_VerifySignature.CallOpts, _hashedMessage, _v, _r, _s)
}
