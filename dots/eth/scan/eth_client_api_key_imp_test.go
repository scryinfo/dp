package scan

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	apiKey = "" //not commit
	url    = "http://api-cn.etherscan.com/api"
)

func TestEthClientApiKeyImp_BlockByNumber(t *testing.T) {
	conn := NewEthClientApiKeyImp(url, apiKey, "")
	block, err := conn.BlockByNumber(context.Background(), nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, block)
	block, err = conn.BlockByNumber(context.Background(), big.NewInt(3))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, block)
}

func TestEthClientApiKeyImp_TransactionByHash(t *testing.T) {
	conn := NewEthClientApiKeyImp(url, apiKey, "")
	hash := common.HexToHash("0x43caf8ba9b1f9c79f41ddf8be41da8ca84fad919b50ed48a4071ff3d7c35814e")
	tx, p, err := conn.TransactionByHash(context.Background(), hash)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, tx)
	assert.Equal(t, false, p)

	hash = common.HexToHash("0x96d72b5c39f54a725f0367e1d67f25b56938f0e35508571f93476e2022ec5a8c")
	tx, p, err = conn.TransactionByHash(context.Background(), hash)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, tx)
	assert.Equal(t, false, p)
}

func TestEthClientApiKeyImp_TransactionReceipt(t *testing.T) {
	conn := NewEthClientApiKeyImp(url, apiKey, "")
	hash := common.HexToHash("0x43caf8ba9b1f9c79f41ddf8be41da8ca84fad919b50ed48a4071ff3d7c35814e")
	receipt, err := conn.TransactionReceipt(context.Background(), hash)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, receipt)
	assert.Equal(t, uint64(1), receipt.Status)
	assert.Equal(t, uint64(21000), receipt.GasUsed)
	assert.Equal(t, big.NewInt(9739302), receipt.BlockNumber)

	hash = common.HexToHash("0x96d72b5c39f54a725f0367e1d67f25b56938f0e35508571f93476e2022ec5a8c")
	receipt, err = conn.TransactionReceipt(context.Background(), hash)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, receipt)
	assert.Equal(t, uint64(0), receipt.Status)
	assert.Equal(t, uint64(74615), receipt.GasUsed)
	assert.Equal(t, big.NewInt(9739320), receipt.BlockNumber)
}

func TestEthClientApiKeyImp_BalanceAt(t *testing.T) {
	conn := NewEthClientApiKeyImp(url, apiKey, "")
	address := common.HexToAddress("0xec610F53d0A191b17D7984c6D31598b2181Ace15")
	value, err := conn.BalanceAt(context.Background(), address, nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, value)

	address = common.HexToAddress("0xDc6e10fbAcf109Efb74E0864CDCE4876C7E729bF")
	value, err = conn.BalanceAt(context.Background(), address, nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, value)
}

func TestEthClientApiKeyImp_NonceAt(t *testing.T) {
	conn := NewEthClientApiKeyImp(url, apiKey, "")
	address := common.HexToAddress("0xec610F53d0A191b17D7984c6D31598b2181Ace15")
	nonce, err := conn.NonceAt(context.Background(), address, nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, nonce)

	address = common.HexToAddress("0xDc6e10fbAcf109Efb74E0864CDCE4876C7E729bF")
	nonce, err = conn.NonceAt(context.Background(), address, nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, nonce)
}

func TestEthClientApiKeyImp_PendingNonceAt(t *testing.T) {
	conn := NewEthClientApiKeyImp(url, apiKey, "")
	address := common.HexToAddress("0x0A98fB70939162725aE66E626Fe4b52cFF62c2e5")
	nonce, err := conn.PendingNonceAt(context.Background(), address)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, nonce)

	address = common.HexToAddress("0xDc6e10fbAcf109Efb74E0864CDCE4876C7E729bF")
	nonce, err = conn.PendingNonceAt(context.Background(), address)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, nonce)
}
