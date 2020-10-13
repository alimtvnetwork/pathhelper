package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func IsSlug(path string) bool {
	if pathhelpercore.IsEmptyPath(path){
		return false
	}

	for i, _ := range forbiddenArray {
		if strings.Contains(path, forbiddenArray[i]) {
			return false
		}
	}

	return true
}
