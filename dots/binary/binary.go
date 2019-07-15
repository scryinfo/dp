// Scry Info.  All rights reserved.
// license that can be found in the license file.

package binary

import (
    "context"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/auth"
    "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/currency"
    "github.com/scryinfo/dp/dots/eth/event"
    "github.com/scryinfo/dp/dots/eth/event/execute"
    "github.com/scryinfo/dp/dots/eth/event/listen"
    "github.com/scryinfo/dp/dots/eth/event/subscribe"
    "github.com/scryinfo/dp/dots/grpc"
    "github.com/scryinfo/dp/dots/storage"
    "go.uber.org/zap"
)

const (
    BinTypeId                 = "92fed326-2f2c-40c5-8664-107473238390"
    BinLiveId                 = "92fed326-2f2c-40c5-8664-107473238390"
    startEngineFailed         = "failed to start engine"
    initContractWrapperFailed = "failed to initialize contract interface"
    initAuthServiceFailed     = "failed to initialize authentication service"
    initStorageServiceFailed  = "failed to initialize storage service"
    protocolAbi               = `[{"inputs":[{"name":"_token","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"price","type":"uint256"},{"indexed":false,"name":"despDataId","type":"string"},{"indexed":false,"name":"supportVerify","type":"bool"},{"indexed":false,"name":"users","type":"address[]"}],"name":"DataPublish","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"proofIds","type":"bytes32[]"},{"indexed":false,"name":"needVerify","type":"bool"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"TransactionCreate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"metaDataIdEncSeller","type":"bytes"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"Buy","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"TransactionClose","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"proofIds","type":"bytes32[]"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"VerifiersChosen","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"metaDataIdEncBuyer","type":"bytes"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"ReadyForDownload","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"proofIds","type":"bytes32[]"},{"indexed":false,"name":"metaDataIdEncArbitrator","type":"bytes"},{"indexed":false,"name":"users","type":"address[]"}],"name":"ArbitrationBegin","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"judge","type":"bool"},{"indexed":false,"name":"identify","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"ArbitrationResult","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"users","type":"address[]"}],"name":"RegisterVerifier","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"judge","type":"bool"},{"indexed":false,"name":"comments","type":"string"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"Vote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"verifier","type":"address"},{"indexed":false,"name":"users","type":"address[]"}],"name":"VerifierDisable","type":"event"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"}],"name":"registerAsVerifier","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"judge","type":"bool"},{"name":"comments","type":"string"}],"name":"vote","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"verifierIndex","type":"uint8"},{"name":"credit","type":"uint8"}],"name":"creditsToVerifier","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"judge","type":"bool"}],"name":"arbitrate","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"publishId","type":"string"},{"name":"price","type":"uint256"},{"name":"metaDataIdEncSeller","type":"bytes"},{"name":"proofDataIds","type":"bytes32[]"},{"name":"descDataId","type":"string"},{"name":"supportVerify","type":"bool"}],"name":"publishDataInfo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"publishId","type":"string"},{"name":"startVerify","type":"bool"}],"name":"createTransaction","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"}],"name":"buyData","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"}],"name":"cancelTransaction","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"encryptedMetaDataId","type":"bytes"},{"name":"encryptedMetaDataIds","type":"bytes"}],"name":"reEncryptMetaDataIdBySeller","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"truth","type":"bool"}],"name":"confirmDataTruth","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"txId","type":"uint256"}],"name":"getBuyer","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"txId","type":"uint256"}],"name":"getArbitrators","outputs":[{"name":"","type":"address[]"}],"payable":false,"stateMutability":"view","type":"function"}]`
    tokenAbi                  = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"INITIAL_SUPPLY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
    maxChannelEventNum        = 10000
)

type Binary struct {
    chainWrapper scry.ChainWrapper
    config       BinaryConfig
    contracts    []event.ContractInfo
    subsRepo     *event.Repository
    dataChannel  chan event.Event
    errorChannel chan error
    Executor     *execute.Executor    `dot:""`
    Listener     *listen.Listener     `dot:""`
    Account      *auth.Account        `dot:""`
    Storage      *storage.Ipfs        `dot:""`
    Subscriber   *subscribe.Subscribe `dot:""`
    Grpc         *grpc.BinaryGrpcServer `dot:""`
}

