package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalized3ExpandIf normalized and expand applied if condition meets
func JoinNormalized3ExpandIf(
	isExpandNormalize bool,
	path1, path2, path3 string,
) string {
	joined := normalize.SimpleJoinPath3(
		path1,
		path2,
		path3)

	if !isExpandNormalize {
		return joined
	}

	expand := expandpath.ExpandVariablesIf(isExpandNormalize, joined)

	return normalize.Path(expand)
}
