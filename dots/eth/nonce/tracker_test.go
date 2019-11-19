// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Tests that rely on the real environments of mainet and PostgreSQL.
func Test_Tracker_on_ETH_mainnet_and_PostgreSQL(t *testing.T) {
	// Test main successful scenarios
	t.Run("main_successful_scenarios", func(t *testing.T) {
		// Inputs
		inDriverName := "postgres"
		inDataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		inNodeURL := "https://mainnet.infura.io/v3/d0a6c76d448d4f16baeb8143050a1f2f"
		address := "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"

		// Run tests
		assert := assert.New(t)
		// Get a tracker, connect to the database and the Ethereum network
		tracker := GetTracker(inDriverName, inDataSourceName, inNodeURL)
		t.Logf("GetTracker: dial %v and open %v is a success\n", inNodeURL, inDriverName)
		// Create an account into the database
		err := tracker.CreateAccount(address)
		if assert.NoErrorf(err, "create account %v shouldn't failed.", address) {
			t.Logf("CreateAccount: %v successfully.\n", address)
			// Retrieve nonce to an address
			nonce1, err := tracker.RetrieveNonceAt(address)
			if assert.NoErrorf(err, "retrieve nonce from account %v shouldn't failed.", address) {
				t.Logf("RetrieveNonceAt: nonce1 %v\n", nonce1)
				// Repeat retrieving nonce to an address
				nonce2, err := tracker.RetrieveNonceAt(address)
				if assert.NoErrorf(err, "retrieve nonce from account %v shouldn't failed.", address) {
					t.Logf("RetrieveNonceAt: nonce2 %v\n", nonce2)
					if assert.Equalf(nonce1+1, nonce2, "nonce1 = %v, but nonce2 = %v, nonce1-nonc2 != 1", nonce1, nonce2) {
						t.Logf("nonce1+1 %v == nonce2 %v\n", nonce1+1, nonce2)
					}
				}
				// Restore nonce to an address
				err = tracker.RestoreNonce(address, nonce2)
				if assert.NoErrorf(err, "restore nonce2 %v to address %v shouldn't fail", nonce2, address) {
					t.Logf("RestoreNonce: nonce2 %v\n", nonce2)
					// Again retrieving nonce to the same address
					nonce2_restored, err := tracker.RetrieveNonceAt(address)
					if assert.NoErrorf(err, "retrieve nonce from account %v shouldn't failed.", address) {
						t.Logf("RetrieveNonceAt: nonce2_restored %v\n", nonce2_restored)
						if assert.Equalf(nonce2, nonce2_restored, "nonce2 = %v, but nonce2_restored = %v, nonce2 != nonce2_restored", nonce2, nonce2_restored) {
							t.Logf("nonce2 %v == nonce2_restored %v\n", nonce2, nonce2_restored)
						}
					}
					// Rest nonce at an address
					err = tracker.ResetNonceAt(address)
					if assert.NoErrorf(err, "reset nonce to address %v shouldn't fail", address) {
						t.Logf("ResetNonceAt: address %s successfully. \n", address)
						// Delete an account
						err := tracker.DeleteAccount(address)
						if assert.NoErrorf(err, "delete account %v shouldn't fail", address) {
							t.Logf("DeleteAccount: %v deleted.\n", address)
							// Retrieve nonce from a non-existing account
							nonce_0, err := tracker.RetrieveNonceAt(address)
							if assert.EqualErrorf(err, ErrNonceNotExist.Error(), "There should be no nonce at address absence") {
								t.Logf("RetrieveNonceAt: nonce_0 = %v at address %v absence.\n", nonce_0, address)
								if assert.Equalf(uint64(0), nonce_0, "Expected zero nonce 0, but got actual nonce %v\n", nonce_0) {
									t.Logf("retrieve nonce of 0 from un-existing account %v\n", address)
								}
							}
						}
					}
				}
			}
			err = tracker.Close()
			if assert.NoErrorf(err, "There shouldn't return an error when close a track.") {
				t.Log("Close: Tracker got closed.")
			}
		}
	})

	// Test GetTracker on input errors scenarios
	t.Run("GetTracker_on_input_errors_scenarios", func(t *testing.T) {
		// Test scenarios
		scenarios := []struct{
			desc string
			inDriverName, inDataSourceName, inETHNodeURL string
		} {
			{
				desc: "unknow driver",
				inDriverName: "unknow_driver",},
			{
				desc: "db not exists",
				inDriverName: "postgres",
				inDataSourceName: "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=not_exists",},
			{
				desc: "invalid ETH node URL",
				inDriverName: "postgres",
				inDataSourceName: "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu",
				inETHNodeURL: "https://invalid",},
		}

		// Run tests
		assert := assert.New(t)
		for _, scenario := range scenarios {
			assert.Panicsf(func() {
				GetTracker(scenario.inDriverName, scenario.inDataSourceName, scenario.inETHNodeURL)
			}, "GetTracker: should panic!")
		}
	})

	// Test CreateAccount, DeleteAccount and Close on input errors scenarios
	t.Run("CreateAccount_and_DeleteAccount_on_input_errors_scenarios", func(t *testing.T) {
		// prerequisites
		driverName := "postgres"
		dataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		ethNodeURL := "https://mainnet.infura.io/v3/d0a6c76d448d4f16baeb8143050a1f2f"
		address := "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"
		tracker := GetTracker(driverName, dataSourceName, ethNodeURL)
		tracker.CreateAccount(address) // Create an account as existing beforehand.

		// Run tests
		assert := assert.New(t)
		defer func() {
			assert.NoError(tracker.Close(), "Close: should return no error.") // Release underlying resources
		}()

		defer func(address string) {
			err := tracker.DeleteAccount(address) // Clear up test trash
			assert.NoErrorf(err,"DeleteAccount: should return no error on an existing account %v\n", address)
			err = tracker.DeleteAccount("non_existing_address")
			assert.NoErrorf(err,"DeleteAccount: should return no error on an absent account %v\n", address)
		}(address)

		// Test scenarios for CreateAccount
		scenarios := []struct{
			desc string
			inAddress string
			expectErr bool
		}{
			{
				desc: "invalid Ethereum address",
				inAddress: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643_",
				expectErr: true,},
			{
				desc: "account does already exist",
				inAddress: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
				expectErr: true,},
		}

		// Run tests for CreateAccount
		for _, scenario := range scenarios {
			err := tracker.CreateAccount(scenario.inAddress)
			assert.Equalf(scenario.expectErr, err != nil, "expectEr: %v, but gotErr: %v\n", scenario.expectErr, err)
		}
	})

	// Test RetrieveNonceAt, RestoreNonce, ResetNonceAt on input errors scenarios
	t.Run("RetrieveNonceAt_RestoreNonce_ResetNonceAt_on_input_errors_scenarios", func(t *testing.T) {
		// prerequisites
		driverName := "postgres"
		dataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		ethNodeURL := "https://mainnet.infura.io/v3/d0a6c76d448d4f16baeb8143050a1f2f"
		tracker := GetTracker(driverName, dataSourceName, ethNodeURL)
		defer tracker.Close()

		assert := assert.New(t)

		// Test ErrNonceNottExist on RetrieveNonceAt
		address := "account_not_exists"
		_, err := tracker.RetrieveNonceAt(address)
		assert.Errorf(err, ErrNonceNotExist.Error(), "no nonce data should be retrieved from account %v\n", address)

		// Test ErrAccountNotExist on RestoreNonce and ResetNoneAt
		err = tracker.RestoreNonce(address, 10)
		assert.Errorf(err, ErrAccountNotExist.Error(), "no nonce data should be retrieved from account %v\n", address)

		err = tracker.ResetNonceAt(address)
		assert.Errorf(err, ErrAccountNotExist.Error(), "no nonce data should be retrieved from account %v\n", address)

		// Test ErrInvalidRecycledNonce on RestoreNonce
		address = "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"
		err = tracker.CreateAccount(address)
		if assert.NoErrorf(err, "CreateAccount: should return no error at address %v\n", address) {
			t.Logf("CreateAccount %v\n", address)
			defer func() {
				err = tracker.DeleteAccount(address)
				assert.NoErrorf(err, "DeleteAccount: should return no error at address %v\n", address)
			}()

			nonce, err := tracker.RetrieveNonceAt(address)
			if assert.NoErrorf(err, "RetrieveNonceAt: should return no error at address %v\n", address) {
				t.Logf("RetrieveNonceAt: address %v, returned nonce %v\n", address, nonce)
				invalidRecycledNonce := nonce + 2
				err = tracker.RestoreNonce(address, invalidRecycledNonce)
				if assert.Errorf(err, ErrInvalidRecycledNonce.Error(), "RestoreNonce(%v, %v) should returns an error.\n", address, invalidRecycledNonce) {
					t.Logf("RestoreNonce(%v, %v) returns error '%v' as expected.\n", address, invalidRecycledNonce, ErrInvalidRecycledNonce)
				}
				alreadyExistRecycledNonce := nonce + 1
				err = tracker.RestoreNonce(address, alreadyExistRecycledNonce)
				if assert.Errorf(err, ErrNonceAlreadyExist.Error(), "RestoreNonce(%v, %v) should returns an error.\n", address, alreadyExistRecycledNonce) {
					t.Logf("RestoreNonce(%v, %v) returns error '%v' as expected.\n", address, alreadyExistRecycledNonce, ErrNonceAlreadyExist)
				}
			}
		}
	})
}

