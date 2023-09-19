package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/normalize"
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

func WithTempTest(paths ...string) string {
	return normalize.JoinNormalizedPaths(
		pathsconst.DefaultTempTestDir,
		paths...)
}
