package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
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
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserDownloadsPath, i)
	}
}

func TestGetUserDownloadsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userDownloadsPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserDownloadsPath, i)
	}
}
