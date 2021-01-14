package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var localTempPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetLocalTempPath",
		expected:               "/tmp",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetLocalTempPath",
		expected:               "C:\\Users\\Administrator\\AppData\\Roaming\\local\\temp",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetLocalTempPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range localTempPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetLocalTempPath, i)
	}
}

func TestGetLocalTempPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range localTempPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetLocalTempPath, i)
	}
}
