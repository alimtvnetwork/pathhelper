package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var hostFilePathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetHostFilePath",
		expected:               "/etc/hosts",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetHostFilePath",
		expected:               "C:\\Windows\\System32\\drivers\\etc\\hosts",
		operatingSystem:        enums.Windows,
	},
}

func TestGetHostFilePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range hostFilePathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetHostFilePath)
	}
}

func TestGetHostFilePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range hostFilePathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetHostFilePath)
	}
}
