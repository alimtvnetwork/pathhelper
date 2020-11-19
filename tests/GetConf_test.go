package tests

import (
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
	"testing"
)

var getConfPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConf",
	expected: "/etc/nginx/conf.d",
}

func TestGetConf(t *testing.T) {
	pathTestCaseInternal_linux(t, getConfPathTestCaseData, nginxlinuxpath.GetConf)
}
