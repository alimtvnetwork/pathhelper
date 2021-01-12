package pathhelper

import (
	"gitlab.com/evatix-go/core/constants"
)

// Replace both double slashes to single slash (// -> /, \\ -> \) and finally all slashes to os.Separator
func RemoveAndFixDoubleSeparatorToOsSeparator(path string) string {
	return RemoveAndFixDoubleSeparatorToFinalSeparator(
		constants.PathSeparator,
		path)
}
