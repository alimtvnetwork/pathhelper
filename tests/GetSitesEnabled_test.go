package tests

import (
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
	"testing"
)

var getSitesEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesEnabled",
	expected: "/etc/nginx/sites-enabled",
}

func TestGetSitesEnabled(t *testing.T) {
	pathTestCaseInternal_linux(t, getSitesEnabledPathTestCaseData, nginxlinuxpath.GetSitesEnabled)
}
