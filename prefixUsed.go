package pathhelper

import "strings"

// todo if  needed
func prefixUsed(stringToCheck string) string {
	if strings.HasPrefix(stringToCheck, Prefix) {
		return Prefix
	}

	return Prefix
}
