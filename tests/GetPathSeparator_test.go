package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/constants"
)

type getPathSeparatorTestCaseWrapper struct {
	operatingSystem, expected, expectedMessage string
}

var pathSeparatorTestCaseWrappers = []getPathSeparatorTestCaseWrapper{
	{
		operatingSystem: "Os is windows",
		expected: constants.BackSlash,
		expectedMessage: constants.BackSlash,
	},
	{
		operatingSystem: "OS other than windows",
		expected: constants.ForwardSlash,
		expectedMessage: constants.ForwardSlash,
	},
}

func TestGetPathSeparator(t *testing.T) {
	for _, testCase := range pathSeparatorTestCaseWrappers{
		// Arrange
		testCaseMessage := fmt.Sprintf("(%s) [GetPathSeparator] inputs () expects (%s)", testCase.operatingSystem, testCase.expectedMessage)

		Convey(testCaseMessage, t, func(){
			// Act
			actual := pathhelper.GetPathSeparator()

			// Assert
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldNotBeNil)

			if pathhelper.IsWindows() {
				So(actual[0], ShouldEqual, testCase.expected[0])
			}

			if !pathhelper.IsWindows() {
				So(actual[1], ShouldEqual, testCase.expected[1])
			}
		})
	}
}
