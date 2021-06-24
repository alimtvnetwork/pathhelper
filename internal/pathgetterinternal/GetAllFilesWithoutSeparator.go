package pathgetterinternal

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func GetAllFilesWithoutSeparator(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllFiles(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
