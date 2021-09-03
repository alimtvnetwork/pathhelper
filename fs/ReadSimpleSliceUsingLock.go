package fs

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func ReadSimpleSliceUsingLock(
	filePath string,
) (*corestr.SimpleSlice, *errorwrapper.Wrapper) {
	results := ReadFileLinesUsingLock(filePath)

	if results.HasIssuesOrEmpty() {
		return corestr.EmptySimpleSlice(), results.ErrorWrapper
	}

	return &corestr.SimpleSlice{Items: results.Values}, errnew.EmptyPtr
}
