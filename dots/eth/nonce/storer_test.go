// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Tests that rely on the real environments of PostgreSQL.
func Test_Storer_on_PostgreSQL(t *testing.T) {
    // Test main successful scenarios
	t.Run("main_successful_scenarios", func(t *testing.T) {
		require := require.New(t)

		// Connect the database
		driverName := "postgres"
		dataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		var storer storer
		require.NotPanicsf(func() {
			storer = getStorer(driverName, dataSourceName)
		}, "getStorer(%v, %v) shouldn't panic.\n", driverName, dataSourceName)

		// Create an account if it does not exist yet.
		address := "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"
		nonceIn := uint64(100)
		err := storer.createAccountIfNotExists(address, nonceIn)
		require.NoErrorf(err, "createAccountIfNotExists(%v, %v) shouldn't return err '%v'\n", address, nonceIn, err)

		// Retrieve nonce at an address
		nonceGot, err := storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(nonceIn, nonceGot, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, nonceIn, nonceGot)

		// Repeat retrieving nonce, expecting an incremented nonce.
		nonceGot2, err := storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(nonceGot, nonceGot2-1, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, nonceGot+1, nonceGot2)

		// Restore nonce
		err = storer.restoreNonce(address, nonceGot2)
		require.NoErrorf(err, "restoreNonce(%v, %v) shouldn't return err '%v'\n", address, nonceGot2, err)

		// Retrieve nonce again, expecting the nonce just restored.
		nonceGot, err = storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(nonceGot2, nonceGot, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, nonceGot2, nonceGot)

		// Reset nonce
		resetNonce := uint64(90)
		err = storer.resetNonceAt(address, resetNonce)
		require.NoErrorf(err, "resetNonceAt(%v, %v) shouldn't return err '%v'\n", address, resetNonce, err)

		// Retrieve nonce again, expecting the value just reset.
		nonceGot, err = storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(resetNonce, nonceGot, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, resetNonce, nonceGot)

		// Close the storer
		err = storer.close()
		require.NoErrorf(err, "storer shouldn't return err '%v'\n", err)
	})
}