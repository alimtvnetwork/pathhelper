package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"

func GetWithoutSlash(currentPath string) (baseDir, fileName string) {
	i := pathsplitinternal.LastSlash(currentPath)

	return currentPath[:i], currentPath[i+1:]
}
