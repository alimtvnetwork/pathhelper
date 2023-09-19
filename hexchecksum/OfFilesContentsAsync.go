package hexchecksum

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/hashas"
)

func OfFilesContentsAsync(
	isSortChecksums,
	isSortFileName bool,
	hashMethod hashas.Variant,
	filesPaths ...string,
) *errstr.Result {
	length := len(filesPaths)
	if length == 0 {
		return errstr.Empty.Result()
	}

	sortIf(isSortFileName, filesPaths)

	eachFilesChecksum := EachFilesChecksumListAsync(
		hashMethod,
		filesPaths...)

	if eachFilesChecksum.HasError() {
		return errstr.New.Result.ErrorWrapper(eachFilesChecksum.ErrorWrapper)
	}

	checkSumValuesSlice := eachFilesChecksum.SafeValues()
	sortIf(isSortChecksums, checkSumValuesSlice)

	// success
	return hashMethod.HexSumOfAny(
		checkSumValuesSlice)
}
