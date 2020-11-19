package tests

import (
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
	"testing"
)

var getModulesAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModulesAvailable",
	expected: "/etc/nginx/modules-available",
}

func TestGetModulesAvailable(t *testing.T) {
	pathTestCaseInternal_linux(t, getModulesAvailablePathTestCaseData, nginxlinuxpath.GetModulesAvailable)
}
