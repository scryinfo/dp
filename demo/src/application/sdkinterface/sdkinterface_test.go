package sdkinterface

import (
    "math/big"
    "testing"
)

func TestInitialize(t *testing.T) {
    err := Initialize()
    if err != nil {
        t.Fail()
    }
}

func TestTransferTokenFromDeployer(t *testing.T) {
    Initialize()

    CreateUserWithLogin("111111")

    err := TransferTokenFromDeployer(big.NewInt(100))
    if err != nil {
        t.Fail()
    }
}
