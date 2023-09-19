package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"
)

func GetRoot() string {
	if osconsts.IsWindows {
		return WindowsRoot()
	}

	return UnixRoot()
}
