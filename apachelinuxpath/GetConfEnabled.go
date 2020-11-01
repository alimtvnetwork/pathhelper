package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/conf-enabled as a string
func GetConfEnabled() string {
	return enums.ConfEnabled.GetPrefixCombinedWith(pathhelper.GetApacheLinuxPath())
}
