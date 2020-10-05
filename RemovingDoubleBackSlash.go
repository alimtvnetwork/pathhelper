package pathhelper

import "strings"

func RemovingDoubleBackSlash(path string) string {
	return strings.ReplaceAll(path, DoubleBackSlash, BackSlash)
}
