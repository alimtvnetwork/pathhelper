package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/expandpath"
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
	for i, testCase := range getVariableTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[getVariable] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := normalize.GetVariables(testCase.input)

			// Assert
			Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldHaveSameTypeAs, []string{})

				length := len(actual)

				if length != 0 {
					for i := 0; i < length; i++ {
						So(actual[i], ShouldEqual, testCase.expected[i])
					}
				}
			})
		})
	}
}
