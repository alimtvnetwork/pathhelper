package pathjoin

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

// Fixed
//
// Items are applied with
// normalize,
// skip on empty path given
func Fixed(
	baseDir string,
	paths ...string,
) string {
	if len(paths) == 0 && baseDir == "" {
		return constants.EmptyString
	}

	return JoinBaseDirWithSep(
		true,
		false,
		true,
		osconsts.PathSeparator,
		baseDir,
		paths...)
}
