package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
)

type pathExistTestCaseWrapper struct {
	input, expectedMessage, operatingSystemMessage string
	expected                                       bool
	operatingSystem                                ostype.Variation
}

var pathExistTestCaseWrappers = []pathExistTestCaseWrapper{
	{
		input:                  "",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "any OS",
	},
	{
		input:                  "c:\\sampleSomething",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Windows",
		operatingSystem:        ostype.Windows,
	},
	{
		input:                  "~/home",
		expected:               false,
		expectedMessage:        "true",
		operatingSystemMessage: "OS is Unix",
		operatingSystem:        ostype.Linux,
	},
}

func TestIsPathExist_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range pathExistTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathExist(t, testCaseMessage, testCase, i)
	}
}

func TestIsPathExist_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range pathExistTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathExist(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForIsPathExist(t *testing.T, testCaseMessage string, testCase pathExistTestCaseWrapper, i int) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.IsPathExist(testCase.input)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
