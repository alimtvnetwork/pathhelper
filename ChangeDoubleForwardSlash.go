package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func ChangeDoubleForwardSlash(path, changeSeparator string) string {
	return strings.ReplaceAll(path, constants.DoubleForwardSlash, changeSeparator)
}
