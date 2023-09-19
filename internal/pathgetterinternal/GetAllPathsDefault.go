package pathgetterinternal

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func GetAllPathsDefault(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllPaths(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
