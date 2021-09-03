package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalized3 normalized applied auto
func JoinNormalized3(path1, path2, path3 string) string {
	joined := normalize.SimpleJoinPath3(
		path1,
		path2,
		path3)

	return normalize.Path(joined)
}
