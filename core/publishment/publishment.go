package publishment

import (
	"../../util"
	"../../usermanager"
	"../../storage/ipfsaccess"
	"fmt"
)

func Initialize(nodeAddr string) bool {
	return ipfsaccess.Initialize(nodeAddr)
}

func Publish(metaData *[]byte, proofData *[]byte, descriptionData *[]byte) (string, error)  {
	defer func(){
		if err:=recover(); err!=nil {
			fmt.Println("Failed to publish data, error:", err)
		}
	}()

	//generate publishId
	publishId := util.GenerateUUID()

	//submit meta data
	cidMd, err := ipfsaccess.SaveToIPFS(*metaData)
	if err != nil {
		fmt.Println("Failed to save meta data to ipfs, error: ", err)
		return "", err
	}

	//submit proof data
	cidPd, err := ipfsaccess.SaveToIPFS(*proofData)
	if err != nil {
		fmt.Println("Failed to save proof data to ipfs, error: ", err)
		return "", err
	}

	//submit description data
	cidDd, err := ipfsaccess.SaveToIPFS(*descriptionData)
	if err != nil {
		fmt.Println("Failed to save description data to ipfs, error: ", err)
		return "", err
	}

	//encrypt meta_data_id
	curUser, err := usermanager.GetCurrentUser()
	if err != nil {
		fmt.Println("Failed to get current user, error: ", err)
		return "", err
	}

	secOper, err := curUser.GetSecurityOpertion()
	if err != nil {
		fmt.Println("Failed to get security operation interface of current user, error: ", err)
		return "", err
	}

	b := []byte(cidMd)
	encMetaId, err := secOper.Encrypt(&b)
	if err != nil {
		fmt.Println("Failed to encrypt meta data hash, error: ", err)
		return "", err
	}

	//upload meta_data_id_enc_seller and other cids to contract

	return publishId, nil
}