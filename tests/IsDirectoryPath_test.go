package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type directoryPathTestCaseWrapper struct {
	operatingSystem                                enums.OperatingSystem
	input, operatingSystemMessage, expectedMessage string
	expected                                       bool
}

var directoryPathTestCaseWrappers = []directoryPathTestCaseWrapper{
	{
		input:                  "C:\\users",
		operatingSystem:        enums.Windows,
		operatingSystemMessage: "Windows OS",
		expected:               true,
		expectedMessage:        "true",
	},
	{
		input:                  "/home",
		operatingSystem:        enums.Ubuntu,
		operatingSystemMessage: "Unix os",
		expected:               true,
		expectedMessage:        "true",
	},
}

func TestIsDirectoryPath_Windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range directoryPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [IsDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsDirectoryPath(t, testCaseMessage, testCase)
	}
}

func TestIsDirectoryPath_Unix(t *testing.T) {
	if pathhelper.IsWindows() {
		t.Skip("Unix tests ignored in Windows.")
	}

	for _, testCase := range directoryPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [IsDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsDirectoryPath(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForIsDirectoryPath(t *testing.T, testCaseMessage string, testCase directoryPathTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.IsDirectoryPath(testCase.input)

		// Assert
		So(actual, ShouldNotBeEmpty)
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
