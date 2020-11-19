package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var userDownloadsPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserDownloadsPath",
		expected:               homepath + "/Downloads",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserDownloadsPath",
		expected:               "C:\\Users\\Administrator\\Downloads",
		operatingSystem:        enums.Windows,
	},
}

func TestGetUserDownloadsPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range userDownloadsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserDownloadsPath)
	}
}

func TestGetUserDownloadsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range userDownloadsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserDownloadsPath)
	}
}
