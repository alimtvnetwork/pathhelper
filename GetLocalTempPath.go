package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to local temp directory. If directory doesn't exist it still returns the path as a string.
func GetLocalTempPath() string {
	if IsWindows() {
		return enums.LocalTempWin.CombineWith(GetAppDataPath())
	}
	return enums.LocalTempUnix.CombineWith(GetUserPath())
}
