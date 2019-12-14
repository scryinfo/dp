package preset

import (
	. "bou.ke/monkey"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	PreDef "github.com/scryinfo/dp/dots/app/business/definition"
	cec "github.com/scryinfo/dp/dots/app/business/preset/chain_event"
	mock_dot "github.com/scryinfo/dp/dots/app/business/preset/mock"
	"github.com/scryinfo/dp/dots/app/server"
	"github.com/scryinfo/dp/dots/app/storage"
	"github.com/scryinfo/dp/dots/auth"
	"github.com/scryinfo/dp/dots/binary"
	"github.com/scryinfo/dp/dots/binary/scry"
	mock_scry "github.com/scryinfo/dp/dots/binary/scry/mock"
	"github.com/scryinfo/dp/dots/eth/event"
	"github.com/scryinfo/dp/dots/eth/event/listen"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPreset_Create(t *testing.T) {
	Convey("test Preset.Create (dot.Create)", t, func() {
		Convey("standard input, expect success", func() {
			preIns := &Preset{}

			controller := gomock.NewController(t)
			mockLineObj := mock_dot.NewMockLine(controller)

			output := preIns.Create(mockLineObj)
			So(output, ShouldBeNil)
		})
	})
}

func TestNewPresetDot(t *testing.T) {
	Convey("test Preset config", t, func() {
		Convey("standard input, expect success", func() {
			confBs := []byte{123, 34, 109, 101, 116, 97, 68, 97, 116, 97, 79, 117, 116, 68, 105, 114, 34, 58, 32, 34, 67, 58, 47, 85, 115, 101, 114, 115, 47, 87, 105, 108, 108, 47, 68, 101, 115, 107, 116, 111, 112, 34, 125}

			output, err := newPresetDot(confBs)
			So(err, ShouldBeNil)

			outputAssert, ok := output.(*Preset)
			So(ok, ShouldBeTrue)
			So(outputAssert.config.MetaDataOutDir, ShouldEqual, "C:/Users/Will/Desktop")
		})

		Convey("conf is not []byte", func() {
			confS := `{"metaDataOutDir": "C:/Users/Will/Desktop"}`

			output, err := newPresetDot(confS)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, dot.NewError("dot_error_parameter", "the parameter error "))
		})

		Convey("unmarshal failed", func() {
			var confBs []byte

			output, err := newPresetDot(confBs)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, dot.NewError("dot_error_parameter", "the parameter error "))
		})
	})
}

func TestPreTypeLive(t *testing.T) {
	Convey("test Preset TypeLive", t, func() {
		Convey("standard input, expect success", func() {
			dotIns := &dot.TypeLives{
				Meta: dot.Metadata{
					TypeId: PreTypeId,
					NewDoter: func(conf interface{}) (dot.Dot, error) {
						return newPresetDot(conf)
					},
				},
			}

			Patch(binary.BinTypeLiveWithoutGrpc, func() []*dot.TypeLives {
				return []*dot.TypeLives{dotIns}
			})
			Patch(cec.CBsTypeLive, func() []*dot.TypeLives {
				return []*dot.TypeLives{dotIns}
			})
			defer UnpatchAll()

			output := PreTypeLive()
			So(output[1], ShouldEqual, dotIns)
			So(output[2], ShouldEqual, dotIns)
		})
	})
}

