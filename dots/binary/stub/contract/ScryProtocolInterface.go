// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractinterface

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ScryProtocolABI is the input ABI used to generate the binding from.
const ScryProtocolABI = "[{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"despDataId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"supportVerify\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"needVerify\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"AdvancePurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ConfirmPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"FinishPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"VerifiersChosen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metaDataIdEncBuyer\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ReEncrypt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"metaDataIdEncArbitrator\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ArbitrationBegin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"judge\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ArbitrationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"RegisterVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"judge\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"comments\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"VoteResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"VerifierDisable\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"name\":\"proofDataIds\",\"type\":\"bytes32[]\"},{\"name\":\"descDataId\",\"type\":\"string\"},{\"name\":\"supportVerify\",\"type\":\"bool\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"startVerify\",\"type\":\"bool\"}],\"name\":\"advancePurchase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"confirmPurchase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"cancelPurchase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"encryptedMetaDataId\",\"type\":\"bytes\"},{\"name\":\"encryptedMetaDataIds\",\"type\":\"bytes\"}],\"name\":\"reEncrypt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"truth\",\"type\":\"bool\"}],\"name\":\"confirmData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"}],\"name\":\"registerAsVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"judge\",\"type\":\"bool\"},{\"name\":\"comments\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"verifierIndex\",\"type\":\"uint8\"},{\"name\":\"credit\",\"type\":\"uint8\"}],\"name\":\"gradeToVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"judge\",\"type\":\"bool\"}],\"name\":\"arbitrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"getBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"getArbitrators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ScryProtocol is an auto generated Go binding around an Ethereum contract.
type ScryProtocol struct {
	ScryProtocolCaller     // Read-only binding to the contract
	ScryProtocolTransactor // Write-only binding to the contract
	ScryProtocolFilterer   // Log filterer for contract events
}

// ScryProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScryProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScryProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScryProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScryProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScryProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScryProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScryProtocolSession struct {
	Contract     *ScryProtocol     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScryProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScryProtocolCallerSession struct {
	Contract *ScryProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ScryProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScryProtocolTransactorSession struct {
	Contract     *ScryProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ScryProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScryProtocolRaw struct {
	Contract *ScryProtocol // Generic contract binding to access the raw methods on
}

// ScryProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScryProtocolCallerRaw struct {
	Contract *ScryProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ScryProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScryProtocolTransactorRaw struct {
	Contract *ScryProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScryProtocol creates a new instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocol(address common.Address, backend bind.ContractBackend) (*ScryProtocol, error) {
	contract, err := bindScryProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScryProtocol{ScryProtocolCaller: ScryProtocolCaller{contract: contract}, ScryProtocolTransactor: ScryProtocolTransactor{contract: contract}, ScryProtocolFilterer: ScryProtocolFilterer{contract: contract}}, nil
}

// NewScryProtocolCaller creates a new read-only instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocolCaller(address common.Address, caller bind.ContractCaller) (*ScryProtocolCaller, error) {
	contract, err := bindScryProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScryProtocolCaller{contract: contract}, nil
}

// NewScryProtocolTransactor creates a new write-only instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ScryProtocolTransactor, error) {
	contract, err := bindScryProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactor{contract: contract}, nil
}

// NewScryProtocolFilterer creates a new log filterer instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ScryProtocolFilterer, error) {
	contract, err := bindScryProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScryProtocolFilterer{contract: contract}, nil
}

// bindScryProtocol binds a generic wrapper to an already deployed contract.
func bindScryProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScryProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScryProtocol *ScryProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ScryProtocol.Contract.ScryProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScryProtocol *ScryProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ScryProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScryProtocol *ScryProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ScryProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScryProtocol *ScryProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ScryProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScryProtocol *ScryProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScryProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScryProtocol *ScryProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScryProtocol.Contract.contract.Transact(opts, method, params...)
}

// GetArbitrators is a free data retrieval call binding the contract method 0x3297e591.
//
// Solidity: function getArbitrators(uint256 txId) constant returns(address[])
func (_ScryProtocol *ScryProtocolCaller) GetArbitrators(opts *bind.CallOpts, txId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ScryProtocol.contract.Call(opts, out, "getArbitrators", txId)
	return *ret0, err
}

