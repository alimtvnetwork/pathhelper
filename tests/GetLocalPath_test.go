package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var localPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "LocalPath",
		expected:               "/home/a",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "LocalPath",
		expected:               "C:\\Users\\Administrator\\AppData\\Roaming\\Local",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetLocalPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range localPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.LocalPath, i)
	}
}

func TestGetLocalPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range localPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.LocalPath, i)
	}
}
