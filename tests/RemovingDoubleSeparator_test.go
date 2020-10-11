package tests

import (
	"testing"
	"fmt"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type removingDoubleSeparatorTestCaseWrapper struct {
	input, expected, expectedMessage string
}

var removingDoubleSeparatorTestCaseWrappers = []removingDoubleSeparatorTestCaseWrapper{
	{
		input: "c:\\\\win",
		expected: "c:\\win",
		expectedMessage: "non-empty, non-nil, return of (c:\\win)",
	},
	{
		input: "home//user",
		expected: "home/user",
		expectedMessage: "non-empty, non-nil, return of (home/user)",
	},
}

func TestRemovingDoubleSeparator(t *testing.T) {
	for _, testCase := range removingDoubleSeparatorTestCaseWrappers{
		// Arrange
		testCaseMessage := fmt.Sprintf("[RemovingDoubleSeparator] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.RemovingDoubleSeparator(testCase.input)

			// Assert
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}