// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// globalDB is the singleton of type DB.
var globalDB *sql.DB

// once aims to prevent creating the table nonces twice.
var once = &sync.Once{}

// variables concerning interacting with database.
var (
	// Table nonces holds records of nonces to corresponding addresses, which are required and vital for successful transactions on the Ethereum network.
	// A nonce traverses two statuses in database, which are recycled and updated respectively. Once a nonce is retrieved from the database but fails to use outside,
	// then it must be restored and recycled for the next transaction. The recycled records can be inserted into/deleted from the database and its count is undeterministic,
	// because the externality such as network partition is out of control for database. However it's sufficient and must to have one updated record out there for each address.
	// Ideally when requested for a nonce, if there's no recycled nonce recorded currently, then database will directly provide the current only updated nonce and increase it's value by one.
	// We say a nonce fails at current context when it has been retrieved from the database but eventually fails being sent into the TxPool of the Ethereum network.
	// which should be removed permanently from the database and recycled for the next transaction prior to the only updated one.
	// Currently for the sake of simplicity, clarity and economy,  additional metadata which maybe important but not essential such as updatedAt time won't be recorded in the table nonces.
	// That is to say the table nonces just holds data that is absolutely indispensable for a transaction to succeed.
	createTableNoncesIFNotExistsSQL = func() string {
		return  `CREATE TABLE IF NOT EXISTS nonces (
		address varchar(42) NOT NULL CONSTRAINT address_length_of_42 CHECK (char_length(address) = 42), -- An address string should contain exactly 42 characters, according to the Ethereum network.
		nonce bigint NOT NULL CONSTRAINT natural_nonce CHECK (nonce >= 0), -- A nonce should be a natural number.
		updated boolean NOT NULL, -- The reverse status is recycled, here one boolean colum is perfect to distinguish two nonce statuses.
		CONSTRAINT unique_address_nonce_pair UNIQUE (address, nonce) -- Nonce record to an address shouldn't be duplicate.
	);`
	}
	insertNonceSQL = func(address string, nonce uint64, updated bool) string { return fmt.Sprintf("INSERT INTO nonces (address, nonce, updated) VALUES ('%v', %v, %v)", address, nonce, updated) }
	queryNonceSQL = func(address string, nonce uint64) string { return fmt.Sprintf("SELECT nonce FROM nonces WHERE address = '%v' AND nonce = %v", address, nonce) }
	queryAllNoncesCountSQL = func(address string) string { return fmt.Sprintf("SELECT count(*) FROM nonces WHERE address = '%v'", address) }
	queryTheOnlyUpdatedNonceSQL = func(address string) string { return fmt.Sprintf("SELECT nonce FROM nonces WHERE address = '%v' AND updated = TRUE", address) }
	//queryRecycledNoncesSQL = func(address string) string { return fmt.Sprintf("SELECT nonce FROM nonces WHERE address = '%v' AND updated = FALSE", address) }
	queryMinRecycledNonceSQL = func(address string) string { return fmt.Sprintf("SELECT nonce FROM nonces WHERE nonce = (SELECT min(nonce) FROM nonces WHERE address = '%v' AND updated = FALSE)", address)}
	//queryMinRecycledNonceSQL = func(address string) string { return "SELECT min(nonce) FROM " + " (" + queryRecycledNoncesSQL(address) + ")" }
	deleteAndReturnMinRecycledNonceSQL = func(address string) string { return fmt.Sprintf("DELETE FROM nonces WHERE address = '%v' AND nonce =  + (", address) + queryMinRecycledNonceSQL(address) + ")" + " RETURNING nonce" }
	deleteRecycledNoncesSQL = func(address string) string { return fmt.Sprintf("DELETE FROM nonces WHERE address = '%v' AND updated = FALSE", address) }
	deleteAllNoncesSQL = func(address string) string { return fmt.Sprintf("DELETE FROM nonces WHERE address = '%v'", address) }
	returnAndIncreaseUpdatedNonceSQL = func(address string) string { return fmt.Sprintf("UPDATE nonces SET nonce = nonce + 1 WHERE address = '%v' AND updated = TRUE RETURNING nonce-1", address) }
	updateTheOnlyUpdatedNonceSQL = func(nonce uint64, address string) string { return fmt.Sprintf("UPDATE nonces SET nonce = %v WHERE address = '%v' AND updated = TRUE", nonce, address) }
)

// openDB connects to the database, pings the connection and then creates the table nonces if neccessary.
// We are dead upon any error thrown.
// Currently PostgreSQL and MySQL are surpported underlyingly.
func openDB(driverName, dataSourceName string) *sql.DB {
	if globalDB == nil {
		db, err := open(driverName, dataSourceName)
		if err != nil {
			panic(err)
		}
		globalDB = db
	}

	if err := globalDB.Ping(); err != nil {
		panic(err)
	}

	once.Do(func() {
		if _, err := globalDB.Exec(createTableNoncesIFNotExistsSQL()); err != nil {
			panic(err)
		}
	})

	return globalDB
}
var open = sql.Open

func insertNonce(address string, nonce uint64, updated bool) error {
	_, err := globalDB.Exec(insertNonceSQL(address, nonce, updated))
	return err
}

func nonceDoesExist(address string, nonce uint64) (bool, error) {
	var result uint
	err := globalDB.QueryRow(queryNonceSQL(address, nonce)).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// resetNonceAt updates the only updated nonce record to an address with a given new value, and then deletes all the existing recycled records.
func resetNonceAt(address string, nonce uint64) error {
	tx, err := globalDB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(updateTheOnlyUpdatedNonceSQL(nonce, address))
	if err != nil {
		return tx.Rollback()
	}

	_, err = tx.Exec(deleteRecycledNoncesSQL(address))
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func countAllNoncesAt(address string) (uint, error)  {
	var result uint
	err := globalDB.QueryRow(queryAllNoncesCountSQL(address)).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return result, nil
}

func theOnlyUpdatedNonceAt(address string) (uint64, error) {
	var result uint64
	err := globalDB.QueryRow(queryTheOnlyUpdatedNonceSQL(address)).Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func deleteAndReturnMinRecycledNonce(address string) (uint64, error)  {
	var result uint64
	err := globalDB.QueryRow(deleteAndReturnMinRecycledNonceSQL(address)).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return result, nil
}

func deleteAllNonces(address string) error {
	_, err := globalDB.Exec(deleteAllNoncesSQL(address))
	return err
}

func returnAndIncreaseTheOnlyUpdatedNonce(address string) (uint64, error)  {
	var result uint64
	err := globalDB.QueryRow(returnAndIncreaseUpdatedNonceSQL(address)).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return result, nil
}