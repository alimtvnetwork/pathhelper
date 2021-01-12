package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/sites-available as a string
func GetSitesAvailable() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.SitesAvailable.CombineWith(pathhelper.GetApacheLinuxPath())
}
