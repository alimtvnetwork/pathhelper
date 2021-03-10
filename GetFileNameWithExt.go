package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithExt(currentPath string) (fileName string) {
	i := splitinternal.LastSlash(
		currentPath)

	return currentPath[i+1:]
}
