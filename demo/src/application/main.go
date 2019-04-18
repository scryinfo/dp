package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
	t "github.com/scryinfo/iscap/demo/src/application/transmission"
	"github.com/scryinfo/iscap/demo/src/sdk"
	rlog "github.com/sirupsen/logrus"
	"time"
)

const logpath = "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/iscap/demo/src/log/scry_sdk.log"

var (
	AppName  string
	w        *astilectron.Window
	scryInfo *settings.ScryInfo
	err      error = nil
)

func init() {
	scryInfo, err = settings.LoadSettings()
	err = sdk.Init(scryInfo.Chain.Ethereum.EthNode,
		scryInfo.Services.Keystore,
		scryInfo.Chain.Contracts.ProtocolAddr,
		scryInfo.Chain.Contracts.TokenAddr,
		scryInfo.Services.Ipfs,
		logpath,
		"App demo",
		)
	sdkinterface.SetScryInfo(scryInfo)
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
				var payload interface{}
				time.Sleep(time.Second)
				if err != nil {
					payload = errors.Wrap(err, "App init failed. ")
				} else {
					payload = "Welcome to my go-astilectron demo, and we will prepare accounts list for you. "
				}
				if err := bootstrap.SendMessage(w, "welcome", payload); err != nil {
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
				Width:  astilectron.PtrInt(1200),
				Height: astilectron.PtrInt(750),
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
