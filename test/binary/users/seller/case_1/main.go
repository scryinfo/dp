package main

import (
    "fmt"
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
        dot.Logger().Errorln("create seller failed. ", zap.NamedError("error", err))
    }

    time.Sleep(time.Second * 3)

    CurUser.SubscribeEvent("DataPublish", onPublish)
    CurUser.SubscribeEvent("TransactionCreate", onTransactionCreate)
    CurUser.SubscribeEvent("TransactionClose", onTransactionClose)

    _, err = Chain.Publish(
       utils.MakeTxParams(CurUser.Account().Addr, password),
       big.NewInt(1000),
       []byte("QmVak3K153a6uEuLQh1etXFWV8Zz3yymGEznR299fz7nDe"),
       []string{"QmRkVKnkni12YfoG4GJcAZzUb9uaFQgDRDzQX3tEBFfqQ3", "QmdGm9Kq6idatRRb74vtEQFUUMJJV94xiT9zzY5E4Cf9BD"},
       int32(2),
       "QmNxBaNK9sFzMguAv7yLj3TaLo1NATmqzYfZj6cdahpjpM",
       false,
    )
    if err != nil {
       fmt.Println("seller publish failed. ", err)
    }
}

func onPublish(event event.Event) bool {
    fmt.Println("> Seller received publish event: ", event.String())

    return true
}

func onTransactionCreate(event event.Event) bool {
    fmt.Println("> Seller received tx create event: ", event.String())

    return true
}

func onTransactionClose(event event.Event) bool {
    fmt.Println("> Seller received tx close event: ", event.String())
    fmt.Print("[exit]")

    return true
}
