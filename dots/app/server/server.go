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

// MessageIn
type MessageIn struct {
	Name    string          `json:"Name"`
	Payload json.RawMessage `json:"Payload"`
}

// MessageOut
type MessageOut struct {
	Name    string      `json:"Name"`
	Payload interface{} `json:"Payload,omitempty"`
}

// PresetFunc
type PresetFunc = func(*MessageIn) (interface{}, error)

// EventSendFailed
const EventSendFailed = " event send failed. "
