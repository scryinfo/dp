package contractclient

import (
    "../core/chainevents"
    "../core/chainoperations"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
)


type ContractClient struct {
	Address string
	IdentityStream string
	Password string
	Opts *bind.TransactOpts
}

func NewContractClient(publicKey string,
					identityStream string,
					password string) (*ContractClient, error) {

	opts := chainoperations.BuildTransactOpts(common.HexToAddress(publicKey), identityStream, password)

	return &ContractClient{
		Address: publicKey,
		IdentityStream: identityStream,
		Password: password,
		Opts: opts,
	}, nil
}

func (client *ContractClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
	chainevents.SubscribeExternal(common.HexToAddress(client.Address), eventName, callback)
}
