package transaction

import "../definition"

type TransactionOperationInterface interface {
	PrepareToBuy() (txId int, err error)
	GetDataDescriptionList() (despList *[]definition.DescriptionData, err error)
	GetProofDataList() (proofList *[]definition.ProofData, err error)
	//VoteAndCommentForMetaData(txId string, vote definition.Vote) (err error)
	//GetVoteAndComments() (vote *[]definition.Vote, err error)
	BuyData(txId int) (err error)
	SubmitMetaDataIdEncWithBuyer(txId int, encyptedMetaDataId []byte) (err error)
	ConfirmDataTruth(truth bool) (err error)
	//DoArbitrate() ()
	//VoteToVerifier() ()
}