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
	"github.com/scryinfo/dp/dots/auth"
	"github.com/scryinfo/dp/dots/binary/stub/contract"
	tx "github.com/scryinfo/dp/dots/eth/transaction"
	"github.com/scryinfo/dp/util"
	"go.uber.org/zap"
	"math/big"
	"strconv"
)

type chainWrapperImp struct {
	conn     *ethclient.Client
	protocol *contractinterface.ScryProtocol
	token    *contractinterface.ScryToken
	Tx       *tx.Transaction `dot:"a3e1a88e-f84e-4285-b5ff-54a16fdcd44c"`
	Account  *auth.Account   `dot:"ca1c6ce4-182b-430a-9813-caeccf83f8ab"`
	appId    string
}

// check if 'chainWrapperImp' implements 'ChainWrapper' interface.
var _ ChainWrapper = (*chainWrapperImp)(nil)

// NewChainWrapper create a new chain wrapper
func NewChainWrapper(protocolcontractinterfaceAddress common.Address,
	tokencontractinterfaceAddress common.Address,
	clientConn *ethclient.Client,
	appId string,
) (ChainWrapper, error) {
	var err error
	c := &chainWrapperImp{}

	c.protocol, err = contractinterface.NewScryProtocol(protocolcontractinterfaceAddress, clientConn)
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("failed to initialize protocol contractinterface interface wrapper.", err))
		return nil, err
	}

	c.token, err = contractinterface.NewScryToken(tokencontractinterfaceAddress, clientConn)
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("failed to initialize token contractinterface interface wrapper.", err))
		return nil, err
	}

	c.conn = clientConn
	c.appId = appId

	//load components
	dot.GetDefaultLine().ToInjecter().Inject(&c)

	return c, err
}

func (c *chainWrapperImp) Conn() *ethclient.Client {
	return c.conn
}

