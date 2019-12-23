package cec

import (
	. "bou.ke/monkey"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/app/server"
	"github.com/scryinfo/dp/dots/app/storage"
	ipfs "github.com/scryinfo/dp/dots/storage/ipfs"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCallbacks_Create(t *testing.T) {
	Convey("test Callbacks.Create", t, func() {
		cbIns := &Callbacks{}

		output := cbIns.Create(nil)
		So(output, ShouldBeNil)
	})
}

func TestNewCBsDot(t *testing.T) {
	Convey("test Callbacks config", t, func() {
		Convey("standard input, expect success", func() {
			confBs := []byte{123,34,112,114,111,111,102,115,79,117,116,68,105,114,34,58,32,34,67,58,47,85,115,101,114,115,47,87,105,108,108,47,68,101,115,107,116,111,112,34,125}

			output, err := newCBsDot(confBs)
			So(err, ShouldBeNil)

			outputAssert, ok := output.(*Callbacks)
			So(ok, ShouldBeTrue)
			So(outputAssert.config.ProofsOutDir, ShouldEqual, "C:/Users/Will/Desktop")
		})

		Convey("unmarshal failed", func() {
			var confBs []byte
			output, err := newCBsDot(confBs)
			So(output, ShouldBeNil)
			So(err, ShouldBeError, dot.NewError("dot_error_parameter", "the parameter error "))
		})
	})
}

func TestCBsTypeLive(t *testing.T) {
	Convey("test Preset TypeLive", t, func() {
		Convey("standard input, expect success", func() {
			dotIns := &dot.TypeLives{
				Meta: dot.Metadata{
					TypeId: CBsTypeId,
					NewDoter: func(conf []byte) (dot.Dot, error) {
						return newCBsDot(conf)
					},
				},
			}

			Patch(server.WebSocketTypeLive, func() *dot.TypeLives {
				return dotIns
			})
			defer UnpatchAll()
			Patch(ipfs.IpfsTypeLive, func() *dot.TypeLives {
				return dotIns
			})
			Patch(storage.SQLiteTypeLive, func() *dot.TypeLives {
				return dotIns
			})

			output := CBsTypeLive()
			So(output[1], ShouldEqual, dotIns)
			So(output[2], ShouldEqual, dotIns)
			So(output[3], ShouldEqual, dotIns)
		})
	})
}

func TestUpdateSlice(t *testing.T) {
	Convey("test UpdateSlice", t, func() {
		bs := []byte{91, 34, 105, 116, 101, 109, 49, 34, 44, 34, 105, 116, 101, 109, 50, 34, 44, 34, 105, 116, 101, 109, 51, 34, 93} // {"item1","item2","item3"}
		str := []string{"item2", "elem not exist"}
		mode := []string{"add", "delete", "mode not exist"}

		Convey("standard add input, expect success", func() {
			result, err := UpdateSlice(bs, str[1], mode[0])
			So(string(result), ShouldEqual, string([]byte{91, 34, 105, 116, 101, 109, 49, 34, 44, 34, 105, 116, 101, 109, 50, 34, 44, 34, 105, 116, 101, 109, 51, 34, 44, 34, 101, 108, 101, 109, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 34, 93}))
			So(err, ShouldBeNil)
		})

		Convey("standard delete input, expect success", func() {
			result, err := UpdateSlice(bs, str[0], mode[1])
			So(string(result), ShouldEqual, string([]byte{91, 34, 105, 116, 101, 109, 49, 34, 44, 34, 105, 116, 101, 109, 51, 34, 93})) // {"item1","item3"}
			So(err, ShouldBeNil)
		})

		Convey("json unmarshal failed", func() {
			result, err := UpdateSlice([]byte{105, 116, 101, 109, 49, 105, 116, 101, 109, 50, 105, 116, 101, 109, 51}, str[1], mode[0])
			So(result, ShouldBeNil)
			So(err, ShouldBeError, errors.New("json unmarshal failed. : invalid character 'i' looking for beginning of value"))
		})

		Convey("unknown mode", func() {
			result, err := UpdateSlice(bs, str[0], mode[2])
			So(string(result), ShouldEqual, string(bs))
			So(err, ShouldBeNil)
		})
	})
}

func TestOnPublish(t *testing.T) {
    Convey("test onPublish", t, func() {
        Convey("standard input, ecpect success", func() {
            cbIns := &Callbacks{DB: &storage.SQLite{},WS: &server.WSServer{}}
        })
    })
}
