// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dot/dots/line"
    "github.com/scryinfo/dp/dots/connection/business"
    "github.com/scryinfo/scryg/sutils/ssignal"
    "go.uber.org/zap"
    "os"
)

func main() {
    l, err := line.BuildAndStart(func(l dot.Line) error {
        l.PreAdd(business.BusTypeLive()...)
        return nil
    })
    if err != nil {
        dot.Logger().Errorln("Line init failed. ", zap.NamedError("error", err))
        return
    }

    defer line.StopAndDestroy(l, true)

    ssignal.WaitCtrlC(func(s os.Signal) bool {
        return false //quit
    })
}
