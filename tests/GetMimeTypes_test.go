package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getMimeTypesPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetMimeTypes",
	expected: "/etc/nginx/mime.types",
}

func TestGetMimeTypes(t *testing.T) {
	getPathTestCommonMethod_linux(t, getMimeTypesPathTestCaseData, nginxlinuxpath.GetMimeTypes)
}
