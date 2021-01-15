package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var userVideosPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserVideosPath",
		expected:               homePath + "/Videos",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserVideosPath",
		expected:               "C:\\Users\\Administrator\\Videos",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserVideosPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userVideosPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserVideosPath, i)
	}
}

func TestGetUserVideosPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userVideosPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserVideosPath, i)
	}
}
