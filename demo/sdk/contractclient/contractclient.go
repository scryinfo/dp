package contractclient

import (
    "../core/chainevents"
    "../util/accounts"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
)

type ContractClient struct {
	account *accounts.Account
}

func NewContractClient(publicKey string) (*ContractClient, error) {
    return &ContractClient{
		account: &accounts.Account{publicKey},
	}, nil
}

func CreateContractClient(publicKey string, password string) (*ContractClient, error) {
    account, err := accounts.GetAMInstance().CreateAccount(password)
    if err != nil {
        fmt.Println("failed to create account, error:", err)
        return nil, err
    }

    return &ContractClient{
        account: account,
    }, nil
}

func (client ContractClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
	chainevents.SubscribeExternal(common.HexToAddress(client.account.Address), eventName, callback)
}

func (client ContractClient) Authticate(password string) (bool, error) {
    return accounts.GetAMInstance().AuthAccount(client.account.Address, password)
}

func (client *ContractClient) SwitchAccount(pubkey string) error {
    if !accounts.GetAMInstance().AccountValid(pubkey) {
        return errors.New("The public key is not valid")
    }

    client.account.Address = pubkey
    return nil
}

