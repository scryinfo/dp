// Scry Info.  All rights reserved.
// license that can be found in the license file.

package server

import "encoding/json"

// Server define what a type of connection should implement.
type Server interface {
	ListenAndServe() error

	// send message.
	SendMessage(name string, payload interface{}) error

	// preset function to handle message.
	PresetMsgHandleFuncs(name []string, presetFunc []PresetFunc) error
}

// MessageIn unified structure deserialize msg from client
type MessageIn struct {
	Name    string          `json:"Name"`
	Payload json.RawMessage `json:"Payload"`
}

// MessageOut unified structure serialize msg send to client
type MessageOut struct {
	Name    string      `json:"Name"`
	Payload interface{} `json:"Payload,omitempty"`
}

// PresetFunc preset system functions' handler
type PresetFunc = func(*MessageIn) (interface{}, error)

// EventSendFailed common error extend msg
const EventSendFailed = " event send failed. "
