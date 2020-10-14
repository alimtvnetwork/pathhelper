package tests

import (
	"testing"
	"fmt"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type slugTestCaseWrapper struct{
	input, expectedMessage string
	expected  bool
}

var slugTestCaseWrappers = []slugTestCaseWrapper{
	{
		input: "",
		expected: false,
		expectedMessage: "should not be empty",
	},
	{
		input: "",
		expected: false,
		expectedMessage: "should not be nil",
	},
	{
		input: "xyz_",
		expected: true,
		expectedMessage: "true",
	},
	{
		input: "%&^2093073070271 b21 2987$#&^^&$(*&$(",
		expected: false,
		expectedMessage: "false",
	},
}

func TestIsSlug(t *testing.T) {
	for _, testCase := range slugTestCaseWrappers  {
		// Arrange
		testCaseMessage := fmt.Sprintf("[IsSlug] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.IsSlug(testCase.input)

			// Assert
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
