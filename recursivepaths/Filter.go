package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func Filter(
	filter pathfuncs.Filter,
	rootPath string,
) *errstr.Results {
	return FilterOptions(
		false,
		false,
		filter,
		rootPath)
}
