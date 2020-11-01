package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/sites-enabled as a string
func SitesEnabled() string {
	return enums.SitesEnabled.GetPrefixCombinedWith(pathhelper.GetNginxLinuxPath())
}
