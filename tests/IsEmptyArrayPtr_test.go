package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
	"testing"
)

type isEmptyArrayPtrTestCaseWrapper struct {
	input           []*string
	expected        bool
	expectedMessage string
}

var (
	hello = "hello"
	world = "world"
)
var isEmptyArrayPtrTestCaseWrappers = []isEmptyArrayPtrTestCaseWrapper{
	{
		input:           []*string{},
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           []*string{&hello, &world},
		expected:        false,
		expectedMessage: "false",
	},
}

func TestIsEmptyArrayPtr(t *testing.T) {
	for _, testCase := range isEmptyArrayPtrTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[IsEmptyArrayPtr] inputs (%v) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelpercore.IsEmptyArrayPtr(testCase.input)

			// Assert
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
