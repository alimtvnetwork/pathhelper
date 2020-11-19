package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/conf.d as a string
func GetConf() string {
	if !pathhelper.IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.Conf.CombineWith(pathhelper.GetNginxLinuxPath())
}
