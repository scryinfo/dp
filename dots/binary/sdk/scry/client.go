package scry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryInfo/dp/dots/binary/sdk/core/chainevents"
	"github.com/scryInfo/dp/dots/binary/sdk/util/accounts"
	"math/big"
)

type Client interface {
	Account() *accounts.Account
	SubscribeEvent(eventName string, callback chainevents.EventCallback) error
	UnSubscribeEvent(eventName string) error
	Authenticate(password string) (bool, error)
	TransferEthFrom(from common.Address, password string, value *big.Int, ec *ethclient.Client) error
	TransferTokenFrom(from common.Address, password string, value *big.Int) error
	GetEth(owner common.Address, ec *ethclient.Client) (*big.Int, error)
	GetScryToken(owner common.Address) (*big.Int, error)
}

