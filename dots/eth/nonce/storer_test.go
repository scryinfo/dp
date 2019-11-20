// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Tests that rely on the real environment of PostgreSQL.
func Test_Storer_on_PostgreSQL(t *testing.T) {
    // Test main successful scenarios
	t.Run("main_successful_scenarios", func(t *testing.T) {
		require := require.New(t)

		// Test connecting the database
		driverName := "postgres"
		dataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		var storer storer
		require.NotPanicsf(func() {
			storer = getStorer(driverName, dataSourceName)
		}, "getStorer(%v, %v) shouldn't panic.\n", driverName, dataSourceName)

		// Test creating an account if it does not exist yet.
		address := "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"
		nonceIn := uint64(100)
		err := storer.createAccountIfNotExists(address, nonceIn)
		require.NoErrorf(err, "createAccountIfNotExists(%v, %v) shouldn't return err '%v'\n", address, nonceIn, err)

		// Test retrieving nonce at an address
		nonceGot, err := storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(nonceIn, nonceGot, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, nonceIn, nonceGot)

		// Test repeating retrieving nonce, expecting an incremented nonce.
		nonceGot2, err := storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(nonceGot, nonceGot2-1, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, nonceGot+1, nonceGot2)

		// Test restoring nonce
		err = storer.restoreNonce(address, nonceGot2)
		require.NoErrorf(err, "restoreNonce(%v, %v) shouldn't return err '%v'\n", address, nonceGot2, err)

		// Test retrieving nonce again, expecting the nonce just restored.
		nonceGot, err = storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(nonceGot2, nonceGot, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, nonceGot2, nonceGot)

		// Test reseting nonce
		resetNonce := uint64(90)
		err = storer.resetNonceAt(address, resetNonce)
		require.NoErrorf(err, "resetNonceAt(%v, %v) shouldn't return err '%v'\n", address, resetNonce, err)

		// Test retrieving nonce again, expecting the value just reset.
		nonceGot, err = storer.retrieveNonceAt(address)
		require.NoErrorf(err, "retrieveNonceAt(%v) shouldn't return err '%v'\n", address, err)
		require.Equalf(resetNonce, nonceGot, "retrieveNonceAt(%v), expected %v, but got %v\n ", address, resetNonce, nonceGot)

		// Test closing the storer
		err = storer.close()
		require.NoErrorf(err, "storer shouldn't return err '%v'\n", err)
	})

	// Test exceptional situations

	// Test panics on getStorer
	t.Run("panics_on_getStorer", func(t *testing.T) {
		require := require.New(t)

		// Test connecting the database with unkown driver and inaccessible data source inputs.
		scenarios := []struct{
			desc string
			inDriveName string
			inDataSourceName string
		} {
			{
				desc: "unknown driver",
				inDriveName: "sqlite",
				inDataSourceName: ""},
			{
				desc: "inaccessible data source",
				inDriveName: "postgres",
				inDataSourceName: ""},

		}
		for _, scenario := range scenarios {
			require.Panicsf(func() {
				getStorer(scenario.inDriveName, scenario.inDataSourceName)
			}, "scenario '%v' failed. getStorer(%v, %v) should panic!", scenario.desc, scenario.inDriveName, scenario.inDataSourceName)
		}
	})

	// Test erroneous situations
	t.Run("erroneous_situations", func(t *testing.T) {
		require := require.New(t)

		// Connect the database as prerequisites for following tests.
		driverName := "postgres"
		dataSourceName := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=ljg_cqu"
		var storer storer
		require.NotPanicsf(func() {
			storer = getStorer(driverName, dataSourceName)
		}, "getStorer(%v, %v) shouldn't panic.\n", driverName, dataSourceName)

		// Test ErrAccountAlreadyExist on createAccountIfNotExists.
		address := "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"
		nonceInit := uint64(100)
		err := storer.deleteAccount(address)
		require.NoErrorf(err, "deleteAccount(%v) shouldn't return err '%v'\n", address, err)
		err = storer.createAccountIfNotExists(address, nonceInit)
		require.NoErrorf(err, "createAccountIfNotExists(%v, %v) shouldn't return err '%v'\n", address, nonceInit, err)
		err = storer.createAccountIfNotExists(address, nonceInit)
		require.EqualErrorf(err, ErrAccountAlreadyExist.Error(),"createAccountIfNotExists(%v, %v): expect err '%v' but got err '%v'\n", address, nonceInit, ErrAccountAlreadyExist, err)

		// Test ErrNonceNotExist on retrieveNonceAt
		address = "addres_that_not_exists"
		_, err = storer.retrieveNonceAt(address)
		require.EqualErrorf(err, ErrNonceNotExist.Error(),"retrieveNonceAt(%v): expect err '%v' but got err '%v'\n", address, ErrNonceNotExist, err)

		// Test ErrNonceAlreadyExist on restoreNonce
		scenarios := []struct{
			desc string
			inAddress string
			inRecycledNonce uint64
			expectedErr error
		} {
			{
				desc: "ErrNonceAlreadyExist",
				inAddress: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
				inRecycledNonce: nonceInit,
				expectedErr: ErrNonceAlreadyExist},
			{
				desc: "ErrAccountNotExist",
				inAddress: "addres_that_not_exists",
				inRecycledNonce: nonceInit,
				expectedErr: ErrAccountNotExist},
			{
				desc: "ErrInvalidRecycledNonce",
				inAddress: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
				inRecycledNonce: nonceInit+1,
				expectedErr: ErrInvalidRecycledNonce},
		}
		for _, scenario := range scenarios {
			err = storer.restoreNonce(scenario.inAddress, scenario.inRecycledNonce)
			require.EqualErrorf(err, scenario.expectedErr.Error(),"restoreNonce(%v, %v): expect err '%v' but got err '%v'\n", scenario.inAddress, scenario.inRecycledNonce, scenario.expectedErr, err)
		}

		// Test ErrAcountNotExist on resetNonceAt
		address = "addres_that_not_exists"
		nonceReset := uint64(100)
		err = storer.resetNonceAt(address, nonceReset)
		require.EqualErrorf(err, ErrAccountNotExist.Error(),"resetNonceAt(%v, %v): expect err '%v' but got err '%v'\n", address, nonceReset, ErrAccountNotExist, err)
	})
}