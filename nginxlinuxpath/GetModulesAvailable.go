package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/modules-available as a string
func GetModulesAvailable() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ModulesAvailable.CombineWith(pathhelper.GetNginxLinuxPath())
}
