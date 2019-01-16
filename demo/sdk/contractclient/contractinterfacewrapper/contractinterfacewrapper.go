package contractinterfacewrapper

import (
	"../../contractinterface"
	"../../util/storage/ipfsaccess"
	"../../util/security"
	"../../util/uuid"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/btcsuite/btcutil/base58"
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

func Publish(txOpts *bind.TransactOpts, price *big.Int, metaData []byte, proofDatas [][]byte,
				proofNum int, descriptionData []byte, supportVerify bool) (string, error)  {
	defer func(){
		if err:=recover(); err!=nil {
			fmt.Println("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := uuid.GenerateUUID()

	//submit meta data
	cidMd, err := ipfsaccess.GetInstance().SaveToIPFS(metaData)
	if err != nil {
		fmt.Println("failed to save meta data to ipfs, error: ", err)
		return "", err
	}

	//submit proof data
	cidPds := make([][32]byte, proofNum)
	for	i := 0; i < proofNum; i++ {
		cidPd, err := ipfsaccess.GetInstance().SaveToIPFS(proofDatas[i])

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
	cidDd, err := ipfsaccess.GetInstance().SaveToIPFS(descriptionData)
	if err != nil {
		fmt.Println("failed to save description data to ipfs, error: ", err)
		return "", err
	}

	b := []byte(cidMd)
	secOper := security.SecurityExecutor{}
	encMetaId, err := secOper.Encrypt(&b)
	if err != nil {
		fmt.Println("failed to encrypt meta data hash, error: ", err)
		return "", err
	}

	//upload meta_data_id_enc_seller and other cids to contracts
	tx, err := scryProtocol.PublishDataInfo(txOpts, uuid.GenerateUUID(), publishId, price, *encMetaId, cidPds, cidDd, supportVerify)
	if err != nil {
		fmt.Println("failed to publish data information, error: ", err)
		return "", err
	}

	fmt.Println("publish transaction:" + string(tx.Data()))

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


func PrepareToBuy(txOpts *bind.TransactOpts, publishId string) (error) {
	defer func(){
		if err:=recover(); err!=nil {
			fmt.Println("failed to prepare to buy , error:", err)
		}
	}()

	tx, err := scryProtocol.CreateTransaction(txOpts, uuid.GenerateUUID(), publishId)
	if err == nil {
		fmt.Println("prepareToBuy transaction:" + string(tx.Data()))
	}

	return err
}

func BuyData(txOpts *bind.TransactOpts, txId *big.Int) (error) {
	tx, err := scryProtocol.BuyData(txOpts, uuid.GenerateUUID(), txId)
	if err == nil {
		fmt.Println("BuyData:", tx.Data(), " tx hash:", tx.Hash().String())
	}

	return err
}

func SubmitMetaDataIdEncWithBuyer(txOpts *bind.TransactOpts, txId *big.Int, encyptedMetaDataId []byte) (error) {
	tx, err := scryProtocol.SubmitMetaDataIdEncWithBuyer(txOpts, uuid.GenerateUUID(), txId, encyptedMetaDataId)
	if err == nil {
		fmt.Println("SubmitMetaDataIdEncWithBuyer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ConfirmDataTruth(txOpts *bind.TransactOpts, txId *big.Int, truth bool) (error) {
	tx, err := scryProtocol.ConfirmDataTruth(txOpts, uuid.GenerateUUID(), txId, truth)
	if err == nil {
		fmt.Println("ConfirmDataTruth:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}

func ApproveTransfer(txOpts *bind.TransactOpts, spender common.Address, value *big.Int) (error) {
	tx, err := scryToken.Approve(txOpts, spender, value)
	if err == nil {
		fmt.Println("ApproveTransfer:", string(tx.Data()), " tx hash:", tx.Hash().String())
	}

	return err
}