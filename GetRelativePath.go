package pathhelper

import (
	"strings"
)

func GetRelativePath(fullPath, basePath string) string {
	if !strings.Contains(fullPath, basePath) {
		return ""
	}

	if strings.Compare(fullPath, basePath) == 0 {
		return "Both paths are same"
	}

	return strings.Replace(fullPath, basePath, "", 1)
}
