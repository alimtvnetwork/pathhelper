package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var userVideosPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserVideosPath",
		expected:               homepath + "/Videos",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserVideosPath",
		expected:               "C:\\Users\\Administrator\\Videos",
		operatingSystem:        enums.Windows,
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
