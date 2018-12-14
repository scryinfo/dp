package usermanager

import (
	"../security"
)

type User struct {
	publicKey string
	selfSecurityMgr bool
	securityInterface security.SecurityOperationInterface
}

