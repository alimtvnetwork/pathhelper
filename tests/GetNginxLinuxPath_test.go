package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"testing"
)

var getNginxLinuxTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetNginxLinuxPath",
	expected: "/etc/nginx/",
}

func TestGetNginxLinuxPath(t *testing.T) {
	pathTestCaseInternal_linux(t, getNginxLinuxTestCaseData, pathhelper.GetNginxLinuxPath)
}
