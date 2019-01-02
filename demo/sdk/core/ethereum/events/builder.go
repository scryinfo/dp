package events

import (
	"context"
	"errors"
	"fmt"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"../common/redo"
	abi "../mabi"
	bind "../mabi/mbind"
	"math/big"
	"strings"
	"time"
)

type Event struct {
	BlockNumber uint64
	TxHash      common.Hash
	Address     common.Address
	Name        string
	Data        abi.JSONObj
}

type Progress struct {
	From uint64
	To   uint64
}

func (evt Event) String() string {
	return fmt.Sprintf(
		`block: %v,tx: %s,address: %s,event: %s,data: %s`,
		evt.BlockNumber,
		evt.TxHash.Hex(),
		evt.Address.Hex(),
		evt.Name,
		evt.Data.String(),
	)
}

type Builder struct {
	es       *eventScanner
	interval time.Duration
}

func NewScanBuilder() *Builder {
	return &Builder{
		es: &eventScanner{Contracts: make(contractMap)},
	}
}

func (b *Builder) SetClient(conn *ethclient.Client) *Builder {
	b.es.conn = conn
	return b
}

// set addr to address(0) e.g.common.Address{} to filter any contracts with same abi
func (b *Builder) SetContract(addr common.Address, abi_str string, evt_names ...string) *Builder {
	b.es.Contracts[strings.ToLower(addr.Hex())] = contractMeta{
		contract:  addr,
		abi_str:   abi_str,
		evt_names: evt_names,
	}
	return b
}

func (b *Builder) SetGracefullExit(yes bool) *Builder {
	b.es.GracefullExit = yes
	return b
}

func (b *Builder) SetBlockMargin(margin uint64) *Builder {
	b.es.marginBlock = margin
	return b
}

func (b *Builder) SetFrom(f uint64) *Builder {
	b.es.From = f
	return b
}

func (b *Builder) SetStep(f uint64) *Builder {
	b.es.StepLength = f
	return b
}

func (b *Builder) SetTo(f uint64) *Builder {
	b.es.To = f
	return b
}

func (b *Builder) SetProgressChan(pc chan<- Progress) *Builder {
	b.es.ProgressChan = pc
	return b
}

func (b *Builder) SetDataChan(dataCh chan<- Event, errChan chan<- error) *Builder {
	b.es.DataChan, b.es.ErrChan = dataCh, errChan
	return b
}

func (b *Builder) SetInterval(interval time.Duration) *Builder {
	b.interval = interval
	return b
}

func (b *Builder) BuildAndRun() (*redo.Recipet, error) {
	if err := b.Build(); err != nil {
		return nil, err
	}
	var recipet *redo.Recipet
	if b.es.GracefullExit {
		recipet = redo.PerformSafe(b.es.scan, b.interval)
	} else {
		recipet = redo.Perform(b.es.scan, b.interval)
	}
	return recipet, nil
}

func (b *Builder) Build() error {
	if b.es.DataChan == nil {
		return errors.New("data channel should not be empty")
	}
	if b.es.conn == nil {
		return errors.New("no eth client")
	}
	if len(b.es.Contracts) == 0 {
		return errors.New("no contract address")
	}
	for _, ct := range b.es.Contracts {
		if ct.contract == (common.Address{}) && len(b.es.Contracts) != 1 {
			return errors.New("should only one zero contract")
		}
	}
	if b.interval == time.Duration(0) {
		b.interval = time.Second * 3
	}
	if b.es.StepLength == 0 {
		b.es.StepLength = 1000
	}

	for key, cm := range b.es.Contracts {
		if len(cm.evt_names) == 0 {
			return errors.New("no event names")
		}
		if cm.abi_str == "" {
			return errors.New("need ABI")
		}
		bc, err := bindContract(cm.abi_str, cm.contract, b.es.conn)
		if err != nil {
			return err
		}
		cm.bc = bc
		b.es.Contracts[key] = cm
	}
	return nil
}

type contractMeta struct {
	contract  common.Address
	abi_str   string
	evt_names []string
	bc        *bind.BoundContract
}

