package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// "/etc/nginx/" + directory.Value()
func GetNginxLinuxPathWith(directory knowndir.Alias) string {
	return pathhelper.GetCombinePathsWith(NginxLinuxPath(), directory.Value())
}
