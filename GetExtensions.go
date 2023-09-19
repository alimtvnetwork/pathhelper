package pathhelper

import (
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

func GetExtensions(
	currentPath string,
) (
	dotExt, ext string,
) {
	return splitinternal.GetBothExtension(
		currentPath)
}
