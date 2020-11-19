package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

var userPicturesPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetUserPicturesPath",
		expected:               homepath + "/Pictures",
		operatingSystem:        enums.Ubuntu,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetUserPicturesPath",
		expected:               "C:\\Users\\Administrator\\Pictures",
		operatingSystem:        enums.Windows,
	},
}

func TestGetUserPicturesPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for _, testCase := range userPicturesPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserPicturesPath)
	}
}

func TestGetUserPicturesPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for _, testCase := range userPicturesPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetUserPicturesPath)
	}
}
