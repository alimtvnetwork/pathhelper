package recursivepaths

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func FilterMust(
	rootPath string,
	isContinueOnError bool,
	filter pathfuncs.Filter,
) *[]string {
	normalizePath := normalize.PathUsingSeparator(
		osconsts.PathSeparator,
		rootPath,
		true,
		true)

	paths, errWrappersCollection := recursiveinternal.GetFilterPaths(
		osconsts.PathSeparator,
		normalizePath,
		isContinueOnError,
		filter)

	errWrappersCollection.HandleError()

	return paths
}
