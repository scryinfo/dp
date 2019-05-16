package scry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/scryinfo/dp/dots/binary/sdk/core/chainoperations"
	"math/big"
)

type ChainWrapper interface {
	Publish(txParams *chainoperations.TransactParams, price *big.Int, metaDataID []byte, proofDataIDs []string,
		proofNum int, detailsID string, supportVerify bool) (string, error)
	PrepareToBuy(txParams *chainoperations.TransactParams, publishId string, startVerify bool) error
	BuyData(txParams *chainoperations.TransactParams, txId *big.Int) error
	CancelTransaction(txParams *chainoperations.TransactParams, txId *big.Int) error
	SubmitMetaDataIdEncWithBuyer(txParams *chainoperations.TransactParams, txId *big.Int, encyptedMetaDataId []byte) error
	ConfirmDataTruth(txParams *chainoperations.TransactParams, txId *big.Int, truth bool) error
	ApproveTransfer(txParams *chainoperations.TransactParams, spender common.Address, value *big.Int) error
	Vote(txParams *chainoperations.TransactParams, txId *big.Int, judge bool, comments string) error
	RegisterAsVerifier(txParams *chainoperations.TransactParams) error
	CreditsToVerifier(txParams *chainoperations.TransactParams, txId *big.Int, index uint8, credit uint8) error
	TransferTokens(txParams *chainoperations.TransactParams, to common.Address, value *big.Int) error
	GetTokenBalance(txParams *chainoperations.TransactParams, owner common.Address) (*big.Int, error)
	TransferEth(from common.Address,
		password string,
		to common.Address,
		value *big.Int) (*types.Transaction, error)
	GetEthBalance(owner common.Address) (*big.Int, error)
}
