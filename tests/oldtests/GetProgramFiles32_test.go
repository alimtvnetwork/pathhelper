package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var programFiles32TestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetProgramFiles32",
		expected:               "",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetProgramFiles32",
		expected:               "C:\\\\Program Files",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetProgramFiles32_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range programFiles32TestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GetProgramFiles32, i)
	}
}

func TestGetProgramFiles32_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range programFiles32TestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GetProgramFiles32, i)
	}
}
