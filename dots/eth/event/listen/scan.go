// Scry Info.  All rights reserved.
// license that can be found in the license file.

package listen

import (
    "context"
    "encoding/binary"
    "errors"
    "fmt"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/scryinfo/dp/dots/eth/event"
    "math/big"
    "reflect"
    "strings"
    "time"
)

var (
	errBadBool     = errors.New("abi: improperly encoded boolean value")
	reflectHash    = reflect.TypeOf(common.Hash{})
	reflectAddress = reflect.TypeOf(common.Address{})
	reflectBigInt  = reflect.TypeOf(new(big.Int))
)

type Progress struct {
	From uint64
	To   uint64
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
		contract: addr,
		abiStr:   abi_str,
		evtNames: evt_names,
	}
	return b
}

func (b *Builder) SetGracefulExit(yes bool) *Builder {
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

func (b *Builder) SetDataChan(dataCh chan<- event.Event, errChan chan<- error) *Builder {
	b.es.DataChan, b.es.ErrChan = dataCh, errChan
	return b
}

func (b *Builder) SetInterval(interval time.Duration) *Builder {
	b.interval = interval
	return b
}

func (b *Builder) BuildAndRun() (*Receipt, error) {
	if err := b.Build(); err != nil {
		return nil, err
	}
	var recipet *Receipt
	if b.es.GracefullExit {
		recipet = PerformSafe(b.es.scan, b.interval)
	} else {
		recipet = Perform(b.es.scan, b.interval)
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
		if len(cm.evtNames) == 0 {
			return errors.New("no event names")
		}
		if cm.abiStr == "" {
			return errors.New("need ABI")
		}
		bc, abi, err := bindContract(cm.abiStr, cm.contract, b.es.conn)
		if err != nil {
			return err
		}
		cm.bc = bc
		cm.abi = abi
		b.es.Contracts[key] = cm
	}
	return nil
}

type contractMeta struct {
	contract common.Address
	abiStr   string
	evtNames []string
	bc       *bind.BoundContract
	abi      *abi.ABI
}

func (cm contractMeta) HasEvent(name string) bool {
	for _, evt := range cm.evtNames {
		if evt == name {
			return true
		}
	}
	return false
}

func (cm contractMeta) UnpackLogToJson(out event.JSONObj, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := cm.Unpack(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cm.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return parseTopics(out, indexed, log.Topics[1:])
}

func (cm contractMeta) Unpack(v event.JSONObj, name string, output []byte) (err error) {
	if len(output) == 0 {
		return fmt.Errorf("abi: unmarshalling empty output")
	}

	if method, ok := cm.abi.Methods[name]; ok {
		if len(output)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted output")
		}
		return cm.UnpackArgs(v, output, method.Outputs)
	} else if event, ok := cm.abi.Events[name]; ok {
		return cm.UnpackArgs(v, output, event.Inputs)
	}
	return fmt.Errorf("abi: could not locate named method or event")
}

func (cm contractMeta) UnpackArgs(v event.JSONObj, data []byte, args abi.Arguments) error {
	if cm.isTuple(args) {
		return cm.unpackTuple(v, data, args)
	}
	return cm.unpackAtomic(v, data, args)
}

func (cm contractMeta) isTuple(args abi.Arguments) bool {
	return len(args) > 1
}

func (cm contractMeta) unpackTuple(v event.JSONObj, output []byte, args abi.Arguments) error {
	i, j := -1, 0
	for _, arg := range args {

		if arg.Indexed {
			// can't read, continue
			continue
		}
		i++
		marshalledValue, err := toGoType((i+j)*32, arg.Type, output)
		if err != nil {
			return err
		}
		v.Set(arg.Name, marshalledValue)
	}
	return nil
}

// unpackAtomic unpacks ( hexdata -> go ) a single value
func (cm contractMeta) unpackAtomic(v event.JSONObj, output []byte, args abi.Arguments) error {
	arg := args[0]
	if arg.Indexed {
		return fmt.Errorf("abi: attempting to unpack indexed variable into element")
	}

	marshalledValue, err := toGoType(0, arg.Type, output)
	if err != nil {
		return err
	}
	v.Set(arg.Name, marshalledValue)
	return nil
}

// toGoType parses the output bytes and recursively assigns the value of these bytes
// into a go type with accordance with the ABI spec.
func toGoType(index int, t abi.Type, output []byte) (interface{}, error) {
	if index+32 > len(output) {
		return nil, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), index+32)
	}

	var (
		returnOutput []byte
		begin, end   int
		err          error
	)

	// if we require a length prefix, find the beginning word and size returned.
	if requiresLengthPrefix(t) {
		begin, end, err = lengthPrefixPointsTo(index, output)
		if err != nil {
			return nil, err
		}
	} else {
		returnOutput = output[index : index+32]
	}

	switch t.T {
	case abi.SliceTy:
		return forEachUnpack(t, output, begin, end)
	case abi.ArrayTy:
		return forEachUnpack(t, output, index, t.Size)
	case abi.StringTy: // variable arrays are written at the end of the return bytes
		return string(output[begin : begin+end]), nil
	case abi.IntTy, abi.UintTy:
		return readInteger(t.Kind, returnOutput), nil
	case abi.BoolTy:
		return readBool(returnOutput)
	case abi.AddressTy:
		return common.BytesToAddress(returnOutput), nil
	case abi.HashTy:
		return common.BytesToHash(returnOutput), nil
	case abi.BytesTy:
		return output[begin : begin+end], nil
	case abi.FixedBytesTy:
		return readFixedBytes(t, returnOutput)
	case abi.FunctionTy:
		return readFunctionType(t, returnOutput)
	default:
		return nil, fmt.Errorf("abi: unknown type %v", t.T)
	}
}

