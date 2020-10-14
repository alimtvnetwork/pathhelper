package pathhelper

func GetAbsoluteFromExecutableDirectoryPath(relativePath string, isNormalize bool) string {
	basePath := GetExecutablePath()

	return GetAbsolutePath(basePath, relativePath, isNormalize)
}
