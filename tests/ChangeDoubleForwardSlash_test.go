package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type doubleForwardSlashTestCaseWrapper struct {
	inputPath, inputSeparator, expected string
}

var doubleForwardSlashTestCaseWrappers = []doubleForwardSlashTestCaseWrapper{
	{
		inputPath:      "c://",
		inputSeparator: "\\",
		expected:       "c:\\",
	},
	{
		inputPath:      "c://",
		inputSeparator: "/",
		expected:       "c:/",
	},
	{
		inputPath:      "c:/",
		inputSeparator: "",
		expected:       "c:/",
	},
}

func TestChangeDoubleForwardSlash(t *testing.T) {
	for i, testCase := range doubleForwardSlashTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[ChangeDoubleForwardSlash] inputs (%s, %s) expects (%s)", testCase.inputPath, testCase.inputSeparator, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.ChangeDoubleForwardSlash(testCase.inputPath, testCase.inputSeparator)

			// Assert
			Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
