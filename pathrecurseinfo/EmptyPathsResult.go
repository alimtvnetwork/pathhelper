package pathrecurseinfo

import "gitlab.com/evatix-go/core/coredata/corestr"

func EmptyPathsResult() *PathsResult {
	return &PathsResult{
		ExpandingPaths: corestr.EmptySimpleSlice(),
		IsExist:        false,
		IsFile:         false,
		IsDir:          false,
	}
}
