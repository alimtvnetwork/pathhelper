package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/mods-available as a string
func GetModsAvailable() string {
	return enums.ModsAvailable.GetPrefixCombinedWith(pathhelper.GetApacheLinuxPath())
}
