package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

// ParentDir alias for GetBaseDir
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func ParentDir(currentPath string) (baseDir string) {
	return splitinternal.GetBaseDir(
		currentPath)
}
