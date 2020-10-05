package pathhelper

import "strings"

func GetAbsolutePath(basePath, relativePath string) string {
	if !strings.HasSuffix(basePath, GetPathSeparator()) || !strings.HasPrefix(relativePath, GetPathSeparator()) {
		return RemovingDouble(RemovingDouble(basePath) + GetPathSeparator() + RemovingDouble(relativePath))
	}

	return RemovingDouble(RemovingDouble(basePath) + RemovingDouble(relativePath))
}
