package recursivepaths

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Files(rootPath string, isContinueOnError bool) *errstr.ResultsWithErrorCollection {
	normalizePath := normalize.PathUsingSeparator(
		osconsts.PathSeparator,
		rootPath,
		true,
		true)

	return recursiveinternal.GetFilesPaths(
		osconsts.PathSeparator,
		normalizePath,
		isContinueOnError)
}
