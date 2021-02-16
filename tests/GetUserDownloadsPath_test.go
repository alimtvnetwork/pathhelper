package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var userDownloadsPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "UserDownloadsPath",
		expected:               homePath + "/Downloads",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "UserDownloadsPath",
		expected:               "C:\\Users\\Administrator\\Downloads",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserDownloadsPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userDownloadsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserDownloadsPath, i)
	}
}

func TestGetUserDownloadsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userDownloadsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserDownloadsPath, i)
	}
}
