// Scry Info.  All rights reserved.
// license that can be found in the license file.

package events

import (
	"encoding/json"
)

// JSONObj json obj
type JSONObj map[string]interface{}

// NewJSONObj create a new JSONObj
func NewJSONObj() JSONObj {
	return make(JSONObj)
}

// Set add k-v to map
func (obj JSONObj) Set(name string, val interface{}) {
	obj[name] = val
}

// Get value from map
func (obj JSONObj) Get(name string) interface{} {
	return obj[name]
}

func (obj JSONObj) String() string {
	data, _ := json.Marshal(obj)
	return string(data)
}

// Unmarshal unmarshal
func (obj JSONObj) Unmarshal(objPtr interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, objPtr)
}
