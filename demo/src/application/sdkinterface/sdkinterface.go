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
    "github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
    "io/ioutil"
    "math/big"
    "strings"
)

const (
	failedToInitSDK = "failed to initialize sdk."
)

var (
    curUser *scryclient.ScryClient = nil
    deployer *scryclient.ScryClient = nil
    scryInfo *settings.ScryInfo = nil
    sep = "|"

)

func Initialize() error {
    // load definition
    var err error
    scryInfo, err = settings.LoadSettings()
    if err != nil {
        fmt.Println(failedToInitSDK, err)
        return err
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
        return err
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
        return false, errors.New("failed to call NewScryClient")
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

func ImportAccount(keyJson string, oldPassword string, newPassword string) (*scryclient.ScryClient, error) {
    address, err := accounts.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword)
    if err != nil {
        fmt.Println("failed to import account. error:", err)
        return nil, err
    }

    client := scryclient.NewScryClient(address)
    return client, nil
}

func TransferTokenFromDeployer(token *big.Int) (error) {
    var err error
    if deployer == nil {
        deployer, err = ImportAccount(scryInfo.Chain.Contracts.DeployerKeyJson,
                                        scryInfo.Chain.Contracts.DeployerPassword,
                                        scryInfo.Chain.Contracts.DeployerPassword)
        if err != nil {
            fmt.Println("failed to transfer token, error:", err)
            return err
        }
    }

    if curUser == nil {
        fmt.Println("failed to transfer token, null current user")
        return errors.New("failed to transfer token, null current user")
    }

    txParam := chainoperations.TransactParams{From: common.HexToAddress(deployer.Account.Address),
                                              Password: scryInfo.Chain.Contracts.DeployerPassword,
                                              Value: big.NewInt(0),
                                              Pending: false}
    err = cif.TransferTokens(&txParam, common.HexToAddress(curUser.Account.Address), token)
    if err != nil {
        fmt.Println("failed to transfer token, error:", err)
        return err
    }

    return nil
}

func PublishData(data definition.PubData, password string) (string, error) {
    if curUser == nil {
        fmt.Println("no current user")
        return "", nil
    }

    curUser.SubscribeEvent("DataPublish", onPublish)
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
    fmt.Println("onpublish: ", event)
    return true
}