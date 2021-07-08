package splitinternal

import "gitlab.com/evatix-go/core/constants"

// GetFileNameWithExt reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithExt(currentPath string) (fileName string) {
	i := LastSlash(
		currentPath)

	if i <= constants.Zero {
		return currentPath
	}

	return currentPath[i+1:]
}
