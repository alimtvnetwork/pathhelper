package pathhelpercore

import "strings"

func IsEmptyPath(path string) bool {
	return &path == nil || path == "" || len(strings.TrimSpace(path)) == 0
}

func IsEmptyPathForPtr(path *string) bool {
	return path == nil || &path == nil || IsEmptyPath(*path)
}
