package usermanager

import (
	"../security"
	"errors"
)

var (
	userMap = make(map[string]*User)
	curUser *User = nil
)

func Register(pubKey string, managerSelf bool, securityOperation security.SecurityOperation) (string, error) {
	_, exist := userMap[pubKey]
	if exist {
		return pubKey, errors.New("user already exist: " + pubKey)
	}

	if managerSelf && (securityOperation == nil) {
		return pubKey, errors.New("")
	} else if !managerSelf {
		securityOperation = &security.SecurityExecutor{}
	}

	user := NewUser(pubKey, managerSelf, securityOperation)
	userMap[pubKey] = user

	return pubKey, nil
}

func SetCurrentUser(publicKey string) (string, error) {
	user, exist := userMap[publicKey]
	if exist {
		curUser = user
		return publicKey, nil
	}

	return publicKey, errors.New("user does not exist: " + publicKey)
}

func GetCurrentUser() (*User, error) {
	if curUser == nil {
		return nil, errors.New("Current user is null ")
	}

	return curUser, nil
}

