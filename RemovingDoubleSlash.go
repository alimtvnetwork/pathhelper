package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func RemovingDoubleSlash(path string) string {
	return strings.ReplaceAll(path, constants.DoubleForwardSlash, constants.ForwardSlash)
}
