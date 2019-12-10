package cec

import (
	"github.com/pkg/errors"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

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
