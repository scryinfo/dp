package preset

import (
	. "bou.ke/monkey"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	cec "github.com/scryinfo/dp/dots/app/business/preset/chain_event"
	"github.com/scryinfo/dp/dots/app/server"
	"github.com/scryinfo/dp/dots/app/server/definition"
	"github.com/scryinfo/dp/dots/app/storage"
	DBDef "github.com/scryinfo/dp/dots/app/storage/definition"
	"github.com/scryinfo/dp/dots/auth"
	"github.com/scryinfo/dp/dots/binary"
	"github.com/scryinfo/dp/dots/binary/scry"
	mock_scry "github.com/scryinfo/dp/dots/binary/scry/mock"
	"github.com/scryinfo/dp/dots/eth/event"
	"github.com/scryinfo/dp/dots/eth/event/listen"
	storage2 "github.com/scryinfo/dp/dots/storage/ipfs"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestPreset_Create(t *testing.T) {
	Convey("test MessageInPayload.Create (dot.Create)", t, func() {
		Convey("standard input, expect success", func() {
			preIns := &Preset{}

			output := preIns.Create(nil)
			So(output, ShouldBeNil)
		})
	})
}

func TestNewPresetDot(t *testing.T) {
	Convey("test MessageInPayload config", t, func() {
		Convey("standard input, expect success", func() {
			confBs := []byte{123, 34, 109, 101, 116, 97, 68, 97, 116, 97, 79, 117, 116, 68, 105, 114, 34, 58, 32, 34, 67, 58, 47, 85, 115, 101, 114, 115, 47, 87, 105, 108, 108, 47, 68, 101, 115, 107, 116, 111, 112, 34, 125}

			output, err := newPresetDot(confBs)
			So(err, ShouldBeNil)

			outputAssert, ok := output.(*Preset)
			So(ok, ShouldBeTrue)
			So(outputAssert.config.MetaDataOutDir, ShouldEqual, "C:/Users/Will/Desktop")
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
	Convey("test MessageInPayload TypeLive", t, func() {
		Convey("standard input, expect success", func() {
			dotIns := &dot.TypeLives{
				Meta: dot.Metadata{
					TypeId: PreTypeId,
					NewDoter: func(conf []byte) (dot.Dot, error) {
						return newPresetDot(conf)
					},
				},
			}

			Patch(binary.BinTypeLiveWithoutGrpc, func() []*dot.TypeLives {
				return []*dot.TypeLives{dotIns}
			})
			defer UnpatchAll()
			Patch(cec.CBsTypeLive, func() []*dot.TypeLives {
				return []*dot.TypeLives{dotIns}
			})

			output := PreTypeLive()
			So(output[1], ShouldEqual, dotIns)
			So(output[2], ShouldEqual, dotIns)
		})
	})
}

func TestPreset_LoginVerify(t *testing.T) {
	Convey("test MessageInPayload.LoginVerify", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.NewScryClient, func(s string, cw scry.ChainWrapper) scry.Client {
				return mockClientObj
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Authenticate("123456").Return(true, nil)

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}, Bin: &binary.Binary{}}
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})

			output, err := preIns.LoginVerify(&definition.MessageInPayload{
				Address:  "0x12c7826747f920bc4bb870ff24fbea0ef9ab4948",
				Password: "123456",
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("client is nil", func() {
			Patch(scry.NewScryClient, func(s string, cw scry.ChainWrapper) scry.Client {
				return nil
			})
			defer UnpatchAll()
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			output, err := preIns.LoginVerify(&definition.MessageInPayload{
				Address:  "0x12c7826747f920bc4bb870ff24fbea0ef9ab4948",
				Password: "123456",
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
			defer UnpatchAll()
			mockClientObj.EXPECT().Authenticate("123456").Return(false, errors.New("an error"))
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			output, err := preIns.LoginVerify(&definition.MessageInPayload{
				Address:  "0x12c7826747f920bc4bb870ff24fbea0ef9ab4948",
				Password: "123456",
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
			defer UnpatchAll()
			mockClientObj.EXPECT().Authenticate("123456").Return(false, nil)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			preIns := &Preset{CBs: &cec.Callbacks{}, Bin: &binary.Binary{}}
			output, err := preIns.LoginVerify(&definition.MessageInPayload{
				Address:  "0x12c7826747f920bc4bb870ff24fbea0ef9ab4948",
				Password: "123456",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Login verify failed. "))
		})
	})
}

func TestPreset_CreateNewAccount(t *testing.T) {
	Convey("test MessageInPayload.CreateNewAccount", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.CreateScryClient, func(s string, cw scry.ChainWrapper) (scry.Client, error) {
				return mockClientObj, nil
			})
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(3)
			Patch((*storage.SQLite).Insert, func(*storage.SQLite, interface{}) (int64, error) {
				return 1, nil
			})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{DB: &storage.SQLite{}}}, Bin: &binary.Binary{}}

			output, err := preIns.CreateNewAccount(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldEqual, "0x9693")
			So(err, ShouldBeNil)
		})

		Convey("CreateScryClient failed", func() {
			Patch(scry.CreateScryClient, func(s string, cw scry.ChainWrapper) (scry.Client, error) {
				return nil, errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{DB: &storage.SQLite{}}}, Bin: &binary.Binary{}}
			output, err := preIns.CreateNewAccount(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Create new user failed. : an error"))
		})

		Convey("db insert unexpected", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch(scry.CreateScryClient, func(s string, cw scry.ChainWrapper) (scry.Client, error) {
				return mockClientObj, nil
			})
			defer UnpatchAll()
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return nil
			})
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(3)
			Patch((*storage.SQLite).Insert, func(*storage.SQLite, interface{}) (int64, error) {
				return 0, nil
			})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{DB: &storage.SQLite{}}}, Bin: &binary.Binary{}}
			output, err := preIns.CreateNewAccount(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldEqual, "0x9693")
			So(err, ShouldBeNil)
		})
	})
}

