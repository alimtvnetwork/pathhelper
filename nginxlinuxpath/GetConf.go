package nginxlinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetConf returns /etc/nginx/conf.d as a string
func GetConf() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.Conf.CombineWith(knowndirget.NginxLinuxPath())
}
