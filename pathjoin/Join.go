package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// isNormalizePlusLongPathFix if true then for windows add UNC Path fix
func Join(path1, path2 string, isNormalizePlusLongPathFix bool) string {
	finalPath := path1 + osconsts.PathSeparator + path2

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		finalPath)
}
