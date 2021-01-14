package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var userMusicPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserMusicPath",
		expected:               homePath + "/Music",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserMusicPath",
		expected:               "C:\\Users\\Administrator\\Music",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserMusicPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userMusicPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserMusicPath, i)
	}
}

func TestGetUserMusicPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userMusicPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserMusicPath, i)
	}
}
