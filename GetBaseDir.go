package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

// No slash at the end
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetBaseDir(currentPath string) (baseDir string) {
	i := splitinternal.LastSlash(
		currentPath)

	return currentPath[:i]
}
