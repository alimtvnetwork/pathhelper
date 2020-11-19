package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/conf-enabled as a string
func GetConfEnabled() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ConfEnabled.CombineWith(pathhelper.GetApacheLinuxPath())
}