func onApprove(_ event.Event) bool {
	return true
}

func TestPreset_CurrentUserDataUpdate(t *testing.T) {
	Convey("test MessageInPayload.CurrentUserDataUpdate", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().SubscribeEvent("Approval", gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			Patch((*listen.Listener).SetFromBlock, func(*listen.Listener, uint64) {
				return
			})
			Patch((*Preset).testTransferEthAndTokens, func(*Preset) error {
				return nil
			})

			preIns := &Preset{
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{
					EventNames:   []string{"Approval"},
					EventHandler: []event.Callback{onApprove},
					WS:           &server.WSServer{CurUser: mockClientObj},
				},
				Bin: &binary.Binary{Listener: &listen.Listener{}},
			}

			output, err := preIns.CurrentUserDataUpdate(nil)
			So(output, ShouldEqual, "")
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
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
					WS:           &server.WSServer{CurUser: mockClientObj},
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
					WS:           &server.WSServer{CurUser: mockClientObj},
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
			defer UnpatchAll()
			Patch((*listen.Listener).SetFromBlock, func(*listen.Listener, uint64) {
				return
			})
			Patch((*Preset).testTransferEthAndTokens, func(*Preset) error {
				return errors.New("an error")
			})
			preIns := &Preset{
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs: &cec.Callbacks{
					EventNames:   []string{"Approval"},
					EventHandler: []event.Callback{onApprove},
					WS:           &server.WSServer{CurUser: mockClientObj},
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
	Convey("test MessageInPayload.testTransferEthAndTokens", t, func() {
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
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			preIns := &Preset{
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin:      &binary.Binary{},
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
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
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
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin:      &binary.Binary{},
			}
			output := preIns.testTransferEthAndTokens()
			So(output, ShouldBeError, errors.New("Transfer token from Deployer failed. : an error"))
		})
	})
}

func TestPreset_Logout(t *testing.T) {
	Convey("test MessageInPayload.Logout", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().UnSubscribeEvent("Approval").Return(nil)

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}, EventNames: []string{"Approval"}}}

			output, err := preIns.Logout(nil)
			So(output, ShouldEqual, true) // interface{}.(bool) == bool
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.Logout(nil)
			var outputExpect interface{}
			So(output, ShouldEqual, outputExpect)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("try to unsubscribe known event", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().UnSubscribeEvent("Approval").Return(errors.New("an error"))
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}, EventNames: []string{"Approval"}}}
			output, err := preIns.Logout(nil)
			var outputExpect interface{}
			So(output, ShouldEqual, outputExpect) // interface{}.(bool) == bool
			So(err, ShouldBeError, errors.New("Unsubscribe failed, event:  Approval . : an error"))
		})
	})
}

