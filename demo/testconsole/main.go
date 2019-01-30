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
	protocolContractAddr = "0x4ae12e473d8eaf98d0410555508cc8ba8c3cf7b6"
	tokenContractAddr = "0x77d631c8c87bd3dd0d5631add4d76364662b2159"
	conn *ethclient.Client = nil
    deployerKeyJson = `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
    sellerKeyJson = `{"version":3,"id":"7f01defb-3543-4459-bcc3-3b86197a4e17","address":"2d13d4faba031e66a36a6b307fce2087db55c43d","crypto":{"ciphertext":"ce55ddca8d430b2aef68a5fafda00f37a6df0aad45a6fa50c67920722c5a06b7","cipherparams":{"iv":"68dbb3ec9294ed20129807e46da74d72"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"b88dfead45a58f8d73645a7c458f4130cd51d4511f8b5861a756c840dc50bd66","n":262144,"r":8,"p":1},"mac":"6d90b054f4a1c8e3278fc47119c684fbb63257d3726df26b6f69fcd3f2f087dc"}}`
    buyerKeyJson = `{"version":3,"id":"aaf00f30-3689-499f-9bd3-3ce9fc02c731","address":"3d6f42489e5283c95af70169373d85ba5799bb6f","crypto":{"ciphertext":"d02ee56a66f7688e761bb70b48bee7fdf4c6b1e78e34bc412756a0d94247e8b6","cipherparams":{"iv":"86b9f6ec3f8ae68c4801925b706814ca"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"902a7aec29df2ffe54fc455dfc6c4addc42e02aa495c891fed196c0021d88b0e","n":262144,"r":8,"p":1},"mac":"d2dc1bdd173c8d63403e60838116a985b44e1f3912784316de352bdb99cba471"}}`
    verifier1KeyJson = `{"version":3,"id":"8d13ca05-0c86-40f4-b79e-d13735b8424e","address":"aef420b44068f363fa9905f3fa2d3eb047d8570c","crypto":{"ciphertext":"43237f0c2eebc3a78cacb99134b2a5b2ea436a98415e5d15781c3b6ceb0d1b29","cipherparams":{"iv":"0bd59781851b2dca98f79f1db06519b0"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"2bad47416714e60e41d81437ab5d1c1007842405722eb01826fc7a41ac790b1d","n":262144,"r":8,"p":1},"mac":"e1afb059ca485604c3bed6e768386b15c2722921429d06cd4930069c8c0ca129"}}`
    verifier2KeyJson = `{"version":3,"id":"857cb57c-cdcb-4a19-a3c2-fe5a2ea3b3c2","address":"8c091c18bf57db0896d077b1b778301cab48bc37","crypto":{"ciphertext":"90141512a4b938921f67f943426939b20a21bef08cb9767f09a0ee3bbf33f8f7","cipherparams":{"iv":"568781bb771bb2aedef69a21eedb3e4d"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"c8740dd4d4ba72b5a513bb8986342a47ac93c9c89ea2bb9420d924d78243b343","n":262144,"r":8,"p":1},"mac":"1760cc87aa8eddb269e582cea5213db55c5f615f3527b2941d1d6b6b2ada200d"}}`
    arbitratorKeyJson = `{"version":3,"id":"2ffd7c1d-e948-4f44-ade4-10b8a6619e0a","address":"30c04b0ded7c042c09ad884bdcb8ddb38e536f0e","crypto":{"ciphertext":"8cd9db8c430dd55fbeb56bfd93c91033972229b783c678f6864bcaf4d1291723","cipherparams":{"iv":"471f25a18e9cb8eddf1965d31f2989b8"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"ba77cda872eb0ef19882c099ec47f33dbfc8a27335cb99d1ede4c2ea1963a5a0","n":262144,"r":8,"p":1},"mac":"b2d34b7e879effbe30185e9f6222607733741e94f04444caa608496f4dcbfcad"}}`
    keyPassword = "12345"
    seller *contractclient.ContractClient = nil
    buyer *contractclient.ContractClient = nil
    verifier1 *contractclient.ContractClient = nil
    verifier2 *contractclient.ContractClient = nil
    arbitrator *contractclient.ContractClient = nil
    sleepTime time.Duration = 20000000000
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

    initClients()

    testTxWithoutVerify()

	testTxWithVerify()

	time.Sleep(100000000000000)
}

func testTxWithoutVerify()  {
    fmt.Println("Start testing tx without verification...")

    SellerPublishData(conn, false)

    time.Sleep(sleepTime)

    BuyerApproveTransfer(conn)

    time.Sleep(sleepTime)

    Buy(conn, txId)

    time.Sleep(sleepTime)

    SubmitMetaDataIdEncWithBuyer(conn, txId)

    time.Sleep(sleepTime)

    ConfirmDataTruth(conn, txId)

    fmt.Println("Testing end")
}

func testTxWithVerify()  {
    fmt.Println("Start testing tx with verification...")

    VerifierApproveTransfer(conn, verifier1)

    time.Sleep(sleepTime)

    VerifierApproveTransfer(conn, verifier2)

    time.Sleep(sleepTime)

    VerifierApproveTransfer(conn, arbitrator)

    time.Sleep(sleepTime)


    RegisterAsVerifier(conn, verifier1)

    time.Sleep(sleepTime)

    RegisterAsVerifier(conn, verifier2)

    time.Sleep(sleepTime)

    RegisterAsVerifier(conn, arbitrator)

    time.Sleep(sleepTime)

    SellerPublishData(conn, true)

    time.Sleep(sleepTime)

    BuyerApproveTransfer(conn)

    time.Sleep(sleepTime)

    Buy(conn, txId)

    time.Sleep(sleepTime)

    SubmitMetaDataIdEncWithBuyer(conn, txId)

    time.Sleep(sleepTime)

    ConfirmDataTruth(conn, txId)

    time.Sleep(sleepTime)

    //credit to verifier
    CreditsToVerifier(conn, common.HexToAddress(verifier1.Account.Address))

    CreditsToVerifier(conn, common.HexToAddress(verifier2.Account.Address))

    fmt.Println("Testing end")

}

func initClients()  {
    seller = getClient(sellerKeyJson)
    buyer = getClient(buyerKeyJson)
    verifier1 = getClient(verifier1KeyJson)
    verifier2 = getClient(verifier2KeyJson)
    arbitrator = getClient(arbitratorKeyJson)
}

func getClient(keyJson string) (*contractclient.ContractClient) {
    client, err := contractclient.NewContractClient(getPublicAddress(keyJson))
    if err != nil {
        fmt.Println("failed to create contract client. error:", err)
        return nil
    }

    return client
}

func SellerPublishData(conn *ethclient.Client, supportVerify bool)  {
	seller.SubscribeEvent("DataPublish", onPublish)

	//publish data
	metaData := []byte{'1','2','3','3'}
	proofData := [][]byte{{'4','5','6','3'}, {'2','2', '1'}}
	despData := []byte{'7','8','9','3'}

	txParam := chainoperations.TransactParams{ common.HexToAddress(seller.Account.Address), keyPassword}
	contractinterfacewrapper.Publish(&txParam, big.NewInt(1000), metaData, proofData, 2, despData, supportVerify)
}

func VerifierApproveTransfer(conn *ethclient.Client, verifier *contractclient.ContractClient)  {
    verifier.SubscribeEvent("Approval", onApprovalVerifierTransfer)

    txParam := chainoperations.TransactParams{ common.HexToAddress(verifier.Account.Address), keyPassword}
    err := contractinterfacewrapper.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(10000))
    if err != nil {
        fmt.Println("VerifierApproveTransfer", err)
    }
}

func RegisterAsVerifier(conn *ethclient.Client, verifier *contractclient.ContractClient) {
    verifier.SubscribeEvent("RegisterVerifier", OnRegisterVerifier)

    txParam := chainoperations.TransactParams{ common.HexToAddress(verifier.Account.Address), keyPassword}
    err := contractinterfacewrapper.RegisterAsVerifier(&txParam)
    if err != nil {
        fmt.Println("RegisterAsVerifier", err)
    }
}

func Vote(conn *ethclient.Client, verifier *contractclient.ContractClient) {
    buyer.SubscribeEvent("Vote", onVote)

    txParam := chainoperations.TransactParams{ common.HexToAddress(verifier.Account.Address), keyPassword}
    err := contractinterfacewrapper.Vote(&txParam, txId, true, "This could be real from " + verifier.Account.Address)
    if err != nil {
        fmt.Println("Vote:", err)
    }
}

func CreditsToVerifier(conn *ethclient.Client, to common.Address) {
    buyer.SubscribeEvent("VerifierDisable", onVerifierDisable)

    txParam := chainoperations.TransactParams{ common.HexToAddress(buyer.Account.Address), keyPassword}
    err := contractinterfacewrapper.CreditsToVerifier(&txParam, txId, to, 5)
    if err != nil {
        fmt.Println("CreditsToVerifier:", err)
    }
}

func BuyerApproveTransfer(conn *ethclient.Client)  {
    buyer.SubscribeEvent("Approval", onApprovalBuyerTransfer)

    txParam := chainoperations.TransactParams{ common.HexToAddress(buyer.Account.Address), keyPassword}
    err := contractinterfacewrapper.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(1000))
    if err != nil {
        fmt.Println("BuyerApproveTransfer:", err)
    }
}


func PrepareToBuy(conn *ethclient.Client, publishId string)  {
	buyer.SubscribeEvent("TransactionCreate", onTransactionCreate)
    verifier1.SubscribeEvent("VerifiersChosen", onVerifier1Chosen)
    verifier2.SubscribeEvent("VerifiersChosen", onVerifier2Chosen)
    arbitrator.SubscribeEvent("VerifiersChosen", onVerifier3Chosen)

    txParam := chainoperations.TransactParams{ common.HexToAddress(buyer.Account.Address), keyPassword}
	err := contractinterfacewrapper.PrepareToBuy(&txParam, publishId)
	if err != nil {
		fmt.Println("failed to prepareToBuy, error:", err)
	}
}

func Buy(conn *ethclient.Client, txId *big.Int) {
	seller.SubscribeEvent("Buy", onPurchase)

    txParam := chainoperations.TransactParams{ common.HexToAddress(buyer.Account.Address), keyPassword}
	err := contractinterfacewrapper.BuyData(&txParam, txId)
	if err != nil {
		fmt.Println("failed to buyData, error:", err)
	}
}

func SubmitMetaDataIdEncWithBuyer(conn *ethclient.Client, txId *big.Int)  {
    buyer.SubscribeEvent("ReadyForDownload", onReadyForDownload)

    txParam := chainoperations.TransactParams{ common.HexToAddress(seller.Account.Address), keyPassword}
	err := contractinterfacewrapper.SubmitMetaDataIdEncWithBuyer(&txParam, txId, metaDataIdEncWithBuyer)
	if err != nil {
		fmt.Println("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
	}
}

func ConfirmDataTruth(conn *ethclient.Client, txId *big.Int)  {
    buyer.SubscribeEvent("TransactionClose", onClose)

    txParam := chainoperations.TransactParams{ common.HexToAddress(buyer.Account.Address), keyPassword}
	err := contractinterfacewrapper.ConfirmDataTruth(&txParam, txId, true)
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


func onApprovalBuyerTransfer(event events.Event) bool {
    fmt.Println("onApprovalBuyerTransfer:", event)

    PrepareToBuy(conn, publishId)
    return true
}

func onApprovalVerifierTransfer(event events.Event) bool {
    fmt.Println("onApprovalVerifierTransfer:", event)



    return true
}


func getContracts() ([]chainevents.ContractInfo) {
    protocolEvents := []string{"DataPublish", "TransactionCreate", "RegisterVerifier", "VerifiersChosen", "Vote", "Buy", "ReadyForDownload", "TransactionClose"}
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


func getPublicAddress(keyJson string) (string) {
    publicAddress := chainoperations.DecodeKeystoreAddress([]byte(keyJson))
    return publicAddress
}

func onTransactionCreate(event events.Event) bool {
    fmt.Println("onTransactionCreated:", event)
    txId = event.Data.Get("transactionId").(*big.Int)
    return true
}

func onVerifier1Chosen(event events.Event) bool {
    fmt.Println("onVerifier1Chosen:", event)

    txId = event.Data.Get("transactionId").(*big.Int)
    Vote(conn, verifier1)
    return true
}

func onVerifier2Chosen(event events.Event) bool {
    fmt.Println("onVerifier2Chosen:", event)

    txId = event.Data.Get("transactionId").(*big.Int)
    Vote(conn, verifier2)
    return true
}

func onVerifier3Chosen(event events.Event) bool {
    fmt.Println("onVerifier3Chosen:", event)

    txId = event.Data.Get("transactionId").(*big.Int)
    Vote(conn, arbitrator)
    return true
}

func onPublish(event events.Event) bool {
    fmt.Println("onpublish: ", event)

    publishId = event.Data.Get("publishId").(string)
    return true
}

func OnRegisterVerifier(event events.Event) bool {
    fmt.Println("OnRegisterVerifier: ", event)

    return true
}

func onVote(event events.Event) bool {
    fmt.Println("onVote: ", event)

    return true
}

func onVerifierDisable(event events.Event) bool {
    fmt.Println("onVerifierDisable: ", event)

    return true
}
