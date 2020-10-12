package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type numberTestCaseWrapper struct {
	input, expectedMessage string
	expected               bool
}

var numberTestCaseWrappers = []numberTestCaseWrapper{
	{
		input:           "-521",
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           "521",
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           "-5something21",
		expected:        false,
		expectedMessage: "false",
	},
	{
		input:           "-521/something",
		expected:        false,
		expectedMessage: "false",
	},
}

func TestIsNumber(t *testing.T) {
	for _, testCase := range numberTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[isNumber] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := isNumber(testCase.input)

			// Assert
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
