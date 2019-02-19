package scryclient

import (
    "../core/chainevents"
    "../util/accounts"
    "github.com/ethereum/go-ethereum/common"
    rlog "github.com/sirupsen/logrus"
)

type ScryClient struct {
	Account *accounts.Account
}

func NewScryClient(publicKey string) (*ScryClient, error) {
    return &ScryClient{
		Account: &accounts.Account{publicKey},
	}, nil
}

func CreateScryClient(password string) (*ScryClient, error) {
    account, err := accounts.GetAMInstance().CreateAccount(password)
    if err != nil {
        rlog.Error("failed to create Account, error:", err)
        return nil, err
    }

    return &ScryClient{
        Account: account,
    }, nil
}

func (client ScryClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
	chainevents.SubscribeExternal(common.HexToAddress(client.Account.Address), eventName, callback)
}

func (client ScryClient) Authenticate(password string) (bool, error) {
    return accounts.GetAMInstance().AuthAccount(client.Account.Address, password)
}

func (client ScryClient) transferEthFrom(from common.Address, value uint64) (error) {

}

func (client ScryClient) transferTokenFrom()  {

}