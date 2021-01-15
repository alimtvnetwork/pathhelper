package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// returns /etc/nginx/sites-available as a string
func GetSitesAvailable() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.SitesAvailable.CombineWith(pathhelper.GetNginxLinuxPath())
}
