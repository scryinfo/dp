package nonce

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// globalClient is the singleton of interface client.
var globalClient client

// A client represents the interface to fetch a nonce
// from the Ethereum network to an address at pending state that should be used for the next transactions.
// client interface is implemented by type ethClient.
type client interface {
	pendingNonceAt(address string) (uint64, error)
	close() error
}

// A client defines typed wrapper for the Ethereum RPC API.
type ethClient struct {
	endpoint  string
	pendingNonceFetcher
}

// A pendingNonceFetcher represents a client connecting to the Ethereum network
// for a nonce to an address at pending state that should be used for the next transactions.
type pendingNonceFetcher interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	Close()
}

// getClient connects a Ethereum client to the given URL.
// We are dead if fails to connect to the Ethereum network.
func getClient(endpoint string) client {
	if globalClient == nil {
		ec, err := ethclient.Dial(endpoint)
		if err != nil {
			panic(err)
		}
		globalClient = &ethClient{endpoint,ec}
	}
	return globalClient
}

// pendingNonce returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
// Note that in case of invalid address, 0 and nil will be returned.
func (e *ethClient) pendingNonceAt(address string) (uint64, error) {
	return 	e.PendingNonceAt(context.Background(), common.HexToAddress(address))
}

// close closes the underlying connection to the Ethereum network.
func (e *ethClient) close() error {
	e.Close()
	return nil
}