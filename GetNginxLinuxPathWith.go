package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// "/etc/nginx/" + directory.Value()
func GetNginxLinuxPathWith(directory enums.KnownDirectory) string {
	return GetCombinePathsWith(GetNginxLinuxPath(), directory.Value())
}
