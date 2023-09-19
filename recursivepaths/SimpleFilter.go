package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
)

func SimpleFilter(
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	return SimpleFilterOptions(
		false,
		false,
		filter,
		rootPath)
}
