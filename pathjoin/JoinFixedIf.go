package pathjoin

import (
	"path"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinFixedIf
//
// isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinFixedIf(
	isFixed bool,
	baseDir string,
	paths ...string,
) string {
	if len(paths) == 0 && baseDir == "" {
		return constants.EmptyString
	}

	if len(paths) == 0 && baseDir != "" {
		return normalize.PathUsingSingleIf(
			isFixed,
			baseDir)
	}

	if !isFixed {
		return path.Join(
			baseDir,
			path.Join(paths...))
	}

	return JoinBaseDirWithSep(
		true,
		false,
		true,
		osconsts.PathSeparator,
		baseDir,
		paths...)
}
