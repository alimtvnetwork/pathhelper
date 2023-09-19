package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

func GetDotExtension(
	currentPath string,
) string {
	dotExt, _ := splitinternal.GetBothExtension(
		currentPath)

	return dotExt
}
