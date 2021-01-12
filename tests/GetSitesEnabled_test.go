package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getSitesEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesEnabled",
	expected: "/etc/nginx/sites-enabled",
}

func TestGetSitesEnabled(t *testing.T) {
	getPathTestCommonMethod_linux(t, getSitesEnabledPathTestCaseData, nginxlinuxpath.GetSitesEnabled)
}
