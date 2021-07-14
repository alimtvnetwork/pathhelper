package pathjoin

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
)

func WithTemp(
	isNormalize,
	isExpand bool,
	locations ...string,
) string {
	return JoinBaseDirWithSep(
		false,
		isExpand,
		isNormalize,
		osconsts.PathSeparator,
		os.TempDir(),
		locations...)
}
