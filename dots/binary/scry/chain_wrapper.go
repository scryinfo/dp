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
        proofNum int32, detailsID string, supportVerify bool) (string, error)
    AdvancePurchase(txParams *tx.TxParams, publishId string, startVerify bool) error
    ConfirmPurchase(txParams *tx.TxParams, txId *big.Int) error
    CancelPurchase(txParams *tx.TxParams, txId *big.Int) error
    ReEncrypt(txParams *tx.TxParams, txId *big.Int, encodedData []byte) error
    ConfirmData(txParams *tx.TxParams, txId *big.Int, truth bool) error
    ApproveTransfer(txParams *tx.TxParams, spender common.Address, value *big.Int) error
    Vote(txParams *tx.TxParams, txId *big.Int, judge bool, comments string) error
    RegisterAsVerifier(txParams *tx.TxParams) error
    GradeToVerifier(txParams *tx.TxParams, txId *big.Int, index uint8, credit uint8) error
    Arbitrate(txParams *tx.TxParams, txId *big.Int, judge bool) error

    GetBuyer(txParams *tx.TxParams, txId *big.Int) (string, error)
    GetArbitrators(txParams *tx.TxParams, txId *big.Int) ([]string, error)

    TransferTokens(txParams *tx.TxParams, to common.Address, value *big.Int) error
    GetTokenBalance(txParams *tx.TxParams, owner common.Address) (*big.Int, error)
}
