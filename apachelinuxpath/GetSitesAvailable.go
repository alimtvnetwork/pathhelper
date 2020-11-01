package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/sites-available as a string
func GetSitesAvailable() string {
	return enums.SitesAvailable.GetPrefixCombinedWith(pathhelper.GetApacheLinuxPath())
}
