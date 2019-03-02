package sdkinterface

import (
    "github.com/scryinfo/iscap/demo/src/application/definition"
    "math/big"
    "testing"
    "time"
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
    err := TransferTokenFromDeployer(big.NewInt(40000000))
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
    TransferTokenFromDeployer(big.NewInt(40000000))

    proofData := []string{"proof1", "proof2", "proof3"}
    data := definition.PubData{
        "i am a solid man",
        proofData,
        "evil data",
        10000,
        "seller",
    }

    _, err := PublishData(data, "111111")
    if err != nil {
        t.Fail()
    }

    time.Sleep(100000000)
}
