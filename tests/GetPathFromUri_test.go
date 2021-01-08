package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

type pathFromUriTestCaseWrapper struct {
	givenPath, expected, expectedMessage, operatingSystemMessage string
	isNormalize                                                  bool
	operatingSystem                                              enums.OperatingSystem
}

var pathFromUriTestCaseWrappers = []pathFromUriTestCaseWrapper{
	{
		givenPath:              "file://c:/windows/users/etc/more",
		isNormalize:            true,
		expected:               "c:\\windows\\users\\etc\\more",
		expectedMessage:        "c:\\windows\\users\\etc\\more",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		givenPath:              "c:\\windows\\users\\etc\\more",
		isNormalize:            true,
		expected:               "c:/windows/users/etc/more",
		expectedMessage:        "c:/windows/users/etc/more",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetPathFromUri_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range pathFromUriTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathFromURI] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.givenPath, testCase.isNormalize, testCase.expectedMessage)

		executeTestCaseForGetPathFromUri(t, testCaseMessage, testCase, i)
	}
}

func TestGetPathFromUri_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range pathFromUriTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathFromURI] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.givenPath, testCase.isNormalize, testCase.expectedMessage)

		executeTestCaseForGetPathFromUri(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetPathFromUri(t *testing.T, testCaseMessage string, testCase pathFromUriTestCaseWrapper, i int) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetPathFromUri(testCase.givenPath, testCase.isNormalize)

		// Assert
		Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
