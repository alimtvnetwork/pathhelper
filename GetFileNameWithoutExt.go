package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"
)

// invalid ext should return empty string.
// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func GetFileNameWithoutExt(currentPath string) (filename string) {
	i := pathsplitinternal.LastSlash(
		currentPath)

	filename = currentPath[i+1:]
	indexOfDot := strings.Index(
		filename,
		constants.Dot)

	if indexOfDot > -1 {
		return filename[:indexOfDot]
	}

	return filename
}
