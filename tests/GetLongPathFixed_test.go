package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper"
	"testing"
)

type longPathFixedTestCaseDataWrapper struct {
	input, expected string
}

var longPathFixedTestCaseDataWrappers = []longPathFixedTestCaseDataWrapper{
	{
		input:    "",
		expected: "",
	},
	{
		input:    "\\\\?\\UNC\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\UNC\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\UNC\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
}

func TestGetLongPathFixed(t *testing.T) {
	for _, testCase := range longPathFixedTestCaseDataWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetLongPathFixed] inputs (%s) expects (%s)", testCase.input, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetLongPathFixed(testCase.input)

			// Assert
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
