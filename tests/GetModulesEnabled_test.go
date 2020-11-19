package tests

import (
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
	"testing"
)

var getModulesEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModulesEnabled",
	expected: "/etc/nginx/modules-enabled",
}

func TestGetModulesEnabled(t *testing.T) {
	pathTestCaseInternal_linux(t, getModulesEnabledPathTestCaseData, nginxlinuxpath.GetModulesEnabled)
}
