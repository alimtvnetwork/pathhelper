package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetExecutableDirectory(t *testing.T) {
	Convey("If function is run", t, func() {
		Convey("it should return \"/d/Programming/Practice/GO/src/pathhelper\"", func() {
			So(pathhelper.GetExecutableDirectory(), ShouldNotBeBlank)
			So(pathhelper.GetExecutableDirectory(), ShouldNotBeEmpty)
			So(pathhelper.GetExecutableDirectory(), ShouldNotBeNil)
			So(pathhelper.GetExecutableDirectory(), ShouldContainSubstring, "C:\\Users\\Naureen\\AppData\\Local\\Temp\\go-build")
		})
	})
}
