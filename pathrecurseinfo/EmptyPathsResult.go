package pathrecurseinfo

import "gitlab.com/auk-go/core/coredata/corestr"

func EmptyPathsResult() *PathsResult {
	return &PathsResult{
		ExpandingPaths: corestr.Empty.SimpleSlice(),
		IsExist:        false,
		IsFile:         false,
		IsDir:          false,
	}
}
