package publishment

import (
	"../../util"
	"../../storage/ipfsaccess"
)

func Initialize(nodeAddr string) bool {
	return ipfsaccess.Initialize(nodeAddr)
}

func Publish(metaData *[]byte, proofData *[]byte, descriptionData *[]byte) (string, error)  {
	defer func(){
		if err:=recover(); err!=nil {

		}
	}()

	//generate publishId
	publishId := util.GenerateUUID()

	//submit data

	//encrypt meta_data_id

	//submit meta_data_id_enc_seller
}