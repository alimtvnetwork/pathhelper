package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

var etcPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetEtcPath",
		expected:               "/etc",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetEtcPath",
		expected:               "C:\\Windows\\System32\\drivers\\etc",
		operatingSystem:        enums.Windows,
	},
}

func TestGetEtcPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range etcPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetEtcPath, i)
	}
}

func TestGetEtcPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range etcPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetEtcPath, i)
	}
}
