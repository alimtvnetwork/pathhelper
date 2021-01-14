package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var programDataTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetProgramData",
		expected:               "",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetProgramData",
		expected:               "C:\\\\Program Data",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetProgramData_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range programDataTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetProgramData, i)
	}
}

func TestGetProgramData_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range programDataTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetProgramData, i)
	}
}
