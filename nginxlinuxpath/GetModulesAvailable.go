package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// returns /etc/nginx/modules-available as a string
func GetModulesAvailable() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS")
	}

	return knowndir.ModulesAvailable.CombineWith(knowndirget.NginxLinuxPath())
}
