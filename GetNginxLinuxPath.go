package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// "/etc/nginx/"
func GetNginxLinuxPath() string {
	return enums.NginxLinuxPath.Value()
}
