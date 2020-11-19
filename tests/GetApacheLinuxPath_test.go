package tests

import (
	"gitlab.com/evatix-go/pathhelper"
	"testing"
)

var getApacheLinuxPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetApacheLinuxPath",
	expected: "/etc/apache/",
}

func TestGetApacheLinuxPath(t *testing.T) {
	pathTestCaseInternal_linux(t, getApacheLinuxPathTestCaseData, pathhelper.GetApacheLinuxPath)
}
