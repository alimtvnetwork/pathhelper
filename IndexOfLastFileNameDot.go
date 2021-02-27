package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"
)

// returns -1 if last file name doesn't have any extension
// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func IndexOfLastFileNameDot(currentPath string) (fileName string, index int) {
	i := pathsplitinternal.LastSlash(currentPath)
	fileName = currentPath[i+1:]

	return fileName, strings.Index(
		fileName,
		constants.Dot)
}
