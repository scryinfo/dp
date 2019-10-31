// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import (
	"sync"
)

type Callback func(event Event) bool

type Repository struct {
	MapEventCallback sync.Map
}

func NewRepository() *Repository {
	return &Repository{}
}
