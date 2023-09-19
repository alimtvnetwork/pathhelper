package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

func WithRoot(paths ...string) string {
	return JoinFix(pathsconst.RootDir, paths...)
}
