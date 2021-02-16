package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var programDataTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "ProgramData",
		expected:               "",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "ProgramData",
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

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.ProgramData, i)
	}
}

func TestGetProgramData_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range programDataTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.ProgramData, i)
	}
}
