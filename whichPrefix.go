package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
)

func whichPrefix(stringToCheck string) enums.UriSchemes {
	if strings.HasPrefix(stringToCheck, constants.UriSchemePrefixStandard) {
		return enums.UriSchemePrefixStandard
	}

	if strings.HasPrefix(stringToCheck, constants.UriSchemePrefixTwoSlashes) {
		return enums.UriSchemePrefixTwoSlashes
	}

	return enums.UriUnknown
}