func TestPreset_Publish(t *testing.T) {
	Convey("test MessageInPayload.Publish", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("uuid", nil)

			preIns := &Preset{
				CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin: &binary.Binary{},
			}

			output, err := preIns.Publish(&definition.MessageInPayload{
				Password:      "123456",
				Price:         "1016",
				SupportVerify: true,
				Ids: definition.Ids{
					MetaDataId:   "QmcHXkMXwgvZP56tsUJNtcfedojHkqrDsgkC4fbsBM1zre",
					ProofDataIds: []string{"QmQNnfqE5qC58UvjK3BXDCPZL5V7NBYHTBmRuoBmZwfde6", "QmaZqYMQmwXXDq4oFqUd9jNdJ29Szuf5PwJ4YmC2gSoy8B"},
					DetailsId:    "QmTxCbAHFopZFLh36MENLgVtE59mPNMRVW8RZbMY1hfieK",
				},
			})
			So(output, ShouldEqual, "uuid")
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.Publish(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("atoi failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.Publish(&definition.MessageInPayload{
				Password:      "123456",
				Price:         "",
				SupportVerify: true,
				Ids: definition.Ids{
					MetaDataId:   "QmcHXkMXwgvZP56tsUJNtcfedojHkqrDsgkC4fbsBM1zre",
					ProofDataIds: []string{"QmQNnfqE5qC58UvjK3BXDCPZL5V7NBYHTBmRuoBmZwfde6", "QmaZqYMQmwXXDq4oFqUd9jNdJ29Szuf5PwJ4YmC2gSoy8B"},
					DetailsId:    "QmTxCbAHFopZFLh36MENLgVtE59mPNMRVW8RZbMY1hfieK",
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New(`strconv.Atoi: parsing "": invalid syntax`))
		})

		Convey("publish failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("uuid", errors.New("an error"))
			preIns := &Preset{
				CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin: &binary.Binary{},
			}
			output, err := preIns.Publish(&definition.MessageInPayload{
				Password:      "123456",
				Price:         "1016",
				SupportVerify: true,
				Ids: definition.Ids{
					MetaDataId:   "QmcHXkMXwgvZP56tsUJNtcfedojHkqrDsgkC4fbsBM1zre",
					ProofDataIds: []string{"QmQNnfqE5qC58UvjK3BXDCPZL5V7NBYHTBmRuoBmZwfde6", "QmaZqYMQmwXXDq4oFqUd9jNdJ29Szuf5PwJ4YmC2gSoy8B"},
					DetailsId:    "QmTxCbAHFopZFLh36MENLgVtE59mPNMRVW8RZbMY1hfieK",
				},
			})
			So(output, ShouldEqual, "uuid")
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_AdvancePurchase(t *testing.T) {
	Convey("test MessageInPayload.AdvancePurchase", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: ""}
			})
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockChainWrapperObj.EXPECT().AdvancePurchase(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)

			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{
				CBs: &cec.Callbacks{
					WS:       &server.WSServer{CurUser: mockClientObj},
					FlagChan: flagChan,
				},
				Bin: &binary.Binary{},
			}

			output, err := preIns.AdvancePurchase(&definition.MessageInPayload{
				Password:    "123456",
				StartVerify: true,
				PublishId:   "1575944461549098600-8261551324452937175",
				Price:       "1020",
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.AdvancePurchase(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("change string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{
				CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
			}
			output, err := preIns.AdvancePurchase(&definition.MessageInPayload{
				Password:    "123456",
				StartVerify: true,
				PublishId:   "1575944461549098600-8261551324452937175",
				Price:       "",
			})
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
			defer UnpatchAll()
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: ""}
			})
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{
				CBs: &cec.Callbacks{
					WS:       &server.WSServer{CurUser: mockClientObj},
					FlagChan: flagChan,
				},
				Bin: &binary.Binary{},
			}

			output, err := preIns.AdvancePurchase(&definition.MessageInPayload{
				Password:    "123456",
				StartVerify: true,
				PublishId:   "1575944461549098600-8261551324452937175",
				Price:       "1020",
			})
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
			defer UnpatchAll()
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: ""}
			})
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockChainWrapperObj.EXPECT().AdvancePurchase(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{
				CBs: &cec.Callbacks{
					WS:       &server.WSServer{CurUser: mockClientObj},
					FlagChan: flagChan,
				},
				Bin: &binary.Binary{},
			}
			output, err := preIns.AdvancePurchase(&definition.MessageInPayload{
				Password:    "123456",
				StartVerify: true,
				PublishId:   "1575944461549098600-8261551324452937175",
				Price:       "1020",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Advance purchase failed. : an error"))
		})
	})
}

