package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to bin directory as a string.
func BinPath() string {
	if !osconsts.IsWindows {
		return knowndir.BinUnix.Value()
	}

	binPath := knowndir.Bin.CombineWith(UserPath())

	return binPath
}
