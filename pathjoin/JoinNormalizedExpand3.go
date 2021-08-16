package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalizedExpand3 normalized and expand applied auto
func JoinNormalizedExpand3(
	path1, path2, path3 string,
) string {
	joined := JoinSimple3(path1, path2, path3)
	expand := expandpath.ExpandVariables(joined)

	return normalize.Path(expand)
}
