package recursivepaths

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func AllMust(rootPath string, isContinueOnError bool) *[]string {
	normalizePath := normalize.PathUsingSeparator(
		osconsts.PathSeparator,
		rootPath,
		true)

	paths, errWrappersCollection := recursiveinternal.GetPaths(
		osconsts.PathSeparator,
		normalizePath,
		isContinueOnError)

	errWrappersCollection.Handle()

	return paths
}
