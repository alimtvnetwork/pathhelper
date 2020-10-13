package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type combinedPathUsingConfigInternalTestCaseWrapper struct {
	inputPathConfig                                   *pathhelpercore.PathConfig
	inputPaths                                        []string
	expected, expectedMessage, operatingSystemMessage string
	operatingSystem                                   enums.OperatingSystem
}

var combinedPathUsingConfigInternalTestCaseWrappers = []combinedPathUsingConfigInternalTestCaseWrapper{
	{
		inputPathConfig:        &pathhelpercore.PathConfig{},
		inputPaths:             []string{"C:\\", "somethingelse\\", "etc"},
		expected:               "C:\\somethingelse\\etc",
		expectedMessage:        "C:\\somethingelse\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputPathConfig:        nil,
		inputPaths:             []string{"C:", "somethingelse", "etc"},
		expected:               "C:\\somethingelse\\etc",
		expectedMessage:        "C:\\somethingelse\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        enums.Windows,
	},
	{
		inputPathConfig:        &pathhelpercore.PathConfig{IsNormalize: true},
		inputPaths:             []string{"home\\", "\\somethingelse\\", "\\etc"},
		expected:               "home/somethingelse/etc",
		expectedMessage:        "home/somethingelse/etc",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        enums.Ubuntu,
	},
	{
		inputPathConfig: &pathhelpercore.PathConfig{},
		inputPaths:      []string{},
		expected:        "Empty paths given",
		expectedMessage: "should panic with Empty paths given",
	},
}

func TestGetCombinedPathUsingConfigInternal_Windows(t *testing.T) {
	if !IsWindows() {
		t.Skip("Windows tests ignored in Unix.")
	}

	for _, testCase := range combinedPathUsingConfigInternalTestCaseWrappers {
		// Arrange
		if IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfigInternal]  inputs (pathConfig: %v, paths: %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPathConfig, testCase.inputPaths, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfigInternal(t, testCaseMessage, testCase)
	}
}

func TestGetCombinedPathUsingConfigInternal_Unix(t *testing.T) {
	if IsWindows() {
		t.Skip("Unix tests ignored in Windows.")
	}

	for _, testCase := range combinedPathUsingConfigInternalTestCaseWrappers {
		// Arrange
		if IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfigInternal]  inputs (pathConfig: %v, paths: %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPathConfig, testCase.inputPaths, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfigInternal(t, testCaseMessage, testCase)
	}
}

func executeTestCaseForGetCombinedPathUsingConfigInternal(t *testing.T, testCaseMessage string, testCase combinedPathUsingConfigInternalTestCaseWrapper) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := getCombinedPathUsingConfigInternal(testCase.inputPathConfig, testCase.inputPaths)

		// Assert
		if len(testCase.inputPaths) == 0 {
			So(actual, ShouldPanicWith, "Empty paths given")
		}

		So(actual, ShouldEqual, testCase.expected)
	})
}
