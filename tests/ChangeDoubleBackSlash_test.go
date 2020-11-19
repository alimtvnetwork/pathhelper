package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type doubleBackSlashTestCaseWrapper struct {
	inputPath, inputSeparator, expected string
}

var doubleBackSlashTestCaseWrappers = []doubleBackSlashTestCaseWrapper{
	{
		inputPath:      "c:\\\\",
		inputSeparator: "\\",
		expected:       "c:\\",
	},
	{
		inputPath:      "c:\\\\",
		inputSeparator: "/",
		expected:       "c:/",
	},
	{
		inputPath:      "c:/",
		inputSeparator: "",
		expected:       "c:/",
	},
}

func TestChangeDoubleBackSlash(t *testing.T) {
	for _, testCase := range doubleBackSlashTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[ChangeDoubleBackSlash] inputs (%s, %s) expects (%s)", testCase.inputPath, testCase.inputSeparator, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.ChangeDoubleBackSlash(testCase.inputPath, testCase.inputSeparator)

			// Assert
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
