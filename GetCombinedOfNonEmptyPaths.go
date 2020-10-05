package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func GetCombinedOfNonEmptyPaths(separator string, paths []string) string {
	optimizedPaths := make([]string, 0, len(paths))

	for _, path := range paths {
		if pathhelpercore.IsEmptyPath(path) {
			continue
		}

		optimizedPaths = append(optimizedPaths, path)
	}

	return strings.Join(optimizedPaths, separator)
}
