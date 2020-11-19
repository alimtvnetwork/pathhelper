package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
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

	for _, testCase := range etcPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetEtcPath)
	}
}

func TestGetEtcPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range etcPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetEtcPath)
	}
}
