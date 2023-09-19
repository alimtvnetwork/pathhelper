package oldtests

import (
	"testing"

	"gitlab.com/auk-go/pathhelper/nginxlinuxpath"
)

var getSitesEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesEnabled",
	expected: "/etc/nginx/sites-enabled",
}

func TestGetSitesEnabled(t *testing.T) {
	getPathTestCommonMethodLinux(
		t,
		getSitesEnabledPathTestCaseData,
		nginxlinuxpath.GetSitesEnabled)
}
