package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// isNormalizePlusLongPathFix if true then for windows add UNC Path fix
func JoinWithSep(
	isNormalizePlusLongPathFix bool,
	sep string,
	paths ...string,
) string {
	if len(paths) == 0 {
		return constants.EmptyString
	}

	finalPath := strings.Join(paths, sep)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		finalPath)
}
