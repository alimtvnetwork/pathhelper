package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
)

// SimpleFilterNonRecursiveFiles normalize, expand false
func SimpleFilterNonRecursiveFiles(
	simpleFilter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	return SimpleFilterFilesOptionsAsync(
		false,
		false,
		false,
		simpleFilter,
		rootPath)
}
