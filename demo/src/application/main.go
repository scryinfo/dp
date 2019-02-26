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
	ABOUT  = "test go -> js."
	ABOUT2 = "test <strong>go -> js.</strong><br/> once more"
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
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr("Main menu 1"),
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
								astilog.Infof("About modal has been displayed and payload is %s!", s)
							}); err != nil {
								astilog.Error(errors.Wrap(err, "sending about event failed"))
							}
							return
						},
					},
					{Role: astilectron.MenuItemRoleClose},
				},
			},
			{
				Label: astilectron.PtrStr("Main menu 2"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astilectron.PtrStr("About2"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "about2", ABOUT2, func(m *bootstrap.MessageIn) {
								// Unmarshal payload
								var s string
								if err := json.Unmarshal(m.Payload, &s); err != nil {
									astilog.Error(errors.Wrap(err, "unmarshal payload failed"))
									return
								}
								astilog.Infof("About2 modal has been displayed and payload is %s!", s)
							}); err != nil {
								astilog.Error(errors.Wrap(err, "sending about2 event failed"))
							}
							return
						},
					},
					{Role: astilectron.MenuItemRoleReload},
				},
			},
		},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(3 * time.Second)
				if err := bootstrap.SendMessage(w, "welcome", "Welcome to my go-astilectron demo!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending welcome event failed"))
				}
			}()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				Center: astilectron.PtrBool(true),
				Width:  astilectron.PtrInt(1000),
				Height: astilectron.PtrInt(700),
				WebPreferences: &astilectron.WebPreferences{
					NodeIntegration: astilectron.PtrBool(true),
					WebSecurity:     astilectron.PtrBool(true),
				},
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
