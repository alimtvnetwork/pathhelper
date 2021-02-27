package pathhelper

import "gitlab.com/evatix-go/core"

func GetAbsolutePaths(
	basePath string,
	isLongPathFix, isNormalize bool,
	relativePaths ...string,
) *[]string {
	if relativePaths == nil {
		return core.EmptyStringsPtr()
	}

	return GetAbsolutePathsOfPtr(
		basePath,
		isLongPathFix,
		isNormalize,
		&relativePaths)
}
