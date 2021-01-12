package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// "/etc/nginx/"
func GetNginxLinuxPath() string {
	if !IsUnix() {
		return ""
	}

	return enums.NginxLinuxPath.Value()
}
