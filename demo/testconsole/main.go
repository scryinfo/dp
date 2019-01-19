package main

import (
    "../sdk/contractclient"
    "../sdk/contractclient/contractinterfacewrapper"
    "../sdk/core"
    "../sdk/core/chainevents"
    "../sdk/core/chainoperations"
    "../sdk/core/ethereum/events"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "io/ioutil"
    "math/big"
    "time"
)

var (
	publishId = ""
	txId *big.Int = big.NewInt(0)
	metaDataIdEncWithSeller []byte
	metaDataIdEncWithBuyer []byte
	protocolContractAddr = "0xf1206c57769a4bf8738b07528f1fa7f7239d903b"
	tokenContractAddr = "0x7768ed4e32f9536e40ae1dd22c36f06c25ae67df"
	conn *ethclient.Client = nil
    keyJson = `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
    keyPassword = "12345"
)

func main()  {
    var err error = nil
	conn, err = core.StartEngine("http://127.0.0.1:7545/", getContracts(), "/ip4/127.0.0.1/tcp/5001")
	if err != nil {
		fmt.Println("failed to start engine. error:", err)
		return
	}

	err = contractinterfacewrapper.Initialize(common.HexToAddress(protocolContractAddr),
		common.HexToAddress(tokenContractAddr), conn)
	if err != nil {
		fmt.Println("failed to initialize contract interface, error:", err)
		return
	}

	SellerPublishData(conn)

	time.Sleep(5000000000)

	ApproveTransfer(conn)

	time.Sleep(5000000000)

	Buy(conn, txId)

	time.Sleep(5000000000)

	SubmitMetaDataIdEncWithBuyer(conn, txId)

	time.Sleep(5000000000)

	ConfirmDataTruth(conn, txId)

	time.Sleep(100000000000000)
}

func SellerPublishData(conn *ethclient.Client)  {
    publicAddress := getClientInfo()
	client, err := contractclient.NewContractClient(publicAddress)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return
	}

	client.SubscribeEvent("DataPublish", onPublish)

	//publish data
	metaData := []byte{'1','2','3','3'}
	proofData := [][]byte{{'4','5','6','3'}, {'2','2', '1'}}
	despData := []byte{'7','8','9','3'}

	txParam := chainoperations.TransactParams{ common.HexToAddress(publicAddress), keyPassword}
	contractinterfacewrapper.Publish(&txParam, big.NewInt(1000), metaData, proofData, 2, despData, false)
}

func onPublish(event events.Event) bool {
    fmt.Println("onpublish: ", event)

    publishId = event.Data.Get("publishId").(string)
    return true
}

func ApproveTransfer(conn *ethclient.Client)  {
    publicAddress := getClientInfo()

    client, err := contractclient.NewContractClient(publicAddress)
    if err != nil {
        fmt.Println("failed to create contract client. error:", err)
        return
    }

    client.SubscribeEvent("Approval", onApproval)

    txParam := chainoperations.TransactParams{ common.HexToAddress(publicAddress), keyPassword}
    contractinterfacewrapper.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(1000))
}


func PrepareToBuy(conn *ethclient.Client, publishId string)  {
    publicAddress := getClientInfo()

	client, err := contractclient.NewContractClient(publicAddress)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return
	}

	client.SubscribeEvent("TransactionCreate", onTransactionCreate)

    txParam := chainoperations.TransactParams{ common.HexToAddress(publicAddress), keyPassword}
	err = contractinterfacewrapper.PrepareToBuy(&txParam, publishId)
	if err != nil {
		fmt.Println("failed to prepareToBuy, error:", err)
	}
}

func Buy(conn *ethclient.Client, txId *big.Int) {
    publicAddress := getClientInfo()

	//initialize sdk
	client, err := contractclient.NewContractClient(publicAddress)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return
	}

	client.SubscribeEvent("Purchase", onPurchase)

    txParam := chainoperations.TransactParams{ common.HexToAddress(publicAddress), keyPassword}
	err = contractinterfacewrapper.BuyData(&txParam, txId)
	if err != nil {
		fmt.Println("failed to buyData, error:", err)
	}
}

func SubmitMetaDataIdEncWithBuyer(conn *ethclient.Client, txId *big.Int)  {
    publicAddress := getClientInfo()

	//initialize sdk
	client, err := contractclient.NewContractClient(publicAddress)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return
	}

	client.SubscribeEvent("ReadyForDownload", onReadyForDownload)

    txParam := chainoperations.TransactParams{ common.HexToAddress(publicAddress), keyPassword}
	err = contractinterfacewrapper.SubmitMetaDataIdEncWithBuyer(&txParam, txId, metaDataIdEncWithBuyer)
	if err != nil {
		fmt.Println("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
	}
}

func ConfirmDataTruth(conn *ethclient.Client, txId *big.Int)  {
    publicAddress := getClientInfo()

	//initialize sdk
	client, err := contractclient.NewContractClient(publicAddress)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return
	}

	client.SubscribeEvent("Close", onClose)

    txParam := chainoperations.TransactParams{ common.HexToAddress(publicAddress), keyPassword}
	err = contractinterfacewrapper.ConfirmDataTruth(&txParam, txId, true)
	if err != nil {
		fmt.Println("failed to ConfirmDataTruth, error:", err)
	}
}


func onPurchase(event events.Event) bool {
	fmt.Println("onPurchase:", event)
	metaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
	metaDataIdEncWithBuyer = make([]byte, len(metaDataIdEncWithSeller))
	copy(metaDataIdEncWithBuyer, metaDataIdEncWithSeller)

	return true
}

func onReadyForDownload(event events.Event) bool {
	fmt.Println("onReadyForDownload:", event)
	metaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
	return true
}

func onClose(event events.Event) bool {
	fmt.Println("onClose:", event)
	return true
}


func onApproval(event events.Event) bool {
    fmt.Println("onApproveTransfer:", event)

    PrepareToBuy(conn, publishId)
    return true
}

func getContracts() ([]chainevents.ContractInfo) {
    protocolEvents := []string{"DataPublish", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose"}
    tokenEvents := []string{"Approval"}

    contracts := []chainevents.ContractInfo {
        {protocolContractAddr, getAbiText("./ScryProtocol.abi"), protocolEvents},
        {tokenContractAddr, getAbiText("./ScryToken.abi"),tokenEvents},
    }

    return contracts
}

func getAbiText(fileName string) string {
    abi, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Println("failed to read abi text", err)
        return ""
    }

    return string(abi)
}


func getClientInfo() (string) {
    publicAddress := chainoperations.DecodeKeystoreAddress([]byte(keyJson))
    return publicAddress
}

func onTransactionCreate(event events.Event) bool {
    fmt.Println("onTransactionCreated:", event)
    txId = event.Data.Get("transactionId").(*big.Int)
    return true
}