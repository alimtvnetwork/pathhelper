package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

var binPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetBinPath",
		expected:               "/usr/bin",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetBinPath",
		expected:               "C:\\Users\\Administrator\\bin",
		operatingSystem:        enums.Windows,
	},
}

func TestGetBinPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range binPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetBinPath, i)
	}
}

func TestGetBinPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range binPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetBinPath, i)
	}
}
