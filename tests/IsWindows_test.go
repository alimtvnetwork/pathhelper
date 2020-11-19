package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestIsWindows(t *testing.T) {
	// Arrange
	Convey("function should return true on windows OS", t, func() {
		SkipOnUnix(t)

		// Act
		actual := pathhelper.IsWindows()

		//Assert
		So(actual, ShouldBeTrue)
	})

	Convey("function should return false on unix OS", t, func() {
		SkipOnWindows(t)

		// Act
		actual := pathhelper.IsWindows()

		//Assert
		So(actual, ShouldBeFalse)
	})
}
