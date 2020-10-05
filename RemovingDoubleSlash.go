package pathhelper

import "strings"

func RemovingDoubleSlash(path string) string {
	return strings.ReplaceAll(path, DoubleSlash, Slash)
}
