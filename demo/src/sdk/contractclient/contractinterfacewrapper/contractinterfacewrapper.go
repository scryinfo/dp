package contractinterfacewrapper

import (
    "../../interface/contractinterface"
    op "../../core/chainoperations"
    "../../util/storage/ipfsaccess"
    "../../util/uuid"
    "../../util/accounts"
    "errors"
    "fmt"
    "github.com/btcsuite/btcutil/base58"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
)

var (
	scryProtocol *contractinterface.ScryProtocol = nil
	scryToken *contractinterface.ScryToken = nil
)


func Initialize(protocolContractAddress common.Address, tokenContractAddress common.Address, conn *ethclient.Client) error {
	var err error = nil

	scryProtocol, err = contractinterface.NewScryProtocol(protocolContractAddress, conn)
	if err != nil {
		fmt.Println("failed to initialize protocol contract interface wrapper.", err)
		return err
	}

	scryToken, err = contractinterface.NewScryToken(tokenContractAddress, conn)
	if err != nil {
		fmt.Println("failed to initialize token contract interface wrapper.", err)
		return err
	}

	return nil
}

func Publish(txParams *op.TransactParams, price *big.Int, metaData []byte, proofDatas [][]byte,
				proofNum int, descriptionData []byte, supportVerify bool) (string, error)  {
	defer func(){
		if err:=recover(); err!=nil {
			fmt.Println("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := uuid.GenerateUUID()

	//submit meta data
	cidMd, err := ipfsaccess.GetIAInstance().SaveToIPFS(metaData)
	if err != nil {
		fmt.Println("failed to save meta data to ipfs, error: ", err)
		return "", err
	}

	//submit proof data
	cidPds := make([][32]byte, proofNum)
	for	i := 0; i < proofNum; i++ {
		cidPd, err := ipfsaccess.GetIAInstance().SaveToIPFS(proofDatas[i])

		if err != nil {
			fmt.Println("failed to save proof data to ipfs, error: ", err)
			return "", err
		}

		cidPds[i], err = ipfsHashToBytes32(cidPd)
		if err != nil {
			fmt.Println("failed to convert ipfs hash to bytes32")
			return "", err
		}
	}

	//submit description data
	cidDd, err := ipfsaccess.GetIAInstance().SaveToIPFS(descriptionData)
	if err != nil {
		fmt.Println("failed to save description data to ipfs, error: ", err)
		return "", err
	}

	b := []byte(cidMd)
	encMetaId, err := accounts.GetAMInstance().Encrypt(b, txParams.From.String(), txParams.Password)
	if err != nil {
		fmt.Println("failed to encrypt meta data hash, error: ", err)
		return "", err
	}

	//upload meta_data_id_enc_seller and other cids to contracts
	tx, err := scryProtocol.PublishDataInfo(buildTxOpts(txParams), uuid.GenerateUUID(), publishId, price, encMetaId, cidPds, cidDd, supportVerify)
	if err != nil {
		fmt.Println("failed to publish data information, error: ", err)
		return "", err
	}

	fmt.Println("publish transaction:" + string(tx.Data()) + " hash:" + tx.Hash().String())

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
			fmt.Println("failed to prepare to buy , error:", err)
		}
	}()

	tx, err := scryProtocol.CreateTransaction(buildTxOpts(txParams), uuid.GenerateUUID(), publishId)
	if err == nil {
		fmt.Println("CreateTransaction:" + string(tx.Data()))
	}

	return err
}

func BuyData(txParams *op.TransactParams, txId *big.Int) (error) {
	tx, err := scryProtocol.BuyData(buildTxOpts(txParams), uuid.GenerateUUID(), txId)
	if err == nil {
		fmt.Println("BuyData:", tx.Data(), " tx hash:", tx.Hash().String())
	}

	return err
}

func SubmitMetaDataIdEncWithBuyer(txParams *op.TransactParams, txId *big.Int, encyptedMetaDataId []byte) (error) {
	tx, err := scryProtocol.SubmitMetaDataIdEncWithBuyer(buildTxOpts(txParams), uuid.GenerateUUID(), txId, encyptedMetaDataId)
	if err == nil {
		fmt.Println("SubmitMetaDataIdEncWithBuyer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ConfirmDataTruth(txParams *op.TransactParams, txId *big.Int, truth bool) (error) {
	tx, err := scryProtocol.ConfirmDataTruth(buildTxOpts(txParams), uuid.GenerateUUID(), txId, truth)
	if err == nil {
		fmt.Println("ConfirmDataTruth:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ApproveTransfer(txParams *op.TransactParams, spender common.Address, value *big.Int) (error) {
	tx, err := scryToken.Approve(buildTxOpts(txParams), spender, value)
	if err == nil {
		fmt.Println("ApproveTransfer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func Vote(txParams *op.TransactParams, txId *big.Int, judge bool, comments string) (error) {
    tx, err := scryProtocol.Vote(buildTxOpts(txParams), uuid.GenerateUUID(), txId, judge, comments)
    if err == nil {
        fmt.Println("Vote:", string(tx.Data()), " tx hash:", tx.Hash().String())
    }

    return err
}

func RegisterAsVerifier(txParams *op.TransactParams) (error) {
    tx, err := scryProtocol.RegisterAsVerifier(buildTxOpts(txParams), uuid.GenerateUUID())
    if err == nil {
        fmt.Println("RegisterAsVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
    }

    return err
}

func CreditsToVerifier(txParams *op.TransactParams, txId *big.Int, to common.Address, credit uint8) (error) {
    tx, err := scryProtocol.CreditsToVerifier(buildTxOpts(txParams), uuid.GenerateUUID(), txId, to, credit)
    if err == nil {
        fmt.Println("RegisterAsVerifier:", string(tx.Data()), " tx hash:", tx.Hash().String())
    }

    return err
}

func buildTxOpts(txParams *op.TransactParams) (*bind.TransactOpts) {
    return op.BuildTransactOpts(txParams)
}

func TransferTokens(txParams *op.TransactParams, to common.Address, value *big.Int) (error)  {
    tx, err := scryToken.Transfer(buildTxOpts(txParams), to, value)
    if err == nil {
        fmt.Println("tx:", tx)
        return err
    }

    return err
}

