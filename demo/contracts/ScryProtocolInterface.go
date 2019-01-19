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
const ScryProtocolABI = "[{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"despDataId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"DataPublish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"chosenProofIds\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"supportVerify\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"TransactionCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metaDataIdEncBuyer\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ReadyForDownload\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"TransactionClose\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"name\":\"proofDataIds\",\"type\":\"bytes32[]\"},{\"name\":\"despDataId\",\"type\":\"string\"},{\"name\":\"supportVerify\",\"type\":\"bool\"}],\"name\":\"publishDataInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"}],\"name\":\"createTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"buyData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"encryptedMetaDataId\",\"type\":\"bytes\"}],\"name\":\"submitMetaDataIdEncWithBuyer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"truth\",\"type\":\"bool\"}],\"name\":\"confirmDataTruth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactor) BuyData(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "buyData", seqNo, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolSession) BuyData(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) BuyData(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(seqNo string, txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) ConfirmDataTruth(opts *bind.TransactOpts, seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "confirmDataTruth", seqNo, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(seqNo string, txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolSession) ConfirmDataTruth(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(seqNo string, txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ConfirmDataTruth(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xfecc18b4.
//
// Solidity: function createTransaction(seqNo string, publishId string) returns()
func (_ScryProtocol *ScryProtocolTransactor) CreateTransaction(opts *bind.TransactOpts, seqNo string, publishId string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "createTransaction", seqNo, publishId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xfecc18b4.
//
// Solidity: function createTransaction(seqNo string, publishId string) returns()
func (_ScryProtocol *ScryProtocolSession) CreateTransaction(seqNo string, publishId string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreateTransaction(&_ScryProtocol.TransactOpts, seqNo, publishId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xfecc18b4.
//
// Solidity: function createTransaction(seqNo string, publishId string) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CreateTransaction(seqNo string, publishId string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreateTransaction(&_ScryProtocol.TransactOpts, seqNo, publishId)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(seqNo string, publishId string, price uint256, metaDataIdEncSeller bytes, proofDataIds bytes32[], despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) PublishDataInfo(opts *bind.TransactOpts, seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "publishDataInfo", seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, despDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(seqNo string, publishId string, price uint256, metaDataIdEncSeller bytes, proofDataIds bytes32[], despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolSession) PublishDataInfo(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, despDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(seqNo string, publishId string, price uint256, metaDataIdEncSeller bytes, proofDataIds bytes32[], despDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) PublishDataInfo(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, despDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, despDataId, supportVerify)
}

// SubmitMetaDataIdEncWithBuyer is a paid mutator transaction binding the contract method 0x8ba737ee.
//
// Solidity: function submitMetaDataIdEncWithBuyer(seqNo string, txId uint256, encryptedMetaDataId bytes) returns()
func (_ScryProtocol *ScryProtocolTransactor) SubmitMetaDataIdEncWithBuyer(opts *bind.TransactOpts, seqNo string, txId *big.Int, encryptedMetaDataId []byte) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "submitMetaDataIdEncWithBuyer", seqNo, txId, encryptedMetaDataId)
}

// SubmitMetaDataIdEncWithBuyer is a paid mutator transaction binding the contract method 0x8ba737ee.
//
// Solidity: function submitMetaDataIdEncWithBuyer(seqNo string, txId uint256, encryptedMetaDataId bytes) returns()
func (_ScryProtocol *ScryProtocolSession) SubmitMetaDataIdEncWithBuyer(seqNo string, txId *big.Int, encryptedMetaDataId []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.SubmitMetaDataIdEncWithBuyer(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId)
}

// SubmitMetaDataIdEncWithBuyer is a paid mutator transaction binding the contract method 0x8ba737ee.
//
// Solidity: function submitMetaDataIdEncWithBuyer(seqNo string, txId uint256, encryptedMetaDataId bytes) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) SubmitMetaDataIdEncWithBuyer(seqNo string, txId *big.Int, encryptedMetaDataId []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.SubmitMetaDataIdEncWithBuyer(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId)
}

// ScryProtocolBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the ScryProtocol contract.
type ScryProtocolBuyIterator struct {
	Event *ScryProtocolBuy // Event containing the contract specifics and raw log

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
func (it *ScryProtocolBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolBuy)
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
		it.Event = new(ScryProtocolBuy)
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
func (it *ScryProtocolBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolBuy represents a Buy event raised by the ScryProtocol contract.
type ScryProtocolBuy struct {
	SeqNo               string
	TransactionId       *big.Int
	PublishId           string
	MetaDataIdEncSeller []byte
	Users               []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0xeeaf28668d6d8a951d1c5b873c5e94068abdea0145811cbeaece6c309118f3e3.
//
// Solidity: e Buy(seqNo string, transactionId uint256, publishId string, metaDataIdEncSeller bytes, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterBuy(opts *bind.FilterOpts) (*ScryProtocolBuyIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Buy")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolBuyIterator{contract: _ScryProtocol.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0xeeaf28668d6d8a951d1c5b873c5e94068abdea0145811cbeaece6c309118f3e3.
//
// Solidity: e Buy(seqNo string, transactionId uint256, publishId string, metaDataIdEncSeller bytes, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *ScryProtocolBuy) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Buy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolBuy)
				if err := _ScryProtocol.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ScryProtocolDataPublishIterator is returned from FilterDataPublish and is used to iterate over the raw logs and unpacked data for DataPublish events raised by the ScryProtocol contract.
type ScryProtocolDataPublishIterator struct {
	Event *ScryProtocolDataPublish // Event containing the contract specifics and raw log

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
func (it *ScryProtocolDataPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolDataPublish)
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
		it.Event = new(ScryProtocolDataPublish)
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
func (it *ScryProtocolDataPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolDataPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolDataPublish represents a DataPublish event raised by the ScryProtocol contract.
type ScryProtocolDataPublish struct {
	SeqNo      string
	PublishId  string
	Price      *big.Int
	DespDataId string
	Users      []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataPublish is a free log retrieval operation binding the contract event 0x9d1b954932f5249f38959200c17d585ac2b4897b77acb4c29a8fb4d8473f0e8c.
//
// Solidity: e DataPublish(seqNo string, publishId string, price uint256, despDataId string, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterDataPublish(opts *bind.FilterOpts) (*ScryProtocolDataPublishIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "DataPublish")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolDataPublishIterator{contract: _ScryProtocol.contract, event: "DataPublish", logs: logs, sub: sub}, nil
}

// WatchDataPublish is a free log subscription operation binding the contract event 0x9d1b954932f5249f38959200c17d585ac2b4897b77acb4c29a8fb4d8473f0e8c.
//
// Solidity: e DataPublish(seqNo string, publishId string, price uint256, despDataId string, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchDataPublish(opts *bind.WatchOpts, sink chan<- *ScryProtocolDataPublish) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "DataPublish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolDataPublish)
				if err := _ScryProtocol.contract.UnpackLog(event, "DataPublish", log); err != nil {
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
	SeqNo              string
	TransactionId      *big.Int
	MetaDataIdEncBuyer []byte
	Users              []common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReadyForDownload is a free log retrieval operation binding the contract event 0x28eb60ba9df25f12ccee668d0a745b5e455365b717f0867aa76edcffe7430cf5.
//
// Solidity: e ReadyForDownload(seqNo string, transactionId uint256, metaDataIdEncBuyer bytes, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterReadyForDownload(opts *bind.FilterOpts) (*ScryProtocolReadyForDownloadIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ReadyForDownload")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolReadyForDownloadIterator{contract: _ScryProtocol.contract, event: "ReadyForDownload", logs: logs, sub: sub}, nil
}

// WatchReadyForDownload is a free log subscription operation binding the contract event 0x28eb60ba9df25f12ccee668d0a745b5e455365b717f0867aa76edcffe7430cf5.
//
// Solidity: e ReadyForDownload(seqNo string, transactionId uint256, metaDataIdEncBuyer bytes, users address[])
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

// ScryProtocolTransactionCloseIterator is returned from FilterTransactionClose and is used to iterate over the raw logs and unpacked data for TransactionClose events raised by the ScryProtocol contract.
type ScryProtocolTransactionCloseIterator struct {
	Event *ScryProtocolTransactionClose // Event containing the contract specifics and raw log

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
func (it *ScryProtocolTransactionCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolTransactionClose)
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
		it.Event = new(ScryProtocolTransactionClose)
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
func (it *ScryProtocolTransactionCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolTransactionCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolTransactionClose represents a TransactionClose event raised by the ScryProtocol contract.
type ScryProtocolTransactionClose struct {
	SeqNo         string
	TransactionId *big.Int
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionClose is a free log retrieval operation binding the contract event 0x784a9feb5a6ecee5d955d1b5e294f2b57d9477164fba686c0bde4a40849f2b0c.
//
// Solidity: e TransactionClose(seqNo string, transactionId uint256, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterTransactionClose(opts *bind.FilterOpts) (*ScryProtocolTransactionCloseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "TransactionClose")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactionCloseIterator{contract: _ScryProtocol.contract, event: "TransactionClose", logs: logs, sub: sub}, nil
}

// WatchTransactionClose is a free log subscription operation binding the contract event 0x784a9feb5a6ecee5d955d1b5e294f2b57d9477164fba686c0bde4a40849f2b0c.
//
// Solidity: e TransactionClose(seqNo string, transactionId uint256, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchTransactionClose(opts *bind.WatchOpts, sink chan<- *ScryProtocolTransactionClose) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "TransactionClose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolTransactionClose)
				if err := _ScryProtocol.contract.UnpackLog(event, "TransactionClose", log); err != nil {
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
	SeqNo          string
	TransactionId  *big.Int
	PublishId      string
	ChosenProofIds [32]byte
	SupportVerify  bool
	Users          []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreate is a free log retrieval operation binding the contract event 0x626f6f0cbcf39ba22ef1c481bf5743596d102baea8fed58505e1255496d5b0a0.
//
// Solidity: e TransactionCreate(seqNo string, transactionId uint256, publishId string, chosenProofIds bytes32, supportVerify bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterTransactionCreate(opts *bind.FilterOpts) (*ScryProtocolTransactionCreateIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "TransactionCreate")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactionCreateIterator{contract: _ScryProtocol.contract, event: "TransactionCreate", logs: logs, sub: sub}, nil
}

// WatchTransactionCreate is a free log subscription operation binding the contract event 0x626f6f0cbcf39ba22ef1c481bf5743596d102baea8fed58505e1255496d5b0a0.
//
// Solidity: e TransactionCreate(seqNo string, transactionId uint256, publishId string, chosenProofIds bytes32, supportVerify bool, users address[])
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
