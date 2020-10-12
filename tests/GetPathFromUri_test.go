package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type pathFromUriTestCaseWrapper struct {
	input, expected, expectedMessage, operatingSystemMessage string
	inputBool                                                bool
	operatingSystem                                          enums.OperatingSystem
}

var pathFromUriTestCaseWrappers = []pathFromUriTestCaseWrapper{
	{
		input:                  "file://c:/windows/users/etc/more",
		inputBool:              true,
		expected:               "c:\\windows\\users\\etc\\more",
		expectedMessage:        "c:\\windows\\users\\etc\\more",
		operatingSystemMessage: "OS is windows",
	},
	{
		input:                  "c:\\windows\\users\\etc\\more",
		inputBool:              false,
		expected:               "c:\\windows\\users\\etc\\more",
		expectedMessage:        "c:\\windows\\users\\etc\\more",
		operatingSystemMessage: "OS other than windows",
	},
}

func TestGetPathFromUri(t *testing.T) {
	// Arrange
	for _, testCase := range pathFromUriTestCaseWrappers {
		testCaseMessage := fmt.Sprintf("[getPathFromURI] inputs (%s, %v) expects (%s)", testCase.input, testCase.inputBool, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetPathFromUri(testCase.input, testCase.inputBool)

			// Assert
			So(actual, ShouldEqual, testCase.expected)
		})
	}
}
