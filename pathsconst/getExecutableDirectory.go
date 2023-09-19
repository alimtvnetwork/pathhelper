package pathsconst

import (
	"os"

	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

func getExecutableDirectory() string {
	exe, _ := os.Executable()
	exeDir, _ := splitinternal.GetWithoutSlash(exe)

	return exeDir
}