// reads the integer based on its kind
func readInteger(kind reflect.Kind, b []byte) interface{} {
	switch kind {
	case reflect.Uint8:
		return b[len(b)-1]
	case reflect.Uint16:
		return binary.BigEndian.Uint16(b[len(b)-2:])
	case reflect.Uint32:
		return binary.BigEndian.Uint32(b[len(b)-4:])
	case reflect.Uint64:
		return binary.BigEndian.Uint64(b[len(b)-8:])
	case reflect.Int8:
		return int8(b[len(b)-1])
	case reflect.Int16:
		return int16(binary.BigEndian.Uint16(b[len(b)-2:]))
	case reflect.Int32:
		return int32(binary.BigEndian.Uint32(b[len(b)-4:]))
	case reflect.Int64:
		return int64(binary.BigEndian.Uint64(b[len(b)-8:]))
	default:
		return new(big.Int).SetBytes(b)
	}
}

// reads a bool
func readBool(word []byte) (bool, error) {
	for _, b := range word[:31] {
		if b != 0 {
			return false, errBadBool
		}
	}
	switch word[31] {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, errBadBool
	}
}

// A function type is simply the address with the function selection signature at the end.
// This enforces that standard by always presenting it as a 24-array (address + sig = 24 bytes)
func readFunctionType(t abi.Type, word []byte) (funcTy [24]byte, err error) {
	if t.T != abi.FunctionTy {
		return [24]byte{}, fmt.Errorf("abi: invalid type in call to make function type byte array")
	}
	if garbage := binary.BigEndian.Uint64(word[24:32]); garbage != 0 {
		err = fmt.Errorf("abi: got improperly encoded function type, got %v", word)
	} else {
		copy(funcTy[:], word[0:24])
	}
	return
}

// through reflection, creates a fixed array to be read from
func readFixedBytes(t abi.Type, word []byte) (interface{}, error) {
	if t.T != abi.FixedBytesTy {
		return nil, fmt.Errorf("abi: invalid type in call to make fixed byte array")
	}
	// convert
	array := reflect.New(t.Type).Elem()

	reflect.Copy(array, reflect.ValueOf(word[0:t.Size]))
	return array.Interface(), nil

}

func requiresLengthPrefix(t abi.Type) bool {
	return t.T == abi.StringTy || t.T == abi.BytesTy || t.T == abi.SliceTy
}

// interprets a 32 byte slice as an offset and then determines which indice to look to decode the type.
func lengthPrefixPointsTo(index int, output []byte) (start int, length int, err error) {
	offset := int(binary.BigEndian.Uint64(output[index+24 : index+32]))
	if offset+32 > len(output) {
		return 0, 0, fmt.Errorf("abi: cannot marshal in to go slice: offset %d would go over slice boundary (len=%d)", len(output), offset+32)
	}
	length = int(binary.BigEndian.Uint64(output[offset+24 : offset+32]))
	if offset+32+length > len(output) {
		return 0, 0, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), offset+32+length)
	}
	start = offset + 32

	return
}

// iteratively unpack elements
func forEachUnpack(t abi.Type, output []byte, start, size int) (interface{}, error) {
	if start+32*size > len(output) {
		return nil, fmt.Errorf("abi: cannot marshal in to go array: offset %d would go over slice boundary (len=%d)", len(output), start+32*size)
	}

	// this value will become our slice or our array, depending on the type
	var refSlice reflect.Value
	slice := output[start : start+size*32]

	if t.T == abi.SliceTy {
		// declare our slice
		refSlice = reflect.MakeSlice(t.Type, size, size)
	} else if t.T == abi.ArrayTy {
		// declare our array
		refSlice = reflect.New(t.Type).Elem()
	} else {
		return nil, fmt.Errorf("abi: invalid type in array/slice unpacking stage")
	}

	for i, j := start, 0; j*32 < len(slice); i, j = i+32, j+1 {
		// this corrects the arrangement so that we get all the underlying array values
		if t.Elem.T == abi.ArrayTy && j != 0 {
			i = start + t.Elem.Size*32*j
		}
		inter, err := toGoType(i, *t.Elem, output)
		if err != nil {
			return nil, err
		}
		// append the item to our reflect slice
		refSlice.Index(j).Set(reflect.ValueOf(inter))
	}

	// return the interface
	return refSlice.Interface(), nil
}

