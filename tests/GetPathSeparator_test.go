package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetPathSeparator(t *testing.T) {

	Convey("If given OS is windows", t, func() {

		Convey("it should return \"\\", func() {
			So(pathhelper.GetPathSeparator(), ShouldEqual, "\\")
		})

	})

	Convey("if given OS is windows", t, func() {

		Convey("it should not return \"/\"", func() {
			So(pathhelper.GetPathSeparator(), ShouldNotEqual, "/")
		})

	})

}
