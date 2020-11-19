package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
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
	for i, testCase := range pathFromEnvVariableTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[PathFromEnvVariable] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.PathFromEnvVariable(testCase.input)

			// Assert
			Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() { //todo check equality
				if pathhelpercore.IsEmptyPath(testCase.input) {
					So(actual, ShouldBeEmpty)
				}

				if !pathhelpercore.IsEmptyPath(testCase.input) {
					So(actual, ShouldNotBeEmpty)
				}
			})
		})

	}
}
