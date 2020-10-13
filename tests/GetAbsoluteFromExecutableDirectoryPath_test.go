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
	operatingSystem                                                      enums.OperatingSystem
}

// todo
var absoluteFromExecutableDirectoryPathTestCaseWrappers = []absoluteFromExecutableDirectoryPathTestCaseWrapper{
	{
		inputRelativePath:      "\\whatever",
		expected:               "",
		expectedMessage:        "",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputRelativePath:      "\\whatever",
		expected:               "",
		expectedMessage:        "",
		operatingSystemMessage: "Linux OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetAbsoluteFromExecutableDirectoryPath_Windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

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
	if pathhelper.IsWindows() {
		t.Skip("Unix tests ignored in Windows.")
	}

	for _, testCase := range absoluteFromExecutableDirectoryPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsoluteFromExecutableDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(t *testing.T, testCaseMessage string, testCase absoluteFromExecutableDirectoryPathTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetAbsoluteFromExecutableDirectoryPath(testCase.inputRelativePath)

		// Assert
		So(actual, ShouldNotBeBlank)
		So(actual, ShouldNotBeEmpty)
		So(actual, ShouldNotBeNil)
		So(pathhelper.GetExecutablePath(), ShouldContainSubstring, "C:\\Users\\Naureen\\AppData\\Local\\Temp\\go-build")
	})
}
