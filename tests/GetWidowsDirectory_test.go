package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type windowsDirectoryTestCaseWrapper struct {
	expected, operatingSystemMessage string
	operatingSystem                  enums.OperatingSystem
}

var windowsDirectoryTestCaseWrappers = []windowsDirectoryTestCaseWrapper{
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

	for i, testCase := range windowsDirectoryTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[GetWidowsDirectory] expects (%s)", testCase.operatingSystemMessage, testCase.expected)

		executeTestCaseForGetWidowsDirectory(t, testCaseMessage, testCase, i)
	}
}

func TestGetWidowsDirectory_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range windowsDirectoryTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[GetWidowsDirectory] expects (%s)", testCase.operatingSystemMessage, testCase.expected)

		executeTestCaseForGetWidowsDirectory(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetWidowsDirectory(
	t *testing.T, testCaseMessage string, testCase windowsDirectoryTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetWidowsDirectory()

		// Assert
		Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
