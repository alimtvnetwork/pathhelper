package pathhelper

import (
	"strings"
)

func IsSlug(path string) bool {
	for i, _ := range forbiddenArray {
		if strings.Contains(path, forbiddenArray[i]) {
			return false
		}
	}

	return true
}
