package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func SplitWithoutSlash(currentPath string) (baseDir, fileName string) {
	i := splitinternal.LastSlash(currentPath)

	return currentPath[:i], currentPath[i+1:]
}
