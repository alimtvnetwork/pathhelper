package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

// invalid ext should return empty string.
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithoutExt(currentPath string) (filename string) {
	i := splitinternal.LastSlash(
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
