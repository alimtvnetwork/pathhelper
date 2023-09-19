package pathhelper

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

func GetBaseDirNames(currentPath string) (baseDirNames *[]string) {
	return splitinternal.GetBaseDirNames(currentPath)
}
