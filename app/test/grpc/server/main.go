package main

import (
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dot/dots/line"
    "github.com/scryinfo/dp/dots/binary"
    "github.com/scryinfo/scryg/sutils/ssignal"
    "go.uber.org/zap"
    "os"
)

func main() {
    logger := dot.Logger()
    l, err := line.BuildAndStart(func(l dot.Line) error {
        l.PreAdd(binary.BinTypeLive()...)
        return nil
    })

    if err != nil {
        dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
        return
    }

    dot.SetDefaultLine(l)

    _, err = l.ToInjecter().GetByLiveId(dot.LiveId(binary.BinLiveId))
    if err != nil {
        logger.Errorln("load Binary component failed.")
    }

    defer line.StopAndDestroy(l, true)

    ssignal.WaitCtrlC(func(s os.Signal) bool {
        return false //quit
    })
}
