// Scry Info.  All rights reserved.
// license that can be found in the license file.

package storage

type Storage interface {
    Init(nodeAddr string) error
    Save(value []byte) (string, error)
    Get(key string, outDir string) error
}
