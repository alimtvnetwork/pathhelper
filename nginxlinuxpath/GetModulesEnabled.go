package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/modules-enabled as a string
func GetModulesEnabled() string {
	return enums.ModulesEnabled.GetPrefixCombinedWith(pathhelper.GetNginxLinuxPath())
}
