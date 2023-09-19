package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns path to local temp directory. If directory doesn't exist it still returns the path as a string.
func LocalTempPath() string {
	if osconsts.IsWindows {
		return knowndir.LocalTempWin.CombineWith(AppDataPath())
	}

	return knowndir.LocalTempUnix.CombineWith(UnixRoot())
}
