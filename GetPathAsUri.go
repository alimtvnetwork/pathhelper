package pathhelper

import "strings"

func GetPathAsUri(path string) string {
	if IsWindows() {
		return PrefixForWindowsURI + strings.ReplaceAll(NormalizePath(path), Slash, BackSlash)
	}

	return Prefix + strings.ReplaceAll(NormalizePath(path), BackSlash, Slash)
}
