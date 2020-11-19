package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/modules-enabled as a string
func GetModulesEnabled() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ModulesEnabled.CombineWith(pathhelper.GetNginxLinuxPath())
}
