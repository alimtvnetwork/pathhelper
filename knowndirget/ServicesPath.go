package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns path to Services directory on different platforms.
func GetServicesPath() string {
	if osconsts.IsWindows {
		return knowndir.Services.CombineWith(EtcPath())
	}

	return GetSystemPath()
}
