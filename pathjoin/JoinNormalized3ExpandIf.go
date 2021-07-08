package pathjoin

import "gitlab.com/evatix-go/pathhelper/expandpath"

// JoinNormalized3ExpandIf normalized and expand applied if condition meets
func JoinNormalized3ExpandIf(
	isExpandNormalize bool,
	path1, path2, path3 string,
) string {
	normalized := JoinNormalized3If(
		isExpandNormalize,
		path1,
		path2,
		path3)

	if !isExpandNormalize {
		return normalized
	}

	return expandpath.ExpandVariables(normalized)
}
