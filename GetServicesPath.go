package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to Services directory on different platforms.
func GetServicesPath() string {
	if osconsts.IsWindows {
		return knowndir.Services.CombineWith(GetEtcPath())
	}

	return GetSystemPath()
}
