// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// Test getClient connecting to an ethereum node hosted by Infura.
func Test_getClient(t *testing.T) {
	scenarios := []struct{
		desc string
		inputEndpoint string
		expectPanic bool
	} {
		{"happy path", "https://mainnet.infura.io/v3/d0a6c76d448d4f16baeb8143050a1f2f", false},
		{"input error path", "invalid/endpoint", true},
	}

	assert := assert.New(t)
	for _, scenario := range scenarios {
		f := func() {
			getClient(scenario.inputEndpoint)
		}
		if !scenario.expectPanic {
			assert.NotPanicsf(f, "getClient shouldn't panic on input '%v'\n",scenario.inputEndpoint)
		}
		assert.NotPanicsf(f, "getClient should panic on input '%v'\n",scenario.inputEndpoint)
	}
}

// Test PendingNonceAt on mainnet of the Ethereum network.
func Test_PendingNonceAt_OnMainETHNetwork(t *testing.T) {
    client := getClient("https://mainnet.infura.io/v3/d0a6c76d448d4f16baeb8143050a1f2f")

	scenarios := []struct{
		desc string
		inputAddress string
		expectErr bool
	} {
		{"happy path", "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", false},
		{"input err path","invalid address", false},
	}

	for _, scenario := range scenarios {
		result, resultErr := client.pendingNonceAt(scenario.inputAddress)
		assert.Equal(t, scenario.expectErr, resultErr != nil, scenario.desc)
		fmt.Println(result)
	}
}

// Test PendingNonceAt on stubs
type pendingNonceFetcherStub struct {
	pendingNonce uint64
	err error
}

func (m *pendingNonceFetcherStub) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
    return m.pendingNonce, m.err
}

func (m *pendingNonceFetcherStub) Close() {}

func Test_PendingNonceAt_OnStubs(t *testing.T) {
    scenarios := []struct{
    	desc string
    	inputAddress string
        nonceFetcher *pendingNonceFetcherStub
    	expectedNonce uint64
    	expectErr bool
	} {
		{
			"happy path",
			"0x63acf8339f376f0ff23da0b9265f0a7687ea72ea",
			&pendingNonceFetcherStub{100, nil},
			100,
			false},
		{
			"input err path",
			"",
			&pendingNonceFetcherStub{0, errors.New("not found")},
			0,
			true},
		{
			"system err path",
			"0x63acf8339f376f0ff23da0b9265f0a7687ea72ea",
			&pendingNonceFetcherStub{0, errors.New("something failed")},
			0,
			true},
	}

    for _, scenario := range scenarios {
    	c := &ethClient{"", scenario.nonceFetcher}
    	result, resultErr := c.pendingNonceAt(scenario.inputAddress)
		assert.Equal(t, scenario.expectErr, resultErr != nil, scenario.desc)
    	assert.Equal(t, scenario.expectedNonce, result, scenario.desc)
	}
}