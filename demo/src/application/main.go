package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
	t "github.com/scryinfo/iscap/demo/src/application/transmission"
	"github.com/scryinfo/iscap/demo/src/sdk"
	"github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
	"time"
)

var (
	AppName    string
	w          *astilectron.Window
	scryInfoAS *settings.ScryInfoAS   = nil
	err error = nil
)

func init() {
	if err = sdk.InitLog(); err != nil {
		rlog.Error("",err)
	}
	if scryInfoAS, err = settings.LoadServicesSettings(); err != nil {
		rlog.Error("", err)
	}
	if err = accounts.GetAMInstance().Initialize(scryInfoAS.Services.Keystore); err != nil {
		rlog.Error("failed to initialize account service, error:", err)
	}
	if err = ipfsaccess.GetIAInstance().Initialize(scryInfoAS.Services.Ipfs); err != nil {
		rlog.Error("failed to initialize ipfs. error: ", err)
	}
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
				Label: astilectron.PtrStr("Administrator"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						// for test only, when reset chain anyhow, run this command to reset indexedDB.
						Label: astilectron.PtrStr("reset chain"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "resetChain", ""); err != nil {
								rlog.Error(errors.Wrap(err, "sending reset event failed"))
							}
							return
						},
					},
					{
						Label: astilectron.PtrStr("init data list"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "initDL", ""); err != nil {
								rlog.Error(errors.Wrap(err, "sending initDL event failed"))
							}
							return
						},
					},
					{
						Label: astilectron.PtrStr("init transaction"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "initTx", ""); err != nil {
								rlog.Error(errors.Wrap(err, "sending initMT event failed"))
							}
							return
						},
					},
				},
			},
		},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			t.SetWindow(w)
			go func() {
				time.Sleep(time.Second)
				if err := bootstrap.SendMessage(w, "welcome", "Welcome to my go-astilectron demo! "); err != nil {
					rlog.Error(errors.Wrap(err, "sending welcome event failed"))
				}
			}()
			return nil
		},
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: t.HandleMessages,
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
		rlog.Fatal(errors.Wrap(err, "Running bootstrap failed. "))
	}
}
