package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var gitGlobalPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetGitGlobal",
		expected:               "XDG_CONFIG_HOME/git/config",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetGitGlobal",
		expected:               "C:\\Users\\Administrator\\.gitconfig",
		operatingSystem:        enums.Windows,
	},
}

func TestGetGitGlobal_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range gitGlobalPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetGitGlobal)
	}
}

func TestGetGitGlobal_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range gitGlobalPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetGitGlobal)
	}
}