func parseTopics(out event.JSONObj, fields abi.Arguments, topics []common.Hash) error {
	// Sanity check that the fields and topics match up
	if len(fields) != len(topics) {
		return errors.New("topic/field count mismatch")
	}
	// Iterate over all the fields and reconstruct them from topics
	for _, arg := range fields {
		if !arg.Indexed {
			return errors.New("non-indexed field in topic reconstruction")
		}
		name := arg.Name
		// Try to parse the topic back into the fields based on primitive types
		switch arg.Type.Kind {
		case reflect.Bool:
			if topics[0][common.HashLength-1] == 1 {
				out.Set(name, true)
			} else {
				out.Set(name, false)
			}
		case reflect.Int8:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Int16:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Int32:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Int64:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Uint8:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, uint8(num.Int64()))

		case reflect.Uint16:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, uint16(num.Int64()))

		case reflect.Uint32:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, uint32(num.Int64()))

		case reflect.Uint64:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Uint64())

		default:
			// Ran out of plain primitive types, try custom types
			switch arg.Type.Type {
			case reflectHash: // Also covers all dynamic types
				out.Set(name, topics[0].Hex())

			case reflectAddress:
				var addr common.Address
				copy(addr[:], topics[0][common.HashLength-common.AddressLength:])
				out.Set(name, addr.Hex())

			case reflectBigInt:
				num := new(big.Int).SetBytes(topics[0][:])
				out.Set(name, num)

			default:
				// Ran out of custom types, try the crazies
				switch {
				case arg.Type.T == abi.FixedBytesTy:
					out.Set(name, topics[0][common.HashLength-arg.Type.Size:])

				default:
					return fmt.Errorf("unsupported indexed type: %v", arg.Type)
				}
			}
		}
		topics = topics[1:]
	}
	return nil
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
	DataChan      chan<- event.Event
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

func (es *eventScanner) sendData(evt event.Event) {
	if es.DataChan != nil {
		es.DataChan <- evt
	}
}

func (es *eventScanner) scan(ctx *RedoCtx) {
	newestBn, err := es.NewestBlockNumber()
	if err != nil {
		// not send this err
		if !strings.Contains(err.Error(), "got null header for uncle") {
			es.sendErr(fmt.Errorf("query newest block number fail:%v, will retry later", err))
		}
		return
	}
	if es.From == 0 {
		es.From = newestBn
	}
	var to_bn uint64
	if es.To > 0 && es.To < newestBn {
		to_bn = es.To
	} else {
		to_bn = newestBn
	}
	if es.From > es.To && es.To > 0 {
		ctx.StopRedo()
		return
	}
	if to_bn < es.From {
		return
	}
	if es.From+es.StepLength < to_bn {
		to_bn = es.From + es.StepLength
	}

	fq := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(es.From),
		ToBlock:   new(big.Int).SetUint64(to_bn),
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{},
	}
	fq.Addresses = es.Contracts.Contracts()
	logs, err := es.conn.FilterLogs(context.Background(), fq)
	if err != nil {
		es.sendErr(fmt.Errorf("filter log(%v,%v) err:%v, will retry later", es.From, to_bn, err))
		return
	}
	for _, lg := range logs {
		e := event.NewJSONObj()
		cm, ok := es.Contracts.GetMeta(lg.Address)
		if !ok {
			continue
		}
		name, err := unpackMatchedLog(e, lg, &cm)
		if err != nil {
			es.sendErr(fmt.Errorf("unpack %s log in tx(%s) fail:%v,abadon", name, lg.TxHash.Hex(), err))
			continue
		}
		if !cm.HasEvent(name) {
			continue
		}
		es.sendData(event.Event{
			BlockNumber: lg.BlockNumber,
			TxHash:      lg.TxHash,
			Address:     lg.Address,
			Name:        name,
			Data:        e,
		})
	}
	if es.ProgressChan != nil {
		es.ProgressChan <- Progress{From: es.From, To: to_bn}
	}
	if to_bn < newestBn {
		ctx.StartNextRightNow()
	}
	es.From = to_bn + 1
}

func unpackMatchedLog(out event.JSONObj, log types.Log, meta *contractMeta) (string, error) {
	topicHex := log.Topics[0].Hex()
	for name, evt := range meta.abi.Events {
		if evt.Id().Hex() == topicHex {
			return name, meta.UnpackLogToJson(out, name, log)
		}
	}
	return "", errors.New("can't find matched event")
}

func bindContract(abiStr string, address common.Address, backend bind.ContractBackend) (*bind.BoundContract, *abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, nil, err
	}
	return bind.NewBoundContract(address, parsed, backend, backend, backend), &parsed, nil
}
