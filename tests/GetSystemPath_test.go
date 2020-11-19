package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
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

	for _, testCase := range systemPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetSystemPath)
	}
}

func TestGetSystemPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range systemPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetSystemPath)
	}
}
