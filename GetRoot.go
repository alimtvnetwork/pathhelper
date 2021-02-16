package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

func GetRoot() string {
	if osconsts.IsWindows {
		return knowndirget.WindowsRoot()
	}

	return knowndirget.UnixRoot()
}
