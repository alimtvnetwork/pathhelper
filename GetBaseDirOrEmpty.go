package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

func GetBaseDirOrEmpty(currentPath string) (baseDir string) {
	return splitinternal.GetBaseDirNameOrEmpty(
		currentPath)
}
