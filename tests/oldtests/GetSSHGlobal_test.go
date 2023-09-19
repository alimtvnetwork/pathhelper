package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var homePath = knowndirget.UserPath()

var sshGlobalTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "SSHGlobal",
		expected:               homePath + "/.ssh",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "SSHGlobal",
		expected:               "C:\\Users\\Administrator\\.ssh",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetSSHGlobal_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range sshGlobalTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(
			t,
			testCase,
			knowndirget.SSHGlobal,
			i)
	}
}

func TestGetSSHGlobal_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range sshGlobalTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.SSHGlobal, i)
	}
}
