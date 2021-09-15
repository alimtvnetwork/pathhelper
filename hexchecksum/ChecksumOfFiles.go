package hexchecksum

import "gitlab.com/evatix-go/core/constants"

func ChecksumOfFiles(request *FilesChecksumRequest) *FilesChecksumResult {
	hexOfListing := ChecksumOfFilesList(
		request.Method,
		request.Files)
	filesCount := len(request.Files)

	if hexOfListing.HasError() || !request.IsGenerateContentsChecksum {
		return &FilesChecksumResult{
			HexFilesListChecksum:     hexOfListing.Value,
			HexFilesContentsChecksum: constants.EmptyString,
			Method:                   request.Method,
			ErrorWrapper:             hexOfListing.ErrorWrapper,
			FilesCount:               filesCount,
		}
	}

	hexContentsChecksum := ChecksumOfFilesContentsAsync(
		request.Method,
		request.Files)

	return &FilesChecksumResult{
		HexFilesListChecksum:     hexOfListing.Value,
		HexFilesContentsChecksum: hexContentsChecksum.Value,
		Method:                   request.Method,
		ErrorWrapper:             hexContentsChecksum.ErrorWrapper,
		FilesCount:               filesCount,
	}
}
