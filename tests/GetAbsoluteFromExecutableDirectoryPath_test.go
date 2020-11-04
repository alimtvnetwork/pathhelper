package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type absoluteFromExecutableDirectoryPathTestCaseWrapper struct {
	inputRelativePath, expected, expectedMessage, operatingSystemMessage string
	isLongPathFix, isNormalize                                           bool
	operatingSystem                                                      enums.OperatingSystem
}

// todo
var absoluteFromExecutableDirectoryPathTestCaseWrappers = []absoluteFromExecutableDirectoryPathTestCaseWrapper{
	{
		inputRelativePath:      "\\Users",
		expected:               "C:\\Users",
		expectedMessage:        "C:\\Users",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		inputRelativePath:      "\\Users",
		expected:               "C:\\Users",
		expectedMessage:        "C:\\Users",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		inputRelativePath:      "//\\home//",
		expected:               "/home/",
		expectedMessage:        "/home/",
		operatingSystemMessage: "Linux OS",
		operatingSystem:        enums.Linux,
		isNormalize:            true,
		isLongPathFix:          true,
	},
}

func TestGetAbsoluteFromExecutableDirectoryPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range absoluteFromExecutableDirectoryPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsoluteFromExecutableDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(t, testCaseMessage, testCase)
	}
}

func TestGetAbsoluteFromExecutableDirectoryPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range absoluteFromExecutableDirectoryPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsoluteFromExecutableDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(
	t *testing.T, testCaseMessage string, testCase absoluteFromExecutableDirectoryPathTestCaseWrapper,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetAbsoluteFromExecutableDirectoryPath(
			testCase.inputRelativePath,
			testCase.isLongPathFix,
			testCase.isNormalize)

		// Assert
		So(actual, ShouldNotBeEmpty)
		So(actual, ShouldContainSubstring, testCase.expected)
	})
}
