package recursivepaths

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Files(rootPath string, isContinueOnError bool) (*[]string, *errwrappers.Collection) {
	normalizePath := normalize.PathUsingSeparator(
		osconsts.PathSeparator,
		rootPath,
		true)

	return recursiveinternal.GetFilesPaths(
		osconsts.PathSeparator,
		normalizePath,
		isContinueOnError)
}
