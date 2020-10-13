package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type pathExistTestCaseWrapper struct {
	input, expectedMessage, operatingSystemMessage string
	expected                                       bool
	operatingSystem                                enums.OperatingSystem
}

var pathExistTestCaseWrappers = []pathExistTestCaseWrapper{
	{
		input:                  "",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "any OS",
	},
	{
		input:                  "c:\\sampleSomething",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Windows",
		operatingSystem:        enums.Windows,
	},
	{
		input:                  "~/home",
		expected:               true,
		expectedMessage:        "true",
		operatingSystemMessage: "OS is Unix",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestIsPathExist_Windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range pathExistTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathExist(t, testCaseMessage, testCase)
	}
}

func TestIsPathExist_Unix(t *testing.T) {
	if pathhelper.IsWindows() {
		t.Skip("Unix tests ignored in Windows.")
	}

	for _, testCase := range pathExistTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathExist(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForIsPathExist(t *testing.T, testCaseMessage string, testCase pathExistTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.IsPathExist(testCase.input)

		// Assert
		So(actual, ShouldNotBeEmpty)
		So(actual, ShouldEqual, testCase.expected)
	})
}
