package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

func IsSlug(path string) bool {
	if ispath.Empty(path) {
		return false
	}

	for i := range forbiddenArray {
		if strings.Contains(path, forbiddenArray[i]) {
			return false
		}
	}

	return true
}