func TestPreset_ConfirmPurchase(t *testing.T) {
	Convey("test MessageInPayload.TestPreset_ConfirmPurchase", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ConfirmPurchase(gomock.Any(), gomock.Any()).Return(nil)

			preIns := &Preset{
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}

			output, err := preIns.ConfirmPurchase(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.ConfirmPurchase(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.ConfirmPurchase(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("confirm purchase failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ConfirmPurchase(gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}
			output, err := preIns.ConfirmPurchase(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Confirm purchase failed. "))
		})
	})
}

func TestPreset_ReEncrypt(t *testing.T) {
	Convey("test MessageInPayload.ReEncrypt", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			mockChainWrapperObj.EXPECT().ReEncrypt(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			preIns := &Preset{
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}

			output, err := preIns.ReEncrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.ReEncrypt(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.ReEncrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 0, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.ReEncrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("re-encrypt failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			mockChainWrapperObj.EXPECT().ReEncrypt(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}
			output, err := preIns.ReEncrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Re-encrypt failed. "))
		})
	})
}

func TestPreset_CancelPurchase(t *testing.T) {
	Convey("test MessageInPayload.CancelPurchase", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().CancelPurchase(gomock.Any(), gomock.Any()).Return(nil)

			preIns := &Preset{
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}

			output, err := preIns.CancelPurchase(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.CancelPurchase(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.CancelPurchase(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("cancel purchase failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().CancelPurchase(gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{
				CBs:      &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}
			output, err := preIns.CancelPurchase(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Cancel purchase failed. : an error"))
		})
	})
}

func TestPreset_Decrypt(t *testing.T) {
	Convey("test MessageInPayload.Decrypt", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			Patch((*Preset).getMetaDataFileName, func(*Preset, *DBDef.Transaction, string) (string, error) {
				return "simulate meta data file absolute path", nil
			})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.Decrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldEqual, "simulate meta data file absolute path")
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.Decrypt(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 0, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.Decrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("get meta data file name failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			Patch((*Preset).getMetaDataFileName, func(*Preset, *DBDef.Transaction, string) (string, error) {
				return "", errors.New("an error")
			})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.Decrypt(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
			})
			So(output, ShouldEqual, "")
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_getMetaDataFileName(t *testing.T) {
	Convey("test MessageInPayload.getMetaDataFileName", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*auth.Account).Decrypt, func(*auth.Account, []byte, string, string) ([]byte, error) {
				return []byte{'m', 'e', 't', 'a', 'd', 'a', 't', 'a'}, nil
			})
			defer UnpatchAll()
			Patch((*storage2.Ipfs).Get, func(*storage2.Ipfs, string, string) error {
				return nil
			})
			Patch(os.Rename, func(string, string) error {
				return nil
			})

			preIns := &Preset{
				CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin: &binary.Binary{
					Storage: &storage2.Ipfs{},
					Account: &auth.Account{},
				},
				config: presetConfig{MetaDataOutDir: "D:/dp/download"},
			}

			output, err := preIns.getMetaDataFileName(&DBDef.Transaction{
				Buyer:                       "0x9693",
				MetaDataIdEncWithBuyer:      "simulate ipfs id with buyer",
				MetaDataIdEncWithArbitrator: "simulate ipfs id with arbitrator",
				MetaDataExtension:           ".txt",
			}, "123456")
			So(output, ShouldEqual, "D:/dp/download/metadata.txt")
			So(err, ShouldBeNil)
		})

		Convey("account decrypt failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*auth.Account).Decrypt, func(*auth.Account, []byte, string, string) ([]byte, error) {
				return []byte{'m', 'e', 't', 'a', 'd', 'a', 't', 'a'}, errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{
				CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin: &binary.Binary{Account: &auth.Account{}},
			}
			output, err := preIns.getMetaDataFileName(&DBDef.Transaction{
				Buyer:                       "0x9822",
				MetaDataIdEncWithBuyer:      "simulate ipfs id with buyer",
				MetaDataIdEncWithArbitrator: "simulate ipfs id with arbitrator",
				MetaDataExtension:           ".txt",
			}, "123456")
			So(output, ShouldEqual, "")
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Decrypt encrypted meta data Id failed. "))
		})

		Convey("ipfs get failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*auth.Account).Decrypt, func(*auth.Account, []byte, string, string) ([]byte, error) {
				return []byte{'m', 'e', 't', 'a', 'd', 'a', 't', 'a'}, nil
			})
			defer UnpatchAll()
			Patch((*storage2.Ipfs).Get, func(*storage2.Ipfs, string, string) error {
				return errors.New("an error")
			})
			preIns := &Preset{
				CBs:    &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin:    &binary.Binary{Storage: &storage2.Ipfs{}, Account: &auth.Account{}},
				config: presetConfig{MetaDataOutDir: "D:/dp/download"},
			}
			output, err := preIns.getMetaDataFileName(&DBDef.Transaction{
				Buyer:                       "0x9693",
				MetaDataIdEncWithBuyer:      "simulate ipfs id with buyer",
				MetaDataIdEncWithArbitrator: "simulate ipfs id with arbitrator",
				MetaDataExtension:           ".txt",
			}, "123456")
			So(output, ShouldEqual, "")
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Get meta data from IPFS failed. "))
		})

		Convey("rename meta data file failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*auth.Account).Decrypt, func(*auth.Account, []byte, string, string) ([]byte, error) {
				return []byte{'m', 'e', 't', 'a', 'd', 'a', 't', 'a'}, nil
			})
			defer UnpatchAll()
			Patch((*storage2.Ipfs).Get, func(*storage2.Ipfs, string, string) error {
				return nil
			})
			Patch(os.Rename, func(string, string) error {
				return errors.New("an error")
			})
			preIns := &Preset{
				CBs:    &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}},
				Bin:    &binary.Binary{Storage: &storage2.Ipfs{}, Account: &auth.Account{}},
				config: presetConfig{MetaDataOutDir: "D:/dp/download"},
			}
			output, err := preIns.getMetaDataFileName(&DBDef.Transaction{
				Buyer:                       "0x9693",
				MetaDataIdEncWithBuyer:      "simulate ipfs id with buyer",
				MetaDataIdEncWithArbitrator: "simulate ipfs id with arbitrator",
				MetaDataExtension:           ".txt",
			}, "123456")
			So(output, ShouldEqual, "")
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Add extension to meta data failed. "))
		})
	})
}

