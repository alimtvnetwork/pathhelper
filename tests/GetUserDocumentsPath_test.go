package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var userDocumentsPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserDocumentsPath",
		expected:               homepath + "/Documents",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserDocumentsPath",
		expected:               "C:\\Users\\Administrator\\Documents",
		operatingSystem:        enums.Windows,
	},
}

func TestGetUserDocumentsPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userDocumentsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserDocumentsPath, i)
	}
}

func TestGetUserDocumentsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userDocumentsPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserDocumentsPath, i)
	}
}
