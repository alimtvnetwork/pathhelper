package hexchecksum

import "gitlab.com/evatix-go/core/constants"

func OfFiles(request *FilesRequest) *FilesResult {
	hexOfListing := OfFilesList(
		request.Method,
		request.Files...)
	filesCount := len(request.Files)

	if hexOfListing.HasError() || !request.IsGenerateContentsChecksum || filesCount == 0 {
		return &FilesResult{
			HexFilesListChecksum:     hexOfListing.Value,
			HexFilesContentsChecksum: constants.EmptyString,
			Method:                   request.Method,
			ErrorWrapper:             hexOfListing.ErrorWrapper,
			FilesCount:               filesCount,
		}
	}

	hexContentsChecksum := OfFilesContentsAsync(
		request.Method,
		request.Files...)

	return &FilesResult{
		HexFilesListChecksum:     hexOfListing.Value,
		HexFilesContentsChecksum: hexContentsChecksum.Value,
		Method:                   request.Method,
		ErrorWrapper:             hexContentsChecksum.ErrorWrapper,
		FilesCount:               filesCount,
	}
}
