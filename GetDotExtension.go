package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func GetDotExtension(
	currentPath string,
) string {
	dotExt, _ := splitinternal.GetExtensions(
		currentPath)

	return dotExt
}
