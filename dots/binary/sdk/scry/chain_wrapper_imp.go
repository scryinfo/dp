// Scry Info.  All rights reserved.
// license that can be found in the license file.

package scry

import (
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/binary/sdk/core/chainoperations"
	"github.com/scryinfo/dp/dots/binary/sdk/interface/contract"
	"github.com/scryinfo/dp/dots/binary/sdk/settings"
	"github.com/scryinfo/dp/dots/service"
	"github.com/scryinfo/dp/util"
	"go.uber.org/zap"
	"math/big"
)

type chainWrapperImp struct {
	conn         *ethclient.Client
	scryProtocol *contract.ScryProtocol
	scryToken    *contract.ScryToken
}

func NewChainWrapper(protocolContractAddress common.Address,
	tokenContractAddress common.Address,
	clientConn *ethclient.Client,
) (ChainWrapper, error) {
	var err error = nil
	c := &chainWrapperImp{}

	c.scryProtocol, err = contract.NewScryProtocol(protocolContractAddress, clientConn)
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("failed to initialize protocol contract interface wrapper.", err))
		return nil, err
	}

	c.scryToken, err = contract.NewScryToken(tokenContractAddress, clientConn)
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("failed to initialize token contract interface wrapper.", err))
		return nil, err
	}

	c.conn = clientConn

	return c, err
}

func (c *chainWrapperImp) Publish(txParams *chainoperations.TransactParams, price *big.Int, metaDataID []byte,
	proofDataIDs []string, proofNum int, detailsID string, supportVerify bool) (string, error) {
	logger := dot.Logger()

	defer func() {
		if er := recover(); er != nil {
			logger.Errorln("", zap.Any("failed to publish data, error:", er))
		}
	}()

	//generate publishId
	publishId := util.GenerateUUID()

	pdIDs := make([][32]byte, proofNum)
	var err error = nil
	for i := 0; i < proofNum; i++ {
		pdIDs[i], err = ipfsHashToBytes32(proofDataIDs[i])
		if err != nil {
			logger.Errorln("failed to convert ipfs hash to bytes32")
			return "", err
		}
	}

	encMetaId, err := service.GetAMIns().Encrypt(metaDataID, txParams.From.String())
	if err != nil {
		logger.Errorln("", zap.NamedError("failed to encrypt meta data hash, error: ", err))
		return "", err
	}

	tx, err := c.scryProtocol.PublishDataInfo(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), publishId, price,
		encMetaId, pdIDs, detailsID, supportVerify)
	if err != nil {
		logger.Errorln("", zap.NamedError("failed to publish data information, error: ", err))
		return "", err
	}

	logger.Debugln("publish transaction: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))

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
	byteArray := [34]byte{18, 32}
	copy(byteArray[2:], value[:])
	if len(byteArray) != 34 {
		return "", errors.New("invalid bytes32 value")
	}

	hash := base58.Encode(byteArray[:])
	return hash, nil
}

func (c *chainWrapperImp) PrepareToBuy(txParams *chainoperations.TransactParams, publishId string, startVerify bool) error {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to prepare to buy , error:", er))
		}
	}()

	tx, err := c.scryProtocol.CreateTransaction(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), publishId, startVerify)
	if err == nil {
		dot.Logger().Debugln("CreateTransaction: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) BuyData(txParams *chainoperations.TransactParams, txId *big.Int) error {
	tx, err := c.scryProtocol.BuyData(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId)
	if err == nil {
		dot.Logger().Debugln("BuyData: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) CancelTransaction(txParams *chainoperations.TransactParams, txId *big.Int) error {
	tx, err := c.scryProtocol.CancelTransaction(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId)
	if err == nil {
		dot.Logger().Debugln("CancelTransaction tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) SubmitMetaDataIdEncWithBuyer(txParams *chainoperations.TransactParams, txId *big.Int, encyptedMetaDataId []byte) error {
	tx, err := c.scryProtocol.SubmitMetaDataIdEncWithBuyer(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, encyptedMetaDataId)
	if err == nil {
		dot.Logger().Debugln("SubmitMetaDataIdEncWithBuyer: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) ConfirmDataTruth(txParams *chainoperations.TransactParams, txId *big.Int, truth bool) error {
	tx, err := c.scryProtocol.ConfirmDataTruth(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, truth)
	if err == nil {
		dot.Logger().Debugln("ConfirmDataTruth: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) ApproveTransfer(txParams *chainoperations.TransactParams, spender common.Address, value *big.Int) error {
	tx, err := c.scryToken.Approve(chainoperations.BuildTransactOpts(txParams), spender, value)
	if err == nil {
		dot.Logger().Debugln("ApproveTransfer: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) Vote(txParams *chainoperations.TransactParams, txId *big.Int, judge bool, comments string) error {
	tx, err := c.scryProtocol.Vote(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, judge, comments)
	if err == nil {
		dot.Logger().Debugln("Vote: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))

	}

	return err
}

func (c *chainWrapperImp) RegisterAsVerifier(txParams *chainoperations.TransactParams) error {
	tx, err := c.scryProtocol.RegisterAsVerifier(chainoperations.BuildTransactOpts(txParams), getAppSeqNo())
	if err == nil {
		dot.Logger().Debugln("RegisterAsVerifier: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) CreditsToVerifier(txParams *chainoperations.TransactParams, txId *big.Int, index uint8, credit uint8) error {
	tx, err := c.scryProtocol.CreditsToVerifier(chainoperations.BuildTransactOpts(txParams), getAppSeqNo(), txId, index, credit)
	if err == nil {
		dot.Logger().Debugln("CreditsToVerifier: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) TransferTokens(txParams *chainoperations.TransactParams, to common.Address, value *big.Int) error {
	tx, err := c.scryToken.Transfer(chainoperations.BuildTransactOpts(txParams), to, value)
	if err == nil {
		dot.Logger().Debugln("TransferTokens: tx hash:" + tx.Hash().String(), zap.Binary(" tx data:", tx.Data()))
	}

	return err
}

func (c *chainWrapperImp) GetTokenBalance(txParams *chainoperations.TransactParams, owner common.Address) (*big.Int, error) {
	return c.scryToken.BalanceOf(chainoperations.BuildCallOpts(txParams), owner)
}

func (c *chainWrapperImp) TransferEth(from common.Address,
	password string,
	to common.Address,
	value *big.Int) (*types.Transaction, error) {
	return chainoperations.TransferEth(from, password, to, value, c.conn)
}

func (c *chainWrapperImp) GetEthBalance(owner common.Address) (*big.Int, error) {
	return chainoperations.GetEthBalance(owner, c.conn)
}

func getAppSeqNo() string {
	return settings.GetAppId()
}
