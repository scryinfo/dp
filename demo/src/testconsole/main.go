package main

import (
    "../sdk/contractclient"
    "../sdk/contractclient/contractinterfacewrapper"
    "../sdk/core"
    "../sdk/core/chainevents"
    "../sdk/core/chainoperations"
    "../sdk/core/ethereum/events"
    "../sdk/util/accounts"
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
    keyPassword = "12345"
    deployer *contractclient.ContractClient = nil
    seller *contractclient.ContractClient = nil
    buyer *contractclient.ContractClient = nil
    verifier1 *contractclient.ContractClient = nil
    verifier2 *contractclient.ContractClient = nil
    verifier3 *contractclient.ContractClient = nil
    arbitrator *contractclient.ContractClient = nil
    sleepTime time.Duration = 20000000000
)

func main()  {
    var err error = nil
	conn, err = core.StartEngine("http://127.0.0.1:7545/", "192.168.1.6:48080", getContracts(), "/ip4/127.0.0.1/tcp/5001")
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

	//testAccounts()

    testTxWithoutVerify()

	//testTxWithVerify()

	time.Sleep(100000000000000)
}

func testAccounts()  {
    fmt.Println("Start testing accounts...")

    ac, err := accounts.GetAMInstance().CreateAccount("12345")
    if err != nil {
        fmt.Println("failed to create account, error:", err)
    }

    rv, err := accounts.GetAMInstance().AuthAccount(ac.Address, "12345")
    if err != nil {
        fmt.Println("failed to authenticate account, error:", err)
    }

    if rv {
        fmt.Println("Account authentication passed")
    } else {
        fmt.Println("Account authentication not passed")
    }

    fmt.Println("Test end")
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

    VerifierApproveTransfer(conn, verifier3)

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

    CreditsToVerifier(conn, common.HexToAddress(verifier1.Account.Address))

    CreditsToVerifier(conn, common.HexToAddress(verifier2.Account.Address))

    fmt.Println("Testing end")

}

func initClients()  {
    deployer = getClient(deployerKeyJson)

    var err error
    seller, err = CreateClientWithToken(big.NewInt(1000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    buyer, err = CreateClientWithToken(big.NewInt(1000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }


    verifier1, err = CreateClientWithToken(big.NewInt(1000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    verifier2, err = CreateClientWithToken(big.NewInt(1000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    verifier3, err = CreateClientWithToken(big.NewInt(1000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    arbitrator, err = CreateClientWithToken(big.NewInt(1000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    time.Sleep(sleepTime)
}

func CreateClientWithToken(value *big.Int) (*contractclient.ContractClient, error) {
    client, err := contractclient.CreateContractClient(keyPassword)
    if err != nil {
        fmt.Println("failed to create user, error:", err)
        return nil, err
    }

    txParam := chainoperations.TransactParams{ common.HexToAddress(deployer.Account.Address), keyPassword}
    err = contractinterfacewrapper.TransferTokens(&txParam, common.HexToAddress(client.Account.Address), value)
    if err != nil {
        fmt.Println("failed to transfer token, error:", err)
        return nil, err
    }

    return client, nil
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
