package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

var systemPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetSystemPath",
		expected:               "/etc/systemd/system",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetSystemPath",
		expected:               "C:\\Windows",
		operatingSystem:        enums.Windows,
	},
}

func TestGetSystemPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range systemPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetSystemPath, i)
	}
}

func TestGetSystemPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range systemPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetSystemPath, i)
	}
}
