package hexchecksum

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
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

	if eachFilesChecksumSliceResult.HasError() {
		return EmptyDetailedResultWithErr(
			eachFilesChecksumSliceResult.ErrorWrapper)
	}

	var wholeChecksum, hexOfListing *errstr.Result

	wg := sync.WaitGroup{}
	wg.Add(2)
	eachChecksumValues := eachFilesChecksumSliceResult.
		ValueNonPtr()
	isGenerateChecksum := request.IsGenerateContentsChecksum &&
		len(eachChecksumValues) == filesCount

	go func() {
		checksumsRequest := stringslice.CloneIf(
			isSortChecksum,
			0,
			eachChecksumValues)

		wholeChecksum = OfChecksums(
			isGenerateChecksum,
			request.IsSortFilesChecksum,
			request.Method,
			checksumsRequest...)

		wg.Done()
	}()

	go func() {
		hexOfListing = OfFilesListIf(
			request.IsGenerateFileListChecksum,
			request.Method,
			request.Files...)

		wg.Done()
	}()

	mappedFileToHexChecksum := corestr.NewHashmap(len(eachChecksumValues))

	if isGenerateChecksum {
		for i, fullFilePath := range request.Files {
			mappedFileToHexChecksum.AddOrUpdate(
				fullFilePath,
				eachChecksumValues[i])
		}
	}

	wg.Wait()

	mergedErr := errnew.Merge(
		hexOfListing.ErrorWrapper,
		wholeChecksum.ErrorWrapper)

	if mergedErr.HasError() {
		return EmptyDetailedResultWithErr(
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
