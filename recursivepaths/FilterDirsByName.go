package recursivepaths

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

func FilterDirsByName(
	isRecursive bool,
	rootPath string,
	acceptingDirNames ...string,
) *errstr.Results {
	if len(acceptingDirNames) == 0 {
		return errstr.Empty.Results()
	}

	hashset := corestr.New.Hashset.StringsPtr(&acceptingDirNames)
	dirNameFilterFunc := func(fullPath string) (isTake bool, err error) {
		dirName := splitinternal.GetName(fullPath)

		return hashset.Has(dirName), nil
	}

	return SimpleFilterDirsOptions(
		isRecursive,
		false,
		false,
		dirNameFilterFunc,
		rootPath)
}
