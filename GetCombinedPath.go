package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// @isIgnoreEmptyPath if true then ignore empty string (nil, "", or any empty spaces "  ")
func GetCombinedPath(separator string, isIgnoreEmptyPath bool, isNormalize bool, paths ...string) string {
	pathConfig := &pathhelpercore.PathConfig{
		Separator:         separator,
		IsNormalize:       isNormalize,
		IsIgnoreEmptyPath: isIgnoreEmptyPath,
	}

	return getCombinedPathUsingConfigInternal(
		pathConfig,
		paths)
}
