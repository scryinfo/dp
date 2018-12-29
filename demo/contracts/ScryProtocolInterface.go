// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ScryProtocol

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
const ScryProtocolABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"despDataId\",\"type\":\"string\"}],\"name\":\"Published\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"name\":\"proofDataId\",\"type\":\"string\"},{\"name\":\"despDataId\",\"type\":\"string\"},{\"name\":\"supportVerify\",\"type\":\"bool\"}],\"name\":\"publishDataInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publishId\",\"type\":\"string\"}],\"name\":\"isPublishedDataExisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// IsPublishedDataExisted is a paid mutator transaction binding the contract method 0xf1f93e80.
//
// Solidity: function isPublishedDataExisted(publishId string) returns(bool)
func (_ScryProtocol *ScryProtocolTransactor) IsPublishedDataExisted(opts *bind.TransactOpts, publishId string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "isPublishedDataExisted", publishId)
}

// IsPublishedDataExisted is a paid mutator transaction binding the contract method 0xf1f93e80.
//
// Solidity: function isPublishedDataExisted(publishId string) returns(bool)
func (_ScryProtocol *ScryProtocolSession) IsPublishedDataExisted(publishId string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.IsPublishedDataExisted(&_ScryProtocol.TransactOpts, publishId)
}

// IsPublishedDataExisted is a paid mutator transaction binding the contract method 0xf1f93e80.
//
// Solidity: function isPublishedDataExisted(publishId string) returns(bool)
func (_ScryProtocol *ScryProtocolTransactorSession) IsPublishedDataExisted(publishId string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.IsPublishedDataExisted(&_ScryProtocol.TransactOpts, publishId)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0xd180d67f.
//
// Solidity: function publishDataInfo(publishId string, metaDataIdEncSeller bytes, proofDataId string, despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) PublishDataInfo(opts *bind.TransactOpts, publishId string, metaDataIdEncSeller []byte, proofDataId string, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "publishDataInfo", publishId, metaDataIdEncSeller, proofDataId, despDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0xd180d67f.
//
// Solidity: function publishDataInfo(publishId string, metaDataIdEncSeller bytes, proofDataId string, despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolSession) PublishDataInfo(publishId string, metaDataIdEncSeller []byte, proofDataId string, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, publishId, metaDataIdEncSeller, proofDataId, despDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0xd180d67f.
//
// Solidity: function publishDataInfo(publishId string, metaDataIdEncSeller bytes, proofDataId string, despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) PublishDataInfo(publishId string, metaDataIdEncSeller []byte, proofDataId string, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, publishId, metaDataIdEncSeller, proofDataId, despDataId, supportVerify)
}

// ScryProtocolPublishedIterator is returned from FilterPublished and is used to iterate over the raw logs and unpacked data for Published events raised by the ScryProtocol contract.
type ScryProtocolPublishedIterator struct {
	Event *ScryProtocolPublished // Event containing the contract specifics and raw log

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
func (it *ScryProtocolPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolPublished)
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
		it.Event = new(ScryProtocolPublished)
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
func (it *ScryProtocolPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolPublished represents a Published event raised by the ScryProtocol contract.
type ScryProtocolPublished struct {
	PublishId  string
	Users      string
	DespDataId string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPublished is a free log retrieval operation binding the contract event 0x46a08b125611ec320cc1058dd976ee1adea76edb0e29c87708a4ac583d31bd77.
//
// Solidity: e Published(publishId string, users string, despDataId string)
func (_ScryProtocol *ScryProtocolFilterer) FilterPublished(opts *bind.FilterOpts) (*ScryProtocolPublishedIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Published")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolPublishedIterator{contract: _ScryProtocol.contract, event: "Published", logs: logs, sub: sub}, nil
}

// WatchPublished is a free log subscription operation binding the contract event 0x46a08b125611ec320cc1058dd976ee1adea76edb0e29c87708a4ac583d31bd77.
//
// Solidity: e Published(publishId string, users string, despDataId string)
func (_ScryProtocol *ScryProtocolFilterer) WatchPublished(opts *bind.WatchOpts, sink chan<- *ScryProtocolPublished) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Published")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolPublished)
				if err := _ScryProtocol.contract.UnpackLog(event, "Published", log); err != nil {
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
