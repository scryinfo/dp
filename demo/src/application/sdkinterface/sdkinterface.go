package sdkinterface

import (
    "errors"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
    "github.com/scryinfo/iscap/demo/src/sdk"
    "github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
    "github.com/scryinfo/iscap/demo/src/sdk/core/chainoperations"
    "github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
    "github.com/scryinfo/iscap/demo/src/sdk/scryclient"
    "github.com/scryinfo/iscap/demo/src/application/definition"
    cif "github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
    "io/ioutil"
    "math/big"
    "strings"
)

var (
    curUser *scryclient.ScryClient = nil
    scryInfo *settings.ScryInfo        = nil
    failedToInitSDK = "failed to initialize sdk. "
    sep = "|"
)

func Initialize() error {
    // load definition
    scryInfo, err := settings.LoadSettings()
    if err != nil {
        fmt.Println(failedToInitSDK, err)
        return errors.New(failedToInitSDK)
    }

    // initialization
    contracts := getContracts(scryInfo.Chain.Contracts.ProtocolAddr,
        scryInfo.Chain.Contracts.TokenAddr,
        scryInfo.Chain.Contracts.ProtocolAbiPath,
        scryInfo.Chain.Contracts.TokenAbiPath,
        scryInfo.Chain.Contracts.ProtocolEvents,
        scryInfo.Chain.Contracts.TokenEvents)

    err = sdk.Init(scryInfo.Chain.Ethereum.EthNode,
        scryInfo.Services.Keystore,
        contracts,
        0,
        scryInfo.Services.Ipfs)
    if err != nil {
        fmt.Println(failedToInitSDK, err)
        return errors.New(failedToInitSDK)
    }

    return nil
}

//new user
func CreateUser(password string) (*scryclient.ScryClient, error) {
    client, err := scryclient.CreateScryClient(password)
    if err != nil {
        return nil, err
    }

    return client, nil
}

func CreateUserWithLogin(password string) (*scryclient.ScryClient, error) {
    client, err := scryclient.CreateScryClient(password)
    if err != nil {
        return nil, err
    }

    curUser = client

    return client, nil
}

func UserLogin(address string, password string) (bool, error) {
    client := scryclient.NewScryClient(address)
    if client == nil {
        return false, errors.New("failed to NewScryClient")
    }

    succ, err := client.Authenticate(password)
    if err != nil {
        return false, errors.New("failed to authenticate user")
    }

    if succ {
        curUser = client
    }

    return succ, nil
}

func PublishData(data definition.PubData, password string) (string, error) {
    if curUser == nil {
        return "", nil
    }

    txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
                                             Password: password,
                                             Value: big.NewInt(0),
                                             Pending: false}

    proofs := convertProofs(data)
    return cif.Publish(&txParam,
        big.NewInt(int64(data.Price)),
        []byte(data.MetaData),
        proofs,
        len(data.ProofData),
        []byte(data.DespData),
        true)
}

func convertProofs(data definition.PubData) [][]byte {
    proofs := make([][]byte, len(data.ProofData))
    for _, proof := range data.ProofData {
        proofs = append(proofs, []byte(proof))
    }

    return proofs
}

func getContracts(protocolContractAddr string,
                  tokenContractAddr string,
                  protocolAbiPath string,
                  tokenAbiPath string,
                  protocolEvents string,
                  tokenEvents string) []chainevents.ContractInfo {
    pe := strings.Split(protocolEvents, sep)
    te := strings.Split(tokenEvents, sep)

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, getAbiText(protocolAbiPath), pe},
		{tokenContractAddr, getAbiText(tokenAbiPath), te},
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

func onPublish(event events.Event) bool {
	//if err := bootstrap.SendMessage(w, "onPublish", "Publish event callback from go"); err != nil {
	//	astilog.Error(errors.Wrap(err, "sending onPublish event failed"))
	//}
	return true
}
