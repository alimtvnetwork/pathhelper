package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func GetAbsolutePath(basePath, relativePath string, isLongPathFix, isNormalize bool) string {
	if pathhelpercore.IsEmptyPath(basePath) || pathhelpercore.IsEmptyPath(relativePath) {
		panic(constants.InvalidAnyPathEmptyErrorMessage)
	}

	return GetCombinedPath(
		constants.PathSeparator,
		false,
		isLongPathFix,
		isNormalize,
		basePath,
		relativePath,
	)
}
