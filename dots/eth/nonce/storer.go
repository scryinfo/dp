// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"database/sql"
)

// storer is logically responsible for storage of nonce
type storer struct {
	dao *dao
}

// getStorer returns a storer after creating table nonces if it hasn't exist yet.
func getStorer(db *sql.DB) (*storer, error) {
	dao := getDao(db)
	if err := dao.createTableNoncesIFNotExists(); err != nil {
		return nil, err
	}
	return &storer{dao}, nil
}

// Essentially createAccountIfNotExists populates a record into the tables nonces with initial/updated nonce to an address.
// ErrAccountAlreadyExist will be returned if the accounte has already been created before in storage.
func (s *storer) createAccountIfNotExists(address string, nonce uint64) error {
	count, err := s.dao.countAllNoncesAt(address)
	if err != nil {
		return err
	}

	if count > 0 {
		return ErrAccountAlreadyExist
	}

	return s.dao.insertNonce(address, nonce, true)
}

// Essentially deleteAccount removes all records belonging to the given address in the table nonces.
func (s *storer) deleteAccount(address string) error {
	return s.dao.deleteAllNonces(address)
}

// retrieveNonceAt returns a nonce to the specified address for the next transaction.
//
// Pre-condition: the exact address should be populated beforehand via method of createAccountIfNotExists.
// Post-condition: if there's no recycled nonce restored, the only updated nonce will be returned and increased by one,
// otherwise the minimum recycled nonce will be removed from the database and returned.
//
// Vialation: in case the address is absent, 0 and ErrNonceNotExist will be returned, meanwhile nothing will be changed in database.
func (s *storer) retrieveNonceAt(address string) (uint64, error) {
	count, err := s.dao.countAllNoncesAt(address)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, ErrNonceNotExist
	}

	if count == 1 {
		return s.dao.returnAndIncreaseTheOnlyUpdatedNonce(address)
	}
	return s.dao.deleteAndReturnMinRecycledNonce(address)
}

// restoreNonce inserts a record with a recycled nonce into the table nonces for the next transaction.
//
// Pre-condition: the recycled nonce must be smaller than the updated one in order to get recorded.
// Post-condition: a new record for the recycled nonce will be inserted into the table nonces.
//
// Violation: the recycled nonce will be ignored and restoreNonce won't affect the table nonces.
//
// In case the address is absent in storage, then ErrAccountNotExist will be returned.
// Otherwise if recycledNonce is less than 0 or bigger than the updatedNonce, ErrInvalidRecycledNonce will be returned.
// Besides ErrNonceAlreadyExist will be returned if the nonce has been recored in storage.
func (s *storer) restoreNonce(address string, recycledNonce uint64) error {
	ok, err := s.dao.nonceDoesExist(address, recycledNonce)
	if err != nil {
		return err
	}
	if ok {
		return ErrNonceAlreadyExist
	}

	updatedNonce, err := s.dao.theOnlyUpdatedNonceAt(address)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrAccountNotExist
		}
		return err
	}

	if recycledNonce > updatedNonce || recycledNonce < 0 {
		return ErrInvalidRecycledNonce
	}

	return s.dao.insertNonce(address, recycledNonce, false)
}

// resetNonceAt updates the only updated nonce record to an address with a given new value, and then deletes all the existing recycled records.
// Note that in case the address is absent, ErrAccountNotExist will be returned.
func (s *storer) resetNonceAt(address string, nonce uint64) error {
	_, err := s.dao.theOnlyUpdatedNonceAt(address)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrAccountNotExist
		}
		return err
	}
	return s.dao.resetNonceAt(address, nonce)
}
