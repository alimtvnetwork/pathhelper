package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to System directory on different platforms.
func GetSystemPath() string {
	if osconsts.IsWindows {
		return GetWidowsDirectory()
	}

	return knowndir.SystemUnix.Value()
}
