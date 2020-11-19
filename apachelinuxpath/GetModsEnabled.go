package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/mods-enabled as a string
func GetModsEnabled() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ModsEnabled.CombineWith(pathhelper.GetApacheLinuxPath())
}
