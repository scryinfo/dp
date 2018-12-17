package transaction

import "../definition"

type TransactionExecutor struct {
	txList map[uint64]*Transaction
}

func (business *TransactionExecutor) PrepareToBuy(buyerAddr string, publishId string) (txId int, err error) {
	//request to contract
}

func (business *TransactionExecutor) GetDataDescriptionList() (despList *[]definition.DescriptionData, err error) {

}

func (business *TransactionExecutor) GetProofDataList() (proofList *[]definition.ProofData, err error) {

}

func (business *TransactionExecutor) BuyData(txId int) (err error) {

}

func (business *TransactionExecutor) SubmitMetaDataIdEncWithBuyer(txId int, encyptedMetaDataId []byte) (err error) {

}

func (business *TransactionExecutor) ConfirmDataTruth(truth bool) (err error) {

}