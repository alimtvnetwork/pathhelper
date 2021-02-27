package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"

// returns base dir with slash
func Split(currentPath string) (baseDir, fileName string) {
	i := pathsplitinternal.LastSlash(currentPath)

	return currentPath[:i+1], currentPath[i+1:]
}
