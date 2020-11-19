package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var rootTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetRoot",
		expected:               "/",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetRoot",
		expected:               "C:\\",
		operatingSystem:        enums.Windows,
	},
}

func TestGetRoot_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range rootTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetRoot)
	}
}

func TestGetRoot_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range rootTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetRoot)
	}
}
