package ipfsaccess

import (
	"fmt"
	"github.com/ipfs/go-ipfs-api"
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

func SaveToIPFS()  {

}