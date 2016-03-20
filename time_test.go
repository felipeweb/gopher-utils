package gopher_utils

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"time"
)

func TestDate(t *testing.T) {
	Convey("Convert unix date format to string", t, func() {
		So(Date(1, "DD/MM/YYYY"), ShouldEqual, "31/12/1969")
	})
	Convey("Convert unix date in string format to humman string", t, func() {
		So(DateS("1", "DD/MM/YYYY"), ShouldEqual, "31/12/1969")
	})
	Convey("Convert time object to string", t, func() {
		So(DateT(time.Unix(int64(1), 0), "DD/MM/YYYY"), ShouldEqual, "31/12/1969")
	})
}
