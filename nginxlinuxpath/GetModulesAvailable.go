package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/modules-available as a string
func GetModulesAvailable() string {
	return enums.ModulesAvailable.GetPrefixCombinedWith(pathhelper.GetNginxLinuxPath())
}
