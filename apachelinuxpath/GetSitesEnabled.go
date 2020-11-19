package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/sites-enabled as a string
func GetSitesEnabled() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.SitesEnabled.CombineWith(pathhelper.GetApacheLinuxPath())
}
