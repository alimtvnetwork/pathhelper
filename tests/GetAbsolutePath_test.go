package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetAbsolutePath(t *testing.T) {
	Convey("If function is run", t, func() {
		Convey("if not windows", func() {
			So(pathhelper.IsWindows(), ShouldBeTrue)
			Convey("it should return absolute path", func() {
				So(pathhelper.GetAbsolutePath("c:\\Windows\\", "\\whatever"), ShouldEqual, "c:\\Windows\\whatever")
				So(pathhelper.GetAbsolutePath("c:\\Windows\\", "whatever"), ShouldEqual, "c:\\Windows\\whatever")
				So(pathhelper.GetAbsolutePath("c:\\Windows", "\\whatever"), ShouldEqual, "c:\\Windows\\whatever")
			})
		})
	})

	Convey("If function is run", t, func() {
		Convey("if not windows", func() {
			So(pathhelper.IsWindows(), ShouldBeTrue)
			Convey("it should return absolute path", func() {
				So(pathhelper.GetAbsolutePath("/home/your_user_name/my_script/", "/whatever"), ShouldNotEqual, "/home/your_user_name/my_script/whatever")
				So(pathhelper.GetAbsolutePath("/home/your_user_name/my_script/", "whatever"), ShouldNotEqual, "/home/your_user_name/my_script/whatever")
				So(pathhelper.GetAbsolutePath("/home/your_user_name/my_script", "/whatever"), ShouldNotEqual, "/home/your_user_name/my_script/whatever")
			})

		})
	})

}
