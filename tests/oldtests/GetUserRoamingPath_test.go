package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/ostype"

	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var userRoamingPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "UserRoamingPath",
		expected:               homePath + "/Roaming",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "UserRoamingPath",
		expected:               "C:\\Users\\Administrator\\Roaming",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserRoamingPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userRoamingPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserRoamingPath, i)
	}
}

func TestGetUserRoamingPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userRoamingPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserRoamingPath, i)
	}
}
