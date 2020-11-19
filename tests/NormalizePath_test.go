package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

type normalizePathTestCaseWrapper struct {
	input, expected, expectedMessage, operatingSystemMessage string
	operatingSystem                                          enums.OperatingSystem
}

var normalizePathTestCaseWrappers = []normalizePathTestCaseWrapper{
	{
		input:                  "",
		expected:               "",
		expectedMessage:        "empty return",
		operatingSystemMessage: "Any OS",
	},
	{
		input:                  "c:/windows/system32/etc",
		expected:               "c:/windows/system32/etc",
		expectedMessage:        "non-empty return (c:/windows/system32/etc)",
		operatingSystemMessage: "Any OS",
	},
	{
		input:                  "c:\\windows//system32\\//etc",
		expected:               "c:\\windows\\system32\\etc",
		expectedMessage:        "non-empty return (c:\\windows\\system32\\etc)",
		operatingSystemMessage: "OS is windows",
		operatingSystem:        enums.Windows,
	},
	{
		input:                  "c:\\windows//system32\\//etc",
		expected:               "c:/windows/system32/etc",
		expectedMessage:        "non-empty return (c:/windows/system32/etc)",
		operatingSystemMessage: "OS other than windows",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestNormalizePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range normalizePathTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[NormalizePath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestNormalizePath(t, testCaseMessage, testCase)
	}
}

func TestNormalizePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range normalizePathTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestNormalizePath(t, testCaseMessage, testCase)
	}
}

func executeTestNormalizePath(t *testing.T, testCaseMessage string, testCase normalizePathTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.NormalizePath(testCase.input)

		// Assert
		So(actual, ShouldNotBeEmpty)
		So(actual, ShouldEqual, testCase.expected)
	})
}
