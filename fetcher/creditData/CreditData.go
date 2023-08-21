// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditData

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

// CreditDataMetaData contains all meta data concerning the CreditData contract.
var CreditDataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"whom\",\"type\":\"address[]\"}],\"name\":\"CombinedScoreReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"ReportReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"whom\",\"type\":\"address[]\"}],\"name\":\"RequestCombinedScore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"RequestReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"RequestScore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"ScoreReady\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_disclosure\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sbt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearCreditReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearCreditScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"getCreditReport\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"getCreditScore\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"whom\",\"type\":\"address[]\"}],\"name\":\"requestCombinedCreditScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"requestCreditReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"requestCreditScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"setDisclosure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"setSbt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"setWriter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"whom\",\"type\":\"address[]\"},{\"internalType\":\"uint16\",\"name\":\"score\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"storeCombinedCreditScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"report\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"storeCreditReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"score\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"storeCreditScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CreditDataABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditDataMetaData.ABI instead.
var CreditDataABI = CreditDataMetaData.ABI

// CreditData is an auto generated Go binding around an Ethereum contract.
type CreditData struct {
	CreditDataCaller     // Read-only binding to the contract
	CreditDataTransactor // Write-only binding to the contract
	CreditDataFilterer   // Log filterer for contract events
}

// CreditDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditDataSession struct {
	Contract     *CreditData       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditDataCallerSession struct {
	Contract *CreditDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CreditDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditDataTransactorSession struct {
	Contract     *CreditDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CreditDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditDataRaw struct {
	Contract *CreditData // Generic contract binding to access the raw methods on
}

// CreditDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditDataCallerRaw struct {
	Contract *CreditDataCaller // Generic read-only contract binding to access the raw methods on
}

// CreditDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditDataTransactorRaw struct {
	Contract *CreditDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditData creates a new instance of CreditData, bound to a specific deployed contract.
func NewCreditData(address common.Address, backend bind.ContractBackend) (*CreditData, error) {
	contract, err := bindCreditData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditData{CreditDataCaller: CreditDataCaller{contract: contract}, CreditDataTransactor: CreditDataTransactor{contract: contract}, CreditDataFilterer: CreditDataFilterer{contract: contract}}, nil
}

// NewCreditDataCaller creates a new read-only instance of CreditData, bound to a specific deployed contract.
func NewCreditDataCaller(address common.Address, caller bind.ContractCaller) (*CreditDataCaller, error) {
	contract, err := bindCreditData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditDataCaller{contract: contract}, nil
}

// NewCreditDataTransactor creates a new write-only instance of CreditData, bound to a specific deployed contract.
func NewCreditDataTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditDataTransactor, error) {
	contract, err := bindCreditData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditDataTransactor{contract: contract}, nil
}

// NewCreditDataFilterer creates a new log filterer instance of CreditData, bound to a specific deployed contract.
func NewCreditDataFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditDataFilterer, error) {
	contract, err := bindCreditData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditDataFilterer{contract: contract}, nil
}

// bindCreditData binds a generic wrapper to an already deployed contract.
func bindCreditData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditDataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditData *CreditDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditData.Contract.CreditDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditData *CreditDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditData.Contract.CreditDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditData *CreditDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditData.Contract.CreditDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditData *CreditDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditData *CreditDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditData *CreditDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditData.Contract.contract.Transact(opts, method, params...)
}

