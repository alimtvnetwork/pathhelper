package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"
)

// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func SplitWithoutSlash(currentPath string) (baseDir, fileName string) {
	i := pathsplitinternal.LastSlash(currentPath)

	return currentPath[:i], currentPath[i+1:]
}
