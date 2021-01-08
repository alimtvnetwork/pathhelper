package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var homepath = pathhelper.GetUserPath()

var sshGlobalTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetSSHGlobal",
		expected:               homepath + "/.ssh",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetSSHGlobal",
		expected:               "C:\\Users\\Administrator\\.ssh",
		operatingSystem:        enums.Windows,
	},
}

func TestGetSSHGlobal_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range sshGlobalTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetSSHGlobal, i)
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
