package ispath

import "gitlab.com/evatix-go/pathhelper/internal/consts"

func Slug(path string) bool {
	if Empty(path) {
		return false
	}

	for _, char := range path {
		if consts.SlugHashset.Has(string(char)) {
			return false
		}
	}

	return true
}
