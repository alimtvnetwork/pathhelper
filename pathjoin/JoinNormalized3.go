package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalized3 normalized applied auto
func JoinNormalized3(path1, path2, path3 string) string {
	finalPath := path1 +
		osconsts.PathSeparator +
		path2 +
		osconsts.PathSeparator +
		path3

	return normalize.Path(finalPath)
}
