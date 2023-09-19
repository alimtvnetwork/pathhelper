package pathhelper

import (
	"strings"

	"gitlab.com/auk-go/core/errcore"

	"gitlab.com/auk-go/pathhelper/internal/ispathinternal"
	"gitlab.com/auk-go/pathhelper/normalize"
	"gitlab.com/auk-go/pathhelper/pathhelpercore"
)

func getCombinedPathUsingConfigInternal(
	pathConfig *pathhelpercore.PathConfig,
	paths []string,
) string {
	if ispathinternal.EmptyArray(paths) {
		panic(errcore.InvalidEmptyPathType)
	}

	pathConfig = pathhelpercore.NewDefaultPathConfigOrExisting(pathConfig)
	var combinedPath string

	if !pathConfig.IsIgnoreEmptyPath {
		combinedPath = strings.Join(paths, pathConfig.Separator)
	} else {
		combinedPath = GetCombinedOfNonEmptyPaths(pathConfig.Separator, paths)
	}

	finalPath := normalize.PathUsingSeparatorIf(
		pathConfig.IsLongPathFix,
		pathConfig.IsLongPathFix,
		pathConfig.IsNormalize,
		pathConfig.Separator,
		combinedPath)

	return finalPath
}
