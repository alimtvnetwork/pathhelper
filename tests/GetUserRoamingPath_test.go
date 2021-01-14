package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var userRoamingPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserRoamingPath",
		expected:               homePath + "/Roaming",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserRoamingPath",
		expected:               "C:\\Users\\Administrator\\Roaming",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserRoamingPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userRoamingPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserRoamingPath, i)
	}
}

func TestGetUserRoamingPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userRoamingPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserRoamingPath, i)
	}
}
