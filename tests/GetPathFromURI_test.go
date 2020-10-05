package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetPathFromURI(t *testing.T) {

	Convey("If given OS is windows", t, func() {
		So(pathhelper.IsWindows(), ShouldBeTrue)
		Convey("if GetPathFromUri is run", func() {

			Convey("it should return", func() {
				So(pathhelper.GetPathFromUri("file://c:/windows/users/etc/more"), ShouldEqual, "c:\\windows\\users\\etc\\more")
			})

		})

	})

	Convey("if given OS is not windows", t, func() {

		Convey("it should return", func() {

			Convey("it should return", func() {
				So(pathhelper.GetPathFromUri("file://c:/windows/users/etc/more"), ShouldNotEqual, "c:/windows/users/etc/more")
			})

		})

	})
}
