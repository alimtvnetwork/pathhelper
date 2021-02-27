package recursivepaths

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func FilesMust(rootPath string, isContinueOnError bool) *[]string {
	normalizePath := normalize.PathUsingSeparator(
		osconsts.PathSeparator,
		rootPath,
		true,
		true)

	paths := recursiveinternal.GetFilesPaths(
		osconsts.PathSeparator,
		normalizePath,
		isContinueOnError)

	paths.ErrorWrappers.HandleError()

	return paths.Values
}
