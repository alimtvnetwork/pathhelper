package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

type combinePathsWithTestCaseDataWrapper struct {
	inputPath1, inputPath2, inputPath3 string
	operatingSystem                    enums.OperatingSystem
	expected, operatingSystemMessage   string
}

var combinePathsWithTestCaseWrappers = []combinePathsWithTestCaseDataWrapper{
	{
		inputPath1:             "something",
		inputPath2:             "somethingElse",
		inputPath3:             "otherThings",
		operatingSystem:        enums.Windows,
		operatingSystemMessage: "Windows OS",
		expected:               "something\\somethingElse\\otherThings",
	},
	{
		inputPath1:             "",
		inputPath2:             "",
		inputPath3:             "",
		operatingSystem:        enums.Windows,
		operatingSystemMessage: "Windows OS",
		expected:               "\\",
	},
	{
		inputPath1:             "something",
		inputPath2:             "somethingElse",
		inputPath3:             "otherThings",
		operatingSystem:        enums.Ubuntu,
		operatingSystemMessage: "Windows OS",
		expected:               "something/somethingElse/otherThings",
	},
}

func TestGetCombinePathsWith_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range combinePathsWithTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinePathsWith] inputs (%s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath1, testCase.inputPath2, testCase.inputPath3, testCase.expected)

		executeTestForGetCombinePathsWith(t, testCaseMessage, testCase)
	}
}

func TestGetCombinePathsWith_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range combinePathsWithTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinePathsWith] inputs (%s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath1, testCase.inputPath2, testCase.inputPath3, testCase.expected)

		executeTestForGetCombinePathsWith(t, testCaseMessage, testCase)
	}
}

func executeTestForGetCombinePathsWith(t *testing.T, testCaseMessage string, testCase combinePathsWithTestCaseDataWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetCombinePathsWith(
			testCase.inputPath1,
			testCase.inputPath2,
			testCase.inputPath3)

		// Assert
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
