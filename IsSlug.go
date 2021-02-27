package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

func IsSlug(path string) bool {
	if ispath.Empty(path) {
		return false
	}

	for i := range slugForbiddenArray {
		if strings.Contains(path, slugForbiddenArray[i]) {
			return false
		}
	}

	return true
}
