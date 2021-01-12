package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getModulesAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModulesAvailable",
	expected: "/etc/nginx/modules-available",
}

func TestGetModulesAvailable(t *testing.T) {
	getPathTestCommonMethod_linux(t, getModulesAvailablePathTestCaseData, nginxlinuxpath.GetModulesAvailable)
}
