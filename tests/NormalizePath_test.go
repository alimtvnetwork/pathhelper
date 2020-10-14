package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type normalizePathTestCaseWrapper struct {
	input, expected, expectedMessage, operatingSystem string
}

var normalizePathTestCaseWrappers = []normalizePathTestCaseWrapper{
	{
		input:           "",
		expected:        "",
		expectedMessage: "empty return",
		operatingSystem: "Any OS",
	},
	{
		input:           "c:/windows/system32/etc",
		expected:        "c:/windows/system32/etc",
		expectedMessage: "non-empty return (c:/windows/system32/etc)",
		operatingSystem: "Any OS",
	},
	{
		input:           "c:\\windows//system32\\//etc",
		expected:        "c:\\windows\\system32\\etc",
		expectedMessage: "non-empty return (c:\\windows\\system32\\etc)",
		operatingSystem: "OS is windows",
	},
	{
		input:           "c:\\windows//system32\\//etc",
		expected:        "c:/windows/system32/etc",
		expectedMessage: "non-empty return (c:/windows/system32/etc)",
		operatingSystem: "OS other than windows",
	},
}

func TestNormalizePath(t *testing.T) {
	for _, testCase := range normalizePathTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("(%s) [NormalizePath] inputs (%s) expects (%s)", testCase.operatingSystem, testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.NormalizePath(testCase.input)

			// Assert
			if pathhelpercore.IsEmptyPath(testCase.input) {
				So(actual, ShouldBeEmpty)
			}

			if !pathhelpercore.IsEmptyPath(testCase.input) {
				So(actual, ShouldNotBeEmpty)

				if !pathhelper.IsWindows() {
					So(actual[2], ShouldEqual, testCase.expected[2])
				}

				if !pathhelper.IsWindows() {
					So(actual[3], ShouldEqual, testCase.expected[3])
				}
			}
		})
	}
}
