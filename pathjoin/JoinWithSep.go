package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinWithSep isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinWithSep(
	isSkipEmpty,
	isExpandEnvVariables,
	isNormalizePlusLongPathFix bool,
	sep string,
	paths ...string,
) string {
	if len(paths) == 0 {
		return constants.EmptyString
	}

	if isSkipEmpty {
		paths = stringslice.
			NonEmptySlice(paths)
	}

	finalPath := strings.Join(
		paths,
		sep)

	expand := expandpath.ExpandVariablesIf(
		isExpandEnvVariables,
		finalPath)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
