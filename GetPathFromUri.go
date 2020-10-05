package pathhelper

import "strings"

func GetPathFromUri(path string) string {
	if IsWindows() {
		return strings.ReplaceAll(NormalizePath(path), Prefix, "")
	}

	// Todo  check prefix difference for other OS
	return strings.ReplaceAll(NormalizePath(path), Prefix, "")
}
