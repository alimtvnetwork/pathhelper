package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper"
)

type getSlugTestCaseWrapper struct {
	inputPath       string
	inputSeparator  string
	expected        string
	expectedMessage string
}

var slugWrappers = []getSlugTestCaseWrapper{
	{
		inputPath:       "",
		inputSeparator:  constants.Underscore,
		expected:        "",
		expectedMessage: "empty",
	},
	{
		inputPath:       "_20971-b21-2987_",
		inputSeparator:  constants.Underscore,
		expected:        "_20971-b21-2987_",
		expectedMessage: "non-empty return",
	},
	{
		inputPath:       "_20971-b21-2987_",
		inputSeparator:  constants.Underscore,
		expected:        "_20971-b21-2987_",
		expectedMessage: "_20971-b21-2987_",
	},
	{
		inputPath:       "%&^2093073070271 b21 2987$#&^^&$(*&$(",
		inputSeparator:  constants.Underscore,
		expected:        "_2093073070271_b21_2987_",
		expectedMessage: "_2093073070271_b21_2987_",
	},
	{
		inputPath:       "%&^2093*73070271 b21 2987$#&^^&$(*&$(",
		inputSeparator:  constants.Dash,
		expected:        "-2093-73070271-b21-2987-",
		expectedMessage: "-2093-73070271-b21-2987-",
	},
}

func TestGetSlug(t *testing.T) {
	for i, testCase := range slugWrappers {
		{
			// Arrange
			testCaseMessage := fmt.Sprintf("[getWindowsBuild] inputs (%s, %s) expects (%s)", testCase.inputPath, testCase.inputSeparator, testCase.expectedMessage)

			Convey(testCaseMessage, t, func() {
				// Act
				actual := pathhelper.GetSlug(testCase.inputPath, testCase.inputSeparator)

				// Assert
				Convey(GetAssertMessage(actual, testCase.expected, i), func() {
					if testCase.inputPath == "" {
						So(actual, ShouldBeEmpty)
					}

					if testCase.inputPath != "" {
						So(actual, ShouldNotBeEmpty)
						So(actual, ShouldEqual, testCase.expected)
					}
				})
			})
		}
	}
}
