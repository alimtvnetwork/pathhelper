package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func getCombinedPathUsingConfigInternal(
	pathConfig *pathhelpercore.PathConfig,
	paths []string,
) string {
	if pathhelpercore.IsEmptyArray(paths) {
		panic(constants.InvalidEmptyPathErrorMessage)
	}

	if pathConfig == nil {
		pathConfig = pathhelpercore.NewDefaultPathConfig()
	}

	var combinedPath string

	if !pathConfig.IsIgnoreEmptyPath {
		combinedPath = strings.Join(paths, pathConfig.Separator)
	} else {
		combinedPath = GetCombinedOfNonEmptyPaths(pathConfig.Separator, paths)
	}

	if pathConfig.IsNormalize {
		combinedPath = NormalizePath(combinedPath)
	}

	return combinedPath
}
