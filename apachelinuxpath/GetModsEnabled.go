package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// returns /etc/apache/mods-enabled as a string
func GetModsEnabled() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.ModsEnabled.CombineWith(knowndirget.ApacheLinuxPath())
}
