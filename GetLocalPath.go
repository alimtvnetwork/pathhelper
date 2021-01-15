package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to local directory in windows. Otherwise returns Users directory.
func GetLocalPath() string {
	if osconsts.IsWindows {
		return knowndir.Local.CombineWith(GetAppDataPath())
	}

	return GetUserPath()
}
