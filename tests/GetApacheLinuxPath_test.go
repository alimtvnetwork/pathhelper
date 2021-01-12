package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
)

var getApacheLinuxPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetApacheLinuxPath",
	expected: "/etc/apache/",
}

func TestGetApacheLinuxPath(t *testing.T) {
	getPathTestCommonMethod_linux(t, getApacheLinuxPathTestCaseData, pathhelper.GetApacheLinuxPath)
}
