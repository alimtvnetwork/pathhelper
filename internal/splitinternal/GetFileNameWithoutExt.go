package splitinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// invalid ext should return empty string.
// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func GetFileNameWithoutExt(currentPath string) (filename string) {
	i := LastSlash(
		currentPath)

	filename = currentPath[i+1:]
	indexOfDot := strings.Index(
		filename,
		constants.Dot)

	if indexOfDot > constants.InvalidNotFoundCase {
		return filename[:indexOfDot]
	}

	return filename
}
