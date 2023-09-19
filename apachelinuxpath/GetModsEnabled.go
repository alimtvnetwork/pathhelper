package apachelinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetModsEnabled returns /etc/apache/mods-enabled as a string
func GetModsEnabled() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.ModsEnabled.CombineWith(knowndirget.ApacheLinuxPath())
}
