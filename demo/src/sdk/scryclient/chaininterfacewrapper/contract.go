package chaininterfacewrapper

import (
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	op "github.com/scryInfo/dp/demo/src/sdk/core/chainoperations"
	"github.com/scryInfo/dp/demo/src/sdk/interface/contractinterface"
	"github.com/scryInfo/dp/demo/src/sdk/settings"
	"github.com/scryInfo/dp/demo/src/sdk/util/accounts"
	"github.com/scryInfo/dp/demo/src/sdk/util/uuid"
	rlog "github.com/sirupsen/logrus"
	"math/big"
)

var (
	conn         *ethclient.Client
	scryProtocol *contractinterface.ScryProtocol
	scryToken    *contractinterface.ScryToken
)

func Initialize(protocolContractAddress common.Address, tokenContractAddress common.Address, clientConn *ethclient.Client) error {
	var err error = nil

	scryProtocol, err = contractinterface.NewScryProtocol(protocolContractAddress, clientConn)
	if err != nil {
		rlog.Error("failed to initialize protocol contract interface wrapper.", err)
		return err
	}

	scryToken, err = contractinterface.NewScryToken(tokenContractAddress, clientConn)
	if err != nil {
		rlog.Error("failed to initialize token contract interface wrapper.", err)
		return err
	}

	conn = clientConn
	return nil
}

func Publish(txParams *op.TransactParams, price *big.Int, metaDataID []byte, proofDataIDs []string,
	proofNum int, detailsID string, supportVerify bool) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := uuid.GenerateUUID()

	pdIDs := make([][32]byte, proofNum)
	var err error = nil
	for i := 0;i < proofNum;i++ {
		pdIDs[i], err = ipfsHashToBytes32(proofDataIDs[i])
		if err != nil {
			rlog.Error("failed to convert ipfs hash to bytes32")
			return "", err
		}
	}

	encMetaId, err := accounts.GetAMInstance().Encrypt(metaDataID, txParams.From.String())
	if err != nil {
		rlog.Error("failed to encrypt meta data hash, error: ", err)
		return "", err
	}

	//upload meta_data_id_enc_seller and other cids to contracts
	tx, err := scryProtocol.PublishDataInfo(op.BuildTransactOpts(txParams), getAppSeqNo(), publishId, price,
		encMetaId, pdIDs, detailsID, supportVerify)
	if err != nil {
		rlog.Error("failed to publish data information, error: ", err)
		return "", err
	}

	rlog.Debug("publish transaction:" + string(tx.Data()) + " hash:" + tx.Hash().String())

	return publishId, nil
}

func ipfsHashToBytes32(src string) ([32]byte, error) {
	hashArray1 := base58.Decode(src)
	if len(hashArray1) != 34 {
		var nilArray [32]byte
		return nilArray, errors.New("invalid ipfs hash")
	}

	var hashArray2 [32]byte
	copy(hashArray2[:], hashArray1[2:])
	return hashArray2, nil
}

func Bytes32ToIpfsHash(value [32]byte) (string, error) {
    byteArray := [34]byte {18, 32}
    copy(byteArray[2:], value[:])
    if len(byteArray) != 34 {
        return "", errors.New("invalid bytes32 value")
    }

    hash := base58.Encode(byteArray[:])
    return hash, nil
}

func PrepareToBuy(txParams *op.TransactParams, publishId string, startVerify bool) error {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to prepare to buy , error:", err)
		}
	}()

	tx, err := scryProtocol.CreateTransaction(op.BuildTransactOpts(txParams), getAppSeqNo(), publishId, startVerify)
	if err == nil {
		rlog.Debug("CreateTransaction:" + string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func BuyData(txParams *op.TransactParams, txId *big.Int) error {
	tx, err := scryProtocol.BuyData(op.BuildTransactOpts(txParams), getAppSeqNo(), txId)
	if err == nil {
		rlog.Debug("BuyData:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func CancelTransaction(txParams *op.TransactParams, txId *big.Int) error {
	tx, err := scryProtocol.CancelTransaction(op.BuildTransactOpts(txParams), getAppSeqNo(), txId)
	if err == nil {
		rlog.Debug("CancelTransaction", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func SubmitMetaDataIdEncWithBuyer(txParams *op.TransactParams, txId *big.Int, encyptedMetaDataId []byte) error {
	tx, err := scryProtocol.SubmitMetaDataIdEncWithBuyer(op.BuildTransactOpts(txParams), getAppSeqNo(), txId, encyptedMetaDataId)
	if err == nil {
		rlog.Debug("SubmitMetaDataIdEncWithBuyer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ConfirmDataTruth(txParams *op.TransactParams, txId *big.Int, truth bool) error {
	tx, err := scryProtocol.ConfirmDataTruth(op.BuildTransactOpts(txParams), getAppSeqNo(), txId, truth)
	if err == nil {
		rlog.Debug("ConfirmDataTruth:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ApproveTransfer(txParams *op.TransactParams, spender common.Address, value *big.Int) error {
	tx, err := scryToken.Approve(op.BuildTransactOpts(txParams), spender, value)
	if err == nil {
		rlog.Debug("ApproveTransfer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func Vote(txParams *op.TransactParams, txId *big.Int, judge bool, comments string) error {
	tx, err := scryProtocol.Vote(op.BuildTransactOpts(txParams), getAppSeqNo(), txId, judge, comments)
	if err == nil {
		rlog.Debug("Vote:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func RegisterAsVerifier(txParams *op.TransactParams) error {
	tx, err := scryProtocol.RegisterAsVerifier(op.BuildTransactOpts(txParams), getAppSeqNo())
	if err == nil {
		rlog.Debug("RegisterAsVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func CreditsToVerifier(txParams *op.TransactParams, txId *big.Int, index uint8, credit uint8) error {
	tx, err := scryProtocol.CreditsToVerifier(op.BuildTransactOpts(txParams), getAppSeqNo(), txId, index, credit)
	if err == nil {
		rlog.Debug("CreditsToVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func TransferTokens(txParams *op.TransactParams, to common.Address, value *big.Int) error {
	tx, err := scryToken.Transfer(op.BuildTransactOpts(txParams), to, value)
	if err == nil {
		rlog.Debug("tx:", tx)
		return err
	}

	return err
}

func GetTokenBalance(txParams *op.TransactParams, owner common.Address) (*big.Int, error) {
	return scryToken.BalanceOf(op.BuildCallOpts(txParams), owner)
}

func TransferEth(from common.Address,
	password string,
	to common.Address,
	value *big.Int) (*types.Transaction, error) {
	return op.TransferEth(from, password, to, value, conn)
}

func GetEthBalance(owner common.Address) (*big.Int, error) {
	return op.GetEthBalance(owner, conn)
}

func getAppSeqNo() string {
    return settings.GetAppId()
}