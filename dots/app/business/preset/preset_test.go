package preset

import (
    "github.com/golang/mock/gomock"
    "github.com/pkg/errors"
    cec "github.com/scryinfo/dp/dots/app/business/preset/chain_event"
    mock_scry "github.com/scryinfo/dp/dots/binary/scry/mock"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestPreset_Logout(t *testing.T) {
    Convey("test Preset.Logout", t, func() {
        Convey("standard input, except success", func() {
            controller := gomock.NewController(t)
            mockClientObj := mock_scry.NewMockClient(controller)
            mockClientObj.EXPECT().UnSubscribeEvent("Approval").Return(nil)

            preIns := &Preset{CBs: &cec.Callbacks{CurUser:mockClientObj, EventNames: []string{"Approval"}}}
            output, err := preIns.Logout(nil)

            So(output, ShouldEqual, true) // interface{}.(bool) == bool
            So(err, ShouldBeNil)
        })

        Convey("current user is nil", func() {
            preIns := &Preset{CBs:&cec.Callbacks{}}
            output, err := preIns.Logout(nil)

            var outputExpect interface{}
            So(output, ShouldEqual, outputExpect)
            So(err, ShouldBeError, errors.New("Current user is nil. "))
        })

        Convey("try to unsubscribe known event", func() {
            controller := gomock.NewController(t)
            mockClientObj := mock_scry.NewMockClient(controller)
            mockClientObj.EXPECT().UnSubscribeEvent("Approval").Return(errors.New("a new error"))

            preIns := &Preset{CBs: &cec.Callbacks{CurUser:mockClientObj, EventNames: []string{"Approval"}}}
            output, err := preIns.Logout(nil)

            var outputExpect interface{}
            So(output, ShouldEqual, outputExpect) // interface{}.(bool) == bool
            So(err, ShouldBeError, errors.New("Unsubscribe failed, event:  Approval . : a new error"))
        })
    })
}
