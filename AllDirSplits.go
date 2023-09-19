package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

func AllDirSplits(currentPath string) (baseDirNames *[]string) {
	return splitinternal.GetAllSplits(currentPath)
}
