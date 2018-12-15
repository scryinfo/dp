package transaction

import "../definition"

type TransactionExecutor struct {
	
}

func (business *TransactionExecutor) PrepareToBuy() (txId int, err error) {

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