package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var userPicturesPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserPicturesPath",
		expected:               homePath + "/Pictures",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserPicturesPath",
		expected:               "C:\\Users\\Administrator\\Pictures",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserPicturesPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userPicturesPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserPicturesPath, i)
	}
}

func TestGetUserPicturesPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userPicturesPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserPicturesPath, i)
	}
}
