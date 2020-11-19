package tests

import (
	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
	"testing"
)

var getConfEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConfEnabled",
	expected: "/etc/apache/conf-enabled",
}

func TestGetConfEnabled(t *testing.T) {
	pathTestCaseInternal_linux(t, getConfEnabledPathTestCaseData, apachelinuxpath.GetConfEnabled)
}
