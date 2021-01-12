package pathhelper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func GetAbsolutePath(
	basePath,
	relativePath string,
	isLongPathFix, isNormalize bool,
) string {
	if pathhelpercore.IsEmptyPath(basePath) || pathhelpercore.IsEmptyPath(relativePath) {
		panic(msgtype.InvalidEmptyPathErrorMessage)
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
