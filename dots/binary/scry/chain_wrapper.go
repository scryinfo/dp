// Scry Info.  All rights reserved.
// license that can be found in the license file.

package scry

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    tx "github.com/scryinfo/dp/dots/eth/transaction"
    "math/big"
)

type ChainWrapper interface {
    Conn() *ethclient.Client
    Publish(txParams *tx.TxParams, price *big.Int, metaDataID []byte, proofDataIDs []string,
        proofNum int, detailsID string, supportVerify bool) (string, error)
    PrepareToBuy(txParams *tx.TxParams, publishId string, startVerify bool) error
    BuyData(txParams *tx.TxParams, txId *big.Int) error
    CancelTransaction(txParams *tx.TxParams, txId *big.Int) error
    SubmitMetaDataIdEncWithBuyer(txParams *tx.TxParams, txId *big.Int, encyptedMetaDataId []byte) error
    ConfirmDataTruth(txParams *tx.TxParams, txId *big.Int, truth bool) error
    ApproveTransfer(txParams *tx.TxParams, spender common.Address, value *big.Int) error
    Vote(txParams *tx.TxParams, txId *big.Int, judge bool, comments string) error
    RegisterAsVerifier(txParams *tx.TxParams) error
    CreditsToVerifier(txParams *tx.TxParams, txId *big.Int, index uint8, credit uint8) error
    TransferTokens(txParams *tx.TxParams, to common.Address, value *big.Int) error
    GetTokenBalance(txParams *tx.TxParams, owner common.Address) (*big.Int, error)
}
