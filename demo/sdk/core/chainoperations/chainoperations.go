package chainoperations

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "math/big"
    "strings"
    "../../util/security"
)

var (
    secExec security.CryptExecutor
)

type TransactParams struct {
    From common.Address
    Password string
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
		Value:   big.NewInt(0),
		Context: context.Background(),
	}

	return opts
}

func SignTransaction(signer types.Signer, address common.Address,
                        transaction *types.Transaction, password string) (*types.Transaction, error) {
    h := signer.Hash(transaction)
    sign, err := secExec.SignTransaction(h[:], address.String(), password)
    if err != nil {
        return nil, err
    }

    return transaction.WithSignature(signer, sign)
}

func BuildCallOpts(pendingState bool, from common.Address, context context.Context) (*bind.CallOpts) {
	opts := &bind.CallOpts{
		pendingState,
		from,
		context,
	}

	return opts
}