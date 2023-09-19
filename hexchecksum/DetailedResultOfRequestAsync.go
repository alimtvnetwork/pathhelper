package hexchecksum

import (
	"gitlab.com/auk-go/asynchelper/syncparallel"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errnew"
)

func DetailedResultOfRequestAsync(
	request *FilesRequest,
) *DetailedResult {
	filesCount := len(request.Files)

	if filesCount == 0 {
		return EmptyDetailedResult()
	}

	isSortChecksum := request.IsSortFilesChecksum
	request.SortFileNamesBasedOnCondition()

	eachFilesChecksumSliceResult := EachFilesChecksumListAsyncIf(
		request.IsGenerateContentsChecksum,
		false,
		request.Method,
		request.Files...)

	if request.IsExitOnError() && eachFilesChecksumSliceResult.HasError() {
		return EmptyDetailedResultWithErr(
			request.Method,
			eachFilesChecksumSliceResult.ErrorWrapper)
	}

	var wholeChecksum, hexOfListing *errstr.Result
	eachChecksumValues := eachFilesChecksumSliceResult.
		SafeValues()
	isGenerateChecksum := request.IsGenerateContentsChecksum &&
		len(eachChecksumValues) == filesCount
	var mappedFileToHexChecksum *corestr.Hashmap

	syncparallel.Tasks(
		func() {
			checksumsRequest := stringslice.CloneIf(
				isSortChecksum,
				0,
				eachChecksumValues)

			wholeChecksum = OfChecksums(
				isGenerateChecksum,
				request.IsSortFilesChecksum,
				request.Method,
				checksumsRequest...)
		},
		func() {
			hexOfListing = OfFilesListIf(
				request.IsGenerateFileListChecksum,
				request.Method,
				request.Files...)
		},
		func() {
			mappedFileToHexChecksum = corestr.New.Hashmap.Cap(len(eachChecksumValues))

			if isGenerateChecksum && len(eachChecksumValues) > 0 {
				for i, fullFilePath := range request.Files {
					mappedFileToHexChecksum.AddOrUpdate(
						fullFilePath,
						eachChecksumValues[i])
				}
			}
		})

	mergedErr := errnew.Merge.New(
		hexOfListing.ErrorWrapper,
		wholeChecksum.ErrorWrapper)

	if mergedErr.HasError() {
		return EmptyDetailedResultWithErr(
			request.Method,
			mergedErr)
	}

	return &DetailedResult{
		FilesResult: FilesResult{
			HexFilesListChecksum:     hexOfListing.Value,
			HexFilesContentsChecksum: wholeChecksum.Value,
			FilesCount:               filesCount,
			Method:                   request.Method,
			ErrorWrapper:             nil,
		},
		Hashmap: mappedFileToHexChecksum,
	}
}
