package nginxlinux

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

const (
	Conf             enums.KnownDirectory = "/etc/nginx/conf.d"
	ModulesAvailable enums.KnownDirectory = "/etc/nginx/modules-available"
	ModulesEnabled   enums.KnownDirectory = "/etc/nginx/modules-enabled"
	SitesAvailable   enums.KnownDirectory = "/etc/nginx/sites-available"
	SitesEnabled     enums.KnownDirectory = "/etc/nginx/sites-enabled"
	MimeTypes        enums.KnownDirectory = "/etc/nginx/mime.types"
)