func (c *chainWrapperImp) Publish(txParams *tx.TxParams, price *big.Int, metaDataID []byte,
	proofDataIDs []string, proofNum int32, detailsID string, supportVerify bool) (string, error) {
	logger := dot.Logger()

	defer func() {
		if er := recover(); er != nil {
			logger.Errorln("", zap.Any("failed to publish data, error:", er))
		}
	}()

	//generate publishId
	publishId := util.GenerateUUID()

	pdIDs := make([][32]byte, proofNum)
	var err error
	for i := int32(0); i < proofNum; i++ {
		pdIDs[i], err = ipfsHashToBytes32(proofDataIDs[i])
		if err != nil {
			logger.Errorln("failed to convert ipfs hash to bytes32")
			return "", err
		}
	}

	encMetaId, err := c.Account.Encrypt(metaDataID, txParams.From.String())
	if err != nil {
		logger.Errorln("", zap.NamedError("failed to encrypt meta data hash, error: ", err))
		return "", err
	}

	t, err := c.protocol.Publish(c.Tx.BuildTransactOpts(txParams), c.appId, publishId, price,
		encMetaId, pdIDs, detailsID, supportVerify)
	if err != nil {
		logger.Errorln("", zap.NamedError("failed to publish data information, error: ", err))
		return "", err
	}

	logger.Debugln("publish Tx: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))

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

// Bytes32ToIpfsHash bytes32 to ipfs hash
func Bytes32ToIpfsHash(value [32]byte) (string, error) {
	byteArray := [34]byte{18, 32}
	copy(byteArray[2:], value[:])
	if len(byteArray) != 34 {
		return "", errors.New("invalid bytes32 value")
	}

	hash := base58.Encode(byteArray[:])
	return hash, nil
}

func (c *chainWrapperImp) AdvancePurchase(txParams *tx.TxParams, publishId string, startVerify bool) error {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to AdvancePurchase, error:", er))
		}
	}()

	t, err := c.protocol.AdvancePurchase(c.Tx.BuildTransactOpts(txParams), c.appId, publishId, startVerify)
	if err == nil {
		dot.Logger().Debugln("AdvancePurchase: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) ConfirmPurchase(txParams *tx.TxParams, txId *big.Int) error {
	t, err := c.protocol.ConfirmPurchase(c.Tx.BuildTransactOpts(txParams), c.appId, txId)
	if err == nil {
		dot.Logger().Debugln("ConfirmPurchase: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) CancelPurchase(txParams *tx.TxParams, txId *big.Int) error {
	t, err := c.protocol.CancelPurchase(c.Tx.BuildTransactOpts(txParams), c.appId, txId)
	if err == nil {
		dot.Logger().Debugln("CancelPurchase tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) ReEncrypt(
	txParams *tx.TxParams,
	txId *big.Int,
	encodedData []byte,
) (err error) {
	buyer, err := c.protocol.GetBuyer(c.Tx.BuildCallOpts(txParams), txId)
	if err != nil {
		dot.Logger().Errorln("chainWrapperImp::ReEncryptMetaDataId", zap.Error(err))
		return err
	}

	if buyer == common.HexToAddress("0x0") {
		e := "invalid buyer address"
		dot.Logger().Errorln("chainWrapperImp::ReEncryptMetaDataId", zap.String("error:", e))
		return errors.New(e)
	}

	edb, err := c.Account.ReEncrypt(encodedData, txParams.From.String(), buyer.String(), txParams.Password)
	if err != nil {
		dot.Logger().Errorln("chainWrapperImp::ReEncryptMetaDataId", zap.Error(err))
		return err
	}

	//re-encrypt with arbitrators public key
	arbitrators, err := c.protocol.GetArbitrators(c.Tx.BuildCallOpts(txParams), txId)
	if err != nil {
		dot.Logger().Errorln("chainWrapperImp::ReEncryptMetaDataId", zap.Error(err))
		return err
	}
	var edaList []byte
	for _, ab := range arbitrators {
		if ab == common.HexToAddress("0x0") {
			e := "invalid arbitrator address"
			dot.Logger().Errorln("chainWrapperImp::ReEncryptMetaDataId", zap.String("error:", e))
			return errors.New(e)
		}

		eda, err := c.Account.ReEncrypt(encodedData, txParams.From.String(), ab.String(), txParams.Password)
		if err != nil {
			dot.Logger().Errorln("chainWrapperImp::ReEncryptMetaDataId", zap.Error(err))
			return err
		}

		edaList = append(edaList, eda...)
	}

	//submit
	t, err := c.protocol.ReEncrypt(c.Tx.BuildTransactOpts(txParams), c.appId, txId, edb, edaList)
	if err == nil {
		dot.Logger().Debugln("ReEncrypt: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) Arbitrate(txParams *tx.TxParams, txId *big.Int, judge bool) error {
	t, err := c.protocol.Arbitrate(c.Tx.BuildTransactOpts(txParams), c.appId, txId, judge)
	if err == nil {
		dot.Logger().Debugln("Arbitrate: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) GetBuyer(txParams *tx.TxParams, txId *big.Int) (string, error) {
	buyer, err := c.protocol.GetBuyer(c.Tx.BuildCallOpts(txParams), txId)
	if err == nil {
		dot.Logger().Debugln("Get buyer, buyer: " + buyer.String())
	}

	return buyer.String(), err
}

func (c *chainWrapperImp) GetArbitrators(txParams *tx.TxParams, txId *big.Int) ([]string, error) {
	arbitratorsAddrs, err := c.protocol.GetArbitrators(c.Tx.BuildCallOpts(txParams), txId)

	var arbitrators = make([]string, len(arbitratorsAddrs))
	for i := 0; i < len(arbitratorsAddrs); i++ {
		arbitrators[i] = arbitratorsAddrs[i].String()
	}

	if err == nil {
		dot.Logger().Debugln("Get arbitrator:", zap.Strings("arbitrators", arbitrators))
	}

	return arbitrators, err
}

func (c *chainWrapperImp) ConfirmData(txParams *tx.TxParams, txId *big.Int, truth bool) error {
	t, err := c.protocol.ConfirmData(c.Tx.BuildTransactOpts(txParams), c.appId, txId, truth)
	if err == nil {
		dot.Logger().Debugln("ConfirmData: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) ApproveTransfer(txParams *tx.TxParams, spender common.Address, value *big.Int) error {
	t, err := c.token.Approve(c.Tx.BuildTransactOpts(txParams), spender, value)
	if err == nil {
		dot.Logger().Debugln("ApproveTransfer: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) Vote(txParams *tx.TxParams, txId *big.Int, judge bool, comments string) error {
	t, err := c.protocol.Vote(c.Tx.BuildTransactOpts(txParams), c.appId, txId, judge, comments)
	if err == nil {
		dot.Logger().Debugln("Vote: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))

	}

	return err
}

func (c *chainWrapperImp) RegisterAsVerifier(txParams *tx.TxParams) error {
	t, err := c.protocol.RegisterAsVerifier(c.Tx.BuildTransactOpts(txParams), c.appId)
	if err == nil {
		dot.Logger().Debugln("RegisterAsVerifier: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) GradeToVerifier(txParams *tx.TxParams, txId *big.Int, index uint8, credit uint8) error {
	t, err := c.protocol.GradeToVerifier(c.Tx.BuildTransactOpts(txParams), c.appId, txId, index, credit)
	if err == nil {
		dot.Logger().Debugln("GradeToVerifier: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) TransferTokens(txParams *tx.TxParams, to common.Address, value *big.Int) error {
	t, err := c.token.Transfer(c.Tx.BuildTransactOpts(txParams), to, value)
	if err == nil {
		dot.Logger().Debugln("TransferTokens: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return err
}

func (c *chainWrapperImp) GetTokenBalance(txParams *tx.TxParams, owner common.Address) (*big.Int, error) {
	return c.token.BalanceOf(c.Tx.BuildCallOpts(txParams), owner)
}

func (c *chainWrapperImp) ModifyContractParam(txParams *tx.TxParams, paramName, paramValue string) (err error) {
	var t *types.Transaction
	switch paramName {
	case "VerifierNum":
		value, _ := strconv.Atoi(paramValue)
		t, err = c.protocol.ModifyVerifierNum(c.Tx.BuildTransactOpts(txParams), uint8(value))
	default:
		err = errors.New("Unknown contract param or param is not allowed to modify. ")
		return
	}
	if err == nil {
		dot.Logger().Debugln("ModifyContractParam: tx hash:"+t.Hash().String(), zap.Binary(" tx data:", t.Data()))
	}

	return
}
