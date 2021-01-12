package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

var localPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetLocalPath",
		expected:               "/home/a",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetLocalPath",
		expected:               "C:\\Users\\Administrator\\AppData\\Roaming\\Local",
		operatingSystem:        enums.Windows,
	},
}

func TestGetLocalPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range localPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetLocalPath, i)
	}
}

func TestGetLocalPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range localPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetLocalPath, i)
	}
}
