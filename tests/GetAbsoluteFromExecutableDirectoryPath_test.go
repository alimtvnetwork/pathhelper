package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetAbsoluteFromExecutableDirectoryPath(t *testing.T) {

	Convey("If function is run", t, func() {

		Convey("it should return absolute path", func() {
			So(pathhelper.GetAbsoluteFromExecutableDirectoryPath("\\whatever"), ShouldNotBeBlank)
			So(pathhelper.GetAbsoluteFromExecutableDirectoryPath("\\whatever"), ShouldNotBeEmpty)
			So(pathhelper.GetAbsoluteFromExecutableDirectoryPath("\\whatever"), ShouldNotBeNil)
			So(pathhelper.GetExecutableDirectory(), ShouldContainSubstring, "C:\\Users\\Naureen\\AppData\\Local\\Temp\\go-build")
		})

	})

}
