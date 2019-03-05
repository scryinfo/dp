package sdkinterface

import (
    "github.com/scryinfo/iscap/demo/src/application/definition"
    "github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
    "math/big"
    "testing"
    "time"
)

var (
    onPublish = false
    onTxCreate = false
    onApprove = false
)

func TestInitialize(t *testing.T) {
    err := Initialize()
    if err != nil {
        t.Fail()
    }
}

//notice: deployer_keyjson and deployer_password \
//need to be changed to yours before testing
func TestTransferTokenFromDeployer(t *testing.T) {
    Initialize()
    CreateUserWithLogin("111111")
    err := TransferTokenFromDeployer(big.NewInt(400))
    if err != nil {
        t.Fail()
    }
}

func TestUserLogin(t *testing.T) {
    Initialize()
    CreateUserWithLogin("111111")
    login, err := UserLogin(curUser.Account.Address, "111111")
    if err != nil {
        t.Fail()
    }

    if !login {
        t.Fail()
    }
}


func TestPublishData(t *testing.T) {
    Initialize()
    CreateUserWithLogin("111111")

    proofData := []string{"proof1", "proof2", "proof3"}
    data := definition.PubData{
        MetaData: "i am a solid man",
        ProofData: proofData,
        DespData: "evil data",
        Price: 10000,
        Seller: "seller",
        Password: "111111",
    }

    _, err := PublishData(data, OnPublish)
    if err != nil {
        t.Fail()
    }

    time.Sleep(10000000000)

    if !onPublish {
        t.Fail()
    }
}

func OnPublish(event events.Event) bool {
    onPublish = true
    return true
}

func TestApproveTransferForBuying(t *testing.T) {
    password := "111111"

    Initialize()
    CreateUserWithLogin(password)

    ApproveTransferForBuying(password, OnApprove)

    time.Sleep(10000000000)

    if !onApprove {
        t.Fail()
    }
}

func OnApprove(event events.Event) bool {
    onApprove = true
    return true
}

func TestCreateTransaction(t *testing.T) {
    password := "111111"

    Initialize()

    //seller
    CreateUserWithLogin(password)
    proofData := []string{"proof1", "proof2", "proof3"}
    data := definition.PubData{
        MetaData: "i am a solid man",
        ProofData: proofData,
        DespData: "evil data",
        Price: 1000,
        Seller: "seller",
        Password: "111111",
        SupportVerify: false,
    }

    publishId, err := PublishData(data, OnPublish)

    //buyer
    CreateUserWithLogin(password)
    TransferTokenFromDeployer(big.NewInt(1600))
    ApproveTransferForBuying(password, OnApprove)

    err = CreateTransaction(publishId, password, OnTxCreate)
    if err != nil {
        t.Fail()
    }

    time.Sleep(10000000000)

    if !onTxCreate {
        t.Fail()
    }
}

func OnTxCreate(event events.Event) bool {
    onTxCreate = true
    return true
}