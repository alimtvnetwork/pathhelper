package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func HasPathIssues(stringToCheck string) bool {
	hasPrefix := strings.HasPrefix(stringToCheck, constants.UriSchemePrefixStandard) || strings.HasPrefix(stringToCheck, constants.UriSchemePrefixTwoSlashes)
	hasSlashAndBackSlash := strings.Contains(stringToCheck, constants.ForwardSlash) && strings.Contains(stringToCheck, constants.BackSlash)
	hasDouble := strings.Contains(stringToCheck, constants.DoubleBackSlash) || strings.Contains(stringToCheck, constants.DoubleForwardSlash)

	return hasPrefix || hasSlashAndBackSlash || hasDouble
}
