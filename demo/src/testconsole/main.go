package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iscap/demo/src/sdk"
	"github.com/iscap/demo/src/sdk/core/chainevents"
	"github.com/iscap/demo/src/sdk/core/chainoperations"
	"github.com/iscap/demo/src/sdk/core/ethereum/events"
	"github.com/iscap/demo/src/sdk/scryclient"
	cif "github.com/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/iscap/demo/src/sdk/util/accounts"
	"io/ioutil"
	"math/big"
	"time"
)

var (
	publishId                        = ""
	txId                    *big.Int = big.NewInt(0)
	metaDataIdEncWithSeller []byte
	metaDataIdEncWithBuyer  []byte
	protocolContractAddr                           = "0x4ae12e473d8eaf98d0410555508cc8ba8c3cf7b6"
	tokenContractAddr                              = "0x77d631c8c87bd3dd0d5631add4d76364662b2159"
	deployerKeyJson                                = `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
	keyPassword                                    = "12345"
	deployer                *scryclient.ScryClient = nil
	seller                  *scryclient.ScryClient = nil
	buyer                   *scryclient.ScryClient = nil
	verifier1               *scryclient.ScryClient = nil
	verifier2               *scryclient.ScryClient = nil
	verifier3               *scryclient.ScryClient = nil
	arbitrator              *scryclient.ScryClient = nil
	sleepTime               time.Duration          = 20000000000
)

func main() {
	//note: asServiceAddr is the host of key management service which is outside
	err := sdk.Init("http://127.0.0.1:7545/", "192.168.1.6:48080", getContracts(), 0, "/ip4/127.0.0.1/tcp/5001", common.HexToAddress(protocolContractAddr), common.HexToAddress(tokenContractAddr))
	if err != nil {
		fmt.Println("failed to initialize sdk, error:", err)
		return
	}

	initClients()

	testAccounts()

	testTransferEth()

	//testTxWithoutVerify()

	testTxWithVerify()

	time.Sleep(100000000000000)
}

func testAccounts() {
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

func testTxWithoutVerify() {
	fmt.Println("Start testing tx without verification...")

	SellerPublishData(false)

	time.Sleep(sleepTime)

	BuyerApproveTransfer()

	time.Sleep(sleepTime)

	Buy(txId)

	time.Sleep(sleepTime)

	SubmitMetaDataIdEncWithBuyer(txId)

	time.Sleep(sleepTime)

	ConfirmDataTruth(txId)

	fmt.Println("Testing end")
}

func testTxWithVerify() {
	fmt.Println("Start testing tx with verification...Note: please restart the test chain before running this case")

	VerifierApproveTransfer(verifier1)

	time.Sleep(sleepTime)

	VerifierApproveTransfer(verifier2)

	time.Sleep(sleepTime)

	VerifierApproveTransfer(verifier3)

	time.Sleep(sleepTime)

	RegisterAsVerifier(verifier1)

	time.Sleep(sleepTime)

	RegisterAsVerifier(verifier2)

	time.Sleep(sleepTime)

	RegisterAsVerifier(verifier3)

	time.Sleep(sleepTime)

	SellerPublishData(true)

	time.Sleep(sleepTime)

	BuyerApproveTransfer()

	time.Sleep(sleepTime)

	Buy(txId)

	time.Sleep(sleepTime)

	SubmitMetaDataIdEncWithBuyer(txId)

	time.Sleep(sleepTime)

	ConfirmDataTruth(txId)

	time.Sleep(sleepTime)

	fmt.Println("Testing end")

}

func testTransferEth() {
	balanceBefore, err := cif.GetEthBalance(common.HexToAddress(deployer.Account.Address))
	if err != nil {
		fmt.Println("failed to transfer eth. error:", err)
		panic(err)
	}

	fmt.Println("buyer's balance before transfer:", balanceBefore)

	//transfer
	_, err = cif.TransferEth(common.HexToAddress(deployer.Account.Address),
		keyPassword,
		common.HexToAddress(buyer.Account.Address),
		big.NewInt(10))
	if err != nil {
		fmt.Println("failed to transfer eth. error:", err)
		panic(err)
	}

	balanceAfter, err := cif.GetEthBalance(common.HexToAddress(deployer.Account.Address))
	if err != nil {
		fmt.Println("failed to transfer eth. error:", err)
		panic(err)
	}

	fmt.Println("buyer's balance after transfer:", balanceAfter)

	sum := big.NewInt(0)
	(*big.Int).Add(sum, balanceAfter, big.NewInt(10))

	if sum.String() != balanceBefore.String() {
		fmt.Println("failed to transfer eth.")
		panic("failed to transfer eth.")
	}

	fmt.Println("Testing end")
}

func initClients() {
	var err error
	deployer, err = ImportAccount(deployerKeyJson, keyPassword, keyPassword)
	if err != nil {
		fmt.Println("failed to init clients, error:", err)
		panic(err)
	}

	seller, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("failed to init clients, error:", err)
		panic(err)
	}

	buyer, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("failed to init clients, error:", err)
		panic(err)
	}

	verifier1, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("failed to init clients, error:", err)
		panic(err)
	}

	verifier2, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("failed to init clients, error:", err)
		panic(err)
	}

	verifier3, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("failed to init clients, error:", err)
		panic(err)
	}

	time.Sleep(sleepTime)
}

func CreateClientWithToken(token *big.Int, eth *big.Int) (*scryclient.ScryClient, error) {
	client, err := scryclient.CreateScryClient(keyPassword)
	if err != nil {
		fmt.Println("failed to create user, error:", err)
		return nil, err
	}

	_, err = cif.TransferEth(common.HexToAddress(deployer.Account.Address),
		keyPassword,
		common.HexToAddress(client.Account.Address),
		big.NewInt(0))
	if err != nil {
		fmt.Println("failed to transfer token, error:", err)
		return nil, err
	}

	txParam := chainoperations.TransactParams{common.HexToAddress(deployer.Account.Address), keyPassword, big.NewInt(0), false}
	err = cif.TransferTokens(&txParam, common.HexToAddress(client.Account.Address), token)
	if err != nil {
		fmt.Println("failed to transfer token, error:", err)
		return nil, err
	}

	return client, nil
}

func getClient(keyJson string) *scryclient.ScryClient {
	client, err := scryclient.NewScryClient(getPublicAddress(keyJson))
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return nil
	}

	return client
}

func ImportAccount(keyJson string, oldPassword string, newPassword string) (*scryclient.ScryClient, error) {
	address, err := accounts.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword)
	if err != nil {
		fmt.Println("failed to import account. error:", err)
		return nil, err
	}

	client, err := scryclient.NewScryClient(address)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return nil, err
	}

	return client, nil
}

func SellerPublishData(supportVerify bool) {
	seller.SubscribeEvent("DataPublish", onPublish)

	//publish data
	metaData := []byte{'1', '2', '3', '3'}
	proofData := [][]byte{{'4', '5', '6', '3'}, {'2', '2', '1'}}
	despData := []byte{'7', '8', '9', '3'}

	txParam := chainoperations.TransactParams{common.HexToAddress(seller.Account.Address), keyPassword, big.NewInt(0), false}
	cif.Publish(&txParam, big.NewInt(1000), metaData, proofData, 2, despData, supportVerify)
}

func VerifierApproveTransfer(verifier *scryclient.ScryClient) {
	verifier.SubscribeEvent("Approval", onApprovalVerifierTransfer)

	txParam := chainoperations.TransactParams{common.HexToAddress(verifier.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(10000))
	if err != nil {
		fmt.Println("VerifierApproveTransfer", err)
	}
}

func RegisterAsVerifier(verifier *scryclient.ScryClient) {
	verifier.SubscribeEvent("RegisterVerifier", OnRegisterVerifier)

	txParam := chainoperations.TransactParams{common.HexToAddress(verifier.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.RegisterAsVerifier(&txParam)
	if err != nil {
		fmt.Println("RegisterAsVerifier", err)
	}
}

func Vote(verifier *scryclient.ScryClient) {
	buyer.SubscribeEvent("Vote", onVote)

	txParam := chainoperations.TransactParams{common.HexToAddress(verifier.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.Vote(&txParam, txId, true, "This could be real from "+verifier.Account.Address)
	if err != nil {
		fmt.Println("Vote:", err)
	}
}

func CreditsToVerifier(to common.Address) {
	buyer.SubscribeEvent("VerifierDisable", onVerifierDisable)

	txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.CreditsToVerifier(&txParam, txId, to, 5)
	if err != nil {
		fmt.Println("CreditsToVerifier:", err)
	}
}

func BuyerApproveTransfer() {
	buyer.SubscribeEvent("Approval", onApprovalBuyerTransfer)

	txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(1000))
	if err != nil {
		fmt.Println("BuyerApproveTransfer:", err)
	}
}

func PrepareToBuy(publishId string) {
	buyer.SubscribeEvent("TransactionCreate", onTransactionCreate)
	verifier1.SubscribeEvent("VerifiersChosen", onVerifier1Chosen)
	verifier2.SubscribeEvent("VerifiersChosen", onVerifier2Chosen)
	verifier3.SubscribeEvent("VerifiersChosen", onVerifier3Chosen)

	txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.PrepareToBuy(&txParam, publishId)
	if err != nil {
		fmt.Println("failed to prepareToBuy, error:", err)
	}
}

func Buy(txId *big.Int) {
	seller.SubscribeEvent("Buy", onPurchase)

	txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.BuyData(&txParam, txId)
	if err != nil {
		fmt.Println("failed to buyData, error:", err)
	}
}

func SubmitMetaDataIdEncWithBuyer(txId *big.Int) {
	buyer.SubscribeEvent("ReadyForDownload", onReadyForDownload)

	txParam := chainoperations.TransactParams{common.HexToAddress(seller.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.SubmitMetaDataIdEncWithBuyer(&txParam, txId, metaDataIdEncWithBuyer)
	if err != nil {
		fmt.Println("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
	}
}

func ConfirmDataTruth(txId *big.Int) {
	buyer.SubscribeEvent("TransactionClose", onClose)

	txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), keyPassword, big.NewInt(0), false}
	err := cif.ConfirmDataTruth(&txParam, txId, true)
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

	PrepareToBuy(publishId)
	return true
}

func onApprovalVerifierTransfer(event events.Event) bool {
	fmt.Println("onApprovalVerifierTransfer:", event)

	return true
}

func getContracts() []chainevents.ContractInfo {
	protocolEvents := []string{"DataPublish", "TransactionCreate", "RegisterVerifier", "VerifiersChosen", "Vote", "Buy", "ReadyForDownload", "TransactionClose"}
	tokenEvents := []string{"Approval"}

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, getAbiText("./ScryProtocol.abi"), protocolEvents},
		{tokenContractAddr, getAbiText("./ScryToken.abi"), tokenEvents},
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

func getPublicAddress(keyJson string) string {
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
	Vote(verifier1)
	return true
}

func onVerifier2Chosen(event events.Event) bool {
	fmt.Println("onVerifier2Chosen:", event)

	txId = event.Data.Get("transactionId").(*big.Int)
	Vote(verifier2)
	return true
}

func onVerifier3Chosen(event events.Event) bool {
	fmt.Println("onVerifier3Chosen:", event)

	txId = event.Data.Get("transactionId").(*big.Int)
	Vote(verifier3)
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
