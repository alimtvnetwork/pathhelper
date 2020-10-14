package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type removeFromPathTestCaseWrapper struct {
	inputPath, expected, expectedMessage, operatingSystemMessage string
	inputBool                                                    bool
	operatingSystem                                              enums.OperatingSystem
}

var removingArray = []string{"/"}

var removeFromPathTestCaseWrappers = []removeFromPathTestCaseWrapper{
	{
		inputPath:              "",
		expected:               "",
		expectedMessage:        "empty return",
		operatingSystemMessage: "Any OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputPath:              "c:\\win\\etc",
		inputBool:              true,
		expected:               "c:\\win\\etc",
		expectedMessage:        "c:\\win\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputPath:              "c:\\\\win\\\\etc",
		inputBool:              true,
		expected:               "c:/win/etc",
		expectedMessage:        "c:/win/etc",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestRemoveFromPath_windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range removeFromPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [RemoveFromPath] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath, testCase.inputBool, testCase.expectedMessage)

		executeTestCaseForRemoveFromPath(t, testCaseMessage, testCase)
	}
}

func TestRemoveFromPath_unix(t *testing.T) {
	if pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range removeFromPathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [RemoveFromPath] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath, testCase.inputBool, testCase.expectedMessage)

		executeTestCaseForRemoveFromPath(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForRemoveFromPath(t *testing.T, testCaseMessage string, testCase removeFromPathTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.RemoveFromPath(testCase.inputPath, &removingArray, testCase.inputBool)

		// Assert
		if pathhelpercore.IsEmptyPath(testCase.inputPath) {
			So(actual, ShouldBeEmpty)
		}

		if !pathhelpercore.IsEmptyPath(testCase.inputPath) {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		}
	})
}
