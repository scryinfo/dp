package chainevents

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"../ethereum/events"
	"time"
)

type ContractInfo struct {
    Address string
    Abi string
    EventNames []string
}

func ListenEvent(conn *ethclient.Client, contracts []ContractInfo,
			interval time.Duration,	dataChannel chan events.Event,
			errorChannel chan error) bool {
	rv := true
	fmt.Println("start listening events...")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to listen event. error:", err)
			rv = false
		}
	}()

    if len(contracts) == 0 {
        fmt.Println("zero contracts")
        return false
    }

    builder := events.NewScanBuilder()
    for _, v := range contracts {
        builder.SetContract(common.HexToAddress(v.Address), v.Abi, v.EventNames...)
    }

	recp, err := builder.SetClient(conn).
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
