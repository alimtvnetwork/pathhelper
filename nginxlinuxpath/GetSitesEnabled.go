package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/sites-enabled as a string
func GetSitesEnabled() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.SitesEnabled.CombineWith(pathhelper.GetNginxLinuxPath())
}