// Disclosure is a free data retrieval call binding the contract method 0xe8e3005a.
//
// Solidity: function _disclosure() view returns(address)
func (_CreditData *CreditDataCaller) Disclosure(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditData.contract.Call(opts, &out, "_disclosure")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Disclosure is a free data retrieval call binding the contract method 0xe8e3005a.
//
// Solidity: function _disclosure() view returns(address)
func (_CreditData *CreditDataSession) Disclosure() (common.Address, error) {
	return _CreditData.Contract.Disclosure(&_CreditData.CallOpts)
}

// Disclosure is a free data retrieval call binding the contract method 0xe8e3005a.
//
// Solidity: function _disclosure() view returns(address)
func (_CreditData *CreditDataCallerSession) Disclosure() (common.Address, error) {
	return _CreditData.Contract.Disclosure(&_CreditData.CallOpts)
}

// Sbt is a free data retrieval call binding the contract method 0xce6ecc0d.
//
// Solidity: function _sbt() view returns(address)
func (_CreditData *CreditDataCaller) Sbt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditData.contract.Call(opts, &out, "_sbt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sbt is a free data retrieval call binding the contract method 0xce6ecc0d.
//
// Solidity: function _sbt() view returns(address)
func (_CreditData *CreditDataSession) Sbt() (common.Address, error) {
	return _CreditData.Contract.Sbt(&_CreditData.CallOpts)
}

// Sbt is a free data retrieval call binding the contract method 0xce6ecc0d.
//
// Solidity: function _sbt() view returns(address)
func (_CreditData *CreditDataCallerSession) Sbt() (common.Address, error) {
	return _CreditData.Contract.Sbt(&_CreditData.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x4c04896f.
//
// Solidity: function _writer() view returns(address)
func (_CreditData *CreditDataCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditData.contract.Call(opts, &out, "_writer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Writer is a free data retrieval call binding the contract method 0x4c04896f.
//
// Solidity: function _writer() view returns(address)
func (_CreditData *CreditDataSession) Writer() (common.Address, error) {
	return _CreditData.Contract.Writer(&_CreditData.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x4c04896f.
//
// Solidity: function _writer() view returns(address)
func (_CreditData *CreditDataCallerSession) Writer() (common.Address, error) {
	return _CreditData.Contract.Writer(&_CreditData.CallOpts)
}

// GetCreditReport is a free data retrieval call binding the contract method 0xfd94ab3d.
//
// Solidity: function getCreditReport(address whom) view returns(string, uint256)
func (_CreditData *CreditDataCaller) GetCreditReport(opts *bind.CallOpts, whom common.Address) (string, *big.Int, error) {
	var out []interface{}
	err := _CreditData.contract.Call(opts, &out, "getCreditReport", whom)

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCreditReport is a free data retrieval call binding the contract method 0xfd94ab3d.
//
// Solidity: function getCreditReport(address whom) view returns(string, uint256)
func (_CreditData *CreditDataSession) GetCreditReport(whom common.Address) (string, *big.Int, error) {
	return _CreditData.Contract.GetCreditReport(&_CreditData.CallOpts, whom)
}

// GetCreditReport is a free data retrieval call binding the contract method 0xfd94ab3d.
//
// Solidity: function getCreditReport(address whom) view returns(string, uint256)
func (_CreditData *CreditDataCallerSession) GetCreditReport(whom common.Address) (string, *big.Int, error) {
	return _CreditData.Contract.GetCreditReport(&_CreditData.CallOpts, whom)
}

// GetCreditScore is a free data retrieval call binding the contract method 0xd3dd2bdf.
//
// Solidity: function getCreditScore(address whom) view returns(uint16, uint256, address[])
func (_CreditData *CreditDataCaller) GetCreditScore(opts *bind.CallOpts, whom common.Address) (uint16, *big.Int, []common.Address, error) {
	var out []interface{}
	err := _CreditData.contract.Call(opts, &out, "getCreditScore", whom)

	if err != nil {
		return *new(uint16), *new(*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return out0, out1, out2, err

}

// GetCreditScore is a free data retrieval call binding the contract method 0xd3dd2bdf.
//
// Solidity: function getCreditScore(address whom) view returns(uint16, uint256, address[])
func (_CreditData *CreditDataSession) GetCreditScore(whom common.Address) (uint16, *big.Int, []common.Address, error) {
	return _CreditData.Contract.GetCreditScore(&_CreditData.CallOpts, whom)
}

// GetCreditScore is a free data retrieval call binding the contract method 0xd3dd2bdf.
//
// Solidity: function getCreditScore(address whom) view returns(uint16, uint256, address[])
func (_CreditData *CreditDataCallerSession) GetCreditScore(whom common.Address) (uint16, *big.Int, []common.Address, error) {
	return _CreditData.Contract.GetCreditScore(&_CreditData.CallOpts, whom)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreditData *CreditDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditData.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreditData *CreditDataSession) Owner() (common.Address, error) {
	return _CreditData.Contract.Owner(&_CreditData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreditData *CreditDataCallerSession) Owner() (common.Address, error) {
	return _CreditData.Contract.Owner(&_CreditData.CallOpts)
}

// ClearCreditReport is a paid mutator transaction binding the contract method 0xa4478d70.
//
// Solidity: function clearCreditReport() returns()
func (_CreditData *CreditDataTransactor) ClearCreditReport(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "clearCreditReport")
}

// ClearCreditReport is a paid mutator transaction binding the contract method 0xa4478d70.
//
// Solidity: function clearCreditReport() returns()
func (_CreditData *CreditDataSession) ClearCreditReport() (*types.Transaction, error) {
	return _CreditData.Contract.ClearCreditReport(&_CreditData.TransactOpts)
}

// ClearCreditReport is a paid mutator transaction binding the contract method 0xa4478d70.
//
// Solidity: function clearCreditReport() returns()
func (_CreditData *CreditDataTransactorSession) ClearCreditReport() (*types.Transaction, error) {
	return _CreditData.Contract.ClearCreditReport(&_CreditData.TransactOpts)
}

// ClearCreditScore is a paid mutator transaction binding the contract method 0x44efed47.
//
// Solidity: function clearCreditScore() returns()
func (_CreditData *CreditDataTransactor) ClearCreditScore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "clearCreditScore")
}

// ClearCreditScore is a paid mutator transaction binding the contract method 0x44efed47.
//
// Solidity: function clearCreditScore() returns()
func (_CreditData *CreditDataSession) ClearCreditScore() (*types.Transaction, error) {
	return _CreditData.Contract.ClearCreditScore(&_CreditData.TransactOpts)
}

// ClearCreditScore is a paid mutator transaction binding the contract method 0x44efed47.
//
// Solidity: function clearCreditScore() returns()
func (_CreditData *CreditDataTransactorSession) ClearCreditScore() (*types.Transaction, error) {
	return _CreditData.Contract.ClearCreditScore(&_CreditData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreditData *CreditDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreditData *CreditDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreditData.Contract.RenounceOwnership(&_CreditData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreditData *CreditDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreditData.Contract.RenounceOwnership(&_CreditData.TransactOpts)
}

// RequestCombinedCreditScore is a paid mutator transaction binding the contract method 0x72f26c0f.
//
// Solidity: function requestCombinedCreditScore(address[] whom) returns()
func (_CreditData *CreditDataTransactor) RequestCombinedCreditScore(opts *bind.TransactOpts, whom []common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "requestCombinedCreditScore", whom)
}

// RequestCombinedCreditScore is a paid mutator transaction binding the contract method 0x72f26c0f.
//
// Solidity: function requestCombinedCreditScore(address[] whom) returns()
func (_CreditData *CreditDataSession) RequestCombinedCreditScore(whom []common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.RequestCombinedCreditScore(&_CreditData.TransactOpts, whom)
}

// RequestCombinedCreditScore is a paid mutator transaction binding the contract method 0x72f26c0f.
//
// Solidity: function requestCombinedCreditScore(address[] whom) returns()
func (_CreditData *CreditDataTransactorSession) RequestCombinedCreditScore(whom []common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.RequestCombinedCreditScore(&_CreditData.TransactOpts, whom)
}

// RequestCreditReport is a paid mutator transaction binding the contract method 0x7ba474f4.
//
// Solidity: function requestCreditReport(address whom) returns()
func (_CreditData *CreditDataTransactor) RequestCreditReport(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "requestCreditReport", whom)
}

// RequestCreditReport is a paid mutator transaction binding the contract method 0x7ba474f4.
//
// Solidity: function requestCreditReport(address whom) returns()
func (_CreditData *CreditDataSession) RequestCreditReport(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.RequestCreditReport(&_CreditData.TransactOpts, whom)
}

// RequestCreditReport is a paid mutator transaction binding the contract method 0x7ba474f4.
//
// Solidity: function requestCreditReport(address whom) returns()
func (_CreditData *CreditDataTransactorSession) RequestCreditReport(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.RequestCreditReport(&_CreditData.TransactOpts, whom)
}

// RequestCreditScore is a paid mutator transaction binding the contract method 0x2684f936.
//
// Solidity: function requestCreditScore(address whom) returns()
func (_CreditData *CreditDataTransactor) RequestCreditScore(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "requestCreditScore", whom)
}

// RequestCreditScore is a paid mutator transaction binding the contract method 0x2684f936.
//
// Solidity: function requestCreditScore(address whom) returns()
func (_CreditData *CreditDataSession) RequestCreditScore(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.RequestCreditScore(&_CreditData.TransactOpts, whom)
}

// RequestCreditScore is a paid mutator transaction binding the contract method 0x2684f936.
//
// Solidity: function requestCreditScore(address whom) returns()
func (_CreditData *CreditDataTransactorSession) RequestCreditScore(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.RequestCreditScore(&_CreditData.TransactOpts, whom)
}

// SetDisclosure is a paid mutator transaction binding the contract method 0xd0cf4a5b.
//
// Solidity: function setDisclosure(address whom) returns()
func (_CreditData *CreditDataTransactor) SetDisclosure(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "setDisclosure", whom)
}

// SetDisclosure is a paid mutator transaction binding the contract method 0xd0cf4a5b.
//
// Solidity: function setDisclosure(address whom) returns()
func (_CreditData *CreditDataSession) SetDisclosure(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.SetDisclosure(&_CreditData.TransactOpts, whom)
}

// SetDisclosure is a paid mutator transaction binding the contract method 0xd0cf4a5b.
//
// Solidity: function setDisclosure(address whom) returns()
func (_CreditData *CreditDataTransactorSession) SetDisclosure(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.SetDisclosure(&_CreditData.TransactOpts, whom)
}

// SetSbt is a paid mutator transaction binding the contract method 0x228dc360.
//
// Solidity: function setSbt(address whom) returns()
func (_CreditData *CreditDataTransactor) SetSbt(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "setSbt", whom)
}

// SetSbt is a paid mutator transaction binding the contract method 0x228dc360.
//
// Solidity: function setSbt(address whom) returns()
func (_CreditData *CreditDataSession) SetSbt(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.SetSbt(&_CreditData.TransactOpts, whom)
}

// SetSbt is a paid mutator transaction binding the contract method 0x228dc360.
//
// Solidity: function setSbt(address whom) returns()
func (_CreditData *CreditDataTransactorSession) SetSbt(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.SetSbt(&_CreditData.TransactOpts, whom)
}

// SetWriter is a paid mutator transaction binding the contract method 0x39e20523.
//
// Solidity: function setWriter(address whom) returns()
func (_CreditData *CreditDataTransactor) SetWriter(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "setWriter", whom)
}

// SetWriter is a paid mutator transaction binding the contract method 0x39e20523.
//
// Solidity: function setWriter(address whom) returns()
func (_CreditData *CreditDataSession) SetWriter(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.SetWriter(&_CreditData.TransactOpts, whom)
}

// SetWriter is a paid mutator transaction binding the contract method 0x39e20523.
//
// Solidity: function setWriter(address whom) returns()
func (_CreditData *CreditDataTransactorSession) SetWriter(whom common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.SetWriter(&_CreditData.TransactOpts, whom)
}

// StoreCombinedCreditScore is a paid mutator transaction binding the contract method 0xbe68d212.
//
// Solidity: function storeCombinedCreditScore(address[] whom, uint16 score, uint256 expiry) returns()
func (_CreditData *CreditDataTransactor) StoreCombinedCreditScore(opts *bind.TransactOpts, whom []common.Address, score uint16, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "storeCombinedCreditScore", whom, score, expiry)
}

// StoreCombinedCreditScore is a paid mutator transaction binding the contract method 0xbe68d212.
//
// Solidity: function storeCombinedCreditScore(address[] whom, uint16 score, uint256 expiry) returns()
func (_CreditData *CreditDataSession) StoreCombinedCreditScore(whom []common.Address, score uint16, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.Contract.StoreCombinedCreditScore(&_CreditData.TransactOpts, whom, score, expiry)
}

// StoreCombinedCreditScore is a paid mutator transaction binding the contract method 0xbe68d212.
//
// Solidity: function storeCombinedCreditScore(address[] whom, uint16 score, uint256 expiry) returns()
func (_CreditData *CreditDataTransactorSession) StoreCombinedCreditScore(whom []common.Address, score uint16, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.Contract.StoreCombinedCreditScore(&_CreditData.TransactOpts, whom, score, expiry)
}

// StoreCreditReport is a paid mutator transaction binding the contract method 0xb5fdc555.
//
// Solidity: function storeCreditReport(address whom, string report, uint256 expiry) returns()
func (_CreditData *CreditDataTransactor) StoreCreditReport(opts *bind.TransactOpts, whom common.Address, report string, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "storeCreditReport", whom, report, expiry)
}

// StoreCreditReport is a paid mutator transaction binding the contract method 0xb5fdc555.
//
// Solidity: function storeCreditReport(address whom, string report, uint256 expiry) returns()
func (_CreditData *CreditDataSession) StoreCreditReport(whom common.Address, report string, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.Contract.StoreCreditReport(&_CreditData.TransactOpts, whom, report, expiry)
}

// StoreCreditReport is a paid mutator transaction binding the contract method 0xb5fdc555.
//
// Solidity: function storeCreditReport(address whom, string report, uint256 expiry) returns()
func (_CreditData *CreditDataTransactorSession) StoreCreditReport(whom common.Address, report string, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.Contract.StoreCreditReport(&_CreditData.TransactOpts, whom, report, expiry)
}

// StoreCreditScore is a paid mutator transaction binding the contract method 0x183090fb.
//
// Solidity: function storeCreditScore(address whom, uint16 score, uint256 expiry) returns()
func (_CreditData *CreditDataTransactor) StoreCreditScore(opts *bind.TransactOpts, whom common.Address, score uint16, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "storeCreditScore", whom, score, expiry)
}

// StoreCreditScore is a paid mutator transaction binding the contract method 0x183090fb.
//
// Solidity: function storeCreditScore(address whom, uint16 score, uint256 expiry) returns()
func (_CreditData *CreditDataSession) StoreCreditScore(whom common.Address, score uint16, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.Contract.StoreCreditScore(&_CreditData.TransactOpts, whom, score, expiry)
}

// StoreCreditScore is a paid mutator transaction binding the contract method 0x183090fb.
//
// Solidity: function storeCreditScore(address whom, uint16 score, uint256 expiry) returns()
func (_CreditData *CreditDataTransactorSession) StoreCreditScore(whom common.Address, score uint16, expiry *big.Int) (*types.Transaction, error) {
	return _CreditData.Contract.StoreCreditScore(&_CreditData.TransactOpts, whom, score, expiry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreditData *CreditDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CreditData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreditData *CreditDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.TransferOwnership(&_CreditData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreditData *CreditDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreditData.Contract.TransferOwnership(&_CreditData.TransactOpts, newOwner)
}

// CreditDataCombinedScoreReadyIterator is returned from FilterCombinedScoreReady and is used to iterate over the raw logs and unpacked data for CombinedScoreReady events raised by the CreditData contract.
type CreditDataCombinedScoreReadyIterator struct {
	Event *CreditDataCombinedScoreReady // Event containing the contract specifics and raw log

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
func (it *CreditDataCombinedScoreReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataCombinedScoreReady)
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
		it.Event = new(CreditDataCombinedScoreReady)
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
func (it *CreditDataCombinedScoreReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataCombinedScoreReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataCombinedScoreReady represents a CombinedScoreReady event raised by the CreditData contract.
type CreditDataCombinedScoreReady struct {
	Whom []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCombinedScoreReady is a free log retrieval operation binding the contract event 0x23dc457501204682403cb2911e231f6302c96c9adc4deab8f4a1d030292e6b11.
//
// Solidity: event CombinedScoreReady(address[] whom)
func (_CreditData *CreditDataFilterer) FilterCombinedScoreReady(opts *bind.FilterOpts) (*CreditDataCombinedScoreReadyIterator, error) {

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "CombinedScoreReady")
	if err != nil {
		return nil, err
	}
	return &CreditDataCombinedScoreReadyIterator{contract: _CreditData.contract, event: "CombinedScoreReady", logs: logs, sub: sub}, nil
}

// WatchCombinedScoreReady is a free log subscription operation binding the contract event 0x23dc457501204682403cb2911e231f6302c96c9adc4deab8f4a1d030292e6b11.
//
// Solidity: event CombinedScoreReady(address[] whom)
func (_CreditData *CreditDataFilterer) WatchCombinedScoreReady(opts *bind.WatchOpts, sink chan<- *CreditDataCombinedScoreReady) (event.Subscription, error) {

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "CombinedScoreReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataCombinedScoreReady)
				if err := _CreditData.contract.UnpackLog(event, "CombinedScoreReady", log); err != nil {
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

// ParseCombinedScoreReady is a log parse operation binding the contract event 0x23dc457501204682403cb2911e231f6302c96c9adc4deab8f4a1d030292e6b11.
//
// Solidity: event CombinedScoreReady(address[] whom)
func (_CreditData *CreditDataFilterer) ParseCombinedScoreReady(log types.Log) (*CreditDataCombinedScoreReady, error) {
	event := new(CreditDataCombinedScoreReady)
	if err := _CreditData.contract.UnpackLog(event, "CombinedScoreReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CreditData contract.
type CreditDataOwnershipTransferredIterator struct {
	Event *CreditDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreditDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataOwnershipTransferred)
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
		it.Event = new(CreditDataOwnershipTransferred)
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
func (it *CreditDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataOwnershipTransferred represents a OwnershipTransferred event raised by the CreditData contract.
type CreditDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreditData *CreditDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreditDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreditDataOwnershipTransferredIterator{contract: _CreditData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreditData *CreditDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreditDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataOwnershipTransferred)
				if err := _CreditData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CreditData *CreditDataFilterer) ParseOwnershipTransferred(log types.Log) (*CreditDataOwnershipTransferred, error) {
	event := new(CreditDataOwnershipTransferred)
	if err := _CreditData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditDataReportReadyIterator is returned from FilterReportReady and is used to iterate over the raw logs and unpacked data for ReportReady events raised by the CreditData contract.
type CreditDataReportReadyIterator struct {
	Event *CreditDataReportReady // Event containing the contract specifics and raw log

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
func (it *CreditDataReportReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataReportReady)
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
		it.Event = new(CreditDataReportReady)
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
func (it *CreditDataReportReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataReportReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataReportReady represents a ReportReady event raised by the CreditData contract.
type CreditDataReportReady struct {
	Whom common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReportReady is a free log retrieval operation binding the contract event 0x13028f428eca8468fdfde5aef5e9db553805d34bdffd0f6b909f163756851420.
//
// Solidity: event ReportReady(address whom)
func (_CreditData *CreditDataFilterer) FilterReportReady(opts *bind.FilterOpts) (*CreditDataReportReadyIterator, error) {

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "ReportReady")
	if err != nil {
		return nil, err
	}
	return &CreditDataReportReadyIterator{contract: _CreditData.contract, event: "ReportReady", logs: logs, sub: sub}, nil
}

// WatchReportReady is a free log subscription operation binding the contract event 0x13028f428eca8468fdfde5aef5e9db553805d34bdffd0f6b909f163756851420.
//
// Solidity: event ReportReady(address whom)
func (_CreditData *CreditDataFilterer) WatchReportReady(opts *bind.WatchOpts, sink chan<- *CreditDataReportReady) (event.Subscription, error) {

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "ReportReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataReportReady)
				if err := _CreditData.contract.UnpackLog(event, "ReportReady", log); err != nil {
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

// ParseReportReady is a log parse operation binding the contract event 0x13028f428eca8468fdfde5aef5e9db553805d34bdffd0f6b909f163756851420.
//
// Solidity: event ReportReady(address whom)
func (_CreditData *CreditDataFilterer) ParseReportReady(log types.Log) (*CreditDataReportReady, error) {
	event := new(CreditDataReportReady)
	if err := _CreditData.contract.UnpackLog(event, "ReportReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditDataRequestCombinedScoreIterator is returned from FilterRequestCombinedScore and is used to iterate over the raw logs and unpacked data for RequestCombinedScore events raised by the CreditData contract.
type CreditDataRequestCombinedScoreIterator struct {
	Event *CreditDataRequestCombinedScore // Event containing the contract specifics and raw log

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
func (it *CreditDataRequestCombinedScoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataRequestCombinedScore)
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
		it.Event = new(CreditDataRequestCombinedScore)
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
func (it *CreditDataRequestCombinedScoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataRequestCombinedScoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataRequestCombinedScore represents a RequestCombinedScore event raised by the CreditData contract.
type CreditDataRequestCombinedScore struct {
	Whom []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestCombinedScore is a free log retrieval operation binding the contract event 0x1f1a6e4720813f454c338d7c777dd2209479b9041b4b48c7e537126c7556790c.
//
// Solidity: event RequestCombinedScore(address[] whom)
func (_CreditData *CreditDataFilterer) FilterRequestCombinedScore(opts *bind.FilterOpts) (*CreditDataRequestCombinedScoreIterator, error) {

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "RequestCombinedScore")
	if err != nil {
		return nil, err
	}
	return &CreditDataRequestCombinedScoreIterator{contract: _CreditData.contract, event: "RequestCombinedScore", logs: logs, sub: sub}, nil
}

// WatchRequestCombinedScore is a free log subscription operation binding the contract event 0x1f1a6e4720813f454c338d7c777dd2209479b9041b4b48c7e537126c7556790c.
//
// Solidity: event RequestCombinedScore(address[] whom)
func (_CreditData *CreditDataFilterer) WatchRequestCombinedScore(opts *bind.WatchOpts, sink chan<- *CreditDataRequestCombinedScore) (event.Subscription, error) {

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "RequestCombinedScore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataRequestCombinedScore)
				if err := _CreditData.contract.UnpackLog(event, "RequestCombinedScore", log); err != nil {
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

// ParseRequestCombinedScore is a log parse operation binding the contract event 0x1f1a6e4720813f454c338d7c777dd2209479b9041b4b48c7e537126c7556790c.
//
// Solidity: event RequestCombinedScore(address[] whom)
func (_CreditData *CreditDataFilterer) ParseRequestCombinedScore(log types.Log) (*CreditDataRequestCombinedScore, error) {
	event := new(CreditDataRequestCombinedScore)
	if err := _CreditData.contract.UnpackLog(event, "RequestCombinedScore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditDataRequestReportIterator is returned from FilterRequestReport and is used to iterate over the raw logs and unpacked data for RequestReport events raised by the CreditData contract.
type CreditDataRequestReportIterator struct {
	Event *CreditDataRequestReport // Event containing the contract specifics and raw log

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
func (it *CreditDataRequestReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataRequestReport)
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
		it.Event = new(CreditDataRequestReport)
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
func (it *CreditDataRequestReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataRequestReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataRequestReport represents a RequestReport event raised by the CreditData contract.
type CreditDataRequestReport struct {
	Whom common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestReport is a free log retrieval operation binding the contract event 0x7dee5acec616daf5cfb291b08b352fef9d591af1d65989cea96d7737af09a579.
//
// Solidity: event RequestReport(address whom)
func (_CreditData *CreditDataFilterer) FilterRequestReport(opts *bind.FilterOpts) (*CreditDataRequestReportIterator, error) {

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "RequestReport")
	if err != nil {
		return nil, err
	}
	return &CreditDataRequestReportIterator{contract: _CreditData.contract, event: "RequestReport", logs: logs, sub: sub}, nil
}

// WatchRequestReport is a free log subscription operation binding the contract event 0x7dee5acec616daf5cfb291b08b352fef9d591af1d65989cea96d7737af09a579.
//
// Solidity: event RequestReport(address whom)
func (_CreditData *CreditDataFilterer) WatchRequestReport(opts *bind.WatchOpts, sink chan<- *CreditDataRequestReport) (event.Subscription, error) {

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "RequestReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataRequestReport)
				if err := _CreditData.contract.UnpackLog(event, "RequestReport", log); err != nil {
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

// ParseRequestReport is a log parse operation binding the contract event 0x7dee5acec616daf5cfb291b08b352fef9d591af1d65989cea96d7737af09a579.
//
// Solidity: event RequestReport(address whom)
func (_CreditData *CreditDataFilterer) ParseRequestReport(log types.Log) (*CreditDataRequestReport, error) {
	event := new(CreditDataRequestReport)
	if err := _CreditData.contract.UnpackLog(event, "RequestReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditDataRequestScoreIterator is returned from FilterRequestScore and is used to iterate over the raw logs and unpacked data for RequestScore events raised by the CreditData contract.
type CreditDataRequestScoreIterator struct {
	Event *CreditDataRequestScore // Event containing the contract specifics and raw log

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
func (it *CreditDataRequestScoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataRequestScore)
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
		it.Event = new(CreditDataRequestScore)
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
func (it *CreditDataRequestScoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataRequestScoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataRequestScore represents a RequestScore event raised by the CreditData contract.
type CreditDataRequestScore struct {
	Whom common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestScore is a free log retrieval operation binding the contract event 0x7d52ce2b13e9530cb9fac8e2bc4a9ef0802b03b37c57ba434baa2e4913ce7330.
//
// Solidity: event RequestScore(address whom)
func (_CreditData *CreditDataFilterer) FilterRequestScore(opts *bind.FilterOpts) (*CreditDataRequestScoreIterator, error) {

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "RequestScore")
	if err != nil {
		return nil, err
	}
	return &CreditDataRequestScoreIterator{contract: _CreditData.contract, event: "RequestScore", logs: logs, sub: sub}, nil
}

// WatchRequestScore is a free log subscription operation binding the contract event 0x7d52ce2b13e9530cb9fac8e2bc4a9ef0802b03b37c57ba434baa2e4913ce7330.
//
// Solidity: event RequestScore(address whom)
func (_CreditData *CreditDataFilterer) WatchRequestScore(opts *bind.WatchOpts, sink chan<- *CreditDataRequestScore) (event.Subscription, error) {

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "RequestScore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataRequestScore)
				if err := _CreditData.contract.UnpackLog(event, "RequestScore", log); err != nil {
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

// ParseRequestScore is a log parse operation binding the contract event 0x7d52ce2b13e9530cb9fac8e2bc4a9ef0802b03b37c57ba434baa2e4913ce7330.
//
// Solidity: event RequestScore(address whom)
func (_CreditData *CreditDataFilterer) ParseRequestScore(log types.Log) (*CreditDataRequestScore, error) {
	event := new(CreditDataRequestScore)
	if err := _CreditData.contract.UnpackLog(event, "RequestScore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditDataScoreReadyIterator is returned from FilterScoreReady and is used to iterate over the raw logs and unpacked data for ScoreReady events raised by the CreditData contract.
type CreditDataScoreReadyIterator struct {
	Event *CreditDataScoreReady // Event containing the contract specifics and raw log

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
func (it *CreditDataScoreReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditDataScoreReady)
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
		it.Event = new(CreditDataScoreReady)
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
func (it *CreditDataScoreReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditDataScoreReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditDataScoreReady represents a ScoreReady event raised by the CreditData contract.
type CreditDataScoreReady struct {
	Whom common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterScoreReady is a free log retrieval operation binding the contract event 0xa49a15c3360c9ca954e8c133b21366f2ae545e0f5e2c6cc81ca9d62e6186ac9b.
//
// Solidity: event ScoreReady(address whom)
func (_CreditData *CreditDataFilterer) FilterScoreReady(opts *bind.FilterOpts) (*CreditDataScoreReadyIterator, error) {

	logs, sub, err := _CreditData.contract.FilterLogs(opts, "ScoreReady")
	if err != nil {
		return nil, err
	}
	return &CreditDataScoreReadyIterator{contract: _CreditData.contract, event: "ScoreReady", logs: logs, sub: sub}, nil
}

// WatchScoreReady is a free log subscription operation binding the contract event 0xa49a15c3360c9ca954e8c133b21366f2ae545e0f5e2c6cc81ca9d62e6186ac9b.
//
// Solidity: event ScoreReady(address whom)
func (_CreditData *CreditDataFilterer) WatchScoreReady(opts *bind.WatchOpts, sink chan<- *CreditDataScoreReady) (event.Subscription, error) {

	logs, sub, err := _CreditData.contract.WatchLogs(opts, "ScoreReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditDataScoreReady)
				if err := _CreditData.contract.UnpackLog(event, "ScoreReady", log); err != nil {
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

// ParseScoreReady is a log parse operation binding the contract event 0xa49a15c3360c9ca954e8c133b21366f2ae545e0f5e2c6cc81ca9d62e6186ac9b.
//
// Solidity: event ScoreReady(address whom)
func (_CreditData *CreditDataFilterer) ParseScoreReady(log types.Log) (*CreditDataScoreReady, error) {
	event := new(CreditDataScoreReady)
	if err := _CreditData.contract.UnpackLog(event, "ScoreReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
