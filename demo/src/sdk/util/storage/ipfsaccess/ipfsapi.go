// Scry Info.  All rights reserved.
// license that can be found in the license file.

package ipfsaccess

import (
	"errors"
	"github.com/ipfs/go-ipfs-api"
	rlog "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

// IPFSOutDir
var IPFSOutDir = "D:/desktop"
var ipaccessor *IpfsAccessor
var once sync.Once

// GetIAInstance
func GetIAInstance() *IpfsAccessor {
	once.Do(func() {
		ipaccessor = &IpfsAccessor{}
	})

	return ipaccessor
}

// IpfsAccessor
type IpfsAccessor struct {
	sh *shell.Shell
}

// Initialize
func (ia *IpfsAccessor) Initialize(nodeAddr string) error {
	if ia.sh == nil {
		ia.sh = shell.NewShell(nodeAddr)
		if ia.sh == nil {
			return errors.New("IPFS init failed. ")
		}
	}

	return nil
}

// SaveToIPFS
func (ia *IpfsAccessor) SaveToIPFS(content []byte) (string, error) {
	if ia.sh == nil {
		rlog.Error("ipfs api shell is nil")
		return "", errors.New("ipfs api shell is nil")
	}

	return ia.sh.Add(strings.NewReader(string(content)))
}

// GetFromIPFS
func (ia *IpfsAccessor) GetFromIPFS(hash string) error {
	if ia.sh == nil {
		return errors.New("Get from IPFS failed, IPFS-api shell is nil. ")
	}

	return ia.sh.Get(hash, IPFSOutDir)
}
