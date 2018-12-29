package chainevents

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/events"
	"time"
)

func ListenEvent(conn *ethclient.Client, contractAddr string, abi string,
			eventNames string,  interval time.Duration,
			dataChannel chan events.Event, errorChannel chan error) bool {
	rv := true
	fmt.Println("start listening events...")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to listen event. error:", err)
			rv = false
		}
	}()

	builder := events.NewScanBuilder()
	recp, err := builder.SetClient(conn).
		SetContract(common.HexToAddress(contractAddr), abi, eventNames).
		SetFrom(getFromBlock()).
		SetTo(getToBlock()).
		SetGracefullExit(true).
		SetDataChan(dataChannel, errorChannel).
		SetInterval(interval).
		BuildAndRun()
	if err != nil {
		fmt.Println("failed to listen to events.", err)
		return false
	}

	recp.WaitChan()

	return rv
}

func getFromBlock() uint64 {
	return 0
}

func getToBlock() uint64 {
	return 0
}
