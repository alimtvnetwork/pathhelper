package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type pathSeparatorTestCaseWrapper struct {
	operatingSystem                                   enums.OperatingSystem
	operatingSystemMessage, expected, expectedMessage string
}

var pathSeparatorTestCaseWrappers = []pathSeparatorTestCaseWrapper{
	{
		operatingSystem:        enums.Windows,
		operatingSystemMessage: "Os is windows",
		expected:               constants.BackSlash,
		expectedMessage:        constants.BackSlash,
	},
	{
		operatingSystem:        enums.Ubuntu,
		operatingSystemMessage: "Unix os",
		expected:               constants.ForwardSlash,
		expectedMessage:        constants.ForwardSlash,
	},
}

func TestGetPathSeparator_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range pathSeparatorTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathSeparator] inputs () expects (%s)", testCase.operatingSystemMessage, testCase.expectedMessage)

		executeTestCaseForGetPathSeparator(t, testCaseMessage, testCase)
	}
}

func TestGetPathSeparator_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range pathSeparatorTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathSeparator] inputs () expects (%s)", testCase.operatingSystemMessage, testCase.expectedMessage)

		executeTestCaseForGetPathSeparator(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetPathSeparator(t *testing.T, testCaseMessage string, testCase pathSeparatorTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := constants.PathSeparator

		// Assert
		So(actual, ShouldNotBeEmpty)
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
