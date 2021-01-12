package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
)

var getSitesEnabledApachePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesEnabled",
	expected: "/etc/apache/sites-enabled",
}

func TestGetSitesEnabled_Apache(t *testing.T) {
	getPathTestCommonMethod_linux(t, getSitesEnabledApachePathTestCaseData, apachelinuxpath.GetSitesEnabled)
}