// Test the Tracker for safety of concurrent use by multiple goroutines, relying on the real environments of mainet and PostgreSQL
func Test_Tracker_for_concurrent_safety_on_ETH_mainnet_and_PostgreSQL(t *testing.T) {
		// prerequisites
		driverName := "postgres"
		dataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		ethNodeURL := "https://mainnet.infura.io/v3/d0a6c76d448d4f16baeb8143050a1f2f"
		address := "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"
		tracker := GetTracker(driverName, dataSourceName, ethNodeURL)
		tracker.CreateAccount(address)

		assert := assert.New(t)
		var wg = &sync.WaitGroup{}

		// Test concurrent creating the same account for 10 times
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := tracker.CreateAccount(address)
				if assert.Errorf(err, ErrAccountAlreadyExist.Error(), "CreateAccount(%v) should return error '%v'\n", err) {
					t.Logf("CreateAccount(%v) failed expectedly for error '%v'\n", address, err)
				}
			}()
		}

		// Test concurrent retrieving and restoring nonce for 20 times.
		var totalNonces int = 20
		var failedNonces, succeededNonces  = []uint64{}, []uint64{}

		for i := 0; i < totalNonces; i++ {
			goID := "goroutine-" + strconv.Itoa(i)
			wg.Add(1)
			go func(goID string) {
				defer wg.Done()

				// Retrieve a nonce
				nonce, err := tracker.RetrieveNonceAt(address)
				if assert.NoErrorf(err, "RetrieveNonceAt(%v) shouldn't returns error '%v'.", address, err) {
					t.Logf("%v RetrieveNonceAt(%v): %v\n", goID, address, nonce)
					// Imitate the process of a transaction
					n := rand.Intn(1000)
					dur := time.Millisecond * time.Duration(n)
					time.Sleep(dur)
					t.Logf("%v takes %v microseconds trying to Tx %v", goID, dur, nonce)
					if n % 2 == 0 {
						t.Logf("%v Tx %v.", goID, nonce)
						succeededNonces = append(succeededNonces, nonce)
					} else {
						err := tracker.RestoreNonce(address, nonce)
						if assert.NoErrorf(err, "RestoreNonce(%v, %v) shouldn't returns error '%v'.", address, nonce, err) {
							t.Logf("%v RestoreNonce(%v, %v).", goID, address, nonce)
							failedNonces = append(failedNonces, nonce)
						}
					}
				}
			}(goID)
		}

		wg.Wait()
		t.Logf("total nonces: %v %v\n", totalNonces, len(succeededNonces))
		t.Logf("%v succeed nonces: %v\n", len(succeededNonces),  succeededNonces)
		t.Logf("%v failed nonces: %v\n", len(failedNonces),  failedNonces)
		for i := 0; i < len(failedNonces); i++ {
			failedNonce, _ := tracker.RetrieveNonceAt(address)
			t.Logf("retrieved failed nonce: %v\n", failedNonce)
		}
		for i := 0; i < 5; i++ {
			nextNonce, _ := tracker.RetrieveNonceAt(address)
			t.Logf("retrieved next nonce: %v\n", nextNonce)
		}

		// Test concurrent reset nonce to an account for 10 times
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := tracker.ResetNonceAt(address)
				if assert.NoErrorf(err,"ResetNonceAt(%v) shouldn't return error '%v'\n", err) {
					t.Logf("ResetNonceAt(%v) runs as expected\n", address)
				}
			}()
		}
		wg.Wait()

		// Test concurrent deleting an account for 10 times
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := tracker.DeleteAccount(address)
				if assert.NoErrorf(err,"DeleteAccount(%v) shouldn't return error '%v'\n", err) {
					t.Logf("DeleteAccount(%v) runs as expected\n", address)
				}
			}()
		}
		wg.Wait()

		// Test concurrent close the Tracker for 10 times
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := tracker.Close()
				if assert.NoErrorf(err,"Close shouldn't return error '%v'\n", err) {
					t.Log("Close runs as expected\n")
				}
			}()
		}
		wg.Wait()
}