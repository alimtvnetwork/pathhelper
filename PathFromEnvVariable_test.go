package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type pathFromEnvVariableTestCaseWrapper struct {
	input, expected, expectedMessage string
}

var pathFromEnvVariableTestCaseWrappers = []pathFromEnvVariableTestCaseWrapper{
	{
		input:           "",
		expected:        "",
		expectedMessage: "",
	},
	{
		input:           "$home $sys hello world $what",
		expected:        "",
		expectedMessage: "",
	},
	{
		input:           "$ComSpec hello $no",
		expected:        "",
		expectedMessage: "",
	},
}

func TestPathFromEnvVariable(t *testing.T) {
	for _, testCase := range pathFromEnvVariableTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[getWindowsBuild] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := PathFromEnvVariable(testCase.input)

			// Assert
			if pathhelpercore.IsEmptyPath(testCase.input) {
				So(actual, ShouldBeEmpty)
			}

			if !pathhelpercore.IsEmptyPath(testCase.input) {
				So(actual, ShouldNotBeEmpty)
				// if os.LookupEnv("home") { // how to check
				// 	So(actual, ShouldEqual, testCase.expected)
				// }
			}
		})

	}
}
