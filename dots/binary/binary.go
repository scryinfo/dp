// Scry Info.  All rights reserved.
// license that can be found in the license file.

package binary

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/binary/core"
    "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/event"
    "go.uber.org/zap"
)

const (
    BinTypeId = "92fed326-2f2c-40c5-8664-107473238390"
    startEngineFailed         = "failed to start engine"
    initContractWrapperFailed = "failed to initialize contract interface"
    protocolAbi = `[{"inputs":[{"name":"_token","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"users","type":"address[]"}],"name":"RegisterVerifier","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"price","type":"uint256"},{"indexed":false,"name":"despDataId","type":"string"},{"indexed":false,"name":"supportVerify","type":"bool"},{"indexed":false,"name":"users","type":"address[]"}],"name":"DataPublish","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"proofIds","type":"bytes32[]"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"VerifiersChosen","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"proofIds","type":"bytes32[]"},{"indexed":false,"name":"needVerify","type":"bool"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"TransactionCreate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"judge","type":"bool"},{"indexed":false,"name":"comments","type":"string"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"Vote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"publishId","type":"string"},{"indexed":false,"name":"metaDataIdEncSeller","type":"bytes"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"buyer","type":"address"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"Buy","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"metaDataIdEncBuyer","type":"bytes"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"ReadyForDownload","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"transactionId","type":"uint256"},{"indexed":false,"name":"state","type":"uint8"},{"indexed":false,"name":"index","type":"uint8"},{"indexed":false,"name":"users","type":"address[]"}],"name":"TransactionClose","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"seqNo","type":"string"},{"indexed":false,"name":"verifier","type":"address"},{"indexed":false,"name":"users","type":"address[]"}],"name":"VerifierDisable","type":"event"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"}],"name":"registerAsVerifier","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"publishId","type":"string"},{"name":"price","type":"uint256"},{"name":"metaDataIdEncSeller","type":"bytes"},{"name":"proofDataIds","type":"bytes32[]"},{"name":"despDataId","type":"string"},{"name":"supportVerify","type":"bool"}],"name":"publishDataInfo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"publishId","type":"string"},{"name":"startVerify","type":"bool"}],"name":"createTransaction","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"judge","type":"bool"},{"name":"comments","type":"string"}],"name":"vote","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"}],"name":"buyData","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"}],"name":"cancelTransaction","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"encryptedMetaDataId","type":"bytes"}],"name":"submitMetaDataIdEncWithBuyer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"truth","type":"bool"}],"name":"confirmDataTruth","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"deposit","type":"uint256"}],"name":"setVerifierDepositToken","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"num","type":"uint8"}],"name":"setVerifierNum","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"bonus","type":"uint256"}],"name":"setVerifierBonus","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"seqNo","type":"string"},{"name":"txId","type":"uint256"},{"name":"verifierIndex","type":"uint8"},{"name":"credit","type":"uint8"}],"name":"creditsToVerifier","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
    tokenAbi = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"INITIAL_SUPPLY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
)

type Binary struct {
    chain scry.ChainWrapper
    config binaryConfig
    contracts []event.ContractInfo
}

type binaryConfig struct {
    ethNodeAddr string
    protocolAddr string
    tokenAddr string
    keyServiceAddr string
    storageNodeAddr string
    appId string
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

    dConf := &binaryConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Binary{config: *dConf}

    return d, err
}

//Data structure needed when generating newer component
func BinTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{TypeId: BinTypeId,
            NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                return newBinaryDot(conf)
            }},
    }
}

func (c *Binary) Create(l dot.Line) error {
    c.contracts = c.getContracts(c.config.protocolAddr, c.config.tokenAddr)
	return nil
}

func (c *Binary) Start(ignore bool) error {
    conn, err := core.StartEngine(
        c.config.ethNodeAddr,
        c.config.keyServiceAddr,
        c.contracts,
        c.config.storageNodeAddr)
    if err != nil {
        return errors.New(startEngineFailed)
    }

    c.chain, err = scry.NewChainWrapper(
        common.HexToAddress(c.contracts[0].Address),
        common.HexToAddress(c.contracts[1].Address),
        conn)
    if err != nil {
        return errors.New(initContractWrapperFailed)
    }

	return nil
}

func (c *Binary) Stop(ignore bool) error {
	return nil
}

func (c *Binary) Destroy(ignore bool) error {
	return nil
}

func (c *Binary) getContracts(
    protocolAddr string,
    tokenAddr string) []event.ContractInfo {
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
    }
    tokenEvents := []string{"Approval"}

    contracts := []event.ContractInfo{
        { Address: protocolAddr, Abi: protocolAbi, Events: protocolEvents,},
        { Address: tokenAddr, Abi: tokenAbi, Events: tokenEvents },
    }

    return contracts
}

func (c *Binary) StartScan(fromBlock uint64) {
    core.StartScan(fromBlock)
}

func (c *Binary) StartEngine() (*ethclient.Client, error) {
    logger := dot.Logger()

    defer func() {
        if er := recover(); er != nil {
            logger.Errorln("", zap.Any("failed to initialize start engine, error:", er))
        }
    }()

    err := ipfsaccess2.GetIAInstance().Initialize(ipfsNodeAddr)
    if err != nil {
        logger.Errorln("", zap.NamedError("failed to initialize ipfs. error: ", err))
        return nil, err
    }

    connector, err := newConnector(ethNodeAddr)
    if err != nil {
        logger.Errorln("", zap.NamedError("failed to initialize connector. error: ", err))
        return nil, err
    }

    err = accounts.GetAMInstance().Initialize(asServiceAddr)
    if err != nil {
        logger.Errorln("", zap.NamedError("failed to initialize account service, error:", err))
        return nil, err
    }

    chainevents.StartEventProcessing(connector.conn, contracts)

    return connector.conn, nil
}