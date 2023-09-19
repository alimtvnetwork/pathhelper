package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// BinPath Returns path to bin directory as a string.
func BinPath() string {
	if osconsts.IsUnixGroup {
		return knowndir.BinUnix.Value()
	}

	binPath := knowndir.Bin.CombineWith(UserPath())

	return binPath
}
