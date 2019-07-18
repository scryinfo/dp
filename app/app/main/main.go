// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dot/dots/line"
    "github.com/scryinfo/dp/dots/connection"
    "github.com/scryinfo/dp/dots/connection/business"
    "github.com/scryinfo/scryg/sutils/ssignal"
    "go.uber.org/zap"
    "os"
)

func main() {
    l, err := line.BuildAndStart(func(l dot.Line) error {
        //todo
        l.PreAdd(business.BusTypeLive()...)
        return nil
    })
    if err != nil {
        dot.Logger().Errorln("Line init failed. ", zap.NamedError("error", err))
        return
    }

    if err := Init(l); err != nil {
        dot.Logger().Errorln("App init failed. ", zap.NamedError("error", err))
        return
    }

    defer line.StopAndDestroy(l, true)

    ssignal.WaitCtrlC(func(s os.Signal) bool {
        return false //quit
    })
}

func Init(l dot.Line) error {
    logger := dot.Logger()

    d, err := l.ToInjecter().GetByLiveId(dot.LiveId(connection.WebSocketTypeId))
    if err != nil {
        logger.Errorln("load connect component failed. ")
        return errors.New("load connect component failed. ")
    }

    // todo: move 'start web server' to 'after start' stage
    if ws, ok := d.(*connection.WSServer); ok {
        if err = ws.Connect(); err != nil {
            logger.Errorln("WebSocket Connect failed. ", zap.NamedError("error", err))
            return errors.New("WebSocket Connect failed. ")
        }
    } else {
        logger.Errorln("load connect component failed.", zap.Any("dot", d))
        return errors.New("load connect component failed. ")
    }

    return nil
}
