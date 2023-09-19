package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns path to System directory on different platforms.
func GetSystemPath() string {
	if osconsts.IsWindows {
		return WidowsDirectory()
	}

	return knowndir.SystemUnix.Value()
}
