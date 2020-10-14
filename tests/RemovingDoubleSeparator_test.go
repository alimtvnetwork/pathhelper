package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type removingDoubleSeparatorTestCaseWrapper struct {
	input, expected, expectedMessage string
	enums.OperatingSystem
}

var removingDoubleSeparatorTestCaseWrappers = []removingDoubleSeparatorTestCaseWrapper{
	{
		input:           "c:\\\\win",
		expected:        "c:\\win",
		expectedMessage: "non-empty, non-nil, return of (c:\\win)",
		OperatingSystem: enums.Windows,
	},
	{
		input:           "c:\\\\\\win/drive",
		expected:        "c:\\win\\drive",
		expectedMessage: "non-empty, non-nil, return of (c:\\win\\drive)",
		OperatingSystem: enums.Windows,
	},
	{
		input:           "home//user",
		expected:        "home/user",
		expectedMessage: "non-empty, non-nil, return of (home/user)",
		OperatingSystem: enums.Ubuntu,
	},
}

func TestRemoveAndFixDoubleSeparatorToOsSeparator_Windows(t *testing.T) {
	UnixTestsSkipOnWindows(t)

	for _, testCase := range removingDoubleSeparatorTestCaseWrappers {
		if pathhelper.IsUnixCase(testCase.OperatingSystem) {
			continue
		}

		// Arrange
		testCaseMessage := fmt.Sprintf("[RemoveAndFixDoubleSeparatorToOsSeparator] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		// Act , Assert
		Convey(testCaseMessage, t, func() {
			internalTestRemoveAndFixDoubleSeparatorToOsSeparatorActAndAssert(testCase)
		})
	}
}



func TestRemoveAndFixDoubleSeparatorToOsSeparator_Unix(t *testing.T) {
	WindowsTestsSkipOnUnix(t)

	for _, testCase := range removingDoubleSeparatorTestCaseWrappers {
		if pathhelper.IsWindowsCase(testCase.OperatingSystem) {
			continue
		}

		// Arrange
		testCaseMessage := fmt.Sprintf("[RemoveAndFixDoubleSeparatorToOsSeparator] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		// Act , Assert
		Convey(testCaseMessage, t, func() {
			internalTestRemoveAndFixDoubleSeparatorToOsSeparatorActAndAssert(testCase)
		})
	}
}

func internalTestRemoveAndFixDoubleSeparatorToOsSeparatorActAndAssert(
	testCase removingDoubleSeparatorTestCaseWrapper) {
	// Act
	actual := pathhelper.RemoveAndFixDoubleSeparatorToOsSeparator(testCase.input)

	// Assert
	So(actual, ShouldNotBeEmpty)
	So(actual, ShouldNotBeNil)
	So(actual, ShouldEqual, testCase.expected)
}
