package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to Fonts directory on different platforms.
func GetFontsPath() string {
	if osconsts.IsWindows {
		return knowndir.Fonts.CombineWith(GetWidowsDirectory())
	}

	return knowndir.FontsUnix.CombineWith(GetUnixRoot())
}
