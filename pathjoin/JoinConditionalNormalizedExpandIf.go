package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
)

// JoinConditionalNormalizedExpandIf normalized or expand or both apply based on condition
func JoinConditionalNormalizedExpandIf(
	isNormalize,
	isExpand bool,
	path1, path2 string,
) string {
	if isNormalize && isExpand {
		return JoinNormalizedExpand(path1, path2)
	}

	if isNormalize {
		return JoinNormalized(path1, path2)
	}

	finalPath := path1 +
		osconsts.PathSeparator +
		path2

	if isExpand {
		return expandpath.ExpandVariables(finalPath)
	}

	return finalPath
}
