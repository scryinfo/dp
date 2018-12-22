package transactions

import (
	"../../business/contractinterface"
	"github.com/ethereum/go-ethereum/common"
)

func PrepareToBuy(buyerAddr string, publishId string) (txId common.Hash, err error) {
	return contractinterface.PrepareToBuy(buyerAddr, publishId)
}

func BuyData(txId common.Hash) (err error) {
	return contractinterface.BuyData(txId)
}

func SubmitMetaDataIdEncWithBuyer(txId common.Hash, encyptedMetaDataId []byte) (err error) {
	return contractinterface.SubmitMetaDataIdEncWithBuyer(txId, encyptedMetaDataId)
}

func ConfirmDataTruth(txId common.Hash, truth bool) (err error) {
	return contractinterface.ConfirmDataTruth(txId, truth)
}
