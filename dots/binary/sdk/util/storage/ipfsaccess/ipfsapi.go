// Scry Info.  All rights reserved.
// license that can be found in the license file.

package ipfsaccess

import (
	"errors"
	"github.com/ipfs/go-ipfs-api"
	"strings"
	"sync"
)

var ipaccessor *IpfsAccessor = nil
var once sync.Once

func GetIAInstance() *IpfsAccessor {
	once.Do(func() {
		ipaccessor = &IpfsAccessor{}
	})

	return ipaccessor
}

type IpfsAccessor struct {
	sh *shell.Shell
}

func (ia *IpfsAccessor) Initialize(nodeAddr string) error {
	if ia.sh == nil {
		ia.sh = shell.NewShell(nodeAddr)
		if ia.sh == nil {
			return errors.New("IPFS init failed. ")
		}
	}

	return nil
}

func (ia *IpfsAccessor) SaveToIPFS(content []byte) (string, error) {
	if ia.sh == nil {
		return "", errors.New("ipfs api shell is nil")
	}

	return ia.sh.Add(strings.NewReader(string(content)))
}

func (ia *IpfsAccessor) GetFromIPFS(hash string, outDir string) error {
	if ia.sh == nil {
		return errors.New("Get from IPFS failed, IPFS-api shell is nil. ")
	}

	return ia.sh.Get(hash, outDir)
}
