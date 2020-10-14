package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func ChangeDoubleBackSlash(path, changeSeparator string) string {
	return strings.ReplaceAll(path, constants.DoubleBackSlash, changeSeparator)
}
