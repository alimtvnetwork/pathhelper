package pathgetterinternal

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

// GetAllFilesDefault only gives files not nested files
func GetAllFilesDefault(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllFiles(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
