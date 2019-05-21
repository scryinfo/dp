// Scry Info.  All rights reserved.
// license that can be found in the license file.

package chainoperations

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	accounts2 "github.com/scryinfo/dp/dots/binary/sdk/util/accounts"
	"math/big"
	"strings"
)

var (
	GAS_LIMIT uint64 = 21000
)

type TransactParams struct {
	From     common.Address
	Password string
	Value    *big.Int
	Pending  bool
}

func DecodeKeystoreAddress(keyJsonStr []byte) string {
	addr := struct {
		Address string `json:"address"`
	}{}
	if err := json.Unmarshal(keyJsonStr, &addr); err != nil {
		panic(fmt.Sprintf("parse address fail:%v", err))
	}
	if !strings.HasPrefix(addr.Address, "0x") {
		addr.Address = `0x` + addr.Address
	}
	return addr.Address
}

func BuildTransactOpts(txParams *TransactParams) *bind.TransactOpts {
	opts := &bind.TransactOpts{
		From:  txParams.From,
		Nonce: nil,
		Signer: func(signer types.Signer, address common.Address,
			transaction *types.Transaction) (*types.Transaction, error) {
			return SignTransaction(signer, address, transaction, txParams.Password)
		},
		Value:    txParams.Value,
		GasPrice: big.NewInt(0),
		GasLimit: 3000000,
		Context:  context.Background(),
	}

	return opts
}

func SignTransaction(signer types.Signer, address common.Address,
	transaction *types.Transaction, password string) (*types.Transaction, error) {
	h := signer.Hash(transaction)

	var sign []byte
	var err error

	sign, err = accounts2.GetAMInstance().SignTransaction(h[:], address.String(), password)
	if err != nil {
		return nil, err
	}

	return transaction.WithSignature(signer, sign)
}

func BuildCallOpts(txParams *TransactParams) *bind.CallOpts {
	opts := &bind.CallOpts{
		Pending:     txParams.Pending,
		From:        txParams.From,
		BlockNumber: nil,
		Context:     context.Background(),
	}

	return opts
}

func TransferEth(from common.Address,
	password string,
	to common.Address,
	value *big.Int,
	client *ethclient.Client) (*types.Transaction, error) {
	txParam := &TransactParams{from, password, value, false}
	opts := BuildTransactOpts(txParam)
	tx, err := transact(opts, to, client)

	return tx, err
}

func transact(opts *bind.TransactOpts, to common.Address, client *ethclient.Client) (*types.Transaction, error) {
	var err error

	// Ensure a valid value field and resolve the account nonce
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = client.PendingNonceAt(opts.Context, opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}
	// Figure out the gas allowance and gas price values
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(opts.Context)
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		gasLimit = GAS_LIMIT
	}

	// Create the transaction, sign it and schedule it for execution
	var rawTx *types.Transaction
	rawTx = types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	if err := client.SendTransaction(opts.Context, signedTx); err != nil {
		return nil, err
	}

	return signedTx, nil
}

func GetEthBalance(owner common.Address, client *ethclient.Client) (*big.Int, error) {
	return client.BalanceAt(context.Background(), owner, nil)
}
