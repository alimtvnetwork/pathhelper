package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// returns /etc/apache/sites-enabled as a string
func GetSitesEnabled() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.SitesEnabled.CombineWith(pathhelper.GetApacheLinuxPath())
}
