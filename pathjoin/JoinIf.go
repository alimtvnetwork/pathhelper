package pathjoin

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinIf isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinIf(isNormalizePlusLongPathFix bool, path1, path2 string) string {
	finalPath := path1 + osconsts.PathSeparator + path2

	if !isNormalizePlusLongPathFix {
		return finalPath
	}

	return normalize.Path(finalPath)
}
