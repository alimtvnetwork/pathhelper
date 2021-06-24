package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// returns /etc/nginx/mime.types as a string
func GetMimeTypes() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.MimeTypes.CombineWith(knowndirget.NginxLinuxPath())
}
