package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/enums"
)

type whichPrefixTestCaseWrapper struct{
	input, expectedMessage string
	expected enums.UriSchemes
}

var whichPrefixTestCaseWrappers =[]whichPrefixTestCaseWrapper{
	{
		input: "",
		expected: enums.UriUnknown,
		expectedMessage: "UriUnknown",
	},
	{
		input: "file:///",
		expected: enums.UriSchemePrefixStandard,
		expectedMessage: "UriSchemePrefixStandard",
	},
	{
		input: "file://",
		expected: enums.UriSchemePrefixTwoSlashes,
		expectedMessage: "UriSchemePrefixTwoSlashes",
	},
}

func TestWhichPrefix(t *testing.T){
	for _, testCase := range whichPrefixTestCaseWrappers{
		// Arrange
		testCaseMessage := fmt.Sprintf("[whichPrefix] inputs (%s) expects (%s)",testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func(){
			// Act
			actual := whichPrefix(testCase.input)

			// Assert
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
