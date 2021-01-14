package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to etc directory on different platforms.
func GetEtcPath() string {
	if osconsts.IsWindows {
		return knowndir.Etc.CombineWith(GetSystemDriversPath())
	}

	return knowndir.Etc.CombineWith(GetUnixRoot())
}
