package pathjoin

import "gitlab.com/evatix-go/pathhelper/expandpath"

// JoinNormalizedExpand3 normalized and expand applied auto
func JoinNormalizedExpand3(
	path1, path2, path3 string,
) string {
	normalized := JoinNormalized3(path1, path2, path3)

	return expandpath.ExpandVariables(normalized)
}
