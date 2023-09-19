package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
)

func SimpleFilterFullPaths(
	filter pathfuncs.SimpleFilter,
	fullPaths ...string,
) *errstr.Results {
	return pathfuncs.SimpleFilterFullPathsAsync(
		false,
		filter,
		fullPaths...)
}
