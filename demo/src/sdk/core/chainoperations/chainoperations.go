package chainoperations

import (
    "../../util/accounts"
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    rlog "github.com/sirupsen/logrus"
    "math/big"
    "strings"
)

var (
    GAS_LIMIT uint64 = 21000
)


type TransactParams struct {
    From common.Address
    Password string
    Value *big.Int
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
        GasPrice: big.NewInt(0),
		GasLimit: 0,
		Context: context.Background(),
	}

	return opts
}

func SignTransaction(signer types.Signer, address common.Address,
                        transaction *types.Transaction, password string) (*types.Transaction, error) {
    h := signer.Hash(transaction)

    var sign []byte
    var err error

    //hardcode here. will dropped later.
    if strings.ToLower(address.String()) == "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8" {
        keyJson := `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
         key, err:= keystore.DecryptKey([]byte(keyJson), "12345")

        if err != nil {
            rlog.Error("failed to sign transaction, error:", err)
            return nil, err
        }

        sign, err = crypto.Sign(h[:], key.PrivateKey)
    } else {
        sign, err = accounts.GetAMInstance().SignTransaction(h[:], address.String(), password)
        if err != nil {
            return nil, err
        }
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

func TransferEth(from common.Address,
                password string,
                to common.Address,
                value uint64,
                client *ethclient.Client) (*types.Transaction, error) {
    txParam := &TransactParams{from, password, value}
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
