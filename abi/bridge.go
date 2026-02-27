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
	_ = abi.ConvertType
)

// ERC20BridgeClaim is an auto generated low-level Go binding around an user-defined struct.
type ERC20BridgeClaim struct {
	EvmChainId *big.Int
	Token      common.Address
	To         common.Address
	Amount     *big.Int
	Nonce      *big.Int
}

// ERC20BridgeMetaData contains all meta data concerning the ERC20Bridge contract.
var ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cosmosRecipient\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockedAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processedUnlockNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlock\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structERC20Bridge.Claim\",\"components\":[{\"name\":\"evmChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockedAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Locked\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cosmosRecipient\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unlocked\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BridgeMetaData.ABI instead.
var ERC20BridgeABI = ERC20BridgeMetaData.ABI

// ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type ERC20Bridge struct {
	ERC20BridgeCaller     // Read-only binding to the contract
	ERC20BridgeTransactor // Write-only binding to the contract
	ERC20BridgeFilterer   // Log filterer for contract events
}

// ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BridgeSession struct {
	Contract     *ERC20Bridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BridgeCallerSession struct {
	Contract *ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BridgeTransactorSession struct {
	Contract     *ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BridgeRaw struct {
	Contract *ERC20Bridge // Generic contract binding to access the raw methods on
}

// ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BridgeCallerRaw struct {
	Contract *ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BridgeTransactorRaw struct {
	Contract *ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Bridge creates a new instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20Bridge(address common.Address, backend bind.ContractBackend) (*ERC20Bridge, error) {
	contract, err := bindERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Bridge{ERC20BridgeCaller: ERC20BridgeCaller{contract: contract}, ERC20BridgeTransactor: ERC20BridgeTransactor{contract: contract}, ERC20BridgeFilterer: ERC20BridgeFilterer{contract: contract}}, nil
}

// NewERC20BridgeCaller creates a new read-only instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*ERC20BridgeCaller, error) {
	contract, err := bindERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeCaller{contract: contract}, nil
}

// NewERC20BridgeTransactor creates a new write-only instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BridgeTransactor, error) {
	contract, err := bindERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeTransactor{contract: contract}, nil
}

// NewERC20BridgeFilterer creates a new log filterer instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BridgeFilterer, error) {
	contract, err := bindERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeFilterer{contract: contract}, nil
}

// bindERC20Bridge binds a generic wrapper to an already deployed contract.
func bindERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Bridge *ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Bridge.Contract.ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Bridge *ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Bridge *ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Bridge *ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Bridge *ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Bridge *ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) ChainId() (*big.Int, error) {
	return _ERC20Bridge.Contract.ChainId(&_ERC20Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) ChainId() (*big.Int, error) {
	return _ERC20Bridge.Contract.ChainId(&_ERC20Bridge.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "lockNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) LockNonce() (*big.Int, error) {
	return _ERC20Bridge.Contract.LockNonce(&_ERC20Bridge.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) LockNonce() (*big.Int, error) {
	return _ERC20Bridge.Contract.LockNonce(&_ERC20Bridge.CallOpts)
}

// LockedAmount is a free data retrieval call binding the contract method 0x6ab28bc8.
//
// Solidity: function lockedAmount() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) LockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "lockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedAmount is a free data retrieval call binding the contract method 0x6ab28bc8.
//
// Solidity: function lockedAmount() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) LockedAmount() (*big.Int, error) {
	return _ERC20Bridge.Contract.LockedAmount(&_ERC20Bridge.CallOpts)
}

// LockedAmount is a free data retrieval call binding the contract method 0x6ab28bc8.
//
// Solidity: function lockedAmount() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) LockedAmount() (*big.Int, error) {
	return _ERC20Bridge.Contract.LockedAmount(&_ERC20Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) Owner() (common.Address, error) {
	return _ERC20Bridge.Contract.Owner(&_ERC20Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) Owner() (common.Address, error) {
	return _ERC20Bridge.Contract.Owner(&_ERC20Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) Paused() (bool, error) {
	return _ERC20Bridge.Contract.Paused(&_ERC20Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) Paused() (bool, error) {
	return _ERC20Bridge.Contract.Paused(&_ERC20Bridge.CallOpts)
}

// ProcessedUnlockNonces is a free data retrieval call binding the contract method 0x06c8abcc.
//
// Solidity: function processedUnlockNonces(uint256 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) ProcessedUnlockNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "processedUnlockNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedUnlockNonces is a free data retrieval call binding the contract method 0x06c8abcc.
//
// Solidity: function processedUnlockNonces(uint256 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) ProcessedUnlockNonces(arg0 *big.Int) (bool, error) {
	return _ERC20Bridge.Contract.ProcessedUnlockNonces(&_ERC20Bridge.CallOpts, arg0)
}

// ProcessedUnlockNonces is a free data retrieval call binding the contract method 0x06c8abcc.
//
// Solidity: function processedUnlockNonces(uint256 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) ProcessedUnlockNonces(arg0 *big.Int) (bool, error) {
	return _ERC20Bridge.Contract.ProcessedUnlockNonces(&_ERC20Bridge.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _ERC20Bridge.Contract.SupportedTokens(&_ERC20Bridge.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _ERC20Bridge.Contract.SupportedTokens(&_ERC20Bridge.CallOpts, arg0)
}

// UnlockNonce is a free data retrieval call binding the contract method 0xdd926714.
//
// Solidity: function unlockNonce() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) UnlockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "unlockNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockNonce is a free data retrieval call binding the contract method 0xdd926714.
//
// Solidity: function unlockNonce() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) UnlockNonce() (*big.Int, error) {
	return _ERC20Bridge.Contract.UnlockNonce(&_ERC20Bridge.CallOpts)
}

// UnlockNonce is a free data retrieval call binding the contract method 0xdd926714.
//
// Solidity: function unlockNonce() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) UnlockNonce() (*big.Int, error) {
	return _ERC20Bridge.Contract.UnlockNonce(&_ERC20Bridge.CallOpts)
}

// UnlockedAmount is a free data retrieval call binding the contract method 0x4dfefebc.
//
// Solidity: function unlockedAmount() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) UnlockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "unlockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedAmount is a free data retrieval call binding the contract method 0x4dfefebc.
//
// Solidity: function unlockedAmount() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) UnlockedAmount() (*big.Int, error) {
	return _ERC20Bridge.Contract.UnlockedAmount(&_ERC20Bridge.CallOpts)
}

// UnlockedAmount is a free data retrieval call binding the contract method 0x4dfefebc.
//
// Solidity: function unlockedAmount() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) UnlockedAmount() (*big.Int, error) {
	return _ERC20Bridge.Contract.UnlockedAmount(&_ERC20Bridge.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address token, uint256 amount, string cosmosRecipient) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Lock(opts *bind.TransactOpts, token common.Address, amount *big.Int, cosmosRecipient string) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "lock", token, amount, cosmosRecipient)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address token, uint256 amount, string cosmosRecipient) returns()
