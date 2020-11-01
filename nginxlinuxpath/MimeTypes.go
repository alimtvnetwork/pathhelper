package nginxlinuxpath

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// returns /etc/nginx/mime.types as a string
func MimeTypes() string {
	return enums.MimeTypes.GetPrefixCombinedWith(pathhelper.GetNginxLinuxPath())
}
