package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/conf.d as a string
func GetConf() string {
	return enums.Conf.GetPrefixCombinedWith(pathhelper.GetNginxLinuxPath())
}
