package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
)

// SimpleFilterNonRecursiveDirs normalize, expand false
func SimpleFilterNonRecursiveDirs(
	simpleFilter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	return SimpleFilterDirsOptions(
		false,
		false,
		false,
		simpleFilter,
		rootPath)
}
