package pathhelper

import "strings"

func RemovingDouble(path string) string {
	if strings.Contains(path, DoubleBackSlash) {
		path = RemovingDoubleBackSlash(path)
	}

	if strings.Contains(path, DoubleSlash) {
		path = RemovingDoubleSlash(path)
	}

	// checking for more double -- REQUIRED or not
	if strings.Contains(path, DoubleBackSlash) || strings.Contains(path, DoubleSlash) {
		RemovingDouble(path)
	}

	return path
}
