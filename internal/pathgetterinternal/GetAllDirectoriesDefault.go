package pathgetterinternal

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func GetAllDirectoriesDefault(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllDirectories(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
