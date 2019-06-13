package transaction

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/auth"
    "math/big"
)

const (
    TxTypeId       = "a3e1a88e-f84e-4285-b5ff-54a16fdcd44c"
    DefaultGasLimit = 21000
)

type Transaction struct {
    Account *auth.Account `dot:"ca1c6ce4-182b-430a-9813-caeccf83f8ab"`
}

type TxParams struct {
    From     common.Address
    Password string
    Value    *big.Int
    Pending  bool
}

//construct dot
func newTxDot(conf interface{}) (dot.Dot, error) {
    d := &Transaction{}
    return d, nil
}

//Data structure needed when generating newer component
func TxTypeLive() []*dot.TypeLives {
    return []*dot.TypeLives {
        &dot.TypeLives{
            Meta: dot.Metadata{TypeId: TxTypeId,
                NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                    return newTxDot(conf)
                }},
        },
        auth.AccountTypeLive(),
    }
}

func (c *Transaction) Create(l dot.Line) error {
    return nil
}

func (c *Transaction) BuildTransactOpts(txParams *TxParams) *bind.TransactOpts {
    opts := &bind.TransactOpts{
        From:  txParams.From,
        Nonce: nil,
        Signer: func(signer types.Signer, address common.Address,
            transaction *types.Transaction) (*types.Transaction, error) {
            return c.SignTransaction(signer, address, transaction, txParams.Password)
        },
        Value:    txParams.Value,
        GasPrice: big.NewInt(0),
        GasLimit: 3000000,
        Context:  context.Background(),
    }

    return opts
}

func (c *Transaction) SignTransaction(
    signer types.Signer,
    address common.Address,
    transaction *types.Transaction,
    password string,
)(*types.Transaction, error) {
    h := signer.Hash(transaction)

    var sign []byte
    var err error

    sign, err = c.Account.SignTransaction(h[:], address.String(), password)
    if err != nil {
        return nil, err
    }

    return transaction.WithSignature(signer, sign)
}

func (c *Transaction) BuildCallOpts(txParams *TxParams) *bind.CallOpts {
    opts := &bind.CallOpts{
        Pending:     txParams.Pending,
        From:        txParams.From,
        BlockNumber: nil,
        Context:     context.Background(),
    }

    return opts
}

func (c *Transaction) Transact(
        opts *bind.TransactOpts,
        to common.Address,
        client *ethclient.Client,
    ) (*types.Transaction, error) {
    var err error

    // Ensure a valid value field and resolve the interface nonce
    value := opts.Value
    if value == nil {
        value = new(big.Int)
    }
    var nonce uint64
    if opts.Nonce == nil {
        nonce, err = client.PendingNonceAt(opts.Context, opts.From)
        if err != nil {
            return nil, fmt.Errorf("failed to retrieve interface nonce: %v", err)
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
        gasLimit = DefaultGasLimit
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