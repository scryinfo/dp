package utils

import (
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "io"
    "io/ioutil"
    "math/big"
    "os/exec"
    "time"
)

const (
    dir = "D:/GoPath/src/github.com/scryinfo/dp/test/binary/users/"
    resultFile = "D:/GoPath/src/github.com/scryinfo/dp/test/binary/result.txt"

    deployerAddr = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8"
    deployerPwd = "111111"
)

func RecordResult(r []byte) {
    _ = ioutil.WriteFile(resultFile, r, 0777)
}

func Start(identify string) (process exec.Cmd, pipe io.ReadCloser) {
    user := exec.Cmd{
        Path: identify + ".exe",
        Dir: dir + identify,
    }

    pipe, err := user.StdoutPipe()
    if err != nil {
        fmt.Println(identify, "pipe make failed", err)
    }

    if err := user.Start(); err != nil {
        fmt.Println("sub process start failed", err)
    }

    fmt.Println("> Start sub process: ", identify, " pid:", user.Process.Pid)
    
    time.Sleep(time.Second * 3)
    
    return user, pipe
}

func ReadPipes(pipes []io.ReadCloser, ch chan int) {
    for i := range pipes {
        go func(pipe io.ReadCloser, i int, ch chan int) {
            var buffer = make([]byte, 4096)
            for {
                n, err := pipe.Read(buffer)
                if err != nil {
                    if err == io.EOF {
                        fmt.Println("pipe has Closed: ", i)
                        break
                    } else {
                        fmt.Println("Read content failed")
                    }
                }
                if string(buffer[:n]) == "[exit]" {
                    fmt.Println("Node: pipe ", i, " exit. ")
                    ch <- i
                }
                fmt.Println(string(buffer[:n]))
            }
        }(pipes[i], i, ch)
    }
}

func Stop(process exec.Cmd) {
    process.Process.Kill()

    if err := process.Wait(); err != nil {
        fmt.Printf("> Release sub process: %d, with error: %v\n", process.Process.Pid, err)
    } else {
        fmt.Println("> Release sub process: ", process.Process.Pid)
    }
}

func CreateClientWithToken(password string, cw scry.ChainWrapper) (scry.Client, error) {
    client, err := scry.CreateScryClient(password, cw)
    if err != nil {
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

    return client, nil
}

func MakeTxParams(address, password string) *transaction.TxParams {
    return &transaction.TxParams{
        From:     common.HexToAddress(address),
        Password: password,
        Value:    big.NewInt(0),
        Pending:  false,
    }
}
