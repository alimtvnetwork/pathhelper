package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
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

	for _, testCase := range fontsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetFontsPath)
	}
}

func TestGetFontsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range fontsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetFontsPath)
	}
}