// GetArbitrators is a free data retrieval call binding the contract method 0x3297e591.
//
// Solidity: function getArbitrators(uint256 txId) constant returns(address[])
func (_ScryProtocol *ScryProtocolSession) GetArbitrators(txId *big.Int) ([]common.Address, error) {
	return _ScryProtocol.Contract.GetArbitrators(&_ScryProtocol.CallOpts, txId)
}

// GetArbitrators is a free data retrieval call binding the contract method 0x3297e591.
//
// Solidity: function getArbitrators(uint256 txId) constant returns(address[])
func (_ScryProtocol *ScryProtocolCallerSession) GetArbitrators(txId *big.Int) ([]common.Address, error) {
	return _ScryProtocol.Contract.GetArbitrators(&_ScryProtocol.CallOpts, txId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(uint256 txId) constant returns(address)
func (_ScryProtocol *ScryProtocolCaller) GetBuyer(opts *bind.CallOpts, txId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ScryProtocol.contract.Call(opts, out, "getBuyer", txId)
	return *ret0, err
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(uint256 txId) constant returns(address)
func (_ScryProtocol *ScryProtocolSession) GetBuyer(txId *big.Int) (common.Address, error) {
	return _ScryProtocol.Contract.GetBuyer(&_ScryProtocol.CallOpts, txId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(uint256 txId) constant returns(address)
func (_ScryProtocol *ScryProtocolCallerSession) GetBuyer(txId *big.Int) (common.Address, error) {
	return _ScryProtocol.Contract.GetBuyer(&_ScryProtocol.CallOpts, txId)
}

// AdvancePurchase is a paid mutator transaction binding the contract method 0x435c379c.
//
// Solidity: function advancePurchase(string seqNo, string publishId, bool startVerify) returns()
func (_ScryProtocol *ScryProtocolTransactor) AdvancePurchase(opts *bind.TransactOpts, seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "advancePurchase", seqNo, publishId, startVerify)
}

// AdvancePurchase is a paid mutator transaction binding the contract method 0x435c379c.
//
// Solidity: function advancePurchase(string seqNo, string publishId, bool startVerify) returns()
func (_ScryProtocol *ScryProtocolSession) AdvancePurchase(seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.AdvancePurchase(&_ScryProtocol.TransactOpts, seqNo, publishId, startVerify)
}

// AdvancePurchase is a paid mutator transaction binding the contract method 0x435c379c.
//
// Solidity: function advancePurchase(string seqNo, string publishId, bool startVerify) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) AdvancePurchase(seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.AdvancePurchase(&_ScryProtocol.TransactOpts, seqNo, publishId, startVerify)
}

// Arbitrate is a paid mutator transaction binding the contract method 0xdaff1168.
//
// Solidity: function arbitrate(string seqNo, uint256 txId, bool judge) returns()
func (_ScryProtocol *ScryProtocolTransactor) Arbitrate(opts *bind.TransactOpts, seqNo string, txId *big.Int, judge bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "arbitrate", seqNo, txId, judge)
}

// Arbitrate is a paid mutator transaction binding the contract method 0xdaff1168.
//
// Solidity: function arbitrate(string seqNo, uint256 txId, bool judge) returns()
func (_ScryProtocol *ScryProtocolSession) Arbitrate(seqNo string, txId *big.Int, judge bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Arbitrate(&_ScryProtocol.TransactOpts, seqNo, txId, judge)
}

// Arbitrate is a paid mutator transaction binding the contract method 0xdaff1168.
//
// Solidity: function arbitrate(string seqNo, uint256 txId, bool judge) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) Arbitrate(seqNo string, txId *big.Int, judge bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Arbitrate(&_ScryProtocol.TransactOpts, seqNo, txId, judge)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0xe5a7a627.
//
// Solidity: function cancelPurchase(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactor) CancelPurchase(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "cancelPurchase", seqNo, txId)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0xe5a7a627.
//
// Solidity: function cancelPurchase(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolSession) CancelPurchase(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CancelPurchase(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0xe5a7a627.
//
// Solidity: function cancelPurchase(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CancelPurchase(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CancelPurchase(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// ConfirmData is a paid mutator transaction binding the contract method 0xd49c1d0b.
//
// Solidity: function confirmData(string seqNo, uint256 txId, bool truth) returns()
func (_ScryProtocol *ScryProtocolTransactor) ConfirmData(opts *bind.TransactOpts, seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "confirmData", seqNo, txId, truth)
}

// ConfirmData is a paid mutator transaction binding the contract method 0xd49c1d0b.
//
// Solidity: function confirmData(string seqNo, uint256 txId, bool truth) returns()
func (_ScryProtocol *ScryProtocolSession) ConfirmData(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmData(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// ConfirmData is a paid mutator transaction binding the contract method 0xd49c1d0b.
//
// Solidity: function confirmData(string seqNo, uint256 txId, bool truth) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ConfirmData(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmData(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x8d6278d3.
//
// Solidity: function confirmPurchase(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactor) ConfirmPurchase(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "confirmPurchase", seqNo, txId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x8d6278d3.
//
// Solidity: function confirmPurchase(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolSession) ConfirmPurchase(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmPurchase(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x8d6278d3.
//
// Solidity: function confirmPurchase(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ConfirmPurchase(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmPurchase(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// GradeToVerifier is a paid mutator transaction binding the contract method 0x77b85731.
//
// Solidity: function gradeToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) returns()
func (_ScryProtocol *ScryProtocolTransactor) GradeToVerifier(opts *bind.TransactOpts, seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "gradeToVerifier", seqNo, txId, verifierIndex, credit)
}

// GradeToVerifier is a paid mutator transaction binding the contract method 0x77b85731.
//
// Solidity: function gradeToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) returns()
func (_ScryProtocol *ScryProtocolSession) GradeToVerifier(seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.Contract.GradeToVerifier(&_ScryProtocol.TransactOpts, seqNo, txId, verifierIndex, credit)
}

// GradeToVerifier is a paid mutator transaction binding the contract method 0x77b85731.
//
// Solidity: function gradeToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) GradeToVerifier(seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.Contract.GradeToVerifier(&_ScryProtocol.TransactOpts, seqNo, txId, verifierIndex, credit)
}

// Publish is a paid mutator transaction binding the contract method 0x17925290.
//
// Solidity: function publish(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string descDataId, bool supportVerify) returns()
func (_ScryProtocol *ScryProtocolTransactor) Publish(opts *bind.TransactOpts, seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "publish", seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// Publish is a paid mutator transaction binding the contract method 0x17925290.
//
// Solidity: function publish(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string descDataId, bool supportVerify) returns()
func (_ScryProtocol *ScryProtocolSession) Publish(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Publish(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// Publish is a paid mutator transaction binding the contract method 0x17925290.
//
// Solidity: function publish(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string descDataId, bool supportVerify) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) Publish(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Publish(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// ReEncrypt is a paid mutator transaction binding the contract method 0x310ddffc.
//
// Solidity: function reEncrypt(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) returns()
func (_ScryProtocol *ScryProtocolTransactor) ReEncrypt(opts *bind.TransactOpts, seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "reEncrypt", seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// ReEncrypt is a paid mutator transaction binding the contract method 0x310ddffc.
//
// Solidity: function reEncrypt(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) returns()
func (_ScryProtocol *ScryProtocolSession) ReEncrypt(seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ReEncrypt(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// ReEncrypt is a paid mutator transaction binding the contract method 0x310ddffc.
//
// Solidity: function reEncrypt(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ReEncrypt(seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ReEncrypt(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0x93151dd5.
//
// Solidity: function registerAsVerifier(string seqNo) returns()
func (_ScryProtocol *ScryProtocolTransactor) RegisterAsVerifier(opts *bind.TransactOpts, seqNo string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "registerAsVerifier", seqNo)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0x93151dd5.
//
// Solidity: function registerAsVerifier(string seqNo) returns()
func (_ScryProtocol *ScryProtocolSession) RegisterAsVerifier(seqNo string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.RegisterAsVerifier(&_ScryProtocol.TransactOpts, seqNo)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0x93151dd5.
//
// Solidity: function registerAsVerifier(string seqNo) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) RegisterAsVerifier(seqNo string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.RegisterAsVerifier(&_ScryProtocol.TransactOpts, seqNo)
}

// Vote is a paid mutator transaction binding the contract method 0x980da40d.
//
// Solidity: function vote(string seqNo, uint256 txId, bool judge, string comments) returns()
func (_ScryProtocol *ScryProtocolTransactor) Vote(opts *bind.TransactOpts, seqNo string, txId *big.Int, judge bool, comments string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "vote", seqNo, txId, judge, comments)
}

// Vote is a paid mutator transaction binding the contract method 0x980da40d.
//
// Solidity: function vote(string seqNo, uint256 txId, bool judge, string comments) returns()
func (_ScryProtocol *ScryProtocolSession) Vote(seqNo string, txId *big.Int, judge bool, comments string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Vote(&_ScryProtocol.TransactOpts, seqNo, txId, judge, comments)
}

// Vote is a paid mutator transaction binding the contract method 0x980da40d.
//
// Solidity: function vote(string seqNo, uint256 txId, bool judge, string comments) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) Vote(seqNo string, txId *big.Int, judge bool, comments string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Vote(&_ScryProtocol.TransactOpts, seqNo, txId, judge, comments)
}

// ScryProtocolAdvancePurchaseIterator is returned from FilterAdvancePurchase and is used to iterate over the raw logs and unpacked data for AdvancePurchase events raised by the ScryProtocol contract.
type ScryProtocolAdvancePurchaseIterator struct {
	Event *ScryProtocolAdvancePurchase // Event containing the contract specifics and raw log

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
func (it *ScryProtocolAdvancePurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolAdvancePurchase)
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
		it.Event = new(ScryProtocolAdvancePurchase)
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
func (it *ScryProtocolAdvancePurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolAdvancePurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolAdvancePurchase represents a AdvancePurchase event raised by the ScryProtocol contract.
type ScryProtocolAdvancePurchase struct {
	SeqNo         string
	TransactionId *big.Int
	PublishId     string
	ProofIds      [][32]byte
	NeedVerify    bool
	State         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdvancePurchase is a free log retrieval operation binding the contract event 0x01692207fb0ede4de9f01a99e030dd094249ddbd9ad5e5eb0136939fc3fedf6a.
//
// Solidity: event AdvancePurchase(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool needVerify, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterAdvancePurchase(opts *bind.FilterOpts) (*ScryProtocolAdvancePurchaseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "AdvancePurchase")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolAdvancePurchaseIterator{contract: _ScryProtocol.contract, event: "AdvancePurchase", logs: logs, sub: sub}, nil
}

// WatchAdvancePurchase is a free log subscription operation binding the contract event 0x01692207fb0ede4de9f01a99e030dd094249ddbd9ad5e5eb0136939fc3fedf6a.
//
// Solidity: event AdvancePurchase(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool needVerify, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchAdvancePurchase(opts *bind.WatchOpts, sink chan<- *ScryProtocolAdvancePurchase) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "AdvancePurchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolAdvancePurchase)
				if err := _ScryProtocol.contract.UnpackLog(event, "AdvancePurchase", log); err != nil {
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

// ScryProtocolArbitrationBeginIterator is returned from FilterArbitrationBegin and is used to iterate over the raw logs and unpacked data for ArbitrationBegin events raised by the ScryProtocol contract.
type ScryProtocolArbitrationBeginIterator struct {
	Event *ScryProtocolArbitrationBegin // Event containing the contract specifics and raw log

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
func (it *ScryProtocolArbitrationBeginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolArbitrationBegin)
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
		it.Event = new(ScryProtocolArbitrationBegin)
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
func (it *ScryProtocolArbitrationBeginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolArbitrationBeginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolArbitrationBegin represents a ArbitrationBegin event raised by the ScryProtocol contract.
type ScryProtocolArbitrationBegin struct {
	SeqNo                   string
	TransactionId           *big.Int
	PublishId               string
	ProofIds                [][32]byte
	MetaDataIdEncArbitrator []byte
	Users                   []common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterArbitrationBegin is a free log retrieval operation binding the contract event 0x541c5ec36a883a22048ce07131e93f1c295bd6ccfed09aaab4e233e9b07c87b0.
//
// Solidity: event ArbitrationBegin(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bytes metaDataIdEncArbitrator, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterArbitrationBegin(opts *bind.FilterOpts) (*ScryProtocolArbitrationBeginIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ArbitrationBegin")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolArbitrationBeginIterator{contract: _ScryProtocol.contract, event: "ArbitrationBegin", logs: logs, sub: sub}, nil
}

// WatchArbitrationBegin is a free log subscription operation binding the contract event 0x541c5ec36a883a22048ce07131e93f1c295bd6ccfed09aaab4e233e9b07c87b0.
//
// Solidity: event ArbitrationBegin(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bytes metaDataIdEncArbitrator, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchArbitrationBegin(opts *bind.WatchOpts, sink chan<- *ScryProtocolArbitrationBegin) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ArbitrationBegin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolArbitrationBegin)
				if err := _ScryProtocol.contract.UnpackLog(event, "ArbitrationBegin", log); err != nil {
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

// ScryProtocolArbitrationResultIterator is returned from FilterArbitrationResult and is used to iterate over the raw logs and unpacked data for ArbitrationResult events raised by the ScryProtocol contract.
type ScryProtocolArbitrationResultIterator struct {
	Event *ScryProtocolArbitrationResult // Event containing the contract specifics and raw log

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
func (it *ScryProtocolArbitrationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolArbitrationResult)
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
		it.Event = new(ScryProtocolArbitrationResult)
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
func (it *ScryProtocolArbitrationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolArbitrationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolArbitrationResult represents a ArbitrationResult event raised by the ScryProtocol contract.
type ScryProtocolArbitrationResult struct {
	SeqNo         string
	TransactionId *big.Int
	Judge         bool
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterArbitrationResult is a free log retrieval operation binding the contract event 0x740cf4421b94868a990fcc1cb9d09134f9ad70520aaebd6769efe1b4a3fe5283.
//
// Solidity: event ArbitrationResult(string seqNo, uint256 transactionId, bool judge, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterArbitrationResult(opts *bind.FilterOpts) (*ScryProtocolArbitrationResultIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ArbitrationResult")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolArbitrationResultIterator{contract: _ScryProtocol.contract, event: "ArbitrationResult", logs: logs, sub: sub}, nil
}

// WatchArbitrationResult is a free log subscription operation binding the contract event 0x740cf4421b94868a990fcc1cb9d09134f9ad70520aaebd6769efe1b4a3fe5283.
//
// Solidity: event ArbitrationResult(string seqNo, uint256 transactionId, bool judge, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchArbitrationResult(opts *bind.WatchOpts, sink chan<- *ScryProtocolArbitrationResult) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ArbitrationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolArbitrationResult)
				if err := _ScryProtocol.contract.UnpackLog(event, "ArbitrationResult", log); err != nil {
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

// ScryProtocolConfirmPurchaseIterator is returned from FilterConfirmPurchase and is used to iterate over the raw logs and unpacked data for ConfirmPurchase events raised by the ScryProtocol contract.
type ScryProtocolConfirmPurchaseIterator struct {
	Event *ScryProtocolConfirmPurchase // Event containing the contract specifics and raw log

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
func (it *ScryProtocolConfirmPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolConfirmPurchase)
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
		it.Event = new(ScryProtocolConfirmPurchase)
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
func (it *ScryProtocolConfirmPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolConfirmPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolConfirmPurchase represents a ConfirmPurchase event raised by the ScryProtocol contract.
type ScryProtocolConfirmPurchase struct {
	SeqNo               string
	TransactionId       *big.Int
	PublishId           string
	MetaDataIdEncSeller []byte
	State               uint8
	Users               []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConfirmPurchase is a free log retrieval operation binding the contract event 0xedb5479a30db18e1e51ef02fb66fcfca0de324d919fd8b33b4b986d2b318cae9.
//
// Solidity: event ConfirmPurchase(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterConfirmPurchase(opts *bind.FilterOpts) (*ScryProtocolConfirmPurchaseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ConfirmPurchase")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolConfirmPurchaseIterator{contract: _ScryProtocol.contract, event: "ConfirmPurchase", logs: logs, sub: sub}, nil
}

// WatchConfirmPurchase is a free log subscription operation binding the contract event 0xedb5479a30db18e1e51ef02fb66fcfca0de324d919fd8b33b4b986d2b318cae9.
//
// Solidity: event ConfirmPurchase(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchConfirmPurchase(opts *bind.WatchOpts, sink chan<- *ScryProtocolConfirmPurchase) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ConfirmPurchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolConfirmPurchase)
				if err := _ScryProtocol.contract.UnpackLog(event, "ConfirmPurchase", log); err != nil {
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

// ScryProtocolFinishPurchaseIterator is returned from FilterFinishPurchase and is used to iterate over the raw logs and unpacked data for FinishPurchase events raised by the ScryProtocol contract.
type ScryProtocolFinishPurchaseIterator struct {
	Event *ScryProtocolFinishPurchase // Event containing the contract specifics and raw log

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
func (it *ScryProtocolFinishPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolFinishPurchase)
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
		it.Event = new(ScryProtocolFinishPurchase)
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
func (it *ScryProtocolFinishPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolFinishPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolFinishPurchase represents a FinishPurchase event raised by the ScryProtocol contract.
type ScryProtocolFinishPurchase struct {
	SeqNo         string
	TransactionId *big.Int
	State         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinishPurchase is a free log retrieval operation binding the contract event 0x0dc9c811deb4794833b0d19240fe819bc7ee9eedc1df619d9edc3cbd6afff1bd.
//
// Solidity: event FinishPurchase(string seqNo, uint256 transactionId, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterFinishPurchase(opts *bind.FilterOpts) (*ScryProtocolFinishPurchaseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "FinishPurchase")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolFinishPurchaseIterator{contract: _ScryProtocol.contract, event: "FinishPurchase", logs: logs, sub: sub}, nil
}

// WatchFinishPurchase is a free log subscription operation binding the contract event 0x0dc9c811deb4794833b0d19240fe819bc7ee9eedc1df619d9edc3cbd6afff1bd.
//
// Solidity: event FinishPurchase(string seqNo, uint256 transactionId, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchFinishPurchase(opts *bind.WatchOpts, sink chan<- *ScryProtocolFinishPurchase) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "FinishPurchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolFinishPurchase)
				if err := _ScryProtocol.contract.UnpackLog(event, "FinishPurchase", log); err != nil {
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

// ScryProtocolPublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the ScryProtocol contract.
type ScryProtocolPublishIterator struct {
	Event *ScryProtocolPublish // Event containing the contract specifics and raw log

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
func (it *ScryProtocolPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolPublish)
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
		it.Event = new(ScryProtocolPublish)
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
func (it *ScryProtocolPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolPublish represents a Publish event raised by the ScryProtocol contract.
type ScryProtocolPublish struct {
	SeqNo         string
	PublishId     string
	Price         *big.Int
	DespDataId    string
	SupportVerify bool
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x96ea505b20c061143f36aae01dd3cf12c53321879ef4f4e89eee5f8f9b5ab995.
//
// Solidity: event Publish(string seqNo, string publishId, uint256 price, string despDataId, bool supportVerify, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterPublish(opts *bind.FilterOpts) (*ScryProtocolPublishIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolPublishIterator{contract: _ScryProtocol.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x96ea505b20c061143f36aae01dd3cf12c53321879ef4f4e89eee5f8f9b5ab995.
//
// Solidity: event Publish(string seqNo, string publishId, uint256 price, string despDataId, bool supportVerify, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *ScryProtocolPublish) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolPublish)
				if err := _ScryProtocol.contract.UnpackLog(event, "Publish", log); err != nil {
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

// ScryProtocolReEncryptIterator is returned from FilterReEncrypt and is used to iterate over the raw logs and unpacked data for ReEncrypt events raised by the ScryProtocol contract.
type ScryProtocolReEncryptIterator struct {
	Event *ScryProtocolReEncrypt // Event containing the contract specifics and raw log

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
func (it *ScryProtocolReEncryptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolReEncrypt)
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
		it.Event = new(ScryProtocolReEncrypt)
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
func (it *ScryProtocolReEncryptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolReEncryptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolReEncrypt represents a ReEncrypt event raised by the ScryProtocol contract.
type ScryProtocolReEncrypt struct {
	SeqNo              string
	TransactionId      *big.Int
	MetaDataIdEncBuyer []byte
	State              uint8
	Users              []common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReEncrypt is a free log retrieval operation binding the contract event 0x09205ba69360e9b0520ab453653871d6d2e61812aba619c6160cf2e3684141be.
//
// Solidity: event ReEncrypt(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterReEncrypt(opts *bind.FilterOpts) (*ScryProtocolReEncryptIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ReEncrypt")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolReEncryptIterator{contract: _ScryProtocol.contract, event: "ReEncrypt", logs: logs, sub: sub}, nil
}

// WatchReEncrypt is a free log subscription operation binding the contract event 0x09205ba69360e9b0520ab453653871d6d2e61812aba619c6160cf2e3684141be.
//
// Solidity: event ReEncrypt(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchReEncrypt(opts *bind.WatchOpts, sink chan<- *ScryProtocolReEncrypt) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ReEncrypt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolReEncrypt)
				if err := _ScryProtocol.contract.UnpackLog(event, "ReEncrypt", log); err != nil {
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

// ScryProtocolRegisterVerifierIterator is returned from FilterRegisterVerifier and is used to iterate over the raw logs and unpacked data for RegisterVerifier events raised by the ScryProtocol contract.
type ScryProtocolRegisterVerifierIterator struct {
	Event *ScryProtocolRegisterVerifier // Event containing the contract specifics and raw log

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
func (it *ScryProtocolRegisterVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolRegisterVerifier)
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
		it.Event = new(ScryProtocolRegisterVerifier)
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
func (it *ScryProtocolRegisterVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolRegisterVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolRegisterVerifier represents a RegisterVerifier event raised by the ScryProtocol contract.
type ScryProtocolRegisterVerifier struct {
	SeqNo string
	Users []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegisterVerifier is a free log retrieval operation binding the contract event 0x476785064b6fb8cce78cd4377a03177c7bac7803ef345a1eaf34d1dbdbf0e864.
//
// Solidity: event RegisterVerifier(string seqNo, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterRegisterVerifier(opts *bind.FilterOpts) (*ScryProtocolRegisterVerifierIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "RegisterVerifier")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolRegisterVerifierIterator{contract: _ScryProtocol.contract, event: "RegisterVerifier", logs: logs, sub: sub}, nil
}

// WatchRegisterVerifier is a free log subscription operation binding the contract event 0x476785064b6fb8cce78cd4377a03177c7bac7803ef345a1eaf34d1dbdbf0e864.
//
// Solidity: event RegisterVerifier(string seqNo, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchRegisterVerifier(opts *bind.WatchOpts, sink chan<- *ScryProtocolRegisterVerifier) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "RegisterVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolRegisterVerifier)
				if err := _ScryProtocol.contract.UnpackLog(event, "RegisterVerifier", log); err != nil {
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

// ScryProtocolVerifierDisableIterator is returned from FilterVerifierDisable and is used to iterate over the raw logs and unpacked data for VerifierDisable events raised by the ScryProtocol contract.
type ScryProtocolVerifierDisableIterator struct {
	Event *ScryProtocolVerifierDisable // Event containing the contract specifics and raw log

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
func (it *ScryProtocolVerifierDisableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolVerifierDisable)
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
		it.Event = new(ScryProtocolVerifierDisable)
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
func (it *ScryProtocolVerifierDisableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolVerifierDisableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolVerifierDisable represents a VerifierDisable event raised by the ScryProtocol contract.
type ScryProtocolVerifierDisable struct {
	SeqNo    string
	Verifier common.Address
	Users    []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierDisable is a free log retrieval operation binding the contract event 0xd0e2127bc672e5762b6852e277e4d77594e1eecad47e50f41f56d9a87c1f7505.
//
// Solidity: event VerifierDisable(string seqNo, address verifier, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterVerifierDisable(opts *bind.FilterOpts) (*ScryProtocolVerifierDisableIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "VerifierDisable")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolVerifierDisableIterator{contract: _ScryProtocol.contract, event: "VerifierDisable", logs: logs, sub: sub}, nil
}

// WatchVerifierDisable is a free log subscription operation binding the contract event 0xd0e2127bc672e5762b6852e277e4d77594e1eecad47e50f41f56d9a87c1f7505.
//
// Solidity: event VerifierDisable(string seqNo, address verifier, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchVerifierDisable(opts *bind.WatchOpts, sink chan<- *ScryProtocolVerifierDisable) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "VerifierDisable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolVerifierDisable)
				if err := _ScryProtocol.contract.UnpackLog(event, "VerifierDisable", log); err != nil {
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

// ScryProtocolVerifiersChosenIterator is returned from FilterVerifiersChosen and is used to iterate over the raw logs and unpacked data for VerifiersChosen events raised by the ScryProtocol contract.
type ScryProtocolVerifiersChosenIterator struct {
	Event *ScryProtocolVerifiersChosen // Event containing the contract specifics and raw log

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
func (it *ScryProtocolVerifiersChosenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolVerifiersChosen)
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
		it.Event = new(ScryProtocolVerifiersChosen)
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
func (it *ScryProtocolVerifiersChosenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolVerifiersChosenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolVerifiersChosen represents a VerifiersChosen event raised by the ScryProtocol contract.
type ScryProtocolVerifiersChosen struct {
	SeqNo         string
	TransactionId *big.Int
	PublishId     string
	ProofIds      [][32]byte
	State         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVerifiersChosen is a free log retrieval operation binding the contract event 0x8c48707022c17239432366b846b96252d3a8c7f241cbba874edb30ec41c1acb3.
//
// Solidity: event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterVerifiersChosen(opts *bind.FilterOpts) (*ScryProtocolVerifiersChosenIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "VerifiersChosen")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolVerifiersChosenIterator{contract: _ScryProtocol.contract, event: "VerifiersChosen", logs: logs, sub: sub}, nil
}

// WatchVerifiersChosen is a free log subscription operation binding the contract event 0x8c48707022c17239432366b846b96252d3a8c7f241cbba874edb30ec41c1acb3.
//
// Solidity: event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, uint8 state, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchVerifiersChosen(opts *bind.WatchOpts, sink chan<- *ScryProtocolVerifiersChosen) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "VerifiersChosen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolVerifiersChosen)
				if err := _ScryProtocol.contract.UnpackLog(event, "VerifiersChosen", log); err != nil {
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

// ScryProtocolVoteResultIterator is returned from FilterVoteResult and is used to iterate over the raw logs and unpacked data for VoteResult events raised by the ScryProtocol contract.
type ScryProtocolVoteResultIterator struct {
	Event *ScryProtocolVoteResult // Event containing the contract specifics and raw log

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
func (it *ScryProtocolVoteResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolVoteResult)
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
		it.Event = new(ScryProtocolVoteResult)
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
func (it *ScryProtocolVoteResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolVoteResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolVoteResult represents a VoteResult event raised by the ScryProtocol contract.
type ScryProtocolVoteResult struct {
	SeqNo         string
	TransactionId *big.Int
	Judge         bool
	Comments      string
	State         uint8
	Index         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoteResult is a free log retrieval operation binding the contract event 0xd43b482cba5262142d7309862c56fbe78283aecf3d0bf76823a70975a121f814.
//
// Solidity: event VoteResult(string seqNo, uint256 transactionId, bool judge, string comments, uint8 state, uint8 index, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) FilterVoteResult(opts *bind.FilterOpts) (*ScryProtocolVoteResultIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "VoteResult")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolVoteResultIterator{contract: _ScryProtocol.contract, event: "VoteResult", logs: logs, sub: sub}, nil
}

// WatchVoteResult is a free log subscription operation binding the contract event 0xd43b482cba5262142d7309862c56fbe78283aecf3d0bf76823a70975a121f814.
//
// Solidity: event VoteResult(string seqNo, uint256 transactionId, bool judge, string comments, uint8 state, uint8 index, address[] users)
func (_ScryProtocol *ScryProtocolFilterer) WatchVoteResult(opts *bind.WatchOpts, sink chan<- *ScryProtocolVoteResult) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "VoteResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolVoteResult)
				if err := _ScryProtocol.contract.UnpackLog(event, "VoteResult", log); err != nil {
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
