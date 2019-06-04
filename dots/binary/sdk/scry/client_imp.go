// Scry Info.  All rights reserved.
// license that can be found in the license file.

package scry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/binary/sdk/core/chainevents"
	"github.com/scryinfo/dp/dots/binary/sdk/core/chainoperations"
	"github.com/scryinfo/dp/dots/service"
	"go.uber.org/zap"
	"math/big"
)

type clientImp struct {
	account      *service.Account
	chainWrapper ChainWrapper `dot:""`
}

func NewScryClient(publicKey string, chainWrapper ChainWrapper) Client {
	return &clientImp{
		account:      &service.Account{Address: publicKey},
		chainWrapper: chainWrapper,
	}
}

func CreateScryClient(password string, chainWrapper ChainWrapper) (Client, error) {
	account, err := service.GetAMIns().CreateAccount(password)
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("failed to create account, error:", err))
		return nil, err
	}

	return &clientImp{
		account:      account,
		chainWrapper: chainWrapper,
	}, nil
}

func (c *clientImp) Account() *service.Account {
	return c.account
}

func (c *clientImp) SubscribeEvent(eventName string, callback chainevents.EventCallback) error {
	return chainevents.SubscribeExternal(common.HexToAddress(c.account.Address), eventName, callback)
}

func (c *clientImp) UnSubscribeEvent(eventName string) error {
	return chainevents.UnSubscribeExternal(common.HexToAddress(c.account.Address), eventName)
}

func (c *clientImp) Authenticate(password string) (bool, error) {
	return service.GetAMIns().AuthAccount(c.account.Address, password)
}

func (c *clientImp) TransferEthFrom(from common.Address, password string, value *big.Int, ec *ethclient.Client) error {
	tx, err := chainoperations.TransferEth(from, password, common.HexToAddress(c.account.Address), value, ec)
	if err == nil {
		dot.Logger().Debugln("transferEthFrom: " + tx.Hash().String() + string(tx.Data()))
	}

	return err
}

func (c *clientImp) TransferTokenFrom(from common.Address, password string, value *big.Int) error {
	txParam := &chainoperations.TransactParams{From: from, Password: password, Value: value}
	return c.chainWrapper.TransferTokens(txParam,
		common.HexToAddress(c.account.Address),
		value)
}

func (c *clientImp) GetEth(owner common.Address, ec *ethclient.Client) (*big.Int, error) {
	return chainoperations.GetEthBalance(owner, ec)
}

func (c *clientImp) GetScryToken(owner common.Address) (*big.Int, error) {
	from := common.HexToAddress(c.account.Address)
	txParam := &chainoperations.TransactParams{From: from, Pending: true}

	return c.chainWrapper.GetTokenBalance(txParam, owner)
}
