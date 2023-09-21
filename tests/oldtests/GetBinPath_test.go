package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var binPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		expected:               "/usr/bin",
		operatingSystemMessage: "Unix OS",
		funcName:               "BinPath",
		osUserName:             "",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "BinPath",
		osUserName:             "\\Administrator\\",
		expected:               "C:\\Users\\Administrator\\bin",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetBinPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range binPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.BinPath, i)
	}
}

func TestGetBinPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range binPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.BinPath, i)
	}
}
