package recursiveinternal

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func GetPathsWithoutSeparator(
	rootPath string,
	isContinueOnEmpty bool,
) (*[]string, *errwrappers.Collection) {
	if rootPath == "" {
		return core.EmptyStringsPtr(), errwrappers.Empty()
	}

	if !fsinternal.IsExistButNotDirectory(rootPath) {
		return &[]string{}, errwrappers.Empty()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	return getPaths(
		osconsts.PathSeparator,
		rootPath,
		fileInfos,
		err,
		isContinueOnEmpty)
}