func (_ERC20Bridge *ERC20BridgeSession) Lock(token common.Address, amount *big.Int, cosmosRecipient string) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Lock(&_ERC20Bridge.TransactOpts, token, amount, cosmosRecipient)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address token, uint256 amount, string cosmosRecipient) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Lock(token common.Address, amount *big.Int, cosmosRecipient string) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Lock(&_ERC20Bridge.TransactOpts, token, amount, cosmosRecipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Bridge *ERC20BridgeSession) Pause() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Pause(&_ERC20Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Pause(&_ERC20Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Bridge *ERC20BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Bridge *ERC20BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.RenounceOwnership(&_ERC20Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.RenounceOwnership(&_ERC20Bridge.TransactOpts)
}

// SetSupportedToken is a paid mutator transaction binding the contract method 0xe7986466.
//
// Solidity: function setSupportedToken(address token, bool status) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) SetSupportedToken(opts *bind.TransactOpts, token common.Address, status bool) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "setSupportedToken", token, status)
}

// SetSupportedToken is a paid mutator transaction binding the contract method 0xe7986466.
//
// Solidity: function setSupportedToken(address token, bool status) returns()
func (_ERC20Bridge *ERC20BridgeSession) SetSupportedToken(token common.Address, status bool) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetSupportedToken(&_ERC20Bridge.TransactOpts, token, status)
}

// SetSupportedToken is a paid mutator transaction binding the contract method 0xe7986466.
//
// Solidity: function setSupportedToken(address token, bool status) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) SetSupportedToken(token common.Address, status bool) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetSupportedToken(&_ERC20Bridge.TransactOpts, token, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Bridge *ERC20BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.TransferOwnership(&_ERC20Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.TransferOwnership(&_ERC20Bridge.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xf45a12ac.
//
// Solidity: function unlock((uint256,address,address,uint256,uint256) claim) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Unlock(opts *bind.TransactOpts, claim ERC20BridgeClaim) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "unlock", claim)
}

// Unlock is a paid mutator transaction binding the contract method 0xf45a12ac.
//
// Solidity: function unlock((uint256,address,address,uint256,uint256) claim) returns()
func (_ERC20Bridge *ERC20BridgeSession) Unlock(claim ERC20BridgeClaim) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Unlock(&_ERC20Bridge.TransactOpts, claim)
}

