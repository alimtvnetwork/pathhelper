package hexchecksum

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/hashas"
)

func OfFilesListIf(
	isGenerate bool,
	hashMethod hashas.Variant,
	files ...string,
) *errstr.Result {
	if !isGenerate {
		return errstr.Empty.Result()
	}

	jsonResult := corejson.NewPtr(files)

	return hashMethod.
		HexOfJsonResult(jsonResult)
}
