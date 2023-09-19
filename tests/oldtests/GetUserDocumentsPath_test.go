package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var userDocumentsPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "UserDocumentsPath",
		expected:               homePath + "/Documents",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "UserDocumentsPath",
		expected:               "C:\\Users\\Administrator\\Documents",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserDocumentsPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userDocumentsPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserDocumentsPath, i)
	}
}

func TestGetUserDocumentsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userDocumentsPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserDocumentsPath, i)
	}
}
