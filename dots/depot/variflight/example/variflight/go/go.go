package main

import (
	"fmt"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/dp/dots/depot/variflight/golang"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
)

func main() {
	// build line and configure dots
	l, err := line.BuildAndStart(buildNewer)
	if err != nil {
		dot.Logger().Debugln("BuildAndStart failed.", zap.Error(err))
		os.Exit(1)
	}
	defer line.StopAndDestroy(l, false)

	//get GrpcWebSocketServer component
	websocketSrv, err := l.ToInjecter().GetByLiveId(golang.GrpcWebSocketServerTypeId)
	if err != nil {
		dot.Logger().Debugln("GetByliveId(golang.VariFlightServerTypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return fmt.Sprintf("VariFlightServer: %#+v", websocketSrv.(*golang.GrpcWebSocketServer))
	})
	fmt.Println("VariFlightServer component now can work normally.")

	ssignal.WaitCtrlC(nil)
}

func buildNewer(l dot.Line) error {
	if err := l.PreAdd(golang.GrpcWebSocketServerTypeLives()...); err != nil {
		dot.Logger().Debugln("PreAdd failed.", zap.Error(err))
		os.Exit(1)
	}
	return nil
}
