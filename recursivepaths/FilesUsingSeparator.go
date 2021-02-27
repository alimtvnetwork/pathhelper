package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func FilesUsingSeparator(
	separator,
	rootPath string,
	isContinueOnError bool,
) *errstr.ResultsWithErrorCollection {
	normalizePath := normalize.PathUsingSeparator(
		separator,
		rootPath,
		true,
		true)

	return recursiveinternal.GetFilesPaths(
		separator,
		normalizePath,
		isContinueOnError)
}
