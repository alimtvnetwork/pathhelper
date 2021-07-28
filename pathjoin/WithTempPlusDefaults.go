package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
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
		JoinSimpleConditionalNonEmpty(
			pathsconst.TempDir,
			baseDir),
		locations...)
}
