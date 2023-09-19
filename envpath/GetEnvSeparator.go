package envpath

import "gitlab.com/auk-go/core/osconsts"

func GetEnvSeparator() string {
	if osconsts.IsWindows {
		return windowsEnvPathSplitter
	}

	return unixEnvPathSplitter
}
