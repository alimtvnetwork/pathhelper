package pathhelper

import (
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

func GetFilenamePlusBothExtensions(
	currentPath string,
) (
	fileName,
	dotExt,
	ext string,
) {
	return splitinternal.
		GetFilenamePlusBothExtensions(
			currentPath)
}
