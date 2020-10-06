package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func HasPathIssues(stringToCheck string) bool {
	hasPrefix := strings.HasPrefix(stringToCheck, constants.UriSchemePrefixStandard)
	hasSlashAndBackSlash := strings.Contains(stringToCheck, constants.ForwardSlash) && strings.Contains(stringToCheck, constants.BackSlash)

	return hasPrefix || hasSlashAndBackSlash
}
