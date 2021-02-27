package normalize

import "strings"

func IsEmptyPath(path string) bool {
	return &path == nil || path == "" || len(path) == 0 || len(strings.TrimSpace(path)) == 0
}

func IsEmptyPathPtr(path *string) bool {
	return path == nil || &path == nil || IsEmptyPath(*path)
}
