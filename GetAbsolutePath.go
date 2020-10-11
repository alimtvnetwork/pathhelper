package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func GetAbsolutePath(basePath, relativePath string) string {
	if pathhelpercore.IsEmptyPath(basePath) && pathhelpercore.IsEmptyPath(relativePath) {
		panic("Empty paths provided!")
	}

	if !strings.HasSuffix(basePath, GetPathSeparator()) || !strings.HasPrefix(relativePath, GetPathSeparator()) {
		return RemovingDoubleSeparator(RemovingDoubleSeparator(basePath) + GetPathSeparator() + RemovingDoubleSeparator(relativePath))
	}

	return RemovingDoubleSeparator(RemovingDoubleSeparator(basePath) + RemovingDoubleSeparator(relativePath))
}
