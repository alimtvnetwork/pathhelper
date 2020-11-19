package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/conf-available as a string
func GetConfAvailable() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ConfAvailable.CombineWith(pathhelper.GetApacheLinuxPath())
}
