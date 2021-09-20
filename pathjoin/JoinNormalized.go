package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalized normalized applied auto
func JoinNormalized(
	path1, path2 string,
) string {
	if path2 == "" {
		return normalize.Path(path1)
	}

	finalPath := path1 +
		osconsts.PathSeparator +
		path2

	return normalize.Path(finalPath)
}