// Unlock is a paid mutator transaction binding the contract method 0xf45a12ac.
//
// Solidity: function unlock((uint256,address,address,uint256,uint256) claim) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Unlock(claim ERC20BridgeClaim) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Unlock(&_ERC20Bridge.TransactOpts, claim)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Bridge *ERC20BridgeSession) Unpause() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Unpause(&_ERC20Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Unpause(&_ERC20Bridge.TransactOpts)
}

// ERC20BridgeLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the ERC20Bridge contract.
type ERC20BridgeLockedIterator struct {
	Event *ERC20BridgeLocked // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeLocked)
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
		it.Event = new(ERC20BridgeLocked)
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
func (it *ERC20BridgeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeLocked represents a Locked event raised by the ERC20Bridge contract.
type ERC20BridgeLocked struct {
	Token           common.Address
	Sender          common.Address
	CosmosRecipient string
	Amount          *big.Int
	Nonce           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x78a4518c05c7ababe3937649e215d6c585b14a47a7936d03ebf0ec95544d5167.
//
// Solidity: event Locked(address indexed token, address indexed sender, string cosmosRecipient, uint256 amount, uint256 nonce)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterLocked(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*ERC20BridgeLockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Locked", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeLockedIterator{contract: _ERC20Bridge.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x78a4518c05c7ababe3937649e215d6c585b14a47a7936d03ebf0ec95544d5167.
//
// Solidity: event Locked(address indexed token, address indexed sender, string cosmosRecipient, uint256 amount, uint256 nonce)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *ERC20BridgeLocked, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Locked", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeLocked)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x78a4518c05c7ababe3937649e215d6c585b14a47a7936d03ebf0ec95544d5167.
//
// Solidity: event Locked(address indexed token, address indexed sender, string cosmosRecipient, uint256 amount, uint256 nonce)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseLocked(log types.Log) (*ERC20BridgeLocked, error) {
	event := new(ERC20BridgeLocked)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20Bridge contract.
type ERC20BridgeOwnershipTransferredIterator struct {
	Event *ERC20BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeOwnershipTransferred)
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
		it.Event = new(ERC20BridgeOwnershipTransferred)
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
func (it *ERC20BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Bridge contract.
type ERC20BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeOwnershipTransferredIterator{contract: _ERC20Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeOwnershipTransferred)
				if err := _ERC20Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20Bridge *ERC20BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20BridgeOwnershipTransferred, error) {
	event := new(ERC20BridgeOwnershipTransferred)
	if err := _ERC20Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Bridge contract.
type ERC20BridgePausedIterator struct {
	Event *ERC20BridgePaused // Event containing the contract specifics and raw log

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
func (it *ERC20BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgePaused)
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
		it.Event = new(ERC20BridgePaused)
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
func (it *ERC20BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgePaused represents a Paused event raised by the ERC20Bridge contract.
type ERC20BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20BridgePausedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgePausedIterator{contract: _ERC20Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20BridgePaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgePaused)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Bridge *ERC20BridgeFilterer) ParsePaused(log types.Log) (*ERC20BridgePaused, error) {
	event := new(ERC20BridgePaused)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the ERC20Bridge contract.
type ERC20BridgeUnlockedIterator struct {
	Event *ERC20BridgeUnlocked // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeUnlocked)
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
		it.Event = new(ERC20BridgeUnlocked)
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
func (it *ERC20BridgeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeUnlocked represents a Unlocked event raised by the ERC20Bridge contract.
type ERC20BridgeUnlocked struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x711dd5075f1013614358f2433df9598bb468e809350a81120d25be184e3a04af.
//
// Solidity: event Unlocked(address indexed token, address indexed to, uint256 amount, uint256 nonce)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterUnlocked(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ERC20BridgeUnlockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Unlocked", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeUnlockedIterator{contract: _ERC20Bridge.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x711dd5075f1013614358f2433df9598bb468e809350a81120d25be184e3a04af.
//
// Solidity: event Unlocked(address indexed token, address indexed to, uint256 amount, uint256 nonce)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *ERC20BridgeUnlocked, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Unlocked", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeUnlocked)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x711dd5075f1013614358f2433df9598bb468e809350a81120d25be184e3a04af.
//
// Solidity: event Unlocked(address indexed token, address indexed to, uint256 amount, uint256 nonce)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseUnlocked(log types.Log) (*ERC20BridgeUnlocked, error) {
	event := new(ERC20BridgeUnlocked)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20Bridge contract.
type ERC20BridgeUnpausedIterator struct {
	Event *ERC20BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeUnpaused)
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
		it.Event = new(ERC20BridgeUnpaused)
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
func (it *ERC20BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeUnpaused represents a Unpaused event raised by the ERC20Bridge contract.
type ERC20BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20BridgeUnpausedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeUnpausedIterator{contract: _ERC20Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeUnpaused)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseUnpaused(log types.Log) (*ERC20BridgeUnpaused, error) {
	event := new(ERC20BridgeUnpaused)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
