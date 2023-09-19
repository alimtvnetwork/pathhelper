package nginxlinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetModulesEnabled
//
//	returns /etc/nginx/modules-enabled as a string
func GetModulesEnabled() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.ModulesEnabled.CombineWith(knowndirget.NginxLinuxPath())
}
