package main

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
	"time"
)

// Constants
const (
	ABOUT = "test go -> js."
)

// Vars
var (
	AppName string
	w       *astilectron.Window
)

func main() {
	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDefaultPath: "resources/icon.ico",
		},
		Debug: true,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("Tools bar"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr("About"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", ABOUT, func(m *bootstrap.MessageIn) {
							// Unmarshal payload
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								astilog.Error(errors.Wrap(err, "unmarshal payload failed"))
								return
							}
							astilog.Debugf("About modal has been displayed and payload is %s!", s)
						}); err != nil {
							astilog.Error(errors.Wrap(err, "sending about event failed"))
						}
						return
					},
				},
				{
					Label: astilectron.PtrStr("Get accounts"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {

						if err := bootstrap.SendMessage(w, "get", accounts, func(m *bootstrap.MessageIn) {

						}); err != nil {
							astilog.Error(errors.Wrap(err, "sending get account event failed"))
						}
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				Center: astilectron.PtrBool(true),
				Height: astilectron.PtrInt(768),
				Width:  astilectron.PtrInt(1366),
				WebPreferences: &astilectron.WebPreferences{
					NodeIntegration: astilectron.PtrBool(true),
					WebSecurity:     astilectron.PtrBool(false),
				},
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
