package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var path = "something/whatever"

func TestIsEmptyPathForPtr(t *testing.T) {
	// Arrange
	testCaseMessage := fmt.Sprintf("[IsEmptyPathForPtr] inputs (something/whatever) expects (false)")

	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelpercore.IsEmptyPathForPtr(&path)

		// Assert
		So(actual, ShouldBeFalse)
	})
}
