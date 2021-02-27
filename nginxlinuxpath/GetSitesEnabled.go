package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// returns /etc/nginx/sites-enabled as a string
func GetSitesEnabled() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.SitesEnabled.CombineWith(knowndirget.NginxLinuxPath())
}
