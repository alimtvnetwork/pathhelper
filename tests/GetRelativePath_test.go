package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetRelativePath(t *testing.T) {

	Convey("If given full-path does not contain basepath", t, func() {
		Convey("it should return \"\"", func() {
			So(pathhelper.GetRelativePath("c:\\Windows\\whatevr", "ex"), ShouldEqual, "")
		})
	})

	Convey("If given fullpath and basepath are same", t, func() {
		Convey("it should return \"\"", func() {
			So(pathhelper.GetRelativePath("c:\\Windows\\whatevr", "c:\\Windows\\whatevr"), ShouldEqual, "Both paths are same")
		})
	})

	Convey("if given fullpath does not contain basepath", t, func() {
		Convey("it should return basepath replaced with \"\" in the fullpath", func() {
			So(pathhelper.GetRelativePath("c:\\Windows\\whatevr", "c:\\Windows"), ShouldEqual, "\\whatevr")
		})
	})

}
