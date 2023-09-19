package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// "/etc/nginx/"
func NginxLinuxPath() string {
	if osconsts.IsWindows {
		return ""
	}

	return knowndir.NginxLinuxPath.Value()
}
