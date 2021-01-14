package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var homePath = pathhelper.GetUserPath()

var sshGlobalTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetSSHGlobal",
		expected:               homePath + "/.ssh",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetSSHGlobal",
		expected:               "C:\\Users\\Administrator\\.ssh",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetSSHGlobal_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range sshGlobalTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(
			t,
			testCase,
			pathhelper.GetSSHGlobal,
			i)
	}
}

func TestGetSSHGlobal_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range sshGlobalTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetSSHGlobal, i)
	}
}
