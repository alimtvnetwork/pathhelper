package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func RemovingDoubleBackSlash(path string) string {
	return strings.ReplaceAll(path, constants.DoubleBackSlash, constants.BackSlash)
}
