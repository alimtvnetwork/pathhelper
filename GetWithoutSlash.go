package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

func GetWithoutSlash(currentPath string) (baseDir, fileName string) {
	i := splitinternal.LastSlash(currentPath)

	return currentPath[:i], currentPath[i+1:]
}
