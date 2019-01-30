package contractclient

import (
    "../core/chainevents"
    "../core/accounts"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
)

type ContractClient struct {
	Account *accounts.Account
}

func NewContractClient(publicKey string) (*ContractClient, error) {
    return &ContractClient{
		Account: &accounts.Account{publicKey},
	}, nil
}

func CreateContractClient(password string) (*ContractClient, error) {
    account, err := accounts.GetAMInstance().CreateAccount(password)
    if err != nil {
        fmt.Println("failed to create Account, error:", err)
        return nil, err
    }

    return &ContractClient{
        Account: account,
    }, nil
}

func (client ContractClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
	chainevents.SubscribeExternal(common.HexToAddress(client.Account.Address), eventName, callback)
}

func (client ContractClient) Authenticate(password string) (bool, error) {
    return accounts.GetAMInstance().AuthAccount(client.Account.Address, password)
}

func (client *ContractClient) SwitchAccount(pubkey string) error {
    if !accounts.GetAMInstance().AccountValid(pubkey) {
        return errors.New("The public key is not valid")
    }

    client.Account.Address = pubkey
    return nil
}

