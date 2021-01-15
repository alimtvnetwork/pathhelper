package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// "/etc/apache/"
func GetApacheLinuxPath() string {
	if osconsts.IsWindows {
		panic("Path only available for Unix OS") // todo test for panic
	}

	return knowndir.ApacheLinuxPath.Value()
}
