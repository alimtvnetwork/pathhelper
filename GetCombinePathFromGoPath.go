package pathhelper

import (
	"gitlab.com/auk-go/pathhelper/knowndirget"
	"gitlab.com/auk-go/pathhelper/pathhelpercore"
)

// By defaultPathConfig:
// Separator = os.PathSeparator
// IsNormalize = true
// IsIgnoreEmptyPath = false
func GetCombinePathFromGoPath(
	pathConfig *pathhelpercore.PathConfig,
	givenPaths ...string,
) string {
	goPath := knowndirget.GoPath()
	combinedPaths := getCombinedPathUsingConfigInternal(pathConfig, givenPaths)

	return GetCombinedPathUsingConfig(pathConfig, goPath, combinedPaths)
}
