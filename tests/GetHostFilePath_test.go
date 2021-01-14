package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var hostFilePathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetHostFilePath",
		expected:               "/etc/hosts",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetHostFilePath",
		expected:               "C:\\Windows\\System32\\drivers\\etc\\hosts",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetHostFilePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range hostFilePathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetHostFilePath, i)
	}
}

func TestGetHostFilePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range hostFilePathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetHostFilePath, i)
	}
}
