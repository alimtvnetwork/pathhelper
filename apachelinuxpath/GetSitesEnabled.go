package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/sites-enabled as a string
func GetSitesEnabled() string {
	return enums.SitesEnabled.GetPrefixCombinedWith(pathhelper.GetApacheLinuxPath())
}
