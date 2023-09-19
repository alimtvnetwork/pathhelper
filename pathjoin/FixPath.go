package pathjoin

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// FixPath normalized or expand or both apply based on condition
func FixPath(
	isNormalizePlusLogPathFix,
	isExpand bool,
	location string,
) string {
	if location == "" {
		return ""
	}

	expand := expandpath.ExpandVariablesIf(
		isExpand,
		location)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLogPathFix,
		osconsts.PathSeparator,
		expand)
}
