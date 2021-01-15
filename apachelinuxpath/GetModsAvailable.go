package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// returns /etc/apache/mods-available as a string
func GetModsAvailable() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.ModsAvailable.CombineWith(pathhelper.GetApacheLinuxPath())
}
