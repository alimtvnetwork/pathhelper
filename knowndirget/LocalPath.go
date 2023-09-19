package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns path to local directory in windows. Otherwise returns Users directory.
func LocalPath() string {
	if osconsts.IsWindows {
		return knowndir.Local.CombineWith(AppDataPath())
	}

	return UserPath()
}
