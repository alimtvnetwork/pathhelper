package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
)

var getConfAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConfAvailable",
	expected: "/etc/apache/conf-available",
}

func TestGetConfAvailable(t *testing.T) {
	getPathTestCommonMethod_linux(t, getConfAvailablePathTestCaseData, apachelinuxpath.GetConfAvailable)
}
