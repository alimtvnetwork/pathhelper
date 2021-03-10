package normalize

import "strings"

func isEmpty(path string) bool {
	return &path == nil || path == "" || len(path) == 0 || len(strings.TrimSpace(path)) == 0
}
