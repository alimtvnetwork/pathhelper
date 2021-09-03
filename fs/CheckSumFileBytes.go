package fs

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func CheckSumFileBytes(
	hashType hashas.Variant,
	location string,
) *errbyte.Results {
	fileInfo, isExist, _ := chmodhelper.GetPathExistStatExpand(location)

	if !isExist {
		return errbyte.EmptyResultsWithError(
			errnew.PathMessages(
				errtype.PathMissingOrInvalid,
				location,
				"cannot hash non readable file."))
	}

	if fileInfo.IsDir() {
		return errbyte.EmptyResultsWithError(
			errnew.PathMessages(
				errtype.UnexpectedDirectory,
				location,
				"cannot hash directory (use checksummer)."))
	}

	readBytes := ReadFile(location)

	return hashType.SumOfErrorBytes(readBytes)
}
