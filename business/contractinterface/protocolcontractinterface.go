package contractinterface

import "github.com/ethereum/go-ethereum/common"

func PublishDataInfo(publishId string, metaDataIdEncSeller []byte,
	proofDataId string, despDataId string, sellerAddr common.Address, supportVerify bool) (err error) {
    return nil
}

func PrepareToBuy(buyerAddr string, publishId string) (txId common.Hash, err error) {
	return nil, nil
}


func BuyData(txId common.Hash) (err error) {

}

func SubmitMetaDataIdEncWithBuyer(txId common.Hash, encyptedMetaDataId []byte) (err error) {

}

func ConfirmDataTruth(txId common.Hash, truth bool) (err error) {

}
