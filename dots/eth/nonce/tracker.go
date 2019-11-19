package nonce

import (
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

// myTracker is the singleton of interface Tracker.
var globalTrack Tracker

// Tracker stands for the interface to mainly retrieve and restore nonce with the database for transactions onto Ethereum.
//
// Callers must create an account with an address to make sure the Tracker can function well. When creating an account
// its nonce will be synchronized with the Ethereum network simultaneously and silently, and henceforth the nonce to an address
// will be guarded by the database. This means that callers don't have to request the Ethereum network for nonces each time they
// want to send transactions. 
//
// Normally callers can happily obtain a nonce via method RetrieveNonceAt and then the cource of corresponding transaction should go smoothly. 
//
// But in rare case that a retrieved nonce unexpectedly fails, the caller must return it back via method RestoreNonce for recycle, 
// which is vital for the Tracker to keep working fine. We say a nonce fails when it has been retrieved from the database 
// but eventually fails being sent into the TxPool of the Ethereum network.
//
// In a worst-case scenario the Tracker happens to generates inappropriate nonces and thus leads to failure of transactions even under normal use,
// the caller can call ResetNonceAt to re-synchronize a nonce with the Ethereum network at the right time and thus to keep data in consistency.
type Tracker interface {
	CreateAccount(address string) error
	DeleteAccount(address string) error

	RetrieveNonceAt(address string) (nonce uint64, err error)
	RestoreNonce(address string, recycledNonce uint64) error

	ResetNonceAt(address string) error

	Close() error
}

// SQLTracker is an implementation of interface Tracker interface through interacting with RDBS
type SQLTracker struct {
	mu             *sync.Mutex
	driverName     string
	dataSourceName string
	ethNodeURL     string
	storer         storer
	client         client
}

// GetTracker return a tracker with the given driverName, dataSourceName and ethNodeURL.
// The returned Traker is safe for concurrent use by multiple goroutines.
func GetTracker(driverName, dataSourceName, ethNodeURL string) Tracker {
	if globalTrack == nil {
		globalTrack = &SQLTracker{
			mu:             &sync.Mutex{},
			driverName:     driverName,
			dataSourceName: dataSourceName,
			ethNodeURL:     ethNodeURL,
			storer:         getStorer(driverName, dataSourceName),
			client:         getClient(ethNodeURL),
		}
	}
	return globalTrack
}

// CreateAccount creates an account to the given address if it does not exist yet
// with an initial/updated pending nonce, which is synchronized from the Ethereum network.
// CreateAccount will return ErrInvalidAddress when the given address is not a valid hex-encoded Ethereum address,
// and will return ErrAccountAlreadyExist if the account has already been created before.
func (t *SQLTracker) CreateAccount(address string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !common.IsHexAddress(address) {
		return ErrInvalidAddress
	}
	
	nonce, err := t.client.pendingNonceAt(address); 
	if err != nil {
		return err
	}
	
	return t.storer.createAccountIfNotExists(address, nonce)
}

// DeleteAccount deletes the given account and its associated nonces from the database.
// It won't throw an error when the address is absent in storage.
func (t *SQLTracker) DeleteAccount(address string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.storer.deleteAccount(address)
}

// RetrieveNonceAt returns the nonce that should be used for the next transaction.
//
// Indeed the returned nonce maybe a newly updated one or just a recycled one, which is fully reflected in storage.
// Usually callers needn't care about what status a returned nonce is, which are equivalent for a successful transaction.
// 
// Thus RetrieveNonceAt provides the pure value of nonce that can be used directly for the next transaction,
// and callers are free from being annoyed by redundant information on nonce status. 
// 
// Note that in case the address is absent, 0 and ErrNonceNotExist will be returned.
func (t *SQLTracker) RetrieveNonceAt(address string) (nonce uint64, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	return t.storer.retrieveNonceAt(address)
}

// RestoreNonce restores a nonce back into the database for recycle.
// Repeating restoring the same nonce make no sense because the storer won't accept duplicate records.
//
// In case the address is absent in storage, then ErrAccountNotExist will be returned.
// Otherwise if recycledNonce is less than 0 or bigger than the updatedNonce, ErrInvalidRecycledNonce will be returned.
// Besides ErrNonceAlreadyExist will be returned if the nonce has been recored in storage.
func (t *SQLTracker) RestoreNonce(address string, recycledNonce uint64) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.storer.restoreNonce(address, recycledNonce)
}

// ResetNonceAt refreshes the nonce to the given address in storage by synchronizing with the Ethereum network.
// Note that in case the address is absent in storage, ErrAccountNotExist will be returned.
func (t *SQLTracker) ResetNonceAt(address string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	nonce, err := t.client.pendingNonceAt(address); 
	if err != nil {
		return err
	}
	
	return t.storer.resetNonceAt(address, nonce)
}

// Close releases all underlining resources, namely the database and ethclient connections.
// It is rarely necessary to close a Tracker.
func (t *SQLTracker) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.client.close()
	return  t.storer.close()
}