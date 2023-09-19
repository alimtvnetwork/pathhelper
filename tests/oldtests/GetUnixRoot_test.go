package oldtests

import (
	"testing"

	"gitlab.com/auk-go/pathhelper/knowndirget"
)

var unixRootTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "UnixRoot",
	expected: "/",
}

func TestGetUnixRoot(t *testing.T) {
	getPathTestCommonMethodLinux(t, unixRootTestCaseData, knowndirget.UnixRoot)
}
