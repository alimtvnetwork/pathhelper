package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetAbsoluteFromExecutableDirectoryPath(t *testing.T) {

	Convey("If function is run", t, func() {

		Convey("it should return absolute path", func() {
			// Act
			actual := pathhelper.GetAbsoluteFromExecutableDirectoryPath("\\whatever")

			// Assert
			So(actual, ShouldNotBeBlank)
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldNotBeNil)
			So(pathhelper.GetExecutablePath(), ShouldContainSubstring, "C:\\Users\\Naureen\\AppData\\Local\\Temp\\go-build")
		})

	})

}
