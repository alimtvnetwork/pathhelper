package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func AllUsingSeparator(separator, rootPath string, isContinueOnError bool) (*[]string, *errwrappers.Collection) {
	normalizePath := normalize.PathUsingSeparator(
		separator,
		rootPath,
		true,
		true)

	return recursiveinternal.GetPaths(
		separator,
		normalizePath,
		isContinueOnError)
}
