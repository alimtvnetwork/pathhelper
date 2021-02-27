package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// returns /etc/apache/sites-available as a string
func GetSitesAvailable() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.SitesAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
