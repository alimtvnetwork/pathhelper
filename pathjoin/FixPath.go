package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// FixPath normalized or expand or both apply based on condition
func FixPath(
	isNormalizePlusLogPathFix,
	isExpand bool,
	location string,
) string {
	location = normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLogPathFix,
		osconsts.PathSeparator,
		location)

	location = expandpath.ExpandVariablesIf(
		isExpand,
		location)

	return location
}