func TestPreset_ConfirmData(t *testing.T) {
	Convey("test MessageInPayload.ConfirmData", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ConfirmData(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}

			output, err := preIns.ConfirmData(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Confirm:       definition.Confirm{Truth: true},
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.ConfirmData(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.ConfirmData(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
				Confirm:       definition.Confirm{Truth: true},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("confirm data failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ConfirmData(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.ConfirmData(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Confirm:       definition.Confirm{Truth: true},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Confirm data failed. "))
		})
	})
}

func TestPreset_Register(t *testing.T) {
	Convey("test MessageInPayload.Register", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: "simulate protocol contract address"}
			})
			mockChainWrapperObj.EXPECT().RegisterAsVerifier(gomock.Any()).Return(nil)

			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}, FlagChan: flagChan}, Bin: &binary.Binary{}}

			output, err := preIns.Register(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.Register(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("approve failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: "simulate protocol contract address"}
			})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.Register(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Contract transfer token from register failed. "))
		})

		Convey("register as verifier failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ApproveTransfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*binary.Binary).Config, func(*binary.Binary) binary.BinaryConfig {
				return binary.BinaryConfig{ProtocolContractAddr: "simulate protocol contract address"}
			})
			mockChainWrapperObj.EXPECT().RegisterAsVerifier(gomock.Any()).Return(errors.New("an error"))
			flagChan := make(chan bool, 10)
			flagChan <- true
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}, FlagChan: flagChan}, Bin: &binary.Binary{}}
			output, err := preIns.Register(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Register as verifier failed. "))
		})
	})
}

