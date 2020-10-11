package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetPathAsURI(t *testing.T) {
	Convey("If given OS is windows", t, func() {
		So(pathhelper.IsWindows(), ShouldBeTrue)
		Convey("if GetPathAsUri is run", func() {

			Convey("it should return", func() {
				So(pathhelper.GetPathAsUri("c:\\windows", true), ShouldEqual, "file:\\\\c:\\windows")
				So(pathhelper.GetPathAsUri("c:/windows", true), ShouldEqual, "file:\\\\c:\\windows")
			})
		})
	})

	Convey("if given OS is not windows", t, func() {
		So(pathhelper.IsWindows(), ShouldBeTrue)
		Convey("it should return", func() {
			Convey("it should return", func() {
				So(pathhelper.GetPathAsUri("c:/windows", true), ShouldNotEqual, "file://c:/windows")
			})
		})
	})
}
