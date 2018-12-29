package main

import (
	"encoding/json"
	"fmt"
	"github.com/qjpcpu/ethereum/events"
	"io/ioutil"
	"strings"
	"time"
)

func main()  {


	//get description data



	time.Sleep(100000000000000)
}

func SellerPublishData()  {
	keyJson := `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
	keyPassword := "12345"
	publicAddress := DecodeKeystoreAddress([]byte(keyJson))

	//initialize sdk
	usermanager.Register(publicAddress, false, nil)
	usermanager.SetCurrentUser(publicAddress)

	client := contractclient.NewContractClient(publicAddress, keyJson, keyPassword)
	client.Initialize("http://127.0.0.1:7545/",
		"0xf83872acdaca3c9a80199eeeebb0c6a323b75187", getAbiText(), "/ip4/127.0.0.1/tcp/5001")
	chainevents.SubscribeExternal("Published", onPublished)


	//publish data
	metaData := []byte{'1','2','3','3'}
	proofData := []byte{'4','5','6','3'}
	despData := []byte{'7','8','9','3'}
	contractinterfacewrapper.Publish(&metaData, &proofData, &despData, false)
}

func BuyerPrepareToBuy()  {

}

func BuyerDecideToBuy() {

}

func onPublished(event events.Event) bool {
	fmt.Println("onpublish: ", event)
	return true
}

func getAbiText() string {
	abi, err := ioutil.ReadFile("./ScryProtocol.abi")
	if err != nil {
		fmt.Println("failed to read abi text", err)
		return ""
	}

	return string(abi)
}

func DecodeKeystoreAddress(keyJsonStr []byte) string {
	addr := struct {
		Address string `json:"address"`
	}{}
	if err := json.Unmarshal(keyJsonStr, &addr); err != nil {
		panic(fmt.Sprintf("parse address fail:%v", err))
	}
	if !strings.HasPrefix(addr.Address, "0x") {
		addr.Address = `0x` + addr.Address
	}
	return addr.Address
}