func TestPreset_Vote(t *testing.T) {
	Convey("test MessageInPayload.Vote", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Vote(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}

			output, err := preIns.Vote(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Verify: definition.Verify{
					Suggestion: true,
					Comment:    "commit",
				},
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.Vote(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.Vote(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
				Verify: definition.Verify{
					Suggestion: true,
					Comment:    "commit",
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("vote failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Vote(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.Vote(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Verify: definition.Verify{
					Suggestion: true,
					Comment:    "commit",
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Vote failed. "))
		})
	})
}

func TestPreset_GradeToVerifier(t *testing.T) {
	Convey("test MessageInPayload.GradeToVerifier", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().GradeToVerifier(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(2)

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}

			output, err := preIns.GradeToVerifier(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Grade: definition.Grade{
					Verifier1Revert: true,
					Verifier1Grade:  5,
					Verifier2Revert: true,
					Verifier2Grade:  5,
				},
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GradeToVerifier(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.GradeToVerifier(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
				Grade: definition.Grade{
					Verifier1Revert: true,
					Verifier1Grade:  5,
					Verifier2Revert: true,
					Verifier2Grade:  5,
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("grade to verifier1 failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().GradeToVerifier(gomock.Any(), gomock.Any(), uint8(0), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.GradeToVerifier(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Grade: definition.Grade{
					Verifier1Revert: true,
					Verifier1Grade:  5,
					Verifier2Revert: true,
					Verifier2Grade:  5,
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Grade verifier1 failed. "))
		})

		Convey("grade to verifier2 failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().GradeToVerifier(gomock.Any(), gomock.Any(), uint8(0), gomock.Any()).Return(nil)
			mockChainWrapperObj.EXPECT().GradeToVerifier(gomock.Any(), gomock.Any(), uint8(1), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.GradeToVerifier(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Grade: definition.Grade{
					Verifier1Revert: true,
					Verifier1Grade:  5,
					Verifier2Revert: true,
					Verifier2Grade:  5,
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Grade verifier2 failed. "))
		})
	})
}

func TestPreset_Arbitrate(t *testing.T) {
	Convey("test MessageInPayload.Arbitrate", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*Preset).updateAccInfo, func(*Preset, string) error {
				return nil
			})
			defer UnpatchAll()
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			mockChainWrapperObj.EXPECT().Arbitrate(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}

			output, err := preIns.Arbitrate(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Arbitrate:     definition.Arbitrate{ArbitrateResult: true},
			})
			So(output, ShouldEqual, "0")
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.Arbitrate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("set string to *big.Int failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.Arbitrate(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "",
				Arbitrate:     definition.Arbitrate{ArbitrateResult: true},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Set to *big.Int failed. "))
		})

		Convey("update acc info failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*Preset).updateAccInfo, func(*Preset, string) error {
				return errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}}
			output, err := preIns.Arbitrate(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Arbitrate:     definition.Arbitrate{ArbitrateResult: true},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})

		Convey("arbitrate failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*Preset).updateAccInfo, func(*Preset, string) error {
				return nil
			})
			defer UnpatchAll()
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			mockChainWrapperObj.EXPECT().Arbitrate(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.Arbitrate(&definition.MessageInPayload{
				Password:      "123456",
				TransactionId: "0",
				Arbitrate:     definition.Arbitrate{ArbitrateResult: true},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Arbitrate failed. "))
		})
	})
}

func TestPreset_updateAccInfo(t *testing.T) {
	Convey("test MessageInPayload.updateAccInfo", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			Patch(cec.UpdateSlice, func([]byte, string, string) ([]byte, error) {
				return nil, nil
			})
			Patch((*storage.SQLite).Update, func(*storage.SQLite, interface{}, map[string]interface{}, string, ...interface{}) (int64, error) {
				return 1, nil
			})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			err := preIns.updateAccInfo("0")
			So(err, ShouldBeNil)
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 0, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			err := preIns.updateAccInfo("0")
			So(err, ShouldBeNil)
		})

		Convey("update slice failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			Patch(cec.UpdateSlice, func([]byte, string, string) ([]byte, error) {
				return nil, errors.New("an error")
			})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			err := preIns.updateAccInfo("0")
			So(err, ShouldBeError, errors.New("an error"))
		})

		Convey("db update failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			Patch(cec.UpdateSlice, func([]byte, string, string) ([]byte, error) {
				return nil, nil
			})
			Patch((*storage.SQLite).Update, func(*storage.SQLite, interface{}, map[string]interface{}, string, ...interface{}) (int64, error) {
				return 0, nil
			})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			err := preIns.updateAccInfo("0")
			So(err, ShouldBeNil)
		})
	})
}

