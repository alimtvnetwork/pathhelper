package pathjoin

import "gitlab.com/evatix-go/pathhelper/expandpath"

// JoinNormalizedExpand normalized and expand applied auto
func JoinNormalizedExpand(
	path1, path2 string,
) string {
	normalized := JoinNormalized(path1, path2)

	return expandpath.ExpandVariables(normalized)
}
