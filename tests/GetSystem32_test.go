package tests

import (
	"fmt"
	"testing"

	"gitlab.com/evatix-go/pathhelper"

	. "github.com/smartystreets/goconvey/convey"
)

var expectedSystem32 string = "C:\\Windows\\System32"

func TestGetSystem32(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	// Arrange
	testCaseMessage := fmt.Sprintf("[GetSystem32] expects (%s)", expectedSystem32)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetSystem32()

		// Assert
		So(actual, ShouldEqual, expectedSystem32)
	})
}
