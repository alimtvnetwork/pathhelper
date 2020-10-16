package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to local temp directory. If directory doesn't exist it creates the directory and returns the path as a string.
func GetLocalTempPath() string {
	var localTempDir string

	if IsWindows() {
		localTempDir = GetCombinePathWith(GetAppDataPath(), enums.LocalTempWin.Value())
	} else {
		localTempDir = GetCombinePathWith(GetUserPath(), enums.LocalTempUnix.Value())
	}

	CreateDirectoryAll(localTempDir, constants.Perm)

	return localTempDir
}
