package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/knowndirstructure"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetConf returns /etc/nginx/conf.d as a string
func GetConf() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.Conf.CombineWith(knowndirget.NginxLinuxPath())
}

func GetFullDirStructure(
	isNormalize bool,
	currentNginxRoot string,
) *knowndirstructure.NginxApacheDirectory {
	return &knowndirstructure.NginxApacheDirectory{
		Root:             normalizeinternal.FixIf(isNormalize, currentNginxRoot, ""),
		RootConfigFile:   fixPath(isNormalize, currentNginxRoot, NginxRootConfigName),
		ConfigAvailable:  fixPath(isNormalize, currentNginxRoot, ConfigAvailableName),
		ConfigEnabled:    fixPath(isNormalize, currentNginxRoot, ConfigEnabledName),
		SitesAvailable:   fixPath(isNormalize, currentNginxRoot, SitesAvailableName),
		SitesEnabled:     fixPath(isNormalize, currentNginxRoot, SitesEnabledName),
		ExtraConfig:      fixPath(isNormalize, currentNginxRoot, ExtraConfName),
		ModulesAvailable: fixPath(isNormalize, currentNginxRoot, ModulesAvailableName),
		ModulesEnabled:   fixPath(isNormalize, currentNginxRoot, ModulesEnabledName),
	}
}
