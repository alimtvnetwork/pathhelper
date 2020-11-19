package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
	"testing"
)

type isEmptyArrayTestCaseWrapper struct {
	input           []string
	expected        bool
	expectedMessage string
}

var isEmptyArrayTestCaseWrappers = []isEmptyArrayTestCaseWrapper{
	{
		input:           []string{},
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           []string{"hello", "world"},
		expected:        false,
		expectedMessage: "false",
	},
}

func TestIsEmptyArray(t *testing.T) {
	for i, testCase := range isEmptyArrayTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[IsEmptyArray] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelpercore.IsEmptyArray(testCase.input)

			// Assert
			Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
