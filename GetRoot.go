package pathhelper

import "gitlab.com/evatix-go/core/osconsts"

func GetRoot() string {
	if osconsts.IsWindows {
		return GetWindowsRoot()
	}

	return GetUnixRoot()
}
