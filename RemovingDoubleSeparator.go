package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func RemovingDoubleSeparator(path string) string {
	if strings.Contains(path, constants.DoubleBackSlash) {
		path = RemovingDoubleBackSlash(path)
	}

	if strings.Contains(path, constants.DoubleForwardSlash) {
		path = RemovingDoubleSlash(path)
	}

	if strings.Contains(path, constants.DoubleBackSlash) || strings.Contains(path, constants.DoubleForwardSlash) {
		RemovingDoubleSeparator(path)
	}

	return path
}
