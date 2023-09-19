package pathjoin

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinBaseDirWithSep isNormalizePlusLongPathFix if true then for windows add UNC Location fix
//
// Omits baseDir if not given
func JoinBaseDirWithSep(
	isSkipEmpty,
	isExpandEnvVariables,
	isNormalizePlusLongPathFix bool,
	sep string,
	baseDir string,
	paths ...string,
) string {
	if len(paths) == 0 && baseDir == "" {
		return constants.EmptyString
	}

	if len(paths) == 0 && baseDir != "" {
		return JoinConditionalNormalizedExpandIf(
			isNormalizePlusLongPathFix,
			isExpandEnvVariables,
			baseDir,
			constants.EmptyString)
	}

	if isSkipEmpty {
		paths = stringslice.
			NonEmptySlice(paths)
	}

	finalPath := strings.Join(
		paths,
		sep)

	if baseDir != "" {
		finalPath = baseDir +
			sep +
			finalPath
	}

	expand := expandpath.ExpandVariablesIf(
		isExpandEnvVariables,
		finalPath)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