func TestPreset_GetEthBalance(t *testing.T) {
	Convey("test MessageInPayload.GetEthBalance", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockClientObj.EXPECT().GetEth(gomock.Any(), gomock.Any()).Return(big.NewInt(100), nil)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Conn().Return(nil)

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}

			output, err := preIns.GetEthBalance(nil)
			So(output, ShouldBeLessThanOrEqualTo, "100|"+time.Now().String())
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetEthBalance(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("get eth balance failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			mockClientObj.EXPECT().GetEth(gomock.Any(), gomock.Any()).Return(big.NewInt(100), errors.New("an error"))
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().Conn().Return(nil)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.GetEthBalance(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Get eth balance failed. "))
		})
	})
}

func TestPreset_GetTokenBalance(t *testing.T) {
	Convey("test MessageInPayload.GetTokenBalance", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().GetTokenBalance(gomock.Any(), gomock.Any()).Return(big.NewInt(100), nil)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}

			output, err := preIns.GetTokenBalance(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldBeLessThanOrEqualTo, "100|"+time.Now().String())
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetTokenBalance(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("get token balance failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().GetTokenBalance(gomock.Any(), gomock.Any()).Return(big.NewInt(100), errors.New("an error"))
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"}).Times(2)
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj}}, Bin: &binary.Binary{}}
			output, err := preIns.GetTokenBalance(&definition.MessageInPayload{Password: "123456"})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Get token balance failed. "))
		})
	})
}

func TestPreset_IsVerifier(t *testing.T) {
	Convey("test MessageInPayload.IsVerifier", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.IsVerifier(nil)
			So(output, ShouldBeFalse)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.IsVerifier(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 0, nil
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.IsVerifier(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})
	})
}

