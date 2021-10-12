package nginxlinuxpath

import "gitlab.com/evatix-go/core/filemode"

const (
	DefaultRoot           = "/etc/nginx"
	RootConfigName        = "nginx.conf"
	ConfigAvailableName   = "config-available"
	ConfigEnabledName     = "config-enabled"
	SitesBackup           = "sites-backup"
	SitesAvailableName    = "sites-available"
	SitesEnabledName      = "sites-enabled"
	ExtraConfName         = "conf.d"
	ModulesAvailableName  = "modules-available"
	ModulesEnabledName    = "modules-enabled"
	Users                 = "users"
	WildcardDotConfFilter = "*.conf"
	DefaultFileChmod      = filemode.X644
	DefaultDirChmod       = filemode.X644
)
