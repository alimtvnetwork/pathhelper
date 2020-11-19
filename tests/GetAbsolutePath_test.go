package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type absolutePathTestCaseWrapper struct {
	basePath, inputRelativePath, expected, expectedMessage, operatingSystemMessage string
	isLongPathFix, isNormalize                                                     bool
	operatingSystem                                                                enums.OperatingSystem
}

var absolutePathTestCaseWrappers = []absolutePathTestCaseWrapper{
	// todo catch panic
	//{
	//	basePath:               "",
	//	inputRelativePath:      "",
	//	expected:               "",
	//	expectedMessage:        "empty return",
	//	operatingSystemMessage: "Any OS",
	//	operatingSystem:        enums.Any,
	//	isNormalize:            true,
	//	isLongPathFix:          true,
	//},
	{
		basePath:               "c:\\Windows\\//",
		inputRelativePath:      "\\whatever",
		expected:               "c:\\Windows\\whatever",
		expectedMessage:        "non-empty return of (c:\\Windows\\whatever)",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		basePath:               "c:\\\\Windows///",
		inputRelativePath:      "whatever",
		expected:               "c:\\Windows\\whatever",
		expectedMessage:        "non-empty return of (c:\\Windows\\whatever)",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		basePath:               "/home/\\//your_user_name/my_script/",
		inputRelativePath:      "/whatever",
		expected:               "/home/your_user_name/my_script/whatever",
		expectedMessage:        "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		basePath:               "/home/your_user_name/my_script",
		inputRelativePath:      "/whatever",
		expected:               "/home/your_user_name/my_script//whatever",
		expectedMessage:        "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystemMessage: "Unix OS",
		isNormalize:            false,
		isLongPathFix:          true,
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetAbsolutePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range absolutePathTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsolutePath] inputs (%s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.basePath, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestForGetAbsolutePath(t, testCaseMessage, testCase)
	}
}

func TestGetAbsolutePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range absolutePathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsolutePath] inputs (%s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.basePath, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestForGetAbsolutePath(t, testCaseMessage, testCase)
	}
}

func executeTestForGetAbsolutePath(t *testing.T, testCaseMessage string, testCase absolutePathTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetAbsolutePath(
			testCase.basePath,
			testCase.inputRelativePath,
			testCase.isLongPathFix,
			testCase.isNormalize)

		// Assert
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