func TestPreset_GetAccountsList(t *testing.T) {
	Convey("test MessageInPayload.GetAccountsList", t, func() {
		Convey("standard input, expect success", func() {
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{DB: &storage.SQLite{}}}}

			output, err := preIns.GetAccountsList(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("db read failed", func() {
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{DB: &storage.SQLite{}}}}
			output, err := preIns.GetAccountsList(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_GetDataList(t *testing.T) {
	Convey("test MessageInPayload.GetDataList", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.GetDataList(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetDataList(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetDataList(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_GetTxSell(t *testing.T) {
	Convey("test MessageInPayload.GetTxSell", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.GetTxSell(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetTxSell(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxSell(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_GetTxBuy(t *testing.T) {
	Convey("test MessageInPayload.GetTxBuy", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.GetTxBuy(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetTxBuy(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			defer UnpatchAll()
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxBuy(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})
	})
}

func TestPreset_GetTxVerify(t *testing.T) {
	Convey("test MessageInPayload.GetTxVerify", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(_ *storage.SQLite, out interface{}, _, _ string, _ ...interface{}) (int64, error) {
				if outAssert, ok := out.(*DBDef.Account); ok {
					*outAssert = DBDef.Account{Verify: []byte{91, 34, 48, 34, 93}}
				}
				return 1, nil
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.GetTxVerify(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("json unmarshal failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(_ *storage.SQLite, out interface{}, _, _ string, _ ...interface{}) (int64, error) {
				if outAssert, ok := out.(*DBDef.Account); ok {
					*outAssert = DBDef.Account{Verify: []byte{48, ':'}}
				}
				return 1, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxVerify(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("invalid character ':' after top-level value"))
		})

		Convey("db read second time failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(_ *storage.SQLite, out interface{}, _, _ string, _ ...interface{}) (int64, error) {
				if outAssert, ok := out.(*DBDef.Account); ok {
					*outAssert = DBDef.Account{Verify: []byte{91, 34, 48, 34, 93}}
				} else {
					return 1, errors.New("an error")
				}
				return 1, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxVerify(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetTxVerify(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxVerify(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})

		Convey("acc.verify is nil", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxVerify(nil)
			So(output, ShouldEqual, "")
			So(err, ShouldBeNil)
		})
	})
}

func TestPreset_GetTxArbitrate(t *testing.T) {
	Convey("test MessageInPayload.GetTxArbitrate", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(_ *storage.SQLite, out interface{}, _, _ string, _ ...interface{}) (int64, error) {
				if outAssert, ok := out.(*DBDef.Account); ok {
					*outAssert = DBDef.Account{Arbitrate: []byte{91, 34, 48, 34, 93}}
				}
				return 1, nil
			})
			defer UnpatchAll()

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.GetTxArbitrate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("json unmarshal failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(_ *storage.SQLite, out interface{}, _, _ string, _ ...interface{}) (int64, error) {
				if outAssert, ok := out.(*DBDef.Account); ok {
					*outAssert = DBDef.Account{Arbitrate: []byte{48, ':'}}
				}
				return 1, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxArbitrate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("invalid character ':' after top-level value"))
		})

		Convey("db read second time failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(_ *storage.SQLite, out interface{}, _, _ string, _ ...interface{}) (int64, error) {
				if outAssert, ok := out.(*DBDef.Account); ok {
					*outAssert = DBDef.Account{Arbitrate: []byte{91, 34, 48, 34, 93}}
				} else {
					return 1, errors.New("an error")
				}
				return 1, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxArbitrate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.GetTxArbitrate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db read failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxArbitrate(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("an error"))
		})

		Convey("acc.arbitrate is nil", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Read, func(*storage.SQLite, interface{}, string, string, ...interface{}) (int64, error) {
				return 1, nil
			})
			defer UnpatchAll()
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.GetTxArbitrate(nil)
			So(output, ShouldEqual, "")
			So(err, ShouldBeNil)
		})
	})
}

func TestPreset_ModifyNickname(t *testing.T) {
	Convey("test MessageInPayload.ModifyNickname", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Update, func(*storage.SQLite, interface{}, map[string]interface{}, string, ...interface{}) (int64, error) {
				return 1, nil
			})

			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}

			output, err := preIns.ModifyNickname(&definition.MessageInPayload{Nickname: definition.ModifyNickname{Nickname: "nick name"}})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("current user is nil", func() {
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{}}}
			output, err := preIns.ModifyNickname(nil)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.New("Current user is nil. "))
		})

		Convey("db update failed", func() {
			controller := gomock.NewController(t)
			mockClientObj := mock_scry.NewMockClient(controller)
			mockClientObj.EXPECT().Account().Return(&auth.UserAccount{Addr: "0x9693"})
			Patch((*storage.SQLite).Update, func(*storage.SQLite, interface{}, map[string]interface{}, string, ...interface{}) (int64, error) {
				return 1, errors.New("an error")
			})
			preIns := &Preset{CBs: &cec.Callbacks{WS: &server.WSServer{CurUser: mockClientObj, DB: &storage.SQLite{}}}}
			output, err := preIns.ModifyNickname(&definition.MessageInPayload{Nickname: definition.ModifyNickname{Nickname: "nick name"}})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "Modify nickname failed. "))
		})
	})
}

func TestPreset_ModifyContractParam(t *testing.T) {
	Convey("test MessageInPayload.ModifyContractParam", t, func() {
		Convey("standard input, expect success", func() {
			controller := gomock.NewController(t)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ModifyContractParam(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			preIns := &Preset{
				Bin:      &binary.Binary{},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}

			output, err := preIns.ModifyContractParam(&definition.MessageInPayload{
				Contract: definition.ModifyContractParam{
					ParamName:  "1",
					ParamValue: "100",
				},
			})
			So(output, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("modify contract params failed", func() {
			controller := gomock.NewController(t)
			mockChainWrapperObj := mock_scry.NewMockChainWrapper(controller)
			Patch((*binary.Binary).ChainWrapper, func(*binary.Binary) scry.ChainWrapper {
				return mockChainWrapperObj
			})
			defer UnpatchAll()
			mockChainWrapperObj.EXPECT().ModifyContractParam(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("an error"))
			preIns := &Preset{
				Bin:      &binary.Binary{},
				Deployer: &definition.MessageInPayload{Address: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8", Password: "111111"},
			}
			output, err := preIns.ModifyContractParam(&definition.MessageInPayload{
				Contract: definition.ModifyContractParam{
					ParamName:  "1",
					ParamValue: "100",
				},
			})
			So(output, ShouldBeNil)
			So(err, ShouldBeError, errors.Wrap(errors.New("an error"), "modify contract param failed"))
		})
	})
}
