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
    // todo: config file?
    dir        = "D:/GoPath/src/github.com/scryinfo/dp/test/binary/users/"
    resultFile = "D:/GoPath/src/github.com/scryinfo/dp/test/binary/result.txt"

    deployerAddr = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8"
    deployerPwd  = "111111"
)

func RecordResult(ch chan string, caseNum int) {
    var result []string
    for {
        t := <- ch
        result = append(result, t)
        if len(result) == caseNum {
            break
        }
    }
    
    var str string
    for i := range result {
        str += result[i] + "\r\n"
    }
    
    _ = ioutil.WriteFile(resultFile, []byte(str), 0666)
}

func Start(identify string, caseNo string) (user exec.Cmd, pipe io.ReadCloser) {
    user = exec.Cmd{
        Path: "main.exe",
        Dir:  dir + identify + "/" + caseNo,
    }

    var err error

    if pipe, err = user.StdoutPipe(); err != nil {
        fmt.Println("make ", identify, " pipe failed", err)
    }

    if err = user.Start(); err != nil {
        fmt.Println("sub process start failed", err)
    }

    fmt.Println("> Start sub process: ", identify, " pid:", user.Process.Pid)
    
    time.Sleep(time.Second * 3)
    
    return
}

func ReadPipes(pipes []io.ReadCloser, ch chan int) {
    for i := range pipes {
        go func(pipe io.ReadCloser, i int, ch chan int) {
            var buffer = make([]byte, 4096)
            for {
                n, err := pipe.Read(buffer)
                if err != nil {
                    if err == io.EOF {
                        fmt.Println("pipe ", i, " has closed. ")
                        break
                    } else {
                        fmt.Println("Read pipe", i, " failed. ")
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

func Stop(processes []exec.Cmd, ch chan int, chStr chan string, userNum int) {
    for i := 0; i < userNum; i++ {
        t := <-ch
        fmt.Println("Node: receive exit flag. ", t)
        stop(processes[t])
    }

    chStr <- "Test case 1 passed. "
}

func stop(process exec.Cmd) {
    process.Process.Kill()

    if err := process.Wait(); err != nil {
        fmt.Printf("> Release sub process: %d, with error: %v\n", process.Process.Pid, err)
    } else {
        fmt.Println("> Release sub process: ", process.Process.Pid)
    }
}

func CreateClientWithToken(password string, cw scry.ChainWrapper) (client scry.Client, err error) {
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
