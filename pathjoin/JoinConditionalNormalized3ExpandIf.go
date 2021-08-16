package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinConditionalNormalized3ExpandIf normalized or expand or both apply based on condition
func JoinConditionalNormalized3ExpandIf(
	isNormalizePlusLongPathFix,
	isExpand bool,
	path1, path2, path3 string,
) string {
	combined := JoinSimple3ConditionalNonEmpty(
		path1,
		path2,
		path3)

	expand := expandpath.ExpandVariablesIf(
		isExpand,
		combined)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
