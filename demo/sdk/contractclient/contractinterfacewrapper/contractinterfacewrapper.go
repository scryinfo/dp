package contractinterfacewrapper

import (
	"../../contractinterface"
	"../../util/storage/ipfsaccess"
	"../../util/usermanager"
	"../../util/uuid"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	txOpts *bind.TransactOpts = nil
	scryProtocol *contractinterface.ScryProtocol = nil
)

func Initialize(sp *contractinterface.ScryProtocol, opts *bind.TransactOpts)  {
	scryProtocol = sp
	txOpts = opts
}

func Publish(metaData *[]byte, proofData *[]byte, descriptionData *[]byte, supportVerify bool) (string, error)  {
	defer func(){
		if err:=recover(); err!=nil {
			fmt.Println("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := uuid.GenerateUUID()

	//submit meta data
	cidMd, err := ipfsaccess.GetInstance().SaveToIPFS(*metaData)
	if err != nil {
		fmt.Println("failed to save meta data to ipfs, error: ", err)
		return "", err
	}

	//submit proof data
	cidPd, err := ipfsaccess.GetInstance().SaveToIPFS(*proofData)
	if err != nil {
		fmt.Println("failed to save proof data to ipfs, error: ", err)
		return "", err
	}

	//submit description data
	cidDd, err := ipfsaccess.GetInstance().SaveToIPFS(*descriptionData)
	if err != nil {
		fmt.Println("failed to save description data to ipfs, error: ", err)
		return "", err
	}

	//encrypt meta_data_id
	curUser, err := usermanager.GetCurrentUser()
	if err != nil {
		fmt.Println("failed to get current user, error: ", err)
		return "", err
	}

	secOper, err := curUser.GetSecurityOpertion()
	if err != nil {
		fmt.Println("failed to get security operation interface of current user, error: ", err)
		return "", err
	}

	b := []byte(cidMd)
	encMetaId, err := secOper.Encrypt(&b)
	if err != nil {
		fmt.Println("failed to encrypt meta data hash, error: ", err)
		return "", err
	}

	//upload meta_data_id_enc_seller and other cids to contracts
	tx, err := scryProtocol.PublishDataInfo(txOpts, publishId, *encMetaId, cidPd, cidDd, supportVerify)
	if err != nil {
		fmt.Println("failed to publish data information, error: ", err)
		return "", err
	}

	fmt.Println("transaction:" + string(tx.Data()))

	return publishId, nil
}

/*func PrepareToBuy(buyerAddr string, publishId string) (txId common.Hash, err error) {
}

func BuyData(txId common.Hash) (err error) {
}

func SubmitMetaDataIdEncWithBuyer(txId common.Hash, encyptedMetaDataId []byte) (err error) {
}

func ConfirmDataTruth(txId common.Hash, truth bool) (err error) {

}*/