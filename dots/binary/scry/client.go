// Scry Info.  All rights reserved.
// license that can be found in the license file.

package scry

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/scryinfo/dp/dots/auth"
    "github.com/scryinfo/dp/dots/eth/event"
    "math/big"
)

type Client interface {
    Account() *auth.UserAccount
    SubscribeEvent(eventName string, callback event.Callback) error
    UnSubscribeEvent(eventName string) error
    Authenticate(password string) (bool, error)
    TransferEthFrom(from common.Address, password string, value *big.Int, ec *ethclient.Client) error
    TransferTokenFrom(from common.Address, password string, value *big.Int) error
    GetEth(owner common.Address, ec *ethclient.Client) (*big.Int, error)
    GetToken(owner common.Address) (*big.Int, error)
}
