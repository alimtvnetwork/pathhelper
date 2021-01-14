package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
)

var servicesPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetServicesPath",
		expected:               "/etc/systemd/system",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetServicesPath",
		expected:               "C:\\Windows\\System32\\drivers\\etc\\services",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetServicesPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range servicesPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetServicesPath, i)
	}
}

func TestGetServicesPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range servicesPathTestCaseDataWrappers {
		// Arrange
		if pathhelper.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, pathhelper.GetServicesPath, i)
	}
}
