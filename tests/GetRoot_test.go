package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var rootTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetRoot",
		expected:               "/",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetRoot",
		expected:               "C:\\",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetRoot_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range rootTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GetRoot, i)
	}
}

func TestGetRoot_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range rootTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GetRoot, i)
	}
}
