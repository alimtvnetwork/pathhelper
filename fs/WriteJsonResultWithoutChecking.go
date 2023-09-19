package fs

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/errorwrapper"
)

func WriteJsonResultWithoutChecking(
	jsonResult *corejson.Result,
	location string,
) *errorwrapper.Wrapper {
	return WriteFile(
		true,
		location,
		jsonResult.Bytes)
}
