// Scry Info.  All rights reserved.
// license that can be found in the license file.

package server

import "github.com/scryinfo/dp/dots/app/server/definition"

// Server define what a type of connection should implement.
type Server interface {
	ListenAndServe() error

	// send message.
	SendMessage(name string, payload interface{}) error

	// preset function to handle message.
	PresetMsgHandleFuncs(name []string, presetFunc []definition.PresetFunc) error
}

// EventSendFailed common error extend msg
const EventSendFailed = " event send failed. "
