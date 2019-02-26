package chaininterfacewrapper

import (
    op "sdk/core/chainoperations"
    "sdk/interface/contractinterface"
    "sdk/util/accounts"
    "sdk/util/storage/ipfsaccess"
    "sdk/util/uuid"
    "errors"
    "github.com/btcsuite/btcutil/base58"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    rlog "github.com/sirupsen/logrus"
    "math/big"
)

var (
    conn *ethclient.Client = nil
	scryProtocol *contractinterface.ScryProtocol = nil
	scryToken *contractinterface.ScryToken = nil
)


func Initialize(protocolContractAddress common.Address,
    tokenContractAddress common.Address,
    clientConn *ethclient.Client) error {
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

func Publish(txParams *op.TransactParams, price *big.Int, metaData []byte, proofDatas [][]byte,
				proofNum int, descriptionData []byte, supportVerify bool) (string, error)  {
	defer func(){
		if err := recover(); err != nil {
			rlog.Error("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := uuid.GenerateUUID()

	//submit meta data
	cidMd, err := ipfsaccess.GetIAInstance().SaveToIPFS(metaData)
	if err != nil {
		rlog.Error("failed to save meta data to ipfs, error: ", err)
		return "", err
	}

	//submit proof data
	cidPds := make([][32]byte, proofNum)
	for	i := 0; i < proofNum; i++ {
		cidPd, err := ipfsaccess.GetIAInstance().SaveToIPFS(proofDatas[i])

		if err != nil {
			rlog.Error("failed to save proof data to ipfs, error: ", err)
			return "", err
		}

		cidPds[i], err = ipfsHashToBytes32(cidPd)
		if err != nil {
			rlog.Error("failed to convert ipfs hash to bytes32")
			return "", err
		}
	}

	//submit description data
	cidDd, err := ipfsaccess.GetIAInstance().SaveToIPFS(descriptionData)
	if err != nil {
		rlog.Error("failed to save description data to ipfs, error: ", err)
		return "", err
	}

	b := []byte(cidMd)
	encMetaId, err := accounts.GetAMInstance().Encrypt(b, txParams.From.String(), txParams.Password)
	if err != nil {
		rlog.Error("failed to encrypt meta data hash, error: ", err)
		return "", err
	}

	//upload meta_data_id_enc_seller and other cids to contracts
	tx, err := scryProtocol.PublishDataInfo(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), publishId, price, encMetaId, cidPds, cidDd, supportVerify)
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


func PrepareToBuy(txParams *op.TransactParams, publishId string) (error) {
	defer func(){
		if err:=recover(); err!=nil {
			rlog.Error("failed to prepare to buy , error:", err)
		}
	}()

	tx, err := scryProtocol.CreateTransaction(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), publishId)
	if err == nil {
		rlog.Debug("CreateTransaction:" + string(tx.Data()))
	}

	return err
}

func BuyData(txParams *op.TransactParams, txId *big.Int) (error) {
	tx, err := scryProtocol.BuyData(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), txId)
	if err == nil {
		rlog.Debug("BuyData:", tx.Data(), " tx hash:", tx.Hash().String())
	}

	return err
}

func SubmitMetaDataIdEncWithBuyer(txParams *op.TransactParams, txId *big.Int, encyptedMetaDataId []byte) (error) {
	tx, err := scryProtocol.SubmitMetaDataIdEncWithBuyer(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), txId, encyptedMetaDataId)
	if err == nil {
		rlog.Debug("SubmitMetaDataIdEncWithBuyer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ConfirmDataTruth(txParams *op.TransactParams, txId *big.Int, truth bool) (error) {
	tx, err := scryProtocol.ConfirmDataTruth(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), txId, truth)
	if err == nil {
		rlog.Debug("ConfirmDataTruth:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ApproveTransfer(txParams *op.TransactParams, spender common.Address, value *big.Int) (error) {
	tx, err := scryToken.Approve(op.BuildTransactOpts(txParams), spender, value)
	if err == nil {
		rlog.Debug("ApproveTransfer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func Vote(txParams *op.TransactParams, txId *big.Int, judge bool, comments string) (error) {
    tx, err := scryProtocol.Vote(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), txId, judge, comments)
    if err == nil {
        rlog.Debug("Vote:", string(tx.Data()), " tx hash:", tx.Hash().String())
    }

    return err
}

func RegisterAsVerifier(txParams *op.TransactParams) (error) {
    tx, err := scryProtocol.RegisterAsVerifier(op.BuildTransactOpts(txParams), uuid.GenerateUUID())
    if err == nil {
        rlog.Debug("RegisterAsVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
    }

    return err
}

func CreditsToVerifier(txParams *op.TransactParams, txId *big.Int, to common.Address, credit uint8) (error) {
    tx, err := scryProtocol.CreditsToVerifier(op.BuildTransactOpts(txParams), uuid.GenerateUUID(), txId, to, credit)
    if err == nil {
        rlog.Debug("RegisterAsVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
    }

    return err
}

func TransferTokens(txParams *op.TransactParams, to common.Address, value *big.Int) (error)  {
    tx, err := scryToken.Transfer(op.BuildTransactOpts(txParams), to, value)
    if err == nil {
        rlog.Debug("tx:", tx)
        return err
    }

    return err
}

func GetTokenBalance(txParams *op.TransactParams, owner common.Address) (*big.Int, error)  {
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