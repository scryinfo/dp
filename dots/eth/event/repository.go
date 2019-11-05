// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import (
	"sync"
)

// Callback
type Callback func(event Event) bool

// Repository
type Repository struct {
	MapEventCallback sync.Map
}

// NewRepository
func NewRepository() *Repository {
	return &Repository{}
}
