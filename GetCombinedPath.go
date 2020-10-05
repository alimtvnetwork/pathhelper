package pathhelper

import (
	"strings"
)

// @isIgnoreEmptyPath if true then ignore empty string (nil, "", or any empty spaces "  ")
func GetCombinedPath(separator string, isIgnoreEmptyPath bool, paths ...string) string {
	if !isIgnoreEmptyPath {
		return strings.Join(paths, separator)
	}

	return GetCombinedOfNonEmptyPaths(separator, paths)
}
