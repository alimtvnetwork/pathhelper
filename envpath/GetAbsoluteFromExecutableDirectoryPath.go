package envpath

import "gitlab.com/evatix-go/pathhelper"

func GetAbsoluteFromExecutableDirectoryPath(relativePath string, isLongPathFix, isNormalize bool) string {
	basePath := GetExecutablePath()

	return pathhelper.GetAbsolutePath(
		basePath,
		relativePath,
		isLongPathFix,
		isNormalize)
}
