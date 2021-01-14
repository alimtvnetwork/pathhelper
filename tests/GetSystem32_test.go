package tests

import (
	"fmt"
	"testing"

	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/pathhelper"

	. "github.com/smartystreets/goconvey/convey"
)

var expectedSystem32 = "C:\\Windows\\System32"

func TestGetSystem32(t *testing.T) {
	coretests.SkipOnUnix(t)

	// Arrange
	testCaseMessage := fmt.Sprintf("[GetSystem32] expects (%s)", expectedSystem32)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetSystem32()

		// Assert
		So(actual, ShouldEqual, expectedSystem32)
	})
}
