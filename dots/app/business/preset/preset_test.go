package preset

import (
	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	mock_dot "github.com/scryinfo/dot/dot/mock"
	cec "github.com/scryinfo/dp/dots/app/business/preset/chain_event"
	"github.com/scryinfo/dp/dots/binary"
	mock_scry "github.com/scryinfo/dp/dots/binary/scry/mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPreset_Create(t *testing.T) {
	Convey("test Preset.Create (dot.Create)", t, func() {
		Convey("standard input, except success", func() {
			preIns := &Preset{}

			controller := gomock.NewController(t)
			mockLineObj := mock_dot.NewMockLine(controller)

			output := preIns.Create(mockLineObj)
			So(output, ShouldBeNil)
		})
	})
}

func TestPreTypeLive(t *testing.T) {
	Convey("test Preset TypeLive", t, func() {
		Convey("standard input, except success", func() {
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

			output := PreTypeLive()
			So(output[1], ShouldEqual, dotIns)
			So(output[2], ShouldEqual, dotIns)
		})
	})
}

func TestPreset_Logout(t *testing.T) {
	Convey("test Preset.Logout", t, func() {
		Convey("standard input, except success", func() {
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
