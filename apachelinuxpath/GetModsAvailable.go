package apachelinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/apache/mods-available as a string
func GetModsAvailable() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ModsAvailable.CombineWith(pathhelper.GetApacheLinuxPath())
}
