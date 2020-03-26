package scan

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
	"math/big"
)

const (
	ConnectTypeId   = "0c73cc80-ed3c-4dc9-8d4f-bcaa4dad29cd"
	HomesteadSigner = "HomesteadSigner" //SignerType
	EIP155Signer    = "EIP155Signer"    //SignerType
)

type connectConfig struct {
	UrlEth     string `json:"urlEth"` //eth连接参数
	ChainId    int64  `json:"chainId"`
	SignerType string `json:"signerType"` //default is "HomesteadSigner", or "EIP155Signer"
}

//add signer and chain id for eth client
type Connect struct {
	conf            connectConfig
	Ctx             context.Context
	cancelFun       context.CancelFunc
	chainId         big.Int
	homesteadSigner types.HomesteadSigner
	frontierSigner  types.FrontierSigner
	eip155Signer    types.EIP155Signer

	EthClient EthClientI   //eth client
	Signer    types.Signer //signer
}

func (c *Connect) Sender(tx *types.Transaction) (common.Address, error) {
	var a common.Address
	var err error
	if tx.ChainId().Cmp(&c.chainId) == 0 {
		a, err = c.eip155Signer.Sender(tx)
	} else {
		a, err = types.NewEIP155Signer(tx.ChainId()).Sender(tx)
	}
	if err != nil {
		a, err = c.homesteadSigner.Sender(tx)
		if err != nil {
			a, err = c.frontierSigner.Sender(tx)
		}
	}
	return a, err
}

//Stop call cancel fun and close the eth client
func (c *Connect) Stop(ignore bool) error {
	if c.cancelFun != nil {
		c.cancelFun()
	}
	if c.EthClient != nil {
		c.EthClient.Close()
	}
	return nil
}

//Create create eth client, chain id, signer
func (c *Connect) Create(l dot.Line) error {
	var err error
	c.Ctx, c.cancelFun = context.WithCancel(context.Background())
	c.EthClient, err = ethclient.DialContext(c.Ctx, c.conf.UrlEth)
	if err == nil {
		c.chainId = *big.NewInt(c.conf.ChainId)
		c.eip155Signer = types.NewEIP155Signer(&c.chainId)
		if len(c.conf.SignerType) < 1 || c.conf.SignerType != EIP155Signer {
			c.Signer = types.HomesteadSigner{}
		} else {
			c.Signer = c.eip155Signer
		}
	}
	return err
}

func newConnect(conf []byte) (d dot.Dot, err error) {
	dconf := &connectConfig{}
	err = dot.UnMarshalConfig(conf, dconf)
	if err != nil {
		return nil, err
	}

	d = &Connect{conf: *dconf}

	return d, err
}

//ConnectTypeLives
func ConnectTypeLives() []*dot.TypeLives {
	lives := []*dot.TypeLives{
		{
			Meta: dot.Metadata{TypeId: ConnectTypeId, NewDoter: func(conf []byte) (dot dot.Dot, err error) {
				return newConnect(conf)
			}},
		},
	}
	return lives
}

//ConnectConfigTypeLives
func ConnectConfigTypeLives() *dot.ConfigTypeLives {
	return &dot.ConfigTypeLives{
		TypeIdConfig: ConnectTypeId,
		ConfigInfo:   &connectConfig{},
	}
}

//ConnectTest just for test
func ConnectTest(eth EthClientI, signer types.Signer, chainId int64) *Connect {
	connect := &Connect{
		EthClient: eth,
		Signer:    signer,
		chainId:   *big.NewInt(chainId),
	}
	return connect
}
