package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func FilterUsingSeparator(
	separator,
	rootPath string,
	isContinueOnError bool,
	filter pathfuncs.Filter,
) (*[]string, *errwrappers.Collection) {
	normalizePath := normalize.PathUsingSeparator(
		separator,
		rootPath,
		true,
		true)

	return recursiveinternal.GetFilterPaths(
		separator,
		normalizePath,
		isContinueOnError,
		filter)
}
