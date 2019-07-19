// Scry Info.  All rights reserved.
// license that can be found in the license file.

package server

import "encoding/json"

// define what a type of connection should implement.
type Server interface {
    ListenAndServe() error

    // send message.
    SendMessage(name string, payload interface{}) error

    // preset function to handle message.
    PresetMsgHandleFuncs(name []string, presetFunc []PresetFunc) error
}

type MessageIn struct {
    Name    string          `json:"Name"`
    Payload json.RawMessage `json:"Payload"`
}

type MessageOut struct {
    Name    string       `json:"Name"`
    Payload interface{} `json:"Payload,omitempty"`
}

type PresetFunc = func(*MessageIn) (interface{}, error)

const EventSendFailed = " event send failed. "
