package apachelinuxpath

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetConfEnabled returns /etc/apache/conf-enabled as a string
func GetConfEnabled() string {
	if osconsts.IsWindows {
		panic("Location only available for Unix OS")
	}

	return knowndir.ConfEnabled.CombineWith(knowndirget.ApacheLinuxPath())
}
