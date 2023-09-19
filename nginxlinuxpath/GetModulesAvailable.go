package nginxlinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetModulesAvailable
//
//	returns /etc/nginx/modules-available as a string
func GetModulesAvailable() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.ModulesAvailable.CombineWith(
		knowndirget.NginxLinuxPath())
}
