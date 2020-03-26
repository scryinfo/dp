package scan

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
)

//只实现了部分函数
type EthClientApiKeyImp struct {
	url    string
	apiKey string
	module string
}

func NewEthClientApiKeyImp(url string, apikey string, module string) *EthClientApiKeyImp {
	if len(module) < 1 {
		module = "module=proxy"
	}
	re := &EthClientApiKeyImp{
		url:    url,
		apiKey: apikey,
		module: module,
	}
	return re
}

func (e *EthClientApiKeyImp) Close() {

}

func (e *EthClientApiKeyImp) ChainID(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	//http://api-cn.etherscan.com/api?module=proxy&action=eth_getBlockByNumber&tag=0x10d4f&boolean=true&apikey=YourApiKeyToken
	tag := "latest"
	if number != nil {
		tag = "0x" + number.Text(16)
	}
	url := fmt.Sprintf("%s?%s&action=eth_getBlockByNumber&tag=%s&boolean=true&apikey=%s", e.url, e.module, tag, e.apiKey)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	block, err := client.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (e *EthClientApiKeyImp) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	//http://api-cn.etherscan.com/api?module=proxy&action=eth_getTransactionByHash&txhash=0x1e2910a262b1008d0616a0beb24c1a491d78771baa54a33e66065e03b1f46bc1&apikey=YourApiKeyToken
	url := fmt.Sprintf("%s?%s&action=eth_getTransactionByHash&txhash=%s&apikey=%s", e.url, e.module, hash.Hex(), e.apiKey)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, false, err
	}
	tx, p, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, false, err
	}
	return tx, p, err
}

func (e *EthClientApiKeyImp) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	//http://api-cn.etherscan.com/api?module=proxy&action=eth_getTransactionReceipt&txhash=0x1e2910a262b1008d0616a0beb24c1a491d78771baa54a33e66065e03b1f46bc1&apikey=YourApiKeyToken
	url := fmt.Sprintf("%s?%s&action=eth_getTransactionReceipt&txhash=%s&apikey=%s", e.url, e.module, txHash.Hex(), e.apiKey)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return receipt, err
}

func (e *EthClientApiKeyImp) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) NetworkID(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	//http://api-cn.etherscan.com/api?module=account&action=balance&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a&tag=latest&apikey=YourApiKeyToken
	tag := "latest"
	if blockNumber != nil {
		tag = "0x" + blockNumber.Text(16)
	}
	url := fmt.Sprintf("%s?module=account&action=balance&address=%s&tag=%s&apikey=%s", e.url, account.Hex(), tag, e.apiKey)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	//{"status":"1","message":"OK","result":"14020563159546670"}
	var re struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		return nil, err
	}
	value := big.NewInt(0)
	value, ok := value.SetString(re.Result, 10)
	if !ok {
		value = nil
		err = errors.New(re.Result)
	}
	return value, err
}

func (e *EthClientApiKeyImp) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	//http://api-cn.etherscan.com/api?module=proxy&action=eth_getTransactionCount&address=0x2910543af39aba0cd09dbb2d50200b3e800a63d2&tag=latest&apikey=YourApiKeyToken
	tag := "latest"
	if blockNumber != nil {
		tag = "0x" + blockNumber.Text(16)
	}
	url := fmt.Sprintf("%s?%s&action=eth_getTransactionCount&address=%s&tag=%s&apikey=%s", e.url, e.module, account.Hex(), tag, e.apiKey)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return 0, err
	}
	return client.NonceAt(ctx, account, blockNumber)
}

func (e *EthClientApiKeyImp) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	tag := "pending"
	url := fmt.Sprintf("%s?%s&action=eth_getTransactionCount&address=%s&tag=%s&apikey=%s", e.url, e.module, account.Hex(), tag, e.apiKey)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return 0, err
	}
	return client.PendingNonceAt(ctx, account)
}

func (e *EthClientApiKeyImp) PendingTransactionCount(ctx context.Context) (uint, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	panic("implement me")
}

func (e *EthClientApiKeyImp) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	//http://api-cn.etherscan.com/api?module=proxy&action=eth_sendRawTransaction&hex=0xf904808000831cfde080&apikey=YourApiKeyToken
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s?%s&action=eth_sendRawTransaction&hex=%s&apikey=%s", e.url, e.module, hexutil.Encode(data), e.apiKey)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return err
	}
	return client.SendTransaction(ctx, tx)
}
