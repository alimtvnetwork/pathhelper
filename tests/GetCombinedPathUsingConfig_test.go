package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type combinedPathUsingConfigTestCaseWrapper struct {
	inputConfig                                       *pathhelpercore.PathConfig
	inputPaths1, inputPaths2, inputPaths3             string
	expected, expectedMessage, operatingSystemMessage string
	operatingSystem                                   enums.OperatingSystem
}

var combinedPathUsingConfigTestCaseWrappers = []combinedPathUsingConfigTestCaseWrapper{
	{
		inputConfig:            &pathhelpercore.PathConfig{Separator: constants.PathSeparator},
		inputPaths1:            "something",
		inputPaths2:            "more",
		inputPaths3:            "etc",
		expected:               "something\\more\\etc",
		expectedMessage:        "something\\more\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputConfig:            &pathhelpercore.PathConfig{Separator: constants.PathSeparator},
		inputPaths1:            "something",
		inputPaths2:            "more",
		inputPaths3:            "etc",
		expected:               "something/more/etc",
		expectedMessage:        "something/more/etc",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
}

func TestGetCombinedPathUsingConfig_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range combinedPathUsingConfigTestCaseWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfig] inputs (%v, %s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputConfig, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfig(t, testCaseMessage, testCase)
	}
}

func TestGetCombinedPathUsingConfig_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range combinedPathUsingConfigTestCaseWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfig] inputs (%v, %s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputConfig, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfig(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetCombinedPathUsingConfig(
	t *testing.T, testCaseMessage string, testCase combinedPathUsingConfigTestCaseWrapper,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetCombinedPathUsingConfig(testCase.inputConfig, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3)

		// Assert
		So(actual, ShouldNotBeNil)
		So(actual, ShouldEqual, testCase.expected)
	})
}