func TestPreset_LoginVerify(t *testing.T) {
	Convey("test Preset.LoginVerify", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.NewScryClient, func(s string, cw scry.ChainWrapper) scry.Client {
				return mockClientObj
			})
			mockClientObj.EXPECT().Authenticate("123456").Return(true, nil)

			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			defer UnpatchAll()

			output, err := preIns.LoginVerify(&server.MessageIn{
				Payload: []byte{123, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 48, 120, 49, 50, 99, 55, 56, 50, 54, 55, 52, 55, 102, 57, 50, 48, 98, 99, 52, 98, 98, 56, 55, 48, 102, 102, 50, 52, 102, 98, 101, 97, 48, 101, 102, 57, 97, 98, 52, 57, 52, 56, 34, 44, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("json unmarshal failed", func() {
			preIns := &Preset{}
			output, err := preIns.LoginVerify(&server.MessageIn{
				Payload: []byte{123, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 48, 120, 49, 50, 99, 55, 56, 50, 54, 55, 52, 55, 102, 57, 50, 48, 98, 99, 52, 98, 98, 56, 55, 48, 102, 102, 50, 52, 102, 98, 101, 97, 48, 101, 102, 57, 97, 98, 52, 57, 52, 56, 34, 44, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 49, 50, 51, 52, 53, 54, 125},
			})

			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("json: cannot unmarshal number into Go struct field Preset.password of type string"))
		})

		Convey("client is nil", func() {
			Patch(scry.NewScryClient, func(s string, cw scry.ChainWrapper) scry.Client {
				return nil
			})

			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			defer UnpatchAll()

			output, err := preIns.LoginVerify(&server.MessageIn{
				Payload: []byte{123, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 48, 120, 49, 50, 99, 55, 56, 50, 54, 55, 52, 55, 102, 57, 50, 48, 98, 99, 52, 98, 98, 56, 55, 48, 102, 102, 50, 52, 102, 98, 101, 97, 48, 101, 102, 57, 97, 98, 52, 57, 52, 56, 34, 44, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Call NewScryClient failed. "))
		})

		Convey("authenticate failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.NewScryClient, func(s string, cw scry.ChainWrapper) scry.Client {
				return mockClientObj
			})
			mockClientObj.EXPECT().Authenticate("123456").Return(false, errors.New("an error"))

			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			defer UnpatchAll()

			output, err := preIns.LoginVerify(&server.MessageIn{
				Payload: []byte{123, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 48, 120, 49, 50, 99, 55, 56, 50, 54, 55, 52, 55, 102, 57, 50, 48, 98, 99, 52, 98, 98, 56, 55, 48, 102, 102, 50, 52, 102, 98, 101, 97, 48, 101, 102, 57, 97, 98, 52, 57, 52, 56, 34, 44, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Authenticate user information failed. : an error"))
		})

		Convey("login verification failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.NewScryClient, func(s string, cw scry.ChainWrapper) scry.Client {
				return mockClientObj
			})
			mockClientObj.EXPECT().Authenticate("123456").Return(false, nil)

			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			defer UnpatchAll()

			output, err := preIns.LoginVerify(&server.MessageIn{
				Payload: []byte{123, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 48, 120, 49, 50, 99, 55, 56, 50, 54, 55, 52, 55, 102, 57, 50, 48, 98, 99, 52, 98, 98, 56, 55, 48, 102, 102, 50, 52, 102, 98, 101, 97, 48, 101, 102, 57, 97, 98, 52, 57, 52, 56, 34, 44, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Login verify failed. "))
		})
	})
}

func TestPreset_CreateNewAccount(t *testing.T) {
	Convey("test Preset.CreateNewAccount", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.CreateScryClient, func(s string, cw scry.ChainWrapper) (scry.Client, error) {
				return mockClientObj, nil
			})
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(3)
			Patch((*storage.SQLite).Insert, func(*storage.SQLite, interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{DB: &storage.SQLite{}}, Bin: &binary.Binary{}}
			output, err := preIns.CreateNewAccount(&server.MessageIn{
				Payload: []byte{123, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldEqual, "0x9693")
			So(err, ShouldBeNil)
		})

		Convey("json unmarshal failed", func() {
			preIns := &Preset{CBs: &cec.Callbacks{DB: &storage.SQLite{}}, Bin: &binary.Binary{}}
			output, err := preIns.CreateNewAccount(&server.MessageIn{
				Payload: []byte{123, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 49, 50, 51, 52, 53, 54, 125},
			})

			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("json: cannot unmarshal number into Go struct field Preset.password of type string"))
		})

		Convey("CreateScryClient failed", func() {
			Patch(scry.CreateScryClient, func(s string, cw scry.ChainWrapper) (scry.Client, error) {
				return nil, errors.New("an error")
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{DB: &storage.SQLite{}}, Bin: &binary.Binary{}}
			output, err := preIns.CreateNewAccount(&server.MessageIn{
				Payload: []byte{123, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Create new user failed. : an error"))
		})

		Convey("db insert unexpected", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.CreateScryClient, func(s string, cw scry.ChainWrapper) (scry.Client, error) {
				return mockClientObj, nil
			})
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(3)
			Patch((*storage.SQLite).Insert, func(*storage.SQLite, interface{}) (int64, error) {
				return 0, nil
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{DB: &storage.SQLite{}}, Bin: &binary.Binary{}}
			output, err := preIns.CreateNewAccount(&server.MessageIn{
				Payload: []byte{123, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 49, 50, 51, 52, 53, 54, 34, 125},
			})

			So(output, ShouldEqual, "0x9693")
			So(err, ShouldBeNil)
		})
	})
}

func onApprove(_ event.Event) bool {
	return true
}

func TestPreset_CurrentUserDataUpdate(t *testing.T) {
	Convey("test Preset.CurrentUserDataUpdate", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().SubscribeEvent("Approval", gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			Patch((*listen.Listener).SetFromBlock, func(*listen.Listener, uint64) {
				return
			})
			Patch((*Preset).testTransferEthAndTokens, func(*Preset) error {
				return nil
			})
			defer UnpatchAll()

			preIns := &Preset{
				Deployer: &PreDef.Preset{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{
					EventNames:   []string{"Approval"},
					EventHandler: []event.Callback{onApprove},
					CurUser:      mockClientObj,
				},
				Bin: &binary.Binary{Listener: &listen.Listener{}},
			}

			output, err := preIns.CurrentUserDataUpdate(nil)
			So(output, ShouldEqual, "")
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{CurUser: nil}}

			output, err := preIns.CurrentUserDataUpdate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("subscribe event failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().SubscribeEvent(gomock.Any(), gomock.Any()).Return(errors.New("an error"))

			preIns := &Preset{
				CBs: &cec.Callbacks{
					EventNames:   []string{"Approval"},
					EventHandler: []event.Callback{onApprove},
					CurUser:      mockClientObj,
				},
			}

			output, err := preIns.CurrentUserDataUpdate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Subscribe event failed. "))
		})

		Convey("db read unexpected", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().SubscribeEvent("Approval", gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 0, nil
			})
			defer UnpatchAll()

			preIns := &Preset{
				CBs: &cec.Callbacks{
					EventNames:   []string{"Approval"},
					EventHandler: []event.Callback{onApprove},
					CurUser:      mockClientObj,
				},
			}
			output, err := preIns.CurrentUserDataUpdate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("transfer test eth and token failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().SubscribeEvent("Approval", gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			Patch((*listen.Listener).SetFromBlock, func(*listen.Listener, uint64) {
				return
			})
			Patch((*Preset).testTransferEthAndTokens, func(*Preset) error {
				return errors.New("an error")
			})
			defer UnpatchAll()

			preIns := &Preset{
				Deployer: &PreDef.Preset{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{
					EventNames:   []string{"Approval"},
					EventHandler: []event.Callback{onApprove},
					CurUser:      mockClientObj,
				},
				Bin: &binary.Binary{Listener: &listen.Listener{}},
			}

			output, err := preIns.CurrentUserDataUpdate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_testTransferEthAndTokens(t *testing.T) {
	Convey("test Preset.testTransferEthAndTokens", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().TransferEthFrom(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Conn().Return(&ethclient.Client{})
			mockChainWrapperObj.EXPECT().TransferTokens(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"})

			preIns := &Preset{
				Deployer: &PreDef.Preset{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{CurUser: mockClientObj},
				Bin: &binary.Binary{},
			}
			output := preIns.testTransferEthAndTokens()
			So(output, ShouldBeNil)
		})

		Convey("transfer eth failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().TransferEthFrom(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Conn().Return(&ethclient.Client{})

			preIns := &Preset{
				Deployer: &PreDef.Preset{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{CurUser: mockClientObj},
			}
			output := preIns.testTransferEthAndTokens()
			So(output, ShouldBeError, errors.New("Transfer eth from Deployer failed. : an error"))
		})

		Convey("transfer token failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().TransferEthFrom(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Conn().Return(&ethclient.Client{})
			mockChainWrapperObj.EXPECT().TransferTokens(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"})

			preIns := &Preset{
				Deployer: &PreDef.Preset{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{CurUser: mockClientObj},
				Bin: &binary.Binary{},
			}
			output := preIns.testTransferEthAndTokens()
			So(output, ShouldBeError, errors.New("Transfer token from Deployer failed. : an error"))
		})
	})
}

func TestPreset_Logout(t *testing.T) {
	Convey("test Preset.Logout", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().UnSubscribeEvent("Approval").Return(nil)

			preIns := &Preset{CBs: &cec.Callbacks{CurUser: mockClientObj, EventNames: []string{"Approval"}}}
			output, err := preIns.Logout(nil)

			So(output, ShouldEqual, true) // interface{}.(bool) == bool
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{}}
			output, err := preIns.Logout(nil)

			var outputExpect interface{}
			So(output, ShouldEqual, outputExpect)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("try to unsubscribe known event", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().UnSubscribeEvent("Approval").Return(errors.New("a new error"))

			preIns := &Preset{CBs: &cec.Callbacks{CurUser: mockClientObj, EventNames: []string{"Approval"}}}
			output, err := preIns.Logout(nil)

			var outputExpect interface{}
			So(output, ShouldEqual, outputExpect) // interface{}.(bool) == bool
			So(err, ShouldBeError, errors.New("Unsubscribe failed, event:  Approval . : a new error"))
		})
	})
}

func TestPreset_Publish(t *testing.T) {
	Convey("test Preset.Publish", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("uuid", nil)

			preIns := &Preset{
				CBs: &cec.Callbacks{CurUser: mockClientObj},
				Bin: &binary.Binary{},
			}

			output, err := preIns.Publish(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,117,112,112,111,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,112,114,105,99,101,34,58,34,49,48,49,54,34,44,34,73,100,115,34,58,123,34,109,101,116,97,68,97,116,97,73,100,34,58,34,81,109,99,72,88,107,77,88,119,103,118,90,80,53,54,116,115,85,74,78,116,99,102,101,100,111,106,72,107,113,114,68,115,103,107,67,52,102,98,115,66,77,49,122,114,101,34,44,34,112,114,111,111,102,68,97,116,97,73,100,115,34,58,91,34,81,109,81,78,110,102,113,69,53,113,67,53,56,85,118,106,75,51,66,88,68,67,80,90,76,53,86,55,78,66,89,72,84,66,109,82,117,111,66,109,90,119,102,100,101,54,34,44,34,81,109,97,90,113,89,77,81,109,119,88,88,68,113,52,111,70,113,85,100,57,106,78,100,74,50,57,83,122,117,102,53,80,119,74,52,89,109,67,50,103,83,111,121,56,66,34,93,44,34,100,101,116,97,105,108,115,73,100,34,58,34,81,109,84,120,67,98,65,72,70,111,112,90,70,76,104,51,54,77,69,78,76,103,86,116,69,53,57,109,80,78,77,82,86,87,56,82,90,98,77,89,49,104,102,105,101,75,34,125,125}})
			So(output, ShouldEqual, "uuid")
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs:&cec.Callbacks{}}

			output, err := preIns.Publish(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("json unmarshal failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)

			preIns := &Preset{CBs: &cec.Callbacks{CurUser: mockClientObj}}

			output, err := preIns.Publish(&server.MessageIn{Payload:[]byte{34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,117,112,112,111,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,112,114,105,99,101,34,58,34,49,48,49,54,34,44,34,73,100,115,34,58,123,34,109,101,116,97,68,97,116,97,73,100,34,58,34,81,109,99,72,88,107,77,88,119,103,118,90,80,53,54,116,115,85,74,78,116,99,102,101,100,111,106,72,107,113,114,68,115,103,107,67,52,102,98,115,66,77,49,122,114,101,34,44,34,112,114,111,111,102,68,97,116,97,73,100,115,34,58,91,34,81,109,81,78,110,102,113,69,53,113,67,53,56,85,118,106,75,51,66,88,68,67,80,90,76,53,86,55,78,66,89,72,84,66,109,82,117,111,66,109,90,119,102,100,101,54,34,44,34,81,109,97,90,113,89,77,81,109,119,88,88,68,113,52,111,70,113,85,100,57,106,78,100,74,50,57,83,122,117,102,53,80,119,74,52,89,109,67,50,103,83,111,121,56,66,34,93,44,34,100,101,116,97,105,108,115,73,100,34,58,34,81,109,84,120,67,98,65,72,70,111,112,90,70,76,104,51,54,77,69,78,76,103,86,116,69,53,57,109,80,78,77,82,86,87,56,82,90,98,77,89,49,104,102,105,101,75,34,125}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("invalid character ':' after top-level value"))
		})

		Convey("atoi failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)

			preIns := &Preset{CBs: &cec.Callbacks{CurUser: mockClientObj}}

			output, err := preIns.Publish(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,117,112,112,111,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,112,114,105,99,101,34,58,34,34,44,34,73,100,115,34,58,123,34,109,101,116,97,68,97,116,97,73,100,34,58,34,81,109,99,72,88,107,77,88,119,103,118,90,80,53,54,116,115,85,74,78,116,99,102,101,100,111,106,72,107,113,114,68,115,103,107,67,52,102,98,115,66,77,49,122,114,101,34,44,34,112,114,111,111,102,68,97,116,97,73,100,115,34,58,91,34,81,109,81,78,110,102,113,69,53,113,67,53,56,85,118,106,75,51,66,88,68,67,80,90,76,53,86,55,78,66,89,72,84,66,109,82,117,111,66,109,90,119,102,100,101,54,34,44,34,81,109,97,90,113,89,77,81,109,119,88,88,68,113,52,111,70,113,85,100,57,106,78,100,74,50,57,83,122,117,102,53,80,119,74,52,89,109,67,50,103,83,111,121,56,66,34,93,44,34,100,101,116,97,105,108,115,73,100,34,58,34,81,109,84,120,67,98,65,72,70,111,112,90,70,76,104,51,54,77,69,78,76,103,86,116,69,53,57,109,80,78,77,82,86,87,56,82,90,98,77,89,49,104,102,105,101,75,34,125,125}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New(`strconv.Atoi: parsing "": invalid syntax`))
		})

		Convey("publish failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("uuid", errors.New("an error"))

			preIns := &Preset{
				CBs: &cec.Callbacks{CurUser: mockClientObj},
				Bin: &binary.Binary{},
			}

			output, err := preIns.Publish(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,117,112,112,111,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,112,114,105,99,101,34,58,34,49,48,49,54,34,44,34,73,100,115,34,58,123,34,109,101,116,97,68,97,116,97,73,100,34,58,34,81,109,99,72,88,107,77,88,119,103,118,90,80,53,54,116,115,85,74,78,116,99,102,101,100,111,106,72,107,113,114,68,115,103,107,67,52,102,98,115,66,77,49,122,114,101,34,44,34,112,114,111,111,102,68,97,116,97,73,100,115,34,58,91,34,81,109,81,78,110,102,113,69,53,113,67,53,56,85,118,106,75,51,66,88,68,67,80,90,76,53,86,55,78,66,89,72,84,66,109,82,117,111,66,109,90,119,102,100,101,54,34,44,34,81,109,97,90,113,89,77,81,109,119,88,88,68,113,52,111,70,113,85,100,57,106,78,100,74,50,57,83,122,117,102,53,80,119,74,52,89,109,67,50,103,83,111,121,56,66,34,93,44,34,100,101,116,97,105,108,115,73,100,34,58,34,81,109,84,120,67,98,65,72,70,111,112,90,70,76,104,51,54,77,69,78,76,103,86,116,69,53,57,109,80,78,77,82,86,87,56,82,90,98,77,89,49,104,102,105,101,75,34,125,125}})
			So(output, ShouldEqual, "uuid")
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_AdvancePurchase(t *testing.T) {
	Convey("test Preset.AdvancePurchase", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: ""}
			})
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockChainWrapperObj.EXPECT().AdvancePurchase(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"}).Times(2)

			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{
				CBs: &cec.Callbacks{
					CurUser:      mockClientObj,
					FlagChan: flagChan,
				},
				Bin: &binary.Binary{},
			}
			output, err := preIns.AdvancePurchase(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,116,97,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,80,117,98,108,105,115,104,73,100,34,58,34,49,53,55,53,57,52,52,52,54,49,53,52,57,48,57,56,54,48,48,45,56,50,54,49,53,53,49,51,50,52,52,53,50,57,51,55,49,55,53,34,44,34,112,114,105,99,101,34,58,34,49,48,50,48,34,125}})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{}}
			output, err := preIns.AdvancePurchase(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("json unmarshal failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)

			preIns := &Preset{
				CBs: &cec.Callbacks{CurUser:mockClientObj},
			}
			output, err := preIns.AdvancePurchase(&server.MessageIn{Payload:[]byte{34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,116,97,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,80,117,98,108,105,115,104,73,100,34,58,34,49,53,55,53,57,52,52,52,54,49,53,52,57,48,57,56,54,48,48,45,56,50,54,49,53,53,49,51,50,52,52,53,50,57,51,55,49,55,53,34,44,34,112,114,105,99,101,34,58,34,49,48,50,48,34}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("invalid character ':' after top-level value"))
		})

		Convey("change string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)

			preIns := &Preset{
				CBs: &cec.Callbacks{CurUser: mockClientObj},
			}
			output, err := preIns.AdvancePurchase(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,116,97,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,80,117,98,108,105,115,104,73,100,34,58,34,49,53,55,53,57,52,52,52,54,49,53,52,57,48,57,56,54,48,48,45,56,50,54,49,53,53,49,51,50,52,52,53,50,57,51,55,49,55,53,34,44,34,112,114,105,99,101,34,58,34,34,125}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("approve transfer failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: ""}
			})
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"})

			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{
				CBs: &cec.Callbacks{
					CurUser:      mockClientObj,
					FlagChan: flagChan,
				},
				Bin: &binary.Binary{},
			}
			output, err := preIns.AdvancePurchase(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,116,97,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,80,117,98,108,105,115,104,73,100,34,58,34,49,53,55,53,57,52,52,52,54,49,53,52,57,48,57,56,54,48,48,45,56,50,54,49,53,53,49,51,50,52,52,53,50,57,51,55,49,55,53,34,44,34,112,114,105,99,101,34,58,34,49,48,50,48,34,125}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Contract transfer token from buyer failed. : an error"))
		})

		Convey("advance purchase failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: ""}
			})
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockChainWrapperObj.EXPECT().AdvancePurchase(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr:"0x9693"}).Times(2)

			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{
				CBs: &cec.Callbacks{
					CurUser:      mockClientObj,
					FlagChan: flagChan,
				},
				Bin: &binary.Binary{},
			}
			output, err := preIns.AdvancePurchase(&server.MessageIn{Payload:[]byte{123,34,112,97,115,115,119,111,114,100,34,58,34,49,50,51,52,53,54,34,44,34,115,116,97,114,116,86,101,114,105,102,121,34,58,116,114,117,101,44,34,80,117,98,108,105,115,104,73,100,34,58,34,49,53,55,53,57,52,52,52,54,49,53,52,57,48,57,56,54,48,48,45,56,50,54,49,53,53,49,51,50,52,52,53,50,57,51,55,49,55,53,34,44,34,112,114,105,99,101,34,58,34,49,48,50,48,34,125}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Advance purchase failed. : an error"))
		})
	})
}

func TestPreset_ConfirmPurchase(t *testing.T) {
	Convey("test Preset.TestPreset_ConfirmPurchase", t, func() {
		Convey("standard input, expect success", func() {

		})
	})
}
