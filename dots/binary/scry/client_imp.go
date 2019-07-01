// Scry Info.  All rights reserved.
// license that can be found in the license file.

package scry

import (
    "errors"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/auth"
    curr "github.com/scryinfo/dp/dots/eth/currency"
    "github.com/scryinfo/dp/dots/eth/event"
    "github.com/scryinfo/dp/dots/eth/event/subscribe"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "go.uber.org/zap"
    "math/big"
)


type clientImp struct {
	userAccount  *auth.UserAccount
	chainWrapper ChainWrapper
	Subscriber   *subscribe.Subscribe `dot:"5535a065-0d90-46f4-9776-26630676c4c5"`
	Currency     *curr.Currency       `dot:"f76a1aac-ff18-479b-9d51-0166a858bec9"`
    Acct         *auth.Account        `dot:"ca1c6ce4-182b-430a-9813-caeccf83f8ab"`
}

var _ Client = (*clientImp)(nil)

func NewScryClient(publicKey string, chainWrapper ChainWrapper) Client {
	c := &clientImp{
		userAccount:  &auth.UserAccount{Addr: publicKey},
		chainWrapper: chainWrapper,
	}

    err := dot.GetDefaultLine().ToInjecter().Inject(&c)
    if err != nil {
        dot.Logger().Errorln("failed to create client", zap.Error(err))
        return nil
    }

    return c
}

func getAccountComponent() (*auth.Account, error) {
    logger := dot.Logger()

    d, err := dot.GetDefaultLine().ToInjecter().GetByLiveId(dot.LiveId(auth.AccountTypeId))
    if err != nil {
        logger.Errorln("loading Binary component failed")
        return nil, errors.New("loading Binary component failed")
    }

    if a, ok := d.(*auth.Account); ok {
        return a, nil
    } else {
        logger.Errorln("loading Binary component failed")
        return nil, errors.New("loading Binary component failed")
    }
}

func CreateScryClient(password string, chainWrapper ChainWrapper) (Client, error) {
    a, err := getAccountComponent()
    if err != nil {
        return nil, err
    }

	ua, err := a.CreateUserAccount(password)
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("failed to create client, error:", err))
		return nil, err
	}

	c := &clientImp{
        userAccount:  ua,
        chainWrapper: chainWrapper,
    }

    err = dot.GetDefaultLine().ToInjecter().Inject(&c)
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("failed to create client, error:", err))
        return nil, err
    }

	return c, nil
}

func (c *clientImp) Account() *auth.UserAccount {
	return c.userAccount
}

func (c *clientImp) SubscribeEvent(eventName string, callback event.Callback) error {
	return c.Subscriber.Subscribe(common.HexToAddress(c.Account().Addr), eventName, callback)
}

func (c *clientImp) UnSubscribeEvent(eventName string) error {
	return c.Subscriber.UnSubscribe(common.HexToAddress(c.Account().Addr), eventName)
}

func (c *clientImp) Authenticate(password string) (bool, error) {
	return c.Acct.AuthUserAccount(c.Account().Addr, password)
}

func (c *clientImp) TransferEthFrom(from common.Address, password string, value *big.Int, ec *ethclient.Client) error {
	tx, err := c.Currency.TransferEth(from, password, common.HexToAddress(c.Account().Addr), value, ec)
	if err == nil {
		dot.Logger().Debugln("transferEthFrom: " + tx.Hash().String() + string(tx.Data()))
	}

	return err
}

func (c *clientImp) TransferTokenFrom(from common.Address, password string, value *big.Int) error {
	txParam := &transaction.TxParams{From: from, Password: password, Value: value}
	return c.chainWrapper.TransferTokens(txParam,
		common.HexToAddress(c.Account().Addr),
		value)
}

func (c *clientImp) GetEth(owner common.Address, ec *ethclient.Client) (*big.Int, error) {
	return c.Currency.GetEthBalance(owner, ec)
}

func (c *clientImp) GetToken(owner common.Address) (*big.Int, error) {
	from := common.HexToAddress(c.Account().Addr)
	txParam := &transaction.TxParams{From: from, Pending: true}

	return c.chainWrapper.GetTokenBalance(txParam, owner)
}
