package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"

// No slash at the end
// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func GetBaseDir(currentPath string) (baseDir string) {
	i := pathsplitinternal.LastSlash(
		currentPath)

	return currentPath[:i]
}
