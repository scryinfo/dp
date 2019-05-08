package chainevents

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	events2 "github.com/scryInfo/dp/dots/binary/sdk/core/ethereum/events"
	rlog "github.com/sirupsen/logrus"
	"time"
)

var (
	builder *events2.Builder = nil
)

type ContractInfo struct {
	Address string
	Abi     string
	Events  []string
}

func ListenEvent(conn *ethclient.Client, contracts []ContractInfo, fromBlock uint64, interval time.Duration,
	dataChannel chan events2.Event, errorChannel chan error) bool {
	rv := true
	rlog.Info("start listening events...")

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("Failed to listen event. error:", err)
			rv = false
		}
	}()

	if len(contracts) == 0 {
		rlog.Error("invalid contracts parameter")
		return false
	}

	builder = events2.NewScanBuilder()
	for _, v := range contracts {
		builder.SetContract(common.HexToAddress(v.Address), v.Abi, v.Events...)
	}

	recp, err := builder.SetClient(conn).
		SetFrom(fromBlock).
		SetTo(0).
		SetGracefullExit(true).
		SetDataChan(dataChannel, errorChannel).
		SetInterval(interval).
		BuildAndRun()
	if err != nil {
		rlog.Error("failed to listen to events.", err)
		return false
	}

	recp.WaitChan()

	return rv
}

func SetFromBlock(from uint64) {
	if builder != nil {
		builder.SetFrom(from)
	} else {
		rlog.Warn("Failed to set from block because of nil builder.")
	}
}
