package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var userRoamingPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserRoamingPath",
		expected:               homepath + "/Roaming",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserRoamingPath",
		expected:               "C:\\Users\\Administrator\\Roaming",
		operatingSystem:        enums.Windows,
	},
}

func TestGetUserRoamingPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range userRoamingPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserRoamingPath)
	}
}

func TestGetUserRoamingPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range userRoamingPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserRoamingPath)
	}
}
