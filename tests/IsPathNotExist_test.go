package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type pathNotExistTestCaseWrapper struct {
	input, expectedMessage, operatingSystemMessage string
	expected                                       bool
	operatingSystem                                enums.OperatingSystem
}

var pathNotExistTestCaseWrappers = []pathNotExistTestCaseWrapper{
	{
		input:                  "",
		expected:               true,
		expectedMessage:        "true",
		operatingSystemMessage: "any OS",
	},
	{
		input:                  "C:\\Users",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Windows",
		operatingSystem:        enums.Windows,
	},
	{
		input:                  "home/user",
		expected:               true,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Unix",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestIsPathNotExist_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range pathNotExistTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathNotExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathNotExist(t, testCaseMessage, testCase, i)
	}
}

func TestIsPathNotExist_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range pathNotExistTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathNotExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathNotExist(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForIsPathNotExist(t *testing.T, testCaseMessage string, testCase pathNotExistTestCaseWrapper, i int) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.IsPathNotExist(testCase.input)

		// Assert
		Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
