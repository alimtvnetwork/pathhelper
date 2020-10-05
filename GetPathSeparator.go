package pathhelper

import "path/filepath"

func GetPathSeparator() string {
	return string(filepath.Separator)
}
