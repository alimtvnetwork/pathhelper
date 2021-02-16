package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var etcPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "EtcPath",
		expected:               "/etc",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "EtcPath",
		expected:               "C:\\Windows\\System32\\drivers\\etc",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetEtcPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range etcPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.EtcPath, i)
	}
}

func TestGetEtcPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range etcPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.EtcPath, i)
	}
}
