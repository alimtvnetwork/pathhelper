package tests

import (
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
	"testing"
)

var getSitesAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesAvailable",
	expected: "/etc/nginx/sites-available",
}

func TestGetSitesAvailable(t *testing.T) {
	getPathTestCommonMethod_linux(t, getSitesAvailablePathTestCaseData, nginxlinuxpath.GetSitesAvailable)
}
