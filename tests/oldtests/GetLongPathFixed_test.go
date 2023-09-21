package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/pathhelper/normalize"
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
		input:    "\\\\?\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
}

func TestGetLongPathFixed_NonNormalized(t *testing.T) {
	options := normalize.Options{
		IsNormalize:        false,
		IsLongPathFix:      true,
		IsForceLongPathFix: true,
	}

	for i, testCase := range longPathFixedTestCaseDataWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetLongPathFixed] inputs (%s) expects (%s)", testCase.input, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := options.FixPath(testCase.input)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				if actual != testCase.expected {
					fmt.Println("Failed test case :", i)
				}

				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}

var longPathFixedTestCaseDataWrappersNormalized = []longPathFixedTestCaseDataWrapper{
	{
		input:    "",
		expected: "",
	},
	{
		input:    "\\\\?\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
}

func TestGetLongPathFixed_Normalized(t *testing.T) {
	options := normalize.Options{
		IsNormalize:        true,
		IsLongPathFix:      true,
		IsForceLongPathFix: true,
	}

	for i, testCase := range longPathFixedTestCaseDataWrappersNormalized {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetLongPathFixed] inputs (%s) expects (%s)", testCase.input, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := options.FixPath(testCase.input)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				if actual != testCase.expected {
					fmt.Println("Failed test case :", i)
				}

				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
