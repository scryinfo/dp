package publishment

import (
	"../../storage/ipfsaccess"
	"../../usermanager"
	"../../util"
	"../contractinterface"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func Initialize(nodeAddr string) bool {
	return ipfsaccess.Initialize(nodeAddr)
}

func Publish(metaData *[]byte, proofData *[]byte, descriptionData *[]byte, supportVerify bool) (string, error)  {
	defer func(){
		if err:=recover(); err!=nil {
			fmt.Println("failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := util.GenerateUUID()

	//submit meta data
	cidMd, err := ipfsaccess.SaveToIPFS(*metaData)
	if err != nil {
		fmt.Println("failed to save meta data to ipfs, error: ", err)
		return "", err
	}

	//submit proof data
	cidPd, err := ipfsaccess.SaveToIPFS(*proofData)
	if err != nil {
		fmt.Println("failed to save proof data to ipfs, error: ", err)
		return "", err
	}

	//submit description data
	cidDd, err := ipfsaccess.SaveToIPFS(*descriptionData)
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

	//upload meta_data_id_enc_seller and other cids to contract
	user, err := usermanager.GetCurrentUser()
	if err != nil {
		fmt.Println("failed to get current user: ", err)
		return "", err
	}

	err = contractinterface.PublishDataInfo(publishId, *encMetaId,
		cidPd, cidDd, common.HexToAddress(user.GetPublicKey()), supportVerify)
	if err != nil {
		fmt.Errorf("failed to publish data information, error: ", err)
		return "", err
	}

	return publishId, nil
}