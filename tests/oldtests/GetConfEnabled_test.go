package oldtests

import (
	"testing"

	"gitlab.com/auk-go/pathhelper/apachelinuxpath"
)

var getConfEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConfEnabled",
	expected: "/etc/apache/conf-enabled",
}

func TestGetConfEnabled(t *testing.T) {
	getPathTestCommonMethodLinux(t, getConfEnabledPathTestCaseData, apachelinuxpath.GetConfEnabled)
}
