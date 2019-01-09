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
const ScryProtocolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"encryptedIdLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionSeq\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"despDataId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"boardcast\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"chosenProofIds\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"supportVerify\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"boardcast\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"TransactionCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"boardcast\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metaDataIdEncBuyer\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"boardcast\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ReadyForDownload\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"boardcast\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Close\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"name\":\"proofDataIds\",\"type\":\"bytes32[]\"},{\"name\":\"despDataId\",\"type\":\"string\"},{\"name\":\"supportVerify\",\"type\":\"bool\"}],\"name\":\"publishDataInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publishId\",\"type\":\"string\"}],\"name\":\"prepareToBuy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"buyData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"encryptedMetaDataId\",\"type\":\"bytes\"}],\"name\":\"submitMetaDataIdEncWithBuyer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"truth\",\"type\":\"bool\"}],\"name\":\"confirmDataTruth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// EncryptedIdLen is a free data retrieval call binding the contract method 0x28589870.
//
// Solidity: function encryptedIdLen() constant returns(uint256)
func (_ScryProtocol *ScryProtocolCaller) EncryptedIdLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ScryProtocol.contract.Call(opts, out, "encryptedIdLen")
	return *ret0, err
}

// EncryptedIdLen is a free data retrieval call binding the contract method 0x28589870.
//
// Solidity: function encryptedIdLen() constant returns(uint256)
func (_ScryProtocol *ScryProtocolSession) EncryptedIdLen() (*big.Int, error) {
	return _ScryProtocol.Contract.EncryptedIdLen(&_ScryProtocol.CallOpts)
}

// EncryptedIdLen is a free data retrieval call binding the contract method 0x28589870.
//
// Solidity: function encryptedIdLen() constant returns(uint256)
func (_ScryProtocol *ScryProtocolCallerSession) EncryptedIdLen() (*big.Int, error) {
	return _ScryProtocol.Contract.EncryptedIdLen(&_ScryProtocol.CallOpts)
}

// TransactionSeq is a free data retrieval call binding the contract method 0x41a3cf63.
//
// Solidity: function transactionSeq() constant returns(uint256)
func (_ScryProtocol *ScryProtocolCaller) TransactionSeq(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ScryProtocol.contract.Call(opts, out, "transactionSeq")
	return *ret0, err
}

// TransactionSeq is a free data retrieval call binding the contract method 0x41a3cf63.
//
// Solidity: function transactionSeq() constant returns(uint256)
func (_ScryProtocol *ScryProtocolSession) TransactionSeq() (*big.Int, error) {
	return _ScryProtocol.Contract.TransactionSeq(&_ScryProtocol.CallOpts)
}

// TransactionSeq is a free data retrieval call binding the contract method 0x41a3cf63.
//
// Solidity: function transactionSeq() constant returns(uint256)
func (_ScryProtocol *ScryProtocolCallerSession) TransactionSeq() (*big.Int, error) {
	return _ScryProtocol.Contract.TransactionSeq(&_ScryProtocol.CallOpts)
}

// BuyData is a paid mutator transaction binding the contract method 0xca209dce.
//
// Solidity: function buyData(txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactor) BuyData(opts *bind.TransactOpts, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "buyData", txId)
}

// BuyData is a paid mutator transaction binding the contract method 0xca209dce.
//
// Solidity: function buyData(txId uint256) returns()
func (_ScryProtocol *ScryProtocolSession) BuyData(txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0xca209dce.
//
// Solidity: function buyData(txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) BuyData(txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, txId)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0x74f1502b.
//
// Solidity: function confirmDataTruth(txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) ConfirmDataTruth(opts *bind.TransactOpts, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "confirmDataTruth", txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0x74f1502b.
//
// Solidity: function confirmDataTruth(txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolSession) ConfirmDataTruth(txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0x74f1502b.
//
// Solidity: function confirmDataTruth(txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ConfirmDataTruth(txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, txId, truth)
}

// PrepareToBuy is a paid mutator transaction binding the contract method 0xe01e8393.
//
// Solidity: function prepareToBuy(publishId string) returns()
func (_ScryProtocol *ScryProtocolTransactor) PrepareToBuy(opts *bind.TransactOpts, publishId string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "prepareToBuy", publishId)
}

// PrepareToBuy is a paid mutator transaction binding the contract method 0xe01e8393.
//
// Solidity: function prepareToBuy(publishId string) returns()
func (_ScryProtocol *ScryProtocolSession) PrepareToBuy(publishId string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PrepareToBuy(&_ScryProtocol.TransactOpts, publishId)
}

