package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var userMusicPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "UserMusicPath",
		expected:               homePath + "/Music",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "UserMusicPath",
		expected:               "C:\\Users\\Administrator\\Music",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserMusicPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userMusicPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserMusicPath, i)
	}
}

func TestGetUserMusicPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userMusicPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserMusicPath, i)
	}
}
