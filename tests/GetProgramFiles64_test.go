package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var programFiles64TestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetProgramFiles64",
		expected:               "",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetProgramFiles64",
		expected:               "C:\\\\Program Files (x86)",
		operatingSystem:        enums.Windows,
	},
}

func TestGetProgramFiles64_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range programFiles64TestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetProgramFiles64, i)
	}
}

func TestGetProgramFiles64_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range programFiles64TestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetProgramFiles64, i)
	}
}
