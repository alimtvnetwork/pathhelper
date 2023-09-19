package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

// GetFileNameWithExt reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithExt(currentPath string) (fileName string) {
	return splitinternal.GetName(
		currentPath)
}
