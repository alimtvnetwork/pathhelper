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
	// {
	// 	input:           "",
	// 	expected:        true,
	// 	expectedMessage: "true",
	// 	operatingSystemMessage: "any OS",
	// },
	{
		input:                  "c:\\sampleSomething",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Windows",
		operatingSystem:        enums.Windows,
	},
	{
		input:                  "C:\\Users",
		expected:               true,
		expectedMessage:        "true",
		operatingSystemMessage: "OS is Windows",
		operatingSystem:        enums.Windows,
	},
	{
		input:                  "c:\\windows\\etc",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Unix",
		operatingSystem:        enums.Ubuntu,
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
	for _, testCase := range pathExistTestCaseWrappers {
		// Arrange
		fmt.Println(pathhelper.IsUnixCase(testCase.operatingSystem))
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			t.Skip("Unix tests ignored in Windows.")
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathExist(t, testCaseMessage, testCase)
	}
}

func TestIsPathExist_Unix(t *testing.T) {
	for _, testCase := range pathExistTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			t.Skip("Windows tests ignored in Unix.")
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
