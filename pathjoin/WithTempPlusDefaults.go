package pathjoin

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
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
			os.TempDir(),
			baseDir),
		locations...)
}
