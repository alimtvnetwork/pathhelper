package hexchecksum

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func OfFilesContentsAsync(
	isSortChecksums,
	isSortFileName bool,
	hashMethod hashas.Variant,
	filesPaths ...string,
) *errstr.Result {
	length := len(filesPaths)
	if length == 0 {
		return errstr.Empty()
	}

	sortIf(isSortFileName, filesPaths)

	eachFilesChecksum := EachFilesChecksumListAsync(
		hashMethod,
		filesPaths...)

	if eachFilesChecksum.HasError() {
		return errstr.ErrorWrapper(eachFilesChecksum.ErrorWrapper)
	}

	checkSumValuesSlice := eachFilesChecksum.ValueNonPtr()
	sortIf(isSortChecksums, checkSumValuesSlice)

	// success
	return hashMethod.HexSumOfAny(
		checkSumValuesSlice)
}
