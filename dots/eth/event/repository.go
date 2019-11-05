// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import (
	"sync"
)

// Callback contract event handler
type Callback func(event Event) bool

// Repository map match user,event and event handler
type Repository struct {
	MapEventCallback sync.Map
}

// NewRepository create a default Repository
func NewRepository() *Repository {
	return &Repository{}
}
