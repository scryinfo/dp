// Scry Info.  All rights reserved.
// license that can be found in the license file.

package client

import (
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/scryinfo/dot/dot"
)

// const
const (
	ClientTypeId = "bee6ae8e-e8a0-4693-ac65-416f618023d5"
)

// Client is a type wrapper around *ethclient.Client.
type Client struct {
	Config    ConfigClient
	ethClient *ethclient.Client
}

// ConfigClient contains URL that used to connects Ethereum API endpoint.
type ConfigClient struct {
	EthEndpoint string `json:"eth_endpoint"`
}

// ClientTypeLive returns TypeLives for generating Client dot.
func ClientTypeLive() *dot.TypeLives {
	newClient := func(conf []byte) (dot.Dot, error) {
		var confClient ConfigClient
		err := dot.UnMarshalConfig(conf, &confClient)
		if err != nil {
			return nil, err
		}

		return &Client{Config: confClient}, nil
	}
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: ClientTypeId, NewDoter: newClient},
	}
}

// Create creates *ethclient.Client from configuration parameters.
func (c *Client) Create(l dot.Line) error {
	if c.ethClient == nil {
		ec, err := ethclient.Dial(c.Config.EthEndpoint)
		if err != nil {
			dot.Logger().Errorln("Client", zap.Error(err))
			return err
		}
		c.ethClient = ec
	}
	return nil
}

// Destroy closes the underlining *ethclient.Client.
func (c *Client) Destroy(ignore bool) error {
	if c != nil {
		if c.ethClient != nil {
			c.ethClient.Close()
			c.ethClient = nil
		}
	}
	return nil
}

// EthClient returns *ethclient.Client that's encapsulated.
func (c *Client) EthClient() *ethclient.Client {
	return c.ethClient
}
