package tests

import (
	"fmt"
	"gitlab.com/evatix-go/pathhelper"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var expectedWindowsRoot string = "C:\\"

func TestGetWindowsRoot(t *testing.T) {
	SkipOnUnix(t)

	// Arrange
	testCaseMessage := fmt.Sprintf("[GetWindowsRoot] expects (%s)", expectedWindowsRoot)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetWindowsRoot()

		// Assert
		So(actual, ShouldEqual, expectedWindowsRoot)
	})
}
