// Scry Info.  All rights reserved.
// license that can be found in the license file.

package connection

import "github.com/scryinfo/dp/dots/app/settings"

// define what a type of connection should implement.
type Connection interface {
	Connect() error

	// send message.
	SendMessage(name string, payload interface{}) error

	// preset function to handle message.
	AddCallbackFunc(name string, presetFunc settings.PresetFunc)
}
