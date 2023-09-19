package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns path to System drivers directory on different platforms.
func GetSystemDriversPath() string {
	if osconsts.IsWindows {
		return knowndir.Drivers.CombineWith(System32())
	}

	return knowndir.DriversUnix.Value()
}
