package mabi

import (
	"encoding/json"
)

type JSONObj map[string]interface{}

func NewJSONObj() JSONObj {
	return make(JSONObj)
}

func (obj JSONObj) Set(name string, val interface{}) {
	obj[name] = val
}

func (obj JSONObj) Get(name string) interface{} {
	return obj[name]
}

func (obj JSONObj) String() string {
	data, _ := json.Marshal(obj)
	return string(data)
}

func (obj JSONObj) Unmarshal(obj_ptr interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj_ptr)
}