// PrepareToBuy is a paid mutator transaction binding the contract method 0xe01e8393.
//
// Solidity: function prepareToBuy(publishId string) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) PrepareToBuy(publishId string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PrepareToBuy(&_ScryProtocol.TransactOpts, publishId)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0xb5e759c8.
//
// Solidity: function publishDataInfo(publishId string, metaDataIdEncSeller bytes, proofDataIds bytes32[], despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) PublishDataInfo(opts *bind.TransactOpts, publishId string, metaDataIdEncSeller []byte, proofDataIds [][32]byte, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "publishDataInfo", publishId, metaDataIdEncSeller, proofDataIds, despDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0xb5e759c8.
//
// Solidity: function publishDataInfo(publishId string, metaDataIdEncSeller bytes, proofDataIds bytes32[], despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolSession) PublishDataInfo(publishId string, metaDataIdEncSeller []byte, proofDataIds [][32]byte, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, publishId, metaDataIdEncSeller, proofDataIds, despDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0xb5e759c8.
//
// Solidity: function publishDataInfo(publishId string, metaDataIdEncSeller bytes, proofDataIds bytes32[], despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) PublishDataInfo(publishId string, metaDataIdEncSeller []byte, proofDataIds [][32]byte, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, publishId, metaDataIdEncSeller, proofDataIds, despDataId, supportVerify)
}

// SubmitMetaDataIdEncWithBuyer is a paid mutator transaction binding the contract method 0x30544e69.
//
// Solidity: function submitMetaDataIdEncWithBuyer(txId uint256, encryptedMetaDataId bytes) returns()
func (_ScryProtocol *ScryProtocolTransactor) SubmitMetaDataIdEncWithBuyer(opts *bind.TransactOpts, txId *big.Int, encryptedMetaDataId []byte) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "submitMetaDataIdEncWithBuyer", txId, encryptedMetaDataId)
}

// SubmitMetaDataIdEncWithBuyer is a paid mutator transaction binding the contract method 0x30544e69.
//
// Solidity: function submitMetaDataIdEncWithBuyer(txId uint256, encryptedMetaDataId bytes) returns()
func (_ScryProtocol *ScryProtocolSession) SubmitMetaDataIdEncWithBuyer(txId *big.Int, encryptedMetaDataId []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.SubmitMetaDataIdEncWithBuyer(&_ScryProtocol.TransactOpts, txId, encryptedMetaDataId)
}

// SubmitMetaDataIdEncWithBuyer is a paid mutator transaction binding the contract method 0x30544e69.
//
// Solidity: function submitMetaDataIdEncWithBuyer(txId uint256, encryptedMetaDataId bytes) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) SubmitMetaDataIdEncWithBuyer(txId *big.Int, encryptedMetaDataId []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.SubmitMetaDataIdEncWithBuyer(&_ScryProtocol.TransactOpts, txId, encryptedMetaDataId)
}

