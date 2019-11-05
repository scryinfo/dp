// Scry Info.  All rights reserved.
// license that can be found in the license file.

package currency

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/eth/transaction"
	"math/big"
)

const (
	// CurrTypeId
	CurrTypeId = "f76a1aac-ff18-479b-9d51-0166a858bec9"
)

// Currency
type Currency struct {
	Tx *transaction.Transaction `dot:"a3e1a88e-f84e-4285-b5ff-54a16fdcd44c"`
}

//construct dot
func newCurrDot(_ interface{}) (dot.Dot, error) {
	d := &Currency{}
	return d, nil
}

// CurrTypeLive Data structure needed when generating newer component
func CurrTypeLive() []*dot.TypeLives {

	t := []*dot.TypeLives{
		{
			Meta: dot.Metadata{TypeId: CurrTypeId,
				NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
					return newCurrDot(conf)
				}},
		},
	}

	t = append(t, transaction.TxTypeLive()...)
	return t
}

// Create
func (c *Currency) Create(l dot.Line) error {
	return nil
}

// TransferEth
func (c *Currency) TransferEth(
	from common.Address,
	password string,
	to common.Address,
	value *big.Int,
	client *ethclient.Client,
) (*types.Transaction, error) {

	txParam := &transaction.TxParams{
		From:     from,
		Password: password,
		Value:    value,
		Pending:  false,
	}

	opts := c.Tx.BuildTransactOpts(txParam)
	tx, err := c.Tx.Transact(opts, to, client)

	return tx, err
}

// GetEthBalance
func (c *Currency) GetEthBalance(
	owner common.Address,
	client *ethclient.Client,
) (*big.Int, error) {
	return client.BalanceAt(context.Background(), owner, nil)
}
