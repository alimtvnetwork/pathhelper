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
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		input:                  "c:\\windows\\users\\etc\\more",
		inputBool:              true,
		expected:               "c:/windows/users/etc/more",
		expectedMessage:        "c:/windows/users/etc/more",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetPathFromUri_Windows(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range pathFromUriTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathFromURI] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.inputBool, testCase.expectedMessage)

		executeTestCaseForGetPathFromUri(t, testCaseMessage, testCase)
	}
}

func TestGetPathFromUri_Unix(t *testing.T) {
	if pathhelper.IsWindows() {
		t.Skip("Unix tests ignored in Windows.")
	}

	for _, testCase := range pathFromUriTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathFromURI] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.inputBool, testCase.expectedMessage)

		executeTestCaseForGetPathFromUri(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetPathFromUri(t *testing.T, testCaseMessage string, testCase pathFromUriTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetPathFromUri(testCase.input, testCase.inputBool)

		// Assert
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
