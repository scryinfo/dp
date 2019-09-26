package utils

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "math/big"
)

func CreateClientWithEthAndToken(password string, cw scry.ChainWrapper) (client scry.Client, err error) {
    if client, err = scry.CreateScryClient(password, cw); err != nil {
        return nil, err
    }

    dot.Logger().Debugln("client:" + client.Account().Addr)

    err = client.TransferEthFrom(
        common.HexToAddress(deployerAddr),
        deployerPwd,
        big.NewInt(10000000),
        cw.Conn(),
    )
    if err != nil {
        return nil, err
    }

    err = cw.TransferTokens(
        MakeTxParams(deployerAddr, deployerPwd),
        common.HexToAddress(client.Account().Addr),
        big.NewInt(10000000),
    )
    if err != nil {
        return nil, err
    }

    return
}

func MakeTxParams(address, password string) *transaction.TxParams {
    return &transaction.TxParams{
        From:     common.HexToAddress(address),
        Password: password,
        Value:    big.NewInt(0),
        Pending:  false,
    }
}
