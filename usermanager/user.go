package usermanager

import (
	"../security"
	"errors"
)

type User struct {
	publicKey string
	managerSelf bool
	securityOperation security.SecurityOperation
}

func NewUser(pubkey string, mgrSelf bool, secOperation security.SecurityOperation) (*User)  {
	return  &User {
		publicKey: pubkey,
		managerSelf: mgrSelf,
		securityOperation: secOperation,
	}
}

func (user *User) GetSecurityOpertion() (security.SecurityOperation, error) {
	if user == nil {
		return nil, errors.New("Null user ")
	}

	return user.securityOperation, nil
}