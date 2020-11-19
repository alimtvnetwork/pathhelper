package tests

import (
	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
	"testing"
)

var getModsEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModsEnabled",
	expected: "/etc/apache/mods-enabled",
}

func TestGetModsEnabled(t *testing.T) {
	pathTestCaseInternal_linux(t, getModsEnabledPathTestCaseData, apachelinuxpath.GetModsEnabled)
}
