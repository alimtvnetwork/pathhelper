package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
)

var unixRootTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetUnixRoot",
	expected: "/",
}

func TestGetUnixRoot(t *testing.T) {
	getPathTestCommonMethod_linux(t, unixRootTestCaseData, pathhelper.GetUnixRoot)
}
