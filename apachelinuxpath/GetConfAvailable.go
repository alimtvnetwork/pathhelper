package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/conf-available as a string
func GetConfAvailable() string {
	return enums.ConfAvailable.GetPrefixCombinedWith(pathhelper.GetApacheLinuxPath())
}
