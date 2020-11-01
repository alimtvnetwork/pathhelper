package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/mods-enabled as a string
func GetModsEnabled() string {
	return enums.ModsEnabled.GetPrefixCombinedWith(pathhelper.GetApacheLinuxPath())
}
