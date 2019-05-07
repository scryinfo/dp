package scry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryInfo/dp/dots/binary/sdk/core/chainevents"
	"github.com/scryInfo/dp/dots/binary/sdk/core/chainoperations"
	"github.com/scryInfo/dp/dots/binary/sdk/util/accounts"
	rlog "github.com/sirupsen/logrus"
	"math/big"
)

type clientImp struct {
	Account_ *accounts.Account
	Chain    ChainWrapper `dot:""`
}

func NewScryClient(publicKey string) Client {
	return &clientImp{
		Account_: &accounts.Account{publicKey},
	}
}

func CreateScryClient(password string) (Client, error) {
	account, err := accounts.GetAMInstance().CreateAccount(password)
	if err != nil {
		rlog.Error("failed to create Account_, error:", err)
		return nil, err
	}

	return &clientImp{
		Account_: account,
	}, nil
}

func (c *clientImp) Account() *accounts.Account {
	return c.Account_
}

func (c *clientImp) SubscribeEvent(eventName string, callback chainevents.EventCallback) error {
	return chainevents.SubscribeExternal(common.HexToAddress(c.Account_.Address), eventName, callback)
}

func (c *clientImp) UnSubscribeEvent(eventName string) error {
    return chainevents.UnSubscribeExternal(common.HexToAddress(c.Account_.Address), eventName)
}

func (c *clientImp) Authenticate(password string) (bool, error) {
	return accounts.GetAMInstance().AuthAccount(c.Account_.Address, password)
}

func (c *clientImp) TransferEthFrom(from common.Address, password string, value *big.Int, ec *ethclient.Client) error {
	tx, err := chainoperations.TransferEth(from, password, common.HexToAddress(c.Account_.Address), value, ec)
	if err == nil {
		rlog.Debug("transferEthFrom: ", tx.Hash(), tx.Data())
	}

	return err
}

func (c *clientImp) TransferTokenFrom(from common.Address, password string, value *big.Int) error {
	txParam := &chainoperations.TransactParams{From: from, Password: password, Value: value}
	return c.Chain.TransferTokens(txParam,
		common.HexToAddress(c.Account_.Address),
		value)
}

func (c *clientImp) GetEth(owner common.Address, ec *ethclient.Client) (*big.Int, error) {
	return chainoperations.GetEthBalance(owner, ec)
}

func (c *clientImp) GetScryToken(owner common.Address) (*big.Int, error) {
	from := common.HexToAddress(c.Account_.Address)
	txParam := &chainoperations.TransactParams{From: from, Pending: true}

	return c.Chain.GetTokenBalance(txParam, owner)
}
