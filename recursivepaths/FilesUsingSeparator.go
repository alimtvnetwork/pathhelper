package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func FilesUsingSeparator(
	separator,
	rootPath string,
	isContinueOnError bool,
) (*[]string, *errwrappers.Collection) {
	normalizePath := normalize.PathUsingSeparator(
		separator,
		rootPath,
		true)

	return recursiveinternal.GetFilesPaths(
		separator,
		normalizePath,
		isContinueOnError)
}
