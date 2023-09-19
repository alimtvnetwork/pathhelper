package pathhelper

import (
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

// reference example : https://play.golang.org/p/oT6eWNZAeEi
func SplitWithoutSlash(currentPath string) (baseDir, fileName string) {
	i := splitinternal.LastSlash(currentPath)

	return currentPath[:i], currentPath[i+1:]
}
