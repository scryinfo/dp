package ipfsaccess

import (
	"errors"
    rlog "github.com/sirupsen/logrus"
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
			errMsg := "failed to create new shell"
			return errors.New(errMsg)
		}
	}

	return nil
}

func (ia *IpfsAccessor) SaveToIPFS(content []byte) (string, error) {
	if ia.sh == nil {
		rlog.Error("ipfs api shell is nil")
		return "", errors.New("ipfs api shell is nil")
	}

	return ia.sh.Add(strings.NewReader(string(content)))
}