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
    CurUser scry.Client
    Chain scry.ChainWrapper
    Listener *listen.Listener
    protocolAddress common.Address
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

    go testCase()

    ssignal.WaitCtrlC(func(s os.Signal) bool {
        return false //quit
    })
}

func testCase() {
    var err error

    if CurUser, err = utils.CreateClientWithEthAndToken(password, Chain); err != nil {
        dot.Logger().Errorln("create buyer failed. ", zap.NamedError("error", err))
    }

    time.Sleep(time.Second * 3)

    CurUser.SubscribeEvent("DataPublish", onPublish)
    CurUser.SubscribeEvent("TransactionCreate", onTransactionCreate)
    CurUser.SubscribeEvent("TransactionClose", onTransactionClose)
}

func onPublish(event event.Event) bool {
    fmt.Println("> buyer received publish event: ", event.String())

    var err error
    err = Chain.ApproveTransfer(utils.MakeTxParams(CurUser.Account().Addr, password), protocolAddress, big.NewInt(1000))
    if err != nil {
        fmt.Println("buyer approve contract transfer token failed. ", err)
    }

    time.Sleep(time.Second * 3)

    err = Chain.PrepareToBuy(
        utils.MakeTxParams(CurUser.Account().Addr, password),
        event.Data.Get("publishId").(string),
        false,
    )
    if err != nil {
        fmt.Println("buyer pre-buy failed. ", err)
    }

    return true
}

func onTransactionCreate(event event.Event) bool {
    fmt.Println("> buyer received tx create event: ", event.String())

    var err error
    err = Chain.CancelTransaction(utils.MakeTxParams(CurUser.Account().Addr, password), event.Data.Get("transactionId").(*big.Int))
    if err != nil {
        fmt.Println("buyer cancel tx failed. ", err)
    }

    return true
}

func onTransactionClose(event event.Event) bool {
    fmt.Println("> buyer received tx close event: ", event.String())
    fmt.Print("[exit]")

    return true
}
