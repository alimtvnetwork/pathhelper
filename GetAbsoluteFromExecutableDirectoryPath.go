package pathhelper

func GetAbsoluteFromExecutableDirectoryPath(relativePath string) string {
	basePath := GetExecutablePath()

	return GetAbsolutePath(basePath, relativePath)
}
