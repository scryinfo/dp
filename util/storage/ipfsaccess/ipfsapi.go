package ipfsaccess

import (
	"errors"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"strings"
)

var (
	sh *shell.Shell = nil
)

func Initialize(nodeAddr string) bool {
	if sh == nil {
		sh = shell.NewShell(nodeAddr)
		if sh == nil {
			fmt.Println("Failed to create new shell.")
			return false
		}
	}

	return true
}

func SaveToIPFS(content []byte) (string, error) {
	if sh == nil {
		fmt.Println("ipfs api shell is nil")
		return "", errors.New("ipfs api shell is nil")
	}

	return sh.Add(strings.NewReader(string(content)))
}