package pathhelpercore

import "strings"

func IsEmptyPath(path string) bool {
	return &path == nil || path == "" || len(strings.TrimSpace(path)) == 0
}
