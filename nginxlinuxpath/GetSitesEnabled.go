package nginxlinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetSitesEnabled
//
//	returns /etc/nginx/sites-enabled as a string
func GetSitesEnabled() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.SitesEnabled.CombineWith(knowndirget.NginxLinuxPath())
}
