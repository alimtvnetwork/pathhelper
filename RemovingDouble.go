package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func RemovingDouble(path string) string {
	if strings.Contains(path, constants.DoubleBackSlash) {
		path = RemovingDoubleBackSlash(path)
	}

	if strings.Contains(path, constants.DoubleSlash) {
		path = RemovingDoubleSlash(path)
	}

	// checking for more double -- REQUIRED or not
	if strings.Contains(path, constants.DoubleBackSlash) || strings.Contains(path, constants.DoubleSlash) {
		RemovingDouble(path)
	}

	return path
}
