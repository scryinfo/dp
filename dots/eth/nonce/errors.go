package nonce

import (
	"errors"
)

// Callers must create an valid Ethereum account via method CreateAccount
// before get its nonces "tracked" normally, otherwise errors as followings will be encountered.
var (
	// Method CreateAccount will return this error when the given address is not a valid hex-encoded Ethereum address.
	ErrInvalidAddress = errors.New("invalid Ethereum address")
	// Method CreateAccount will return this error when the given address does already exist in storage.
	ErrAccountAlreadyExist = errors.New("account does already exist")
	// Methods RestoreNonce and ResetNonceAt will return this error when an account does non exist in storage.
	ErrAccountNotExist = errors.New("account does not exist.")

	// Method RestoreNonce will return this error when the recycledNonce is less than 0 or bigger than the updatedNonce.
	ErrInvalidRecycledNonce = errors.New("invalid recycled nonce.")
	// Method RestoreNonce will return this error when the recycledNonce has already been recorded into storage.
	ErrNonceAlreadyExist = errors.New("nonce does already exist")
	// Method RetrieveNonceAt will return this error when no nonce to an account found in storage.
	ErrNonceNotExist = errors.New("nonce does not exist.")
)
