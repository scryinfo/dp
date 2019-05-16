package sdkinterface

import (
	settings2 "github.com/scryinfo/dp/app/app/settings"
	chainevents2 "github.com/scryinfo/dp/dots/binary/sdk/core/chainevents"
	"math/big"
)

// wrap sdk interface call.
type SDKWrapper interface {
	CreateUserWithLogin(password string) (string, error)
	UserLogin(address string, password string) (bool, error)
	TransferTokenFromDeployer(token *big.Int) error
	SubscribeEvents(eventName []string, cb ...chainevents2.EventCallback) error
	UnsubscribeEvents(eventName []string) error
	PublishData(data *settings2.PublishData) (string, error)
	ApproveTransferToken(password string, quantity *big.Int) error
	CreateTransaction(publishId string, password string, startVerify bool) error
	Buy(txId string, password string) error
	SubmitMetaDataIdEncWithBuyer(txId string, password, seller, buyer string, metaDataIDEncSeller []byte) error
	CancelTransaction(txId, password string) error
	DecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error)
	ConfirmDataTruth(txId string, password string, truth bool) error
	RegisterAsVerifier(password string) error
	Vote(password, txId string, judge bool, comment string) error
	CreditToVerifiers(creditData *settings2.CreditData) error
}
