package gopher_utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	Convey("Convert unix date format to string", t, func() {
		So(Date(1, "DD/MM/YYYY"), ShouldEqual, "01/01/1970")
	})
	Convey("Convert unix date in string format to humman string", t, func() {
		So(DateS("1", "DD/MM/YYYY"), ShouldEqual, "01/01/1970")
	})
	Convey("Convert time object to string", t, func() {
		So(DateT(time.Unix(int64(1), 0), "DD/MM/YYYY"), ShouldEqual, "01/01/1970")
	})
}
