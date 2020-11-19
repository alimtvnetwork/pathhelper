package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type widowsDirectoryTestCaseWrapper struct {
	expected, operatingSystemMessage string
	operatingSystem                  enums.OperatingSystem
}

var widowsDirectoryTestCaseWrappers = []widowsDirectoryTestCaseWrapper{
	{
		expected:               "C:\\Windows",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		expected:               "",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetWidowsDirectory_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range widowsDirectoryTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[GetWidowsDirectory] expects (%s)", testCase.operatingSystemMessage, testCase.expected)

		executeTestCaseForGetWidowsDirectory(t, testCaseMessage, testCase)
	}
}

func TestGetWidowsDirectory_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range widowsDirectoryTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[GetWidowsDirectory] expects (%s)", testCase.operatingSystemMessage, testCase.expected)

		executeTestCaseForGetWidowsDirectory(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetWidowsDirectory(t *testing.T, testCaseMessage string, testCase widowsDirectoryTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetWidowsDirectory()

		// Assert
		So(actual, ShouldEqual, testCase.expected)
	})
}
