package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type absolutePathTestCaseWrapper struct {
	inputBasepath, inputRelativePath, expected, expectedMessage, operatingSystemMessage string
	operatingSystem                                                                     enums.OperatingSystem
}

var absolutePathTestCaseWrappers = []absolutePathTestCaseWrapper{
	{
		inputBasepath:          "",
		inputRelativePath:      "",
		expected:               "",
		expectedMessage:        "empty return",
		operatingSystemMessage: "Any OS",
	},
	{
		inputBasepath:          "c:\\Windows\\",
		inputRelativePath:      "\\whatever",
		expected:               "c:\\Windows\\whatever",
		expectedMessage:        "non-empty return of (c:\\Windows\\whatever)",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputBasepath:          "c:\\Windows\\",
		inputRelativePath:      "whatever",
		expected:               "c:\\Windows\\whatever",
		expectedMessage:        "non-empty return of (c:\\Windows\\whatever)",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputBasepath:          "/home/your_user_name/my_script/",
		inputRelativePath:      "/whatever",
		expected:               "/home/your_user_name/my_script/whatever",
		expectedMessage:        "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
	{
		inputBasepath:          "/home/your_user_name/my_script",
		inputRelativePath:      "/whatever",
		expected:               "/home/your_user_name/my_script/whatever",
		expectedMessage:        "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetAbsolutePath_Windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range absolutePathTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsolutePath] inputs (%s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputBasepath, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestForGetAbsolutePath(t, testCaseMessage, testCase)
	}
}

func TestGetAbsolutePath_Unix(t *testing.T) {
	if pathhelper.IsWindows() {
		t.Skip("Unix tests ignored in Windows.")
	}

	for _, testCase := range absolutePathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsolutePath] inputs (%s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputBasepath, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestForGetAbsolutePath(t, testCaseMessage, testCase)
	}
}

func executeTestForGetAbsolutePath(t *testing.T, testCaseMessage string, testCase absolutePathTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetAbsolutePath(testCase.inputBasepath, testCase.inputRelativePath)

		// Assert
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
