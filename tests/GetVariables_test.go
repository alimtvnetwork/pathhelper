package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type getVariableTestCaseWrapper struct {
	input, expectedMessage string
	expected               []string
}

var getVariableTestCaseWrappers = []getVariableTestCaseWrapper{
	{
		input:           "$home Hello World $whatever $sample",
		expected:        []string{"home", "whatever", "sample"},
		expectedMessage: "[]string{home, whatever, sample}",
	},
	{
		input:           "home Hello World",
		expected:        nil,
		expectedMessage: "nil",
	},
	{
		input:           "",
		expected:        nil,
		expectedMessage: "nil",
	},
}

func TestGetVariables(t *testing.T) {
	for _, testCase := range getVariableTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[getVariable] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetVariables(testCase.input)

			// Assert
			So(actual, ShouldHaveSameTypeAs, []string{})
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
