package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinSimple3 Doesn't apply normalize
func JoinSimple3(path1, path2, path3 string) string {
	return normalize.SimpleJoinPath3(path1, path2, path3)
}
