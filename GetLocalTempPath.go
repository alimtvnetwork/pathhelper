package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to local temp directory. If directory doesn't exist it still returns the path as a string.
func GetLocalTempPath() string {
	var localTempDir string

	if IsWindows() {
		localTempDir = enums.LocalTempWin.GetPrefixCombinedWith(GetAppDataPath())
	} else {
		localTempDir = enums.LocalTempUnix.GetPrefixCombinedWith(GetUserPath())
	}

	return localTempDir
}
