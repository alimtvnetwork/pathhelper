package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getModulesEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModulesEnabled",
	expected: "/etc/nginx/modules-enabled",
}

func TestGetModulesEnabled(t *testing.T) {
	getPathTestCommonMethod_linux(t, getModulesEnabledPathTestCaseData, nginxlinuxpath.GetModulesEnabled)
}
