package scry

import (
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryInfo/dp/dots/binary/sdk/core/chainoperations"
	"github.com/scryInfo/dp/dots/binary/sdk/interface/contract"
	"github.com/scryInfo/dp/dots/binary/sdk/settings"
	"github.com/scryInfo/dp/dots/binary/sdk/util/accounts"
	"github.com/scryInfo/dp/util"
	rlog "github.com/sirupsen/logrus"
	"math/big"
)

type chainWrapperImp struct {
	conn         *ethclient.Client
	scryProtocol *contract.ScryProtocol
	scryToken    *contract.ScryToken
}


func NewChainWrapper(protocolContractAddress common.Address, tokenContractAddress common.Address, clientConn *ethclient.Client) (ChainWrapper, error) {
	var err error = nil
	c := &chainWrapperImp{}


	c.scryProtocol, err = contract.NewScryProtocol(protocolContractAddress, clientConn)
	if err != nil {
		rlog.Error("failed to initialize protocol contract interface wrapper.", err)
		return nil, err
	}

	c.scryToken, err = contract.NewScryToken(tokenContractAddress, clientConn)
	if err != nil {
		rlog.Error("failed to initialize token contract interface wrapper.", err)
		return nil, err
	}

	c.conn = clientConn

	return c,err
}

func (c *chainWrapperImp)Publish(txParams *chainoperations.TransactParams, price *big.Int, metaDataID []byte, proofDataIDs []string,
	proofNum int, detailsID string, supportVerify bool) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := util.GenerateUUID()

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
	tx, err := c.scryProtocol.PublishDataInfo(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), publishId, price,
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

func (c *chainWrapperImp)PrepareToBuy(txParams *chainoperations.TransactParams, publishId string, startVerify bool) error {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to prepare to buy , error:", err)
		}
	}()

	tx, err := c.scryProtocol.CreateTransaction(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), publishId, startVerify)
	if err == nil {
		rlog.Debug("CreateTransaction:" + string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)BuyData(txParams *chainoperations.TransactParams, txId *big.Int) error {
	tx, err := c.scryProtocol.BuyData(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId)
	if err == nil {
		rlog.Debug("BuyData:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)CancelTransaction(txParams *chainoperations.TransactParams, txId *big.Int) error {
	tx, err := c.scryProtocol.CancelTransaction(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId)
	if err == nil {
		rlog.Debug("CancelTransaction", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)SubmitMetaDataIdEncWithBuyer(txParams *chainoperations.TransactParams, txId *big.Int, encyptedMetaDataId []byte) error {
	tx, err := c.scryProtocol.SubmitMetaDataIdEncWithBuyer(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, encyptedMetaDataId)
	if err == nil {
		rlog.Debug("SubmitMetaDataIdEncWithBuyer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)ConfirmDataTruth(txParams *chainoperations.TransactParams, txId *big.Int, truth bool) error {
	tx, err := c.scryProtocol.ConfirmDataTruth(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, truth)
	if err == nil {
		rlog.Debug("ConfirmDataTruth:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)ApproveTransfer(txParams *chainoperations.TransactParams, spender common.Address, value *big.Int) error {
	tx, err := c.scryToken.Approve(chainoperations.BuildTransactOpts(txParams), spender, value)
	if err == nil {
		rlog.Debug("ApproveTransfer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)Vote(txParams *chainoperations.TransactParams, txId *big.Int, judge bool, comments string) error {
	tx, err := c.scryProtocol.Vote(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, judge, comments)
	if err == nil {
		rlog.Debug("Vote:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)RegisterAsVerifier(txParams *chainoperations.TransactParams) error {
	tx, err := c.scryProtocol.RegisterAsVerifier(chainoperations.BuildTransactOpts(txParams), getAppSeqNo())
	if err == nil {
		rlog.Debug("RegisterAsVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)CreditsToVerifier(txParams *chainoperations.TransactParams, txId *big.Int, index uint8, credit uint8) error {
	tx, err := c.scryProtocol.CreditsToVerifier(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, index, credit)
	if err == nil {
		rlog.Debug("CreditsToVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func (c *chainWrapperImp)TransferTokens(txParams *chainoperations.TransactParams, to common.Address, value *big.Int) error {
	tx, err := c.scryToken.Transfer(chainoperations.BuildTransactOpts(txParams), to, value)
	if err == nil {
		rlog.Debug("tx:", tx)
		return err
	}

	return err
}

func (c *chainWrapperImp)GetTokenBalance(txParams *chainoperations.TransactParams, owner common.Address) (*big.Int, error) {
	return c.scryToken.BalanceOf(chainoperations.BuildCallOpts(txParams), owner)
}

func (c *chainWrapperImp)TransferEth(from common.Address,
	password string,
	to common.Address,
	value *big.Int) (*types.Transaction, error) {
	return chainoperations.TransferEth(from, password, to, value, c.conn)
}

func (c *chainWrapperImp)GetEthBalance(owner common.Address) (*big.Int, error) {
	return chainoperations.GetEthBalance(owner, c.conn)
}

func getAppSeqNo() string {
    return settings.GetAppId()
}