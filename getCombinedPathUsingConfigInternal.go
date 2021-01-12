package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/msgtype"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func getCombinedPathUsingConfigInternal(
	pathConfig *pathhelpercore.PathConfig,
	paths []string,
) string {
	if pathhelpercore.IsEmptyArray(paths) {
		panic(msgtype.InvalidEmptyPathErrorMessage)
	}

	pathConfig = pathhelpercore.NewDefaultPathConfigOrExisting(pathConfig)
	var combinedPath string

	if !pathConfig.IsIgnoreEmptyPath {
		combinedPath = strings.Join(paths, pathConfig.Separator)
	} else {
		combinedPath = GetCombinedOfNonEmptyPaths(pathConfig.Separator, paths)
	}

	finalPath := NormalizePathUsingSeparatorIf(
		pathConfig.IsLongPathFix,
		pathConfig.IsNormalize,
		pathConfig.Separator,
		combinedPath)

	return finalPath
}
