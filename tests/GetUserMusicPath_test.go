package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

var userMusicPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserMusicPath",
		expected:               homepath + "/Music",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserMusicPath",
		expected:               "C:\\Users\\Administrator\\Music",
		operatingSystem:        enums.Windows,
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
