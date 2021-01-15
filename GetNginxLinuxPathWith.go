package pathhelper

import "gitlab.com/evatix-go/pathhelper/knowndir"

// "/etc/nginx/" + directory.Value()
func GetNginxLinuxPathWith(directory knowndir.Alias) string {
	return GetCombinePathsWith(GetNginxLinuxPath(), directory.Value())
}
