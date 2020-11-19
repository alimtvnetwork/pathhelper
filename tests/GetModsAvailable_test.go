package tests

import (
	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
	"testing"
)

var getModsAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModsAvailable",
	expected: "/etc/apache/mods-available",
}

func TestGetModsAvailable(t *testing.T) {
	pathTestCaseInternal_linux(t, getModsAvailablePathTestCaseData, apachelinuxpath.GetModsAvailable)
}
