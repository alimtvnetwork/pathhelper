package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/sites-available as a string
func GetSitesAvailable() string {
	return enums.SitesAvailable.GetPrefixCombinedWith(pathhelper.GetNginxLinuxPath())
}
