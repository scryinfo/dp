// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import "encoding/json"

// JSONObj
type JSONObj map[string]interface{}

// NewJSONObj
func NewJSONObj() JSONObj {
	return make(JSONObj)
}

// Set
func (obj JSONObj) Set(name string, val interface{}) {
	obj[name] = val
}

// Get
func (obj JSONObj) Get(name string) interface{} {
	return obj[name]
}

// String
func (obj JSONObj) String() string {
	data, _ := json.Marshal(obj)
	return string(data)
}

// Unmarshal
func (obj JSONObj) Unmarshal(objPtr interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, objPtr)
}
