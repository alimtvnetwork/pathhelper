package tests

import (
	"testing"
	"fmt"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type gAbsoluteFromExecutableDirectoryPath struct {
	i
}

func TestGetAbsoluteFromExecutableDirectoryPath_Windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range pathSeparatorTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathSeparator] inputs () expects (%s)", testCase.operatingSystemMessage, testCase.expectedMessage)

		executeTestCaseForGetPathSeparator(t, testCaseMessage, testCase)
	}
}
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
