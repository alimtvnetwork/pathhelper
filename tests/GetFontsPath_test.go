package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

var fontsPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetFontsPath",
		expected:               "/usr/share/fonts",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetFontsPath",
		expected:               "C:\\Windows\\Fonts",
		operatingSystem:        enums.Windows,
	},
}

func TestGetFontsPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range fontsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetFontsPath, i)
	}
}

func TestGetFontsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range fontsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetFontsPath, i)
	}
}
