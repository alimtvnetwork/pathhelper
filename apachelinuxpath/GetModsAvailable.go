package apachelinuxpath

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetModsAvailable returns /etc/apache/mods-available as a string
func GetModsAvailable() string {
	if osconsts.IsWindows {
		panic("Location only available for Unix OS")
	}

	return knowndir.ModsAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
