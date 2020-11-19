package tests

import (
	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
	"testing"
)

var getSitesAvailableApachePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesAvailable",
	expected: "/etc/apache/sites-available",
}

func TestGetSitesAvailable_Apache(t *testing.T) {
	pathTestCaseInternal_linux(t, getSitesAvailableApachePathTestCaseData, apachelinuxpath.GetSitesAvailable)
}
