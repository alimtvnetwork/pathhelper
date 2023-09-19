package pathjoin

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/normalize"
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

func WithTempPlusDefaults(
	baseDir string,
	locations ...string,
) string {
	return JoinBaseDirWithSep(
		true,
		true,
		true,
		osconsts.PathSeparator,
		normalize.SimpleJoinPath(
			pathsconst.TempPermanentDir,
			baseDir),
		locations...)
}
