package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func WithTempPlus(
	isNormalize,
	isExpand bool,
	baseDir string,
	locations ...string,
) string {
	return JoinBaseDirWithSep(
		false,
		isExpand,
		isNormalize,
		osconsts.PathSeparator,
		JoinSimple(pathsconst.TempDir, baseDir),
		locations...)
}
