package pathhelper

import (
	"gitlab.com/auk-go/core"
)

func GetRelativePaths(
	rootPath string,
	fullPaths ...string,
) *[]string {
	if fullPaths == nil {
		return core.EmptyStringsPtr()
	}

	return GetRelativePathsOfPtr(
		rootPath,
		&fullPaths)
}
