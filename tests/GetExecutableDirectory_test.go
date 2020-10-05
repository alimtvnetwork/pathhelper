package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetExecutableDirectory(t *testing.T) {
	Convey("If function is run", t, func() {
		Convey("it should return \"/d/Programming/Practice/GO/src/pathhelper\"", func() {
			So(pathhelper.GetExecutablePath(), ShouldNotBeBlank)
			So(pathhelper.GetExecutablePath(), ShouldNotBeEmpty)
			So(pathhelper.GetExecutablePath(), ShouldNotBeNil)
			So(pathhelper.GetExecutablePath(), ShouldContainSubstring, "C:\\Users\\Naureen\\AppData\\Local\\Temp\\go-build")
		})
	})
}
