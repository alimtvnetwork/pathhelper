package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinNormalizedExpand normalized and expand applied auto
func JoinNormalizedExpand(
	path1, path2 string,
) string {
	expand := expandpath.ExpandVariables(
		JoinSimple(path1, path2))

	return normalize.Path(expand)
}
