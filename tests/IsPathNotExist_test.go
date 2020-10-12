package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type isPathNotExistTestCaseWrapper struct {
	input, expectedMessage, operatingSystem string
	expected                                bool
}

var isPathNotExistTestCaseWrappers = []isPathNotExistTestCaseWrapper{
	{
		input:           "",
		expected:        true,
		expectedMessage: "true",
		operatingSystem: "any OS",
	},
	{
		input:           "C:\\Users",
		expected:        false,
		expectedMessage: "false",
		operatingSystem: "OS is windows",
	},
	{
		input:           "home/user",
		expected:        false,
		expectedMessage: "false",
		operatingSystem: "OS is linux",
	},
}

func TestIsPathNotExist(t *testing.T) {
	for _, testCase := range isPathNotExistTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("(%s)[IsPathNotExist] inputs (%s) expects (%s)", testCase.operatingSystem, testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.IsPathNotExist(testCase.input)

			// Assert
			So(actual, ShouldNotBeEmpty)
			if pathhelper.IsWindows() {
				So(actual, ShouldEqual, testCase.expected)
			}
		})
	}
}
