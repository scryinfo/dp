package contractclient

import (
    "../core/chainevents"
    "github.com/ethereum/go-ethereum/accounts"
    "github.com/ethereum/go-ethereum/common"
)


type ContractClient struct {
	Address string
	IdentityStream string
	account accounts.Account
}

func NewContractClient(publicKey string,
					identityStream string) (*ContractClient, error) {

	return &ContractClient{
		Address: publicKey,
		IdentityStream: identityStream,
	}, nil
}

func (client *ContractClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
	chainevents.SubscribeExternal(common.HexToAddress(client.Address), eventName, callback)
}
