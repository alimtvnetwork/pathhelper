package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to local directory in windows. Otherwise returns Users directory.
func GetLocalPath() string {
	if IsWindows() {
		return enums.Local.GetPrefixCombinedWith(GetAppDataPath())
	}

	return GetUserPath()
}
