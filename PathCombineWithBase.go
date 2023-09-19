package pathhelper

import (
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func PathCombineWithBase(
	isNormalize,
	isExpandEnv bool,
	baseDir,
	relativePath string,
) string {
	combinedPath := pathjoin.JoinNormalizedIf(
		isNormalize,
		baseDir,
		relativePath)

	if isExpandEnv {
		combinedPath = expandpath.ExpandVariables(combinedPath)
	}

	return combinedPath
}
