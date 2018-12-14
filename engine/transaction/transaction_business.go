package transaction

import "../definition"

type TransactionBusiness struct {
	
}

func (business *TransactionBusiness) PrepareToBuy() (txId int, err error) {

}

func (business *TransactionBusiness) GetDataDescriptionList() (despList *[]definition.DescriptionData, err error) {

}

func (business *TransactionBusiness) GetProofDataList() (proofList *[]definition.ProofData, err error) {

}

func (business *TransactionBusiness) BuyData(txId int) (err error) {

}

func (business *TransactionBusiness) SubmitMetaDataIdEncWithBuyer(txId int, encyptedMetaDataId []byte) (err error) {

}

func (business *TransactionBusiness) ConfirmDataTruth(truth bool) (err error) {

}