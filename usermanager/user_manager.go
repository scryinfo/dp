package usermanager

import "../security"


type UserManager struct {
	userMap map[string]User
	currentUser User
}

func (mgr *UserManager) Register(pubKey string,
	isSelfSecMgr bool, secInterface security.SecurityOperationInterface)  {

}

func SetCurrentUser(publicKey string)  {

}