// ScryProtocolCloseIterator is returned from FilterClose and is used to iterate over the raw logs and unpacked data for Close events raised by the ScryProtocol contract.
type ScryProtocolCloseIterator struct {
	Event *ScryProtocolClose // Event containing the contract specifics and raw log

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
func (it *ScryProtocolCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolClose)
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
		it.Event = new(ScryProtocolClose)
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
func (it *ScryProtocolCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolClose represents a Close event raised by the ScryProtocol contract.
type ScryProtocolClose struct {
	TransactionId *big.Int
	Boardcast     bool
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClose is a free log retrieval operation binding the contract event 0xa332f85d6f38a0976f840537b79ecff1a59e902c9165a1a30d58cc844de47b31.
//
// Solidity: e Close(transactionId uint256, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterClose(opts *bind.FilterOpts) (*ScryProtocolCloseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolCloseIterator{contract: _ScryProtocol.contract, event: "Close", logs: logs, sub: sub}, nil
}

// WatchClose is a free log subscription operation binding the contract event 0xa332f85d6f38a0976f840537b79ecff1a59e902c9165a1a30d58cc844de47b31.
//
// Solidity: e Close(transactionId uint256, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchClose(opts *bind.WatchOpts, sink chan<- *ScryProtocolClose) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolClose)
				if err := _ScryProtocol.contract.UnpackLog(event, "Close", log); err != nil {
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
	PublishId  string
	DespDataId string
	Boardcast  bool
	Users      []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0xaabf77ef569fa1fabd0d636cf70fa9cdca9b6acb7c1258ac8bbc70ff8c1bdb17.
//
// Solidity: e Publish(publishId string, despDataId string, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterPublish(opts *bind.FilterOpts) (*ScryProtocolPublishIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolPublishIterator{contract: _ScryProtocol.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0xaabf77ef569fa1fabd0d636cf70fa9cdca9b6acb7c1258ac8bbc70ff8c1bdb17.
//
// Solidity: e Publish(publishId string, despDataId string, boardcast bool, users address[])
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

// ScryProtocolPurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the ScryProtocol contract.
type ScryProtocolPurchaseIterator struct {
	Event *ScryProtocolPurchase // Event containing the contract specifics and raw log

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
func (it *ScryProtocolPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolPurchase)
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
		it.Event = new(ScryProtocolPurchase)
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
func (it *ScryProtocolPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolPurchase represents a Purchase event raised by the ScryProtocol contract.
type ScryProtocolPurchase struct {
	TransactionId       *big.Int
	PublishId           string
	MetaDataIdEncSeller []byte
	Boardcast           bool
	Users               []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0x0fbc46e594df1851f28df5833e566c8c7024e215ebf4a4be8ad1b2a7ee410116.
//
// Solidity: e Purchase(transactionId uint256, publishId string, metaDataIdEncSeller bytes, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterPurchase(opts *bind.FilterOpts) (*ScryProtocolPurchaseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Purchase")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolPurchaseIterator{contract: _ScryProtocol.contract, event: "Purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0x0fbc46e594df1851f28df5833e566c8c7024e215ebf4a4be8ad1b2a7ee410116.
//
// Solidity: e Purchase(transactionId uint256, publishId string, metaDataIdEncSeller bytes, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *ScryProtocolPurchase) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Purchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolPurchase)
				if err := _ScryProtocol.contract.UnpackLog(event, "Purchase", log); err != nil {
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

// ScryProtocolReadyForDownloadIterator is returned from FilterReadyForDownload and is used to iterate over the raw logs and unpacked data for ReadyForDownload events raised by the ScryProtocol contract.
type ScryProtocolReadyForDownloadIterator struct {
	Event *ScryProtocolReadyForDownload // Event containing the contract specifics and raw log

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
func (it *ScryProtocolReadyForDownloadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolReadyForDownload)
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
		it.Event = new(ScryProtocolReadyForDownload)
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
func (it *ScryProtocolReadyForDownloadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolReadyForDownloadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolReadyForDownload represents a ReadyForDownload event raised by the ScryProtocol contract.
type ScryProtocolReadyForDownload struct {
	TransactionId      *big.Int
	MetaDataIdEncBuyer []byte
	Boardcast          bool
	Users              []common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReadyForDownload is a free log retrieval operation binding the contract event 0x432e20200e274397185e681b58f4ba63a66921d9064407bdf9f97ac5fbe5bf67.
//
// Solidity: e ReadyForDownload(transactionId uint256, metaDataIdEncBuyer bytes, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterReadyForDownload(opts *bind.FilterOpts) (*ScryProtocolReadyForDownloadIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ReadyForDownload")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolReadyForDownloadIterator{contract: _ScryProtocol.contract, event: "ReadyForDownload", logs: logs, sub: sub}, nil
}

// WatchReadyForDownload is a free log subscription operation binding the contract event 0x432e20200e274397185e681b58f4ba63a66921d9064407bdf9f97ac5fbe5bf67.
//
// Solidity: e ReadyForDownload(transactionId uint256, metaDataIdEncBuyer bytes, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchReadyForDownload(opts *bind.WatchOpts, sink chan<- *ScryProtocolReadyForDownload) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ReadyForDownload")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolReadyForDownload)
				if err := _ScryProtocol.contract.UnpackLog(event, "ReadyForDownload", log); err != nil {
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

// ScryProtocolTransactionCreateIterator is returned from FilterTransactionCreate and is used to iterate over the raw logs and unpacked data for TransactionCreate events raised by the ScryProtocol contract.
type ScryProtocolTransactionCreateIterator struct {
	Event *ScryProtocolTransactionCreate // Event containing the contract specifics and raw log

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
func (it *ScryProtocolTransactionCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolTransactionCreate)
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
		it.Event = new(ScryProtocolTransactionCreate)
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
func (it *ScryProtocolTransactionCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolTransactionCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolTransactionCreate represents a TransactionCreate event raised by the ScryProtocol contract.
type ScryProtocolTransactionCreate struct {
	TransactionId  *big.Int
	PublishId      string
	ChosenProofIds [32]byte
	SupportVerify  bool
	Boardcast      bool
	Users          []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreate is a free log retrieval operation binding the contract event 0x18ef72fd77cecc8fd4b61244c6f7bd40148cb71237eaeb896d3cab05825b284f.
//
// Solidity: e TransactionCreate(transactionId uint256, publishId string, chosenProofIds bytes32, supportVerify bool, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterTransactionCreate(opts *bind.FilterOpts) (*ScryProtocolTransactionCreateIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "TransactionCreate")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactionCreateIterator{contract: _ScryProtocol.contract, event: "TransactionCreate", logs: logs, sub: sub}, nil
}

// WatchTransactionCreate is a free log subscription operation binding the contract event 0x18ef72fd77cecc8fd4b61244c6f7bd40148cb71237eaeb896d3cab05825b284f.
//
// Solidity: e TransactionCreate(transactionId uint256, publishId string, chosenProofIds bytes32, supportVerify bool, boardcast bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchTransactionCreate(opts *bind.WatchOpts, sink chan<- *ScryProtocolTransactionCreate) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "TransactionCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolTransactionCreate)
				if err := _ScryProtocol.contract.UnpackLog(event, "TransactionCreate", log); err != nil {
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
