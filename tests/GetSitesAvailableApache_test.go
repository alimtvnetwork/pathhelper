package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
)

var getSitesAvailableApachePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesAvailable",
	expected: "/etc/apache/sites-available",
}

func TestGetSitesAvailable_Apache(t *testing.T) {
	getPathTestCommonMethod_linux(t, getSitesAvailableApachePathTestCaseData, apachelinuxpath.GetSitesAvailable)
}
