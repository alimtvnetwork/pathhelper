package tests

import (
	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
	"testing"
)

var getConfAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConfAvailable",
	expected: "/etc/apache/conf-available",
}

func TestGetConfAvailable(t *testing.T) {
	pathTestCaseInternal_linux(t, getConfAvailablePathTestCaseData, apachelinuxpath.GetConfAvailable)
}
