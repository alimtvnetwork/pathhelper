package normalize

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func ChangeDoubleForwardSlash(path, changeSeparator string) string {
	return strings.ReplaceAll(path, constants.DoubleForwardSlash, changeSeparator)
}
