package pathhelper

import "strings"

func HasPathIssues(stringToCheck string) bool {
	hasPrefix := strings.HasPrefix(stringToCheck, Prefix)
	hasSlashAndBackSlash := strings.Contains(stringToCheck, Slash) && strings.Contains(stringToCheck, BackSlash)

	return hasPrefix || hasSlashAndBackSlash
}
