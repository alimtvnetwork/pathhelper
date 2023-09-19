package normalize

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func ChangeDoubleBackSlash(path, changeSeparator string) string {
	return strings.ReplaceAll(path, constants.DoubleBackSlash, changeSeparator)
}
