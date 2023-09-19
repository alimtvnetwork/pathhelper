package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// FontsPath
//
// Returns path to Fonts directory on different platforms.
func FontsPath() string {
	if osconsts.IsWindows {
		return knowndir.Fonts.CombineWith(WidowsDirectory())
	}

	return knowndir.FontsUnix.CombineWith(UnixRoot())
}
