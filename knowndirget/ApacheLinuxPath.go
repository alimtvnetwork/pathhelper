package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

// ApacheLinuxPath "/etc/apache/"
func ApacheLinuxPath() string {
	if osconsts.IsWindows {
		panic("Location only available for Unix OS") // todo test for panic
	}

	return knowndir.ApacheLinuxPath.Value()
}
