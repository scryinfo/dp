// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
const ScryProtocolABI = "[{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"}],\"name\":\"registerAsVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"judge\",\"type\":\"bool\"},{\"name\":\"comments\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"verifierIndex\",\"type\":\"uint8\"},{\"name\":\"credit\",\"type\":\"uint8\"}],\"name\":\"creditsToVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"judge\",\"type\":\"bool\"}],\"name\":\"arbitrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"name\":\"proofDataIds\",\"type\":\"bytes32[]\"},{\"name\":\"descDataId\",\"type\":\"string\"},{\"name\":\"supportVerify\",\"type\":\"bool\"}],\"name\":\"publishDataInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"startVerify\",\"type\":\"bool\"}],\"name\":\"createTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"buyData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"cancelTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"encryptedMetaDataId\",\"type\":\"bytes\"},{\"name\":\"encryptedMetaDataIds\",\"type\":\"bytes\"}],\"name\":\"reEncryptMetaDataIdFromSeller\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"truth\",\"type\":\"bool\"}],\"name\":\"confirmDataTruth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"getBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"getArbitrators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactor) BuyData(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "buyData", seqNo, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolSession) BuyData(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) BuyData(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0xcca8f8c3.
//
// Solidity: function cancelTransaction(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactor) CancelTransaction(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "cancelTransaction", seqNo, txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0xcca8f8c3.
//
// Solidity: function cancelTransaction(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolSession) CancelTransaction(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CancelTransaction(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0xcca8f8c3.
//
// Solidity: function cancelTransaction(string seqNo, uint256 txId) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CancelTransaction(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CancelTransaction(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(string seqNo, uint256 txId, bool truth) returns()
func (_ScryProtocol *ScryProtocolTransactor) ConfirmDataTruth(opts *bind.TransactOpts, seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "confirmDataTruth", seqNo, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(string seqNo, uint256 txId, bool truth) returns()
func (_ScryProtocol *ScryProtocolSession) ConfirmDataTruth(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(string seqNo, uint256 txId, bool truth) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ConfirmDataTruth(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xccc2ba76.
//
// Solidity: function createTransaction(string seqNo, string publishId, bool startVerify) returns()
func (_ScryProtocol *ScryProtocolTransactor) CreateTransaction(opts *bind.TransactOpts, seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "createTransaction", seqNo, publishId, startVerify)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xccc2ba76.
//
// Solidity: function createTransaction(string seqNo, string publishId, bool startVerify) returns()
func (_ScryProtocol *ScryProtocolSession) CreateTransaction(seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreateTransaction(&_ScryProtocol.TransactOpts, seqNo, publishId, startVerify)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xccc2ba76.
//
// Solidity: function createTransaction(string seqNo, string publishId, bool startVerify) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CreateTransaction(seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreateTransaction(&_ScryProtocol.TransactOpts, seqNo, publishId, startVerify)
}

// CreditsToVerifier is a paid mutator transaction binding the contract method 0xbd7ff8e5.
//
// Solidity: function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) returns()
func (_ScryProtocol *ScryProtocolTransactor) CreditsToVerifier(opts *bind.TransactOpts, seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "creditsToVerifier", seqNo, txId, verifierIndex, credit)
}

// CreditsToVerifier is a paid mutator transaction binding the contract method 0xbd7ff8e5.
//
// Solidity: function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) returns()
func (_ScryProtocol *ScryProtocolSession) CreditsToVerifier(seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreditsToVerifier(&_ScryProtocol.TransactOpts, seqNo, txId, verifierIndex, credit)
}

// CreditsToVerifier is a paid mutator transaction binding the contract method 0xbd7ff8e5.
//
// Solidity: function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CreditsToVerifier(seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreditsToVerifier(&_ScryProtocol.TransactOpts, seqNo, txId, verifierIndex, credit)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string descDataId, bool supportVerify) returns()
func (_ScryProtocol *ScryProtocolTransactor) PublishDataInfo(opts *bind.TransactOpts, seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "publishDataInfo", seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string descDataId, bool supportVerify) returns()
func (_ScryProtocol *ScryProtocolSession) PublishDataInfo(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string descDataId, bool supportVerify) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) PublishDataInfo(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// ReEncryptMetaDataIdFromSeller is a paid mutator transaction binding the contract method 0x4385617c.
//
// Solidity: function reEncryptMetaDataIdFromSeller(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) returns()
func (_ScryProtocol *ScryProtocolTransactor) ReEncryptMetaDataIdFromSeller(opts *bind.TransactOpts, seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "reEncryptMetaDataIdFromSeller", seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// ReEncryptMetaDataIdFromSeller is a paid mutator transaction binding the contract method 0x4385617c.
//
// Solidity: function reEncryptMetaDataIdFromSeller(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) returns()
func (_ScryProtocol *ScryProtocolSession) ReEncryptMetaDataIdFromSeller(seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ReEncryptMetaDataIdFromSeller(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// ReEncryptMetaDataIdFromSeller is a paid mutator transaction binding the contract method 0x4385617c.
//
// Solidity: function reEncryptMetaDataIdFromSeller(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ReEncryptMetaDataIdFromSeller(seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ReEncryptMetaDataIdFromSeller(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
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