type BinaryConfig struct {
    AppId                string `json:"appId"`
    EthSrvAddr           string `json:"ethServiceAddr"`
    KeySrvAddr           string `json:"keyServiceAddr"`
    StorageSrvAddr       string `json:"storageServiceAddr"`
    ProtocolContractAddr string `json:"protocolContractAddr"`
    TokenContractAddr    string `json:"tokenContractAddr"`
}

//construct dot
func newBinaryDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &BinaryConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Binary{config: *dConf}

    return d, err
}

//Data structure needed when generating newer component
func BinTypeLive() []*dot.TypeLives {
    t := []*dot.TypeLives{
        &dot.TypeLives{
            Meta: dot.Metadata{TypeId: BinTypeId, NewDoter: func(conf interface{}) (dot.Dot, error) {
                return newBinaryDot(conf)
            }},
        },
        execute.ExecutorTypeLive(),
        listen.ListenerTypeLive(),
        auth.AccountTypeLive(),
        storage.IpfsTypeLive(),
        subscribe.SubsTypeLive(),
    }

    t = append(t, currency.CurrTypeLive()...)
    t = append(t, grpc.BinaryGrpcServerTypeLive()...)
    return t
}

func (c *Binary) Create(l dot.Line) error {
    c.contracts = c.getContracts(c.config.ProtocolContractAddr, c.config.TokenContractAddr)
    c.subsRepo = event.NewRepository()
    c.dataChannel = make(chan event.Event, maxChannelEventNum)
    c.errorChannel = make(chan error, 1)

    return nil
}

func (c *Binary) Config() BinaryConfig {
    return c.config
}

func (c *Binary) Start(ignore bool) error {
    c.Subscriber.SetRepo(c.subsRepo)

    conn, err := c.StartEngine()
    if err != nil {
        return errors.New(startEngineFailed)
    }

    c.chainWrapper, err = scry.NewChainWrapper(
        common.HexToAddress(c.contracts[0].Address),
        common.HexToAddress(c.contracts[1].Address),
        conn,
        c.config.AppId,
    )
    if err != nil {
        return errors.New(initContractWrapperFailed)
    }

    err = c.Account.Initialize(c.config.KeySrvAddr)
    if err != nil {
        return errors.New(initAuthServiceFailed)
    }

    err = c.Storage.Initialize(c.config.StorageSrvAddr)
    if err != nil {
        return errors.New(initStorageServiceFailed)
    }

    c.Grpc.SetChainWrapper(c.chainWrapper)

    return nil
}

func (c *Binary) Stop(ignore bool) error {
    return nil
}

func (c *Binary) Destroy(ignore bool) error {
    return nil
}

func (c *Binary) ChainWrapper() scry.ChainWrapper {
    return c.chainWrapper
}

func (c *Binary) getContracts(
    protocolAddr string,
    tokenAddr string,
) []event.ContractInfo {
    protocolEvents := []string{
        "DataPublish",
        "TransactionCreate",
        "RegisterVerifier",
        "VerifiersChosen",
        "Vote",
        "Buy",
        "ReadyForDownload",
        "TransactionClose",
        "VerifierDisable",
        "ArbitrationBegin",
        "ArbitrationResult",
    }
    tokenEvents := []string{"Approval"}

    contracts := []event.ContractInfo{
        {Address: protocolAddr, Abi: protocolAbi, Events: protocolEvents},
        {Address: tokenAddr, Abi: tokenAbi, Events: tokenEvents},
    }

    return contracts
}

func (c *Binary) StartEngine() (*ethclient.Client, error) {
    logger := dot.Logger()

    defer func() {
        if er := recover(); er != nil {
            logger.Errorln("", zap.Any("failed to initialize start engine, error:", er))
        }
    }()

    connector, err := newConnector(c.config.EthSrvAddr)
    if err != nil {
        logger.Errorln("", zap.NamedError("failed to initialize connector. error: ", err))
        return nil, err
    }

    go c.Executor.ExecuteEvents(c.dataChannel, c.subsRepo, c.config.AppId)
    go c.Listener.ListenEvent(
        connector.conn,
        c.contracts,
        0,
        60,
        c.dataChannel,
        c.errorChannel)

    return connector.conn, nil
}

type Connector struct {
    ctx  context.Context
    conn *ethclient.Client
}

func newConnector(ethNodeAddr string) (*Connector, error) {
    cn, err := ethclient.Dial(ethNodeAddr)
    if err != nil {
        return nil, errors.Wrap(err, "Connect to node: "+ethNodeAddr+" failed. ")
    }

    return &Connector{
        ctx:  context.Background(),
        conn: cn,
    }, nil
}
