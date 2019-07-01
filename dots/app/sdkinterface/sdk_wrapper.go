// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sdkinterface

import (
	"github.com/scryinfo/dp/dots/app/settings"
	"math/big"

	"github.com/scryinfo/dp/dots/eth/event"
)

// wrap sdk interface call.
type SDKWrapper interface {
	// auth component
	CreateUserWithLogin(password string) (string, error)
	UserLogin(address string, password string) (bool, error)

	// transfer from deployer
	TransferEthFromDeployer(eth *big.Int) error
	TransferTokenFromDeployer(token *big.Int) error

	// listen component
	SubscribeEvents(eventName []string, cb ...event.Callback) error
	UnsubscribeEvents(eventName []string) error

	// app functions
	PublishData(data *settings.PublishData) (string, error)
	ApproveTransferToken(password string, quantity *big.Int) error
	CreateTransaction(publishId string, password string, startVerify bool) error
	Buy(txId string, password string) error
	ReEncryptMetaDataIdFromSeller(txId string, password, seller string, metaDataIDEncSeller []byte) error
	CancelTransaction(txId, password string) error
	DecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error)
	ConfirmDataTruth(txId string, password string, truth bool) error
	RegisterAsVerifier(password string) error
	Vote(password, txId string, judge bool, comment string) error
	CreditToVerifiers(creditData *settings.CreditData) error
	GetEthBalance(password string) (string, error)
	GetTokenBalance(password string) (string, error)
}
