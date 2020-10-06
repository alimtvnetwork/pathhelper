package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func RemovingDoubleSlash(path string) string {
	return strings.ReplaceAll(path, constants.DoubleSlash, constants.ForwardSlash)
}