func (cm contractMeta) HasEvent(name string) bool {
	for _, evt := range cm.evt_names {
		if evt == name {
			return true
		}
	}
	return false
}

type contractMap map[string]contractMeta

func (cm contractMap) Contracts() []common.Address {
	var arr []common.Address
	for _, e := range cm {
		if e.contract != (common.Address{}) {
			arr = append(arr, e.contract)
		}
	}
	return arr
}

func (cm contractMap) Topics() []common.Hash {
	var arr []common.Hash
	for _, e := range cm {
		for _, evt := range e.evt_names {
			arr = append(arr, e.bc.EventTopic(evt))
		}
	}
	return arr
}

func (cm contractMap) GetMeta(addr common.Address) (contractMeta, bool) {
	meta, ok := cm[strings.ToLower(addr.Hex())]
	if !ok {
		if metaAny, okAny := cm[strings.ToLower((common.Address{}).Hex())]; okAny {
			return metaAny, okAny
		}
	}
	return meta, ok
}

type eventScanner struct {
	conn          *ethclient.Client
	Contracts     contractMap
	From          uint64
	StepLength    uint64
	To            uint64
	DataChan      chan<- Event
	ErrChan       chan<- error
	ProgressChan  chan<- Progress
	GracefullExit bool
	marginBlock   uint64
}

func (es *eventScanner) NewestBlockNumber() (uint64, error) {
	block, err := es.conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return block.Number.Uint64() - es.marginBlock, nil
}

func (es *eventScanner) sendErr(err error) {
	if es.ErrChan != nil && err != nil {
		es.ErrChan <- err
	}
}

func (es *eventScanner) sendData(evt Event) {
	if es.DataChan != nil {
		es.DataChan <- evt
	}
}

func (es *eventScanner) scan(ctx *redo.RedoCtx) {
	newest_bn, err := es.NewestBlockNumber()
	if err != nil {
		// not send this err
		if !strings.Contains(err.Error(), "got null header for uncle") {
			es.sendErr(fmt.Errorf("query newest block number fail:%v, will retry later", err))
		}
		return
	}
	if es.From == 0 {
		es.From = newest_bn
	}
	var to_bn uint64
	if es.To > 0 && es.To < newest_bn {
		to_bn = es.To
	} else {
		to_bn = newest_bn
	}
	if es.From > es.To && es.To > 0 {
		ctx.StopRedo()
		return
	}
	if to_bn <= es.From {
		return
	}
	if es.From+es.StepLength < to_bn {
		to_bn = es.From + es.StepLength
	}
	var topics []common.Hash = es.Contracts.Topics()

	fq := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(es.From),
		ToBlock:   new(big.Int).SetUint64(to_bn),
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{topics},
	}
	fq.Addresses = es.Contracts.Contracts()
	logs, err := es.conn.FilterLogs(context.Background(), fq)
	if err != nil {
		es.sendErr(fmt.Errorf("filter log(%v,%v) err:%v, will retry later", es.From, to_bn, err))
		return
	}
	for _, lg := range logs {
		evt := abi.NewJSONObj()
		cm, ok := es.Contracts.GetMeta(lg.Address)
		if !ok {
			continue
		}
		name, err := cm.bc.UnpackMatchedLog(evt, lg)
		if err != nil {
			es.sendErr(fmt.Errorf("unpack %s log in tx(%s) fail:%v,abadon", name, lg.TxHash.Hex(), err))
			continue
		}
		if !cm.HasEvent(name) {
			continue
		}
		es.sendData(Event{
			BlockNumber: lg.BlockNumber,
			TxHash:      lg.TxHash,
			Address:     lg.Address,
			Name:        name,
			Data:        evt,
		})
	}
	if es.ProgressChan != nil {
		es.ProgressChan <- Progress{From: es.From, To: to_bn}
	}
	if to_bn < newest_bn {
		ctx.StartNextRightNow()
	}
	es.From = to_bn + 1
}

func bindContract(abi_str string, address common.Address, backend bind.ContractBackend) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(abi_str))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, backend, backend, backend), nil
}
