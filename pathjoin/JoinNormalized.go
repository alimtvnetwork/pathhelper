package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// normalized applied auto
func JoinNormalized(path1, path2 string) string {
	finalPath := path1 + osconsts.PathSeparator + path2

	return normalize.Path(finalPath)
}
