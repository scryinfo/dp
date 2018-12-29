package chainoperations

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

func DecodeKeystoreAddress(keyJsonStr []byte) common.Address {
	addr := struct {
		Address string `json:"address"`
	}{}
	if err := json.Unmarshal(keyJsonStr, &addr); err != nil {
		panic(fmt.Sprintf("parse address fail:%v", err))
	}
	if !strings.HasPrefix(addr.Address, "0x") {
		addr.Address = `0x` + addr.Address
	}
	return common.HexToAddress(addr.Address)
}

func BuildTransactOpts(from common.Address, keyJson string, keyPasswd string) *bind.TransactOpts {
	opts := &bind.TransactOpts{
		From:  from,
		Nonce: nil,
		Signer: func(signer types.Signer, addresses common.Address,
			transaction *types.Transaction) (*types.Transaction, error) {
			key, err := keystore.DecryptKey([]byte(keyJson), keyPasswd)
			if err != nil {
				return nil, err
			}
			signTransaction, err := types.SignTx(transaction, signer, key.PrivateKey)
			if err != nil {
				return nil, err
			}
			return signTransaction, nil
		},
		Value:   big.NewInt(0),
		Context: context.Background(),
	}
	return opts
}

func BuildCallOpts(pendingState bool, from common.Address, context context.Context) (*bind.CallOpts) {
	opts := &bind.CallOpts{
		pendingState,
		from,
		context,
	}

	return opts
}