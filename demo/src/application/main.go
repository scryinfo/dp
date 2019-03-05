package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	"time"
)

// Constants
const (
	ABOUT  = "test go -> js."
	ShortMessage = "You have new short-message, remember to checkout it."
)

// Vars
var (
	AppName string
	w       *astilectron.Window
	err error = nil
)

func init() {
	err = sdkinterface.Initialize()
}

func main() {
	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDefaultPath: "resources/icon.ico",
		},
		Debug: true,
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr("Options"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Role: astilectron.MenuItemRoleReload},
					{Role: astilectron.MenuItemRoleClose},
				},
			},
			{
				Label:astilectron.PtrStr("Tests (go -> js)"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astilectron.PtrStr("test sdk init."),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err != nil {
								if err := bootstrap.SendMessage(w, "sdkInit", err.Error()); err != nil {
									astilog.Error(errors.Wrap(err, "sending welcome event failed"))
								}
							}
							return
						},
					},
					{
						Label: astilectron.PtrStr("test send short-message."),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "sendMessage", ShortMessage); err != nil {
								astilog.Error(errors.Wrap(err, "sending welcome event failed"))
							}
							return
						},
					},
				},
			},
		},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(time.Second)
				if err := bootstrap.SendMessage(w, "welcome", "Welcome to my go-astilectron demo!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending welcome event failed"))
				}
			}()
			return nil
		},
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
