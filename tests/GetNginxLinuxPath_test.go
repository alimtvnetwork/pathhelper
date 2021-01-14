package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
)

var getNginxLinuxTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetNginxLinuxPath",
	expected: "/etc/nginx/",
}

func TestGetNginxLinuxPath(t *testing.T) {
	getPathTestCommonMethodLinux(t, getNginxLinuxTestCaseData, pathhelper.GetNginxLinuxPath)
}
