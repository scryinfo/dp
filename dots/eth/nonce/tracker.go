// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"context"
	"sync"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/db/gorms"
	"github.com/scryinfo/dp/dots/eth/client"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// const
const (
	TrackerTypeId = "15fa392f-042e-43f4-9df0-934bdb0fcf78"
)

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
// the caller can call ResetNonceAt to re-synchronize a nonce with the Ethereum network at the right time and thus to keep dao in consistency.
//
// Tracker is designed to be safe for concurrent use by multiple goroutines.
type Tracker interface {
	CreateAccount(address string) error
	DeleteAccount(address string) error

	RetrieveNonceAt(address string) (nonce uint64, err error)
	RestoreNonce(address string, recycledNonce uint64) error

	ResetNonceAt(address string) error
}

// TrackerImp implements Tracker interface, with component dependencies of *gorms.Gorms and *client.Client,
// which must be configured correctly by callers beforehand in order to get TrackerImp work normally as expected.
type TrackerImp struct {
	mu *sync.Mutex

	ClientDot *client.Client `dot:"bee6ae8e-e8a0-4693-ac65-416f618023d5"` // "github.com/scryinfo/dp/dots/eth/client"
	GormsDot  *gorms.Gorms   `dot:"d2b575cd-e38f-4002-b4bd-9dc85fe13fe6"` // "github.com/scryinfo/dot/dots/db/gorms"

	ethClient ethClient // depended to fetch and synchronize nonce from Ethereum network
	storer    *storer   // depended to storer and guard nonce off Ethereum network
}

// ethClient is an interface wrapper around *ethclient.Client.
type ethClient interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

// TrackerLive returns TypeLives for generating TrackerImp dot.
func TrackerLive() *dot.TypeLives {
	newTracker := func(conf []byte) (dot.Dot, error) {
		return &TrackerImp{mu: &sync.Mutex{}}, nil
	}

	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: TrackerTypeId, NewDoter: newTracker},
	}
}

//TrackerLiveAndDep returns TypeLives for generating TrackerImp and its dependent dots.
func TrackerLiveAndDep() []*dot.TypeLives {
	newTracker := func(conf []byte) (dot.Dot, error) {
		return &TrackerImp{mu: &sync.Mutex{}}, nil
	}

	return []*dot.TypeLives{
		&dot.TypeLives{
			Meta: dot.Metadata{TypeId: TrackerTypeId, NewDoter: newTracker},
		},
		client.ClientTypeLive(),
		gorms.TypeLives()[0],
	}
}

// Create injects ClientDot, GormsDot as well as unexported ethClient and storer,
// basing on communicating with the given dot line.
func (t *TrackerImp) Create(l dot.Line) error {
	logger := dot.Logger()

	// configure ClientDot
	clientDot, err := l.ToInjecter().GetByLiveId(client.ClientTypeId)
	if err != nil {
		logger.Errorln("TrackerImp dotGetByLiveId", zap.Error(err))
		return err
	}

	t.ClientDot = clientDot.(*client.Client)

	// configure GormsDot
	gormsDot, err := l.ToInjecter().GetByLiveId(gorms.TypeId)
	if err != nil {
		logger.Errorln("TrackerImp dotGetByLiveId", zap.Error(err))
		return err
	}
	t.GormsDot = gormsDot.(*gorms.Gorms)

	// create ethClient
	t.ethClient = t.ClientDot.EthClient()

	// create storer
	if err := createStorer(t); err != nil {
		logger.Errorln("TrackerImp initStorer", zap.Error(err))
		return err
	}
	return nil
}

// createStorer creates a storer for the given TrackerImp.
// Under the hook table nonces will be created if it doesn't exist yet.
var createStorer = func(t *TrackerImp) error {
	db := t.GormsDot.Db.DB()
	if err := db.Ping(); err != nil {
		return err
	}
	storer, err := getStorer(db)
	if err != nil {
		return err
	}
	t.storer = storer
	return nil
}

// CreateAccount creates an account to the given address if it does not exist yet
// with an initial/updated pending nonce, which is synchronized from the Ethereum network.
// CreateAccount will return ErrInvalidAddress when the given address is not a valid hex-encoded Ethereum address,
// and will return ErrAccountAlreadyExist if the account has already been created before.
func (t *TrackerImp) CreateAccount(address string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !common.IsHexAddress(address) {
		return ErrInvalidAddress
	}

	nonce, err := t.pendingNonceAt(address)
	if err != nil {
		return err
	}

	return t.storer.createAccountIfNotExists(address, nonce)
}

// DeleteAccount deletes the given account and its associated nonces from the database.
// It won't throw an error when the address is absent in storage.
func (t *TrackerImp) DeleteAccount(address string) error {
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
func (t *TrackerImp) RetrieveNonceAt(address string) (nonce uint64, err error) {
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
func (t *TrackerImp) RestoreNonce(address string, recycledNonce uint64) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.storer.restoreNonce(address, recycledNonce)
}

// ResetNonceAt refreshes the nonce to the given address in storage by synchronizing with the Ethereum network.
// Note that in case the address is absent in storage, ErrAccountNotExist will be returned.
func (t *TrackerImp) ResetNonceAt(address string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	nonce, err := t.pendingNonceAt(address)
	if err != nil {
		return err
	}

	return t.storer.resetNonceAt(address, nonce)
}

// pendingNonce returns nonce of the given account at pending state.
// This is the nonce that should be used for the next transaction.
// Note that in case of invalid address, 0 and nil will be returned.
func (t *TrackerImp) pendingNonceAt(address string) (uint64, error) {
	return t.ethClient.PendingNonceAt(context.Background(), common.HexToAddress(address))
}
