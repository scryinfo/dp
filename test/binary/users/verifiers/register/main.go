package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dot/dots/line"
    "github.com/scryinfo/dp/dots/binary"
    "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/event"
    "github.com/scryinfo/dp/dots/eth/event/listen"
    "github.com/scryinfo/dp/test/binary/utils"
    "github.com/scryinfo/scryg/sutils/ssignal"
    "go.uber.org/zap"
    "math/big"
    "os"
    "time"
)

const (
    password = "123456"
)

var (
    CurUser         scry.Client
    Chain           scry.ChainWrapper
    Listener        *listen.Listener
    protocolAddress common.Address

    verifiers []scry.Client
    ch = make(chan int)
)

func main() {
    logger := dot.Logger()
    l, err := line.BuildAndStart(func(l dot.Line) error {
        l.PreAdd(binary.BinTypeLiveWithoutGrpc()...)
        return nil
    })

    if err != nil {
        dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
        return
    }

    d, err := l.ToInjecter().GetByLiveId(dot.LiveId(binary.BinLiveId))
    if err != nil {
        logger.Errorln("load Binary component failed.")
    }

    if bin, ok := d.(*binary.Binary); !ok {
        logger.Errorln("load Binary component failed.", zap.Any("d", d))
    } else {
        Chain = bin.ChainWrapper()
        Listener = bin.Listener
        protocolAddress = common.HexToAddress(bin.Config().ProtocolContractAddr)
    }

    defer line.StopAndDestroy(l, true)

    go register()

    ssignal.WaitCtrlC(func(s os.Signal) bool {
        return false //quit
    })
}

func register() {
    var err error

    for len(verifiers) < 4 {
        if CurUser, err = utils.CreateClientWithEthAndToken(password, Chain); err != nil {
            dot.Logger().Errorln("create verifier failed. ", zap.NamedError("error", err))
        }

        time.Sleep(time.Second * 3)

        CurUser.SubscribeEvent("RegisterVerifier", onRegister)

        err = Chain.ApproveTransfer(utils.MakeTxParams(CurUser.Account().Addr, password), protocolAddress, big.NewInt(10000))

        time.Sleep(time.Second * 3)

        err = Chain.RegisterAsVerifier(utils.MakeTxParams(CurUser.Account().Addr, password))
        if err != nil {
            fmt.Println("verifier register failed. ", err)
        }

        select {
        case <- ch : {
            CurUser.UnSubscribeEvent("RegisterVerifier")
            verifiers = append(verifiers, CurUser)
            CurUser = nil
            fmt.Println("verifier register successful, valid verifier num now is: ", len(verifiers))
        }
        case <- time.After(time.Second * 10) : {
            CurUser.UnSubscribeEvent("RegisterVerifier")
            CurUser = nil
            fmt.Println("verifier register failed with err: timeout. (10s)")
        }
        }
    }

    fmt.Println("[exit]")
}

func onRegister(_ event.Event) bool {
    ch <- 1

    return true
}
