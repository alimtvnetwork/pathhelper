package pathjoin

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinWithBaseDirSep
//
//	isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinWithBaseDirSep(
	isSkipEmpty,
	isExpandEnvVariables,
	isNormalizePlusLongPathFix bool,
	sep string,
	baseDir string,
	relativePaths ...string,
) string {
	if len(relativePaths) == 0 {
		return constants.EmptyString
	}

	if isSkipEmpty {
		relativePaths = stringslice.
			NonEmptySlice(relativePaths)
	}

	finalPath := strings.Join(
		relativePaths,
		sep)

	if baseDir != "" {
		finalPath = baseDir + sep + finalPath
	}

	expand := expandpath.ExpandVariablesIf(
		isExpandEnvVariables,
		finalPath)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
