package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func WithTemp(
	locations ...string,
) string {
	return Fixed(
		pathsconst.TempDir,
		locations...)
}
