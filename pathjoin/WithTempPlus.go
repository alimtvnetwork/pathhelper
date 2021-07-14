package pathjoin

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
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
		JoinSimple(os.TempDir(), baseDir),
		locations...)
}
