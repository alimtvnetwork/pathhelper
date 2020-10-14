package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var pathToCheck = "something/whatever"

func TestIsEmptyPath(t *testing.T) {
	// Arrange
	testCaseMessage := fmt.Sprintf("[IsEmptyPath] inputs (something/whatever) expects (false)")

	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelpercore.IsEmptyPath(pathToCheck)

		// Assert
		So(actual, ShouldBeFalse)
	})
}
