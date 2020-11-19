package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var localTempPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetLocalTempPath",
		expected:               "/home/a/temp",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetLocalTempPath",
		expected:               "C:\\Users\\Administrator\\AppData\\Roaming\\local\\temp",
		operatingSystem:        enums.Windows,
	},
}

func TestGetLocalTempPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range localTempPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetLocalTempPath)
	}
}

func TestGetLocalTempPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range localTempPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetLocalTempPath)
	}
}